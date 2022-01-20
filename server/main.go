package main

import (
	"flag"
	"log"
	"net/http"
	"slipcover/packages/logger"
)

func main() {
	logger.Log("Test")

	port := flag.String("port", "3001", "Running Port for WebApp")

	fs := http.FileServer(http.Dir("./appbuild"))
	http.Handle("/", fs)

	log.Println("Listening on : " + *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
