package main

import (
	"log"

	"github.com/fogo-sh/grackdb/api"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := api.StartApi()
	if err != nil {
		log.Fatalf("Error starting API: %s", err)
	}
}
