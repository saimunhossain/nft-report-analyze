package main

import (
	"fmt"
	"log"
	"net/http"

	Handlers "github.com/LuisAcerv/goeth-api/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2468f3c54a7b498284b55d91676c0913")

	if err != nil {
		fmt.Println(err)
	}

	r := mux.NewRouter()

	r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{Client: client})
	log.Fatal(http.ListenAndServe(":7000", r))
}
