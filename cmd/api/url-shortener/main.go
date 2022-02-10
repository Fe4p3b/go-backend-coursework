package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Fe4p3b/go-backend-coursework/cmd/api/url-shortener/config"
	"github.com/Fe4p3b/go-backend-coursework/internal/app/shortener"
	"github.com/Fe4p3b/go-backend-coursework/internal/handlers"
	"github.com/Fe4p3b/go-backend-coursework/internal/storage/pg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echolog "github.com/labstack/gommon/log"
)

func main() {
	server := echo.New()
	server.Logger.SetLevel(echolog.INFO)

	cfg := &config.Config{}
	if err := cfg.Read(server.Logger); err != nil {
		log.Fatal(err)
	}

	if logLevel, ok := config.LogLevelMap[cfg.LogLevel]; ok {
		server.Logger.SetLevel(logLevel)
	}

	// m := memory.New(map[string]string{})

	db, err := pg.NewConnection(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	s := shortener.New(db, cfg.BaseURL)
	h := handlers.New(s)

	server.GET("/:url", h.Get)
	server.GET("/:url/stats", h.GetStats)
	server.POST("/", h.Post)

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Use(middleware.Gzip())
	server.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:8081",
				"https://gb-backend1-coursework-front.herokuapp.com",
			},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAcceptEncoding},
		},
	))

	go func() {
		if err := server.Start(fmt.Sprintf("%s:%s", cfg.Address, cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	server.Logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
