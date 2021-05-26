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

	openFileLoop(path)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	fmt.Printf("%v by %v", config.Project, config.Author)
}

func openFileLoop(path string) {
	for i := 0; i < 5; i++ {
		file, err := os.Open(path)
		if err != nil {
			fmt.Errorf("failed to open config file, error: %w", err)
			break
		}
		defer file.Close()

		fi, _ := file.Stat()
		fmt.Printf("%v. file is %d bytes long\n", i+1, fi.Size())
	}
}
