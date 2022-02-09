package main

import (
	"log"

	"github.com/Fe4p3b/go-backend-coursework/cmd/api/url-shortener/config"
	"github.com/Fe4p3b/go-backend-coursework/internal/app/shortener"
	"github.com/Fe4p3b/go-backend-coursework/internal/handlers/http"
	"github.com/Fe4p3b/go-backend-coursework/internal/storage/pg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := &config.Config{}
	if err := cfg.Read(); err != nil {
		log.Fatal(err)
	}

	// m := memory.New(map[string]string{})

	db, err := pg.NewConnection(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	s := shortener.New(db, cfg.BaseURL)
	h := http.New(s)
	server := echo.New()
	server.GET("/:url", h.Get)
	server.POST("/", h.Post)

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
	if err := server.Start(cfg.Address + ":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
