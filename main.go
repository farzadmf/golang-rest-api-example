package main

import (
	"log"

	"assignment/pkg/config"
	"assignment/pkg/repository"
	"assignment/pkg/server"
)

func main() {
	cfg, err := config.New()
	check(err)

	repo, err := repository.New(cfg)
	check(err)

	srv := server.New(repo)
	log.Fatal(srv.Run())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
