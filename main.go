package main

import (
	"log"

	"github.com/asynched/blockchain/web"
)

func main() {
	server := web.MakeServer()

	log.SetFlags(log.Ltime | log.Ldate)
	log.Println("[INFO] Starting server on port :8081")
	log.Fatal(server.ListenAndServe())
}
