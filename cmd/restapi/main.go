// Winery is a simple REST API that serves
// information about different wines.
//
// Copyright (c) 2022 Nikita Mezhenskyi. All rights reserved.
// MIT Licensed.
package main

import (
	"log"

	"github.com/nmezhenskyi/go-rest-api-example/internal/webserver"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server := webserver.NewServer()
	err := server.PopulateWithData("data/data_sample.json")
	if err != nil {
		return err
	}
	return server.ListenAndServe("localhost:5000")
}
