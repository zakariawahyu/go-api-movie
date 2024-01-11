package main

import (
	"github.com/zakariawahyu/go-api-movie/config"
	"github.com/zakariawahyu/go-api-movie/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewHttpServer(cfg)
	log.Fatal(s.Run())
}
