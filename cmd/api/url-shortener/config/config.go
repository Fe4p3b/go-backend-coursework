package config

import (
	"encoding/json"
	"flag"

	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	echolog "github.com/labstack/gommon/log"
)

type Config struct {
	Port        string `env:"PORT,required" envDefault:"8080"`
	Address     string `env:"SERVER_ADDRESS,required" envDefault:"0.0.0.0"`
	BaseURL     string `env:"BASE_URL,required" envDefault:"http://localhost:8080"`
	DatabaseDSN string `env:"DATABASE_URL,required" envDefault:""`
	LogLevel    string `env:"LOG_LEVEL"`
}

func (c *Config) Read(logger echo.Logger) error {
	err := env.Parse(c)
	if err != nil {
		return err
	}

	var (
		address     string
		baseURL     string
		databaseDSN string
		port        string
		logLevel    string
	)

	flag.StringVar(&address, "a", "", "Адрес запуска HTTP-сервера")
	flag.StringVar(&port, "p", "", "Порт запуска HTTP-сервера")
	flag.StringVar(&baseURL, "b", "", "Базовый адрес результирующего сокращённого URL")
	flag.StringVar(&databaseDSN, "d", "", "Строка с адресом подключения к БД")
	flag.StringVar(&logLevel, "l", "", "Уровень логирования")
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

	if logLevel != "" {
		c.LogLevel = logLevel
	}

	readCfg, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	logger.Infof("config: %s", readCfg)

	return nil
}

var LogLevelMap = map[string]echolog.Lvl{
	"DEBUG": echolog.DEBUG,
	"INFO":  echolog.INFO,
	"ERROR": echolog.ERROR,
	"WARN":  echolog.WARN,
	"OFF":   echolog.OFF,
}
