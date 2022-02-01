package main

import (
	"log"

	"github.com/Fe4p3b/go-backend-coursework/cmd/api/url-shortener/config"
	"github.com/Fe4p3b/go-backend-coursework/internal/app/shortener"
	"github.com/Fe4p3b/go-backend-coursework/internal/handlers/http"
	"github.com/Fe4p3b/go-backend-coursework/internal/storage/memory"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := &config.Config{}
	if err := cfg.Read(); err != nil {
		log.Fatal(err)
	}

	m := memory.New(map[string]string{})
	s := shortener.New(m, cfg.BaseURL)
	h := http.New(s)
	server := echo.New()
	server.GET("/:url", h.Get)
	server.POST("/", h.Post)

	if err := server.Start(cfg.Address); err != nil {
		log.Fatal(err)
	}
}
