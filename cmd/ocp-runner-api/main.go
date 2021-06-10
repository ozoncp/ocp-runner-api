package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-runner-api/internal/api"
	"github.com/ozoncp/ocp-runner-api/internal/config"
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
	}

	server.RegisterOcpRunnerServiceServer(grpcServer, api.NewRunnerApi())
	log.Info().Str("gRPC server started at", config.GrpcPort).Send()

	if err := grpcServer.Serve(listen); err != nil {
		ec <- err
	}
}
