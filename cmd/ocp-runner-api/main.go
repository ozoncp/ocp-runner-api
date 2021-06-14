package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-runner-api/internal/api"
	"github.com/ozoncp/ocp-runner-api/internal/broker"
	"github.com/ozoncp/ocp-runner-api/internal/config"
	"github.com/ozoncp/ocp-runner-api/internal/repo"
	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp})
}

func main() {
	cfg, err := config.Read("config.yml")
	if err != nil {
		os.Exit(1)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, os.Interrupt)

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	ec := make(chan error)
	go runGrpc(grpcServer, cfg, ec)

	select {
	case err := <-ec:
		log.Error().Err(err).Send()
	case <-sc:
		log.Info().Msg("terminating service...")
	}
}

// runGrpc runs gRPC server
func runGrpc(grpcServer *grpc.Server, config *config.Config, ec chan<- error) {
	listen, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		ec <- err
		return
	}

	db, err := connectDatabase(config)
	if err != nil {
		ec <- err
		return
	}
	defer db.Close()

	prod, cons, err := initBrokers(config)
	if err != nil {
		ec <- err
		return
	}
	defer func() { prod.Close(); cons.Close() }()

	repository := repo.New(db)
	server.RegisterOcpRunnerServiceServer(grpcServer, api.NewRunnerApi(repository, prod))
	log.Info().Str("gRPC server started at", config.GrpcPort).Send()

	if err := grpcServer.Serve(listen); err != nil {
		ec <- err
	}
}

// connectDatabase initializes DB connection
func connectDatabase(config *config.Config) (*sqlx.DB, error) {
	log.Info().Msg("db: connecting...")

	db, err := sqlx.Connect("postgres", config.Database)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Info().Msg("db: successfully connected")
	return db, nil
}

// initBrokers initialize kafka brokers
func initBrokers(config *config.Config) (broker.Producer, broker.Consumer, error) {
	prod := broker.NewProducer()
	if err := prod.Init(config.KafkaBrokers); err != nil {
		return nil, nil, err
	}

	cons := broker.NewConsumer()
	if err := cons.Init(config.KafkaBrokers); err != nil {
		return prod, nil, err
	}
	if err := cons.Subscribe("events"); err != nil {
		return prod, nil, err
	}

	return prod, cons, nil
}
