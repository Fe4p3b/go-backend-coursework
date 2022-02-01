package config

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0:8080"`
	BaseURL string `env:"BASE_URL,required" envDefault:"http://localhost:8080"`
}

func (c *Config) Read() error {
	err := env.Parse(c)
	if err != nil {
		return err
	}

	var (
		address string
		baseURL string
	)

	flag.StringVar(&address, "a", "localhost:8080", "Адрес запуска HTTP-сервера")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "Базовый адрес результирующего сокращённого URL")
	flag.Parse()

	if address != "localhost:8080" {
		c.Address = address
	}

	if baseURL != "http://localhost:8080" {
		c.BaseURL = baseURL
	}

	readCfg, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	log.Printf("config: %s", readCfg)

	return nil
}
