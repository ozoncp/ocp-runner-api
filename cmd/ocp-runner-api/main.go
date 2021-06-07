package main

import (
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"net"
	"os"
	"sync"
	"time"

	"github.com/ozoncp/ocp-runner-api/internal/api"
	"github.com/rs/zerolog/log"

	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
	"google.golang.org/grpc"
)

type Config struct {
	Project    string `yaml:"project"`
	Version    string `yaml:"version"`
	Author     string `yaml:"author"`
	HttpPort   string `yaml:"http_port"`
	GrpcPort   string `yaml:"grpc_port"`
	SwaggerDir string `yaml:"swagger_dir"`
}

var wg = &sync.WaitGroup{}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp})
}

func main() {
	config := &Config{}

	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatal().Str("message", "failed to open config file").Err(err).Send()
		os.Exit(1)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal().Str("message", "failed to parse config file").Err(err).Send()
		os.Exit(1)
	}

	wg.Add(1)
	go func() {
		if err := runGrpc(config); err != nil {
			log.Fatal().Str("message", "failed to start grpc server").Err(err).Send()
			wg.Done()
		}
	}()
	wg.Wait()
}

// runGrpc runs gRPC server
func runGrpc(config *Config) error {
	listen, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.Fatal().Err(err).Send()
		return err
	}

	s := grpc.NewServer()
	server.RegisterOcpRunnerApiServer(s, api.NewRunnerApi())
	log.Info().Str("gRPC server started at", config.GrpcPort).Send()

	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Send()
		return err
	}

	return nil
}
