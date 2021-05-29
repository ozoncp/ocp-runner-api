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
		fmt.Printf("failed to open config file, error: %v\n", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("failed to close file, error: %v\n", closeErr)
		}
	}(file)

	openConfigLoop(path)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Printf("failed to parse config file, error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v by %v", config.Project, config.Author)
}

func openConfigLoop(path string) {
	for i := 0; i < 5; i++ {
		func() {
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("failed to open config file, error: %v\n", err)
				return
			}
			defer func(file *os.File) {
				if closeErr := file.Close(); closeErr != nil {
					fmt.Printf("failed to close file, error: %v\n", closeErr)
				}
			}(file)

			fi, _ := file.Stat()
			fmt.Printf("%v. file is %d bytes long\n", i+1, fi.Size())
			fmt.Println("for test 1")
		}()
	}
}
