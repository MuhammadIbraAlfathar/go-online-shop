package main

import (
	"github.com/MuhammadIbraAlfathar/go-online-shop/external/database"
	"github.com/MuhammadIbraAlfathar/go-online-shop/internal/config"
	"log"
)

func main() {

	filename := "cmd/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Printf("connected database")
	}

}
