package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-runner-api/internal/api"
	"github.com/ozoncp/ocp-runner-api/internal/broker"
	"github.com/ozoncp/ocp-runner-api/internal/config"
	"github.com/ozoncp/ocp-runner-api/internal/metrics"
	"github.com/ozoncp/ocp-runner-api/internal/repo"
	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

const (
	configPath = "config.yml"
)

type context struct {
	config     *config.Config
	repo       repo.Repo
	producer   broker.Producer
	consumer   broker.Consumer
	grpcServer *grpc.Server
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp})
}

func main() {
	ctx := &context{}

	if err := readConfig(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := connectToDatabase(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}
	defer ctx.repo.Close()

	if err := connectToBroker(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}
	defer func() { ctx.producer.Close(); ctx.consumer.Close() }()

	ctx.grpcServer = grpc.NewServer()
	defer ctx.grpcServer.GracefulStop()

	ec := make(chan error)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, os.Interrupt)

	go runGrpc(ctx, ec)
	go runMetrics(ctx, ec)

	select {
	case err := <-ec:
		log.Error().Err(err).Send()
	case <-sc:
		log.Info().Msg("terminating service...")
	}
}

// runGrpc runs gRPC server
func runGrpc(ctx *context, ec chan<- error) {
	listen, err := net.Listen("tcp", ctx.config.GrpcPort)
	if err != nil {
		ec <- err
		return
	}

	server.RegisterOcpRunnerServiceServer(ctx.grpcServer, api.NewRunnerApi(ctx.repo, ctx.producer))
	log.Info().Str("gRPC server started at", ctx.config.GrpcPort).Send()

	if err := ctx.grpcServer.Serve(listen); err != nil {
		ec <- err
	}
}

// runMetrics runs prometheus
func runMetrics(ctx *context, ec chan<- error) {
	metrics.RegisterMetrics()

	http.Handle(ctx.config.MetricsHandle, promhttp.Handler())
	if err := http.ListenAndServe(ctx.config.MetricsPort, nil); err != nil {
		ec <- err
	}
}

// readConfig reads config
func readConfig(ctx *context) error {
	cfg, err := config.Read(configPath)
	if err != nil {
		return err
	}
	ctx.config = cfg
	return nil
}

// connectToDatabase initializes database connection
func connectToDatabase(ctx *context) error {
	log.Info().Msg("connection to database...")

	db, err := sqlx.Connect("postgres", ctx.config.Database)
	if err != nil {
		return err
	}

	log.Info().Msg("successfully connected!")
	ctx.repo = repo.New(db)

	return nil
}

// connectToBroker connects to kafka broker
func connectToBroker(ctx *context) error {
	prod := broker.NewProducer()
	if err := prod.Init(ctx.config.KafkaBrokers); err != nil {
		return err
	}

	cons := broker.NewConsumer()
	if err := cons.Init(ctx.config.KafkaBrokers); err != nil {
		return err
	}
	if err := cons.Subscribe("events"); err != nil {
		return err
	}

	ctx.producer = prod
	ctx.consumer = cons

	return nil
}
