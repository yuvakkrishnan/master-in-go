package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	//Create a Server

	myServer := &http.Server{
		//Set the Server Address
		Addr: "127.0.0.1:8080",
		//define some specific configuration
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//Launch the server

	log.Fatal(myServer.ListenAndServe())
}
