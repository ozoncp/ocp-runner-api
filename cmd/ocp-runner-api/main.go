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

	db, err := connectToDB(cfg)
	if err != nil {
		os.Exit(0)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Send()
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, os.Interrupt)

	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	ec := make(chan error)
	go runGrpc(grpcServer, db, cfg, ec)

	select {
	case err := <-ec:
		log.Error().Err(err).Send()
	case <-sc:
		log.Info().Msg("terminating service...")
	}
}

// runGrpc runs gRPC server
func runGrpc(grpcServer *grpc.Server, db *sqlx.DB, config *config.Config, ec chan<- error) {
	listen, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		ec <- err
		return
	}

	repository := repo.New(db)
	server.RegisterOcpRunnerServiceServer(grpcServer, api.NewRunnerApi(repository))
	log.Info().Str("gRPC server started at", config.GrpcPort).Send()

	if err := grpcServer.Serve(listen); err != nil {
		ec <- err
	}
}

// connectToDB initializes database connection
func connectToDB(config *config.Config) (*sqlx.DB, error) {
	log.Info().Msg("db: connecting...")

	db, err := sqlx.Connect("postgres", config.Database)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("db: successfully connected")
	return db, nil
}
