package main

import (
	"broadcast_service/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handlers.WsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
