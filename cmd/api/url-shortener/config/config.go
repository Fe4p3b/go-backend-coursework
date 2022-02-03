package config

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address     string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0:8080"`
	BaseURL     string `env:"BASE_URL,required" envDefault:"http://localhost:8080"`
	DatabaseDSN string `env:"DATABASE_DSN,required" envDefault:"postgres://gopher:12345@postgres:5432/shortener"`
}

func (c *Config) Read() error {
	err := env.Parse(c)
	if err != nil {
		return err
	}

	var (
		address     string
		baseURL     string
		databaseDSN string
	)

	flag.StringVar(&address, "a", "", "Адрес запуска HTTP-сервера")
	flag.StringVar(&baseURL, "b", "", "Базовый адрес результирующего сокращённого URL")
	flag.StringVar(&databaseDSN, "d", "", "Строка с адресом подключения к БД")
	flag.Parse()

	if address != "" {
		c.Address = address
	}

	if baseURL != "" {
		c.BaseURL = baseURL
	}

	if databaseDSN != "" {
		c.DatabaseDSN = databaseDSN
	}

	readCfg, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	log.Printf("config: %s", readCfg)

	return nil
}
