package config

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port        string `env:"PORT,required" envDefault:"8080"`
	Address     string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0"`
	BaseURL     string `env:"BASE_URL,required" envDefault:"http://localhost:8080"`
	DatabaseDSN string `env:"DATABASE_URL,required" envDefault:"postgres://gopher:12345@postgres:5432/shortener"`
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
		port        string
	)

	flag.StringVar(&address, "a", "", "Адрес запуска HTTP-сервера")
	flag.StringVar(&port, "p", "", "Порт запуска HTTP-сервера")
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

	if port != "" {
		c.Port = port
	}

	readCfg, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	log.Printf("config: %s", readCfg)

	return nil
}
