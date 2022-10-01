package main

import (
	"log"
	"test/internal/app"
	"test/internal/mux"
)

func main() {
	_ = app.Instance("Successful")
	//rabbitmq.Init()
	log.Println("MUX STARTING")
	mux.Init()
}
