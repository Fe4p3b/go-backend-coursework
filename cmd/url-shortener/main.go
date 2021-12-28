package main

import (
	"log"

	"github.com/Fe4p3b/go-backend-coursework/internal/app/shortener"
	"github.com/Fe4p3b/go-backend-coursework/internal/handlers/http"
	"github.com/Fe4p3b/go-backend-coursework/internal/server"
	"github.com/Fe4p3b/go-backend-coursework/internal/storage/memory"
)

func main() {
	m := memory.New(map[string]string{})
	s := shortener.New(m)
	_ = http.NewHttpHandler(s)
	server := server.New(":8080", nil)

	log.Fatal(server.Start())
}
