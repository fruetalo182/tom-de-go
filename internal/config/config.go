package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port      string
	ImagePath string
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	imagePath := os.Getenv("IMAGE_PATH")
	if imagePath == "" {
		return nil, fmt.Errorf("IMAGE_PATH environment variable is not set")
	}

	return &Config{
		Port:      port,
		ImagePath: imagePath,
	}, nil
}