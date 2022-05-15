// Winery is a simple REST API that serves
// information about different wines.
//
// Copyright (c) 2022 Nikita Mezhenskyi. All rights reserved.
// MIT Licensed.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nmezhenskyi/go-rest-api-example/internal/webserver"
)

func main() {
	server := webserver.NewServer()
	err := server.PopulateWithData("data/data_sample.json")
	if err != nil {
		log.Fatalf("Failed to load data sample: %v\n", err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe("localhost:5000")
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v\n", err)
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}
	log.Println("Server has been shutdown")
}
