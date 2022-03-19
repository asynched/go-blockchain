package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asynched/blockchain/domain/chain"
	router "github.com/asynched/blockchain/utils/http"
)

var blockchain = chain.NewBlockChain()

type CreateBlock struct {
	Data string `json:"data"`
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Query for blockchain")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(blockchain)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Creating a new block")

	decoder := json.NewDecoder(r.Body)
	var cb CreateBlock

	err := decoder.Decode(&cb)

	if err != nil {
		log.Println("[ERROR] Failed to decode request body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	blockchain.AddBlock(cb.Data)
	log.Println("[INFO] Successfully created block")

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(blockchain)
}

func MakeServer() http.Server {
	router := router.Route{
		Get:  handleGet,
		Post: handlePost,
	}

	http.HandleFunc("/api/blockchain", router.MakeHandler())

	server := http.Server{
		Addr: ":8081",
	}

	return server
}
