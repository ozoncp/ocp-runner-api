package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Project       string   `yaml:"project"`
	Version       string   `yaml:"version"`
	Author        string   `yaml:"author"`
	HttpPort      string   `yaml:"http_port"`
	GrpcPort      string   `yaml:"grpc_port"`
	MetricsPort   string   `yaml:"metrics_port"`
	MetricsHandle string   `yaml:"metrics_handle"`
	SwaggerDir    string   `yaml:"swagger_dir"`
	Database      string   `yaml:"database"`
	KafkaBrokers  []string `yaml:"kafka_brokers"`
}

// Read reads config from file
func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal().Str("message", "failed to open config file").Err(err).Send()
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal().Str("message", "failed to parse config file").Err(err).Send()
		return nil, err
	}

	return config, nil
}
