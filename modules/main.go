package modules

import (
	"context"
	"fmt"
	"log"
	"math/big"

	Models "github.com/LuisAcerv/goeth-api/models"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetLatestBlock from blockchain
func GetLatestBlock(client ethclient.Client) *Models.Block {
	// We add a recover function from panics to prevent our API from crashing due to an unexpected error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// Query the latest block
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	// Build the response to our model
	_block := &Models.Block{
		BlockNumber:       block.Number().Int64(),
		Timestamp:         block.Time(),
		Difficulty:        block.Difficulty().Uint64(),
		Hash:              block.Hash().String(),
		TransactionsCount: len(block.Transactions()),
		Transactions:      []Models.Transaction{},
	}

	for _, tx := range block.Transactions() {
		_block.Transactions = append(_block.Transactions, Models.Transaction{
			Hash:     	 tx.Hash().String(),
			Type: 	  	 "mint",
			Address:  	 tx.To().String(),
			BlockHash:   tx.Hash().Hex(),
			BlockNumber: block.Number().Int64(),
			To:       	 tx.To().String(),
			From: 		 tx.Hash().Hex(),
			Gas:         tx.Gas(),
			GasPrice: 	 tx.GasPrice().Uint64(),
			Nonce:    	 tx.Nonce(),
			Date: "20220829,00:23:41",
			CollectionName: "Otherdeed for Otherside",
			CollectionAddress: tx.To().Hex(),
		})
	}

	return _block
}