package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	Models "github.com/saimunhossain/nft-report-analyze/models"
	Modules "github.com/saimunhossain/nft-report-analyze/modules"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// ClientHandler ethereum client instance
type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get parameter from url request
	vars := mux.Vars(r)
	module := vars["module"]

	hash := r.URL.Query().Get("hash")
	// Set our response header
	w.Header().Set("Content-Type", "application/json")

	// Handle each request using the module parameter:
	switch module {
	case "get-latest-block":
		_block := Modules.GetLatestBlock(*client.Client)
		json.NewEncoder(w).Encode(_block)

	case "get-tx":
		if hash == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Empty request",
			})
			return
		}
		txHash := common.HexToHash(hash)
		_tx := Modules.GetTxByHash(*client.Client, txHash)

		if _tx != nil {
			json.NewEncoder(w).Encode(_tx)
			return
		}

		json.NewEncoder(w).Encode(&Models.Error{
			Code:    404,
			Message: "Tx Not Found!",
		})

	case "store-tx":
		if hash == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Empty request",
			})
			return
		}
		txHash := common.HexToHash(hash)
		_tx := Modules.StoreTxByHash(*client.Client, txHash)

		if _tx != nil {
			json.NewEncoder(w).Encode(_tx)
			return
		}

		json.NewEncoder(w).Encode(&Models.Error{
			Code:    200,
			Message: "Tx stored successfully!",
		})


	case "send-eth":
		decoder := json.NewDecoder(r.Body)
		var t Models.TransferEthRequest

		err := decoder.Decode(&t)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_hash, err := Modules.TransferEth(*client.Client, t.PrivKey, t.To, t.Amount)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&Models.HashResponse{
			Hash: _hash,
		})
	}

}
