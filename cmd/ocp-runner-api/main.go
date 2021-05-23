package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Project string `yaml:"project"`
	Version string `yaml:"version"`
	Author  string `yaml:"author"`
}

func main() {
	config := &Config{}

	path, _ := filepath.Abs("./config.yml")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Errorf("failed to close file, reason: %w", err)
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	fmt.Printf("%v by %v", config.Project, config.Author)
}
