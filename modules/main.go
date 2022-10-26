package modules

import (
	"context"
	"fmt"
	"log"
	"bytes"
	"io/ioutil"
	"net/http"
	"math/big"
	"encoding/json"

	Models "github.com/saimunhossain/nft-report-analyze/models"
	"github.com/ethereum/go-ethereum/common"
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

// GetTxByHash by a given hash
func GetTxByHash(client ethclient.Client, hash common.Hash) *Models.Transaction {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		fmt.Println(err)
	}

	return &Models.Transaction{
		Hash:     	 tx.Hash().String(),
		Type: 	  	 "mint",
		Address:  	 tx.To().String(),
		BlockHash:   tx.Hash().Hex(),
		To:       	 tx.To().String(),
		From: 		 tx.Hash().Hex(),
		Gas:         tx.Gas(),
		GasPrice: 	 tx.GasPrice().Uint64(),
		Nonce:    	 tx.Nonce(),
		Date: 		 "20220829,00:23:41",
		CollectionName: "Otherdeed for Otherside",
		CollectionAddress: tx.To().Hex(),
		Pending: pending,
	}
}

// StoreTxByHash by a given hash
func StoreTxByHash(client ethclient.Client, hash common.Hash) *Models.Transaction {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		fmt.Println(err)
	}

	nftActions := &Models.Transaction{
		Hash:     	 tx.Hash().String(),
		Type: 	  	 "mint",
		Address:  	 tx.To().String(),
		BlockHash:   tx.Hash().Hex(),
		To:       	 tx.To().String(),
		From: 		 tx.Hash().Hex(),
		Gas:         tx.Gas(),
		GasPrice: 	 tx.GasPrice().Uint64(),
		Nonce:    	 tx.Nonce(),
		Date: 		 "20220829,00:23:41",
		CollectionName: "Otherdeed for Otherside",
		CollectionAddress: tx.To().Hex(),
		Pending: pending,
	}
	responseJson, _ := json.Marshal(nftActions)

	resp, err := http.Get("http://dataworks.gw106.oneitfarm.com/v1/project/blockchain_analytics/new_upload_data_url?table_name=ods_nft")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	type Data struct {
		Url string `json:"url"`
		Token string `json:"token"`
		RawUrl string `json:"raw_url"`
	}
	type Result struct {
		State string `json:"state"`
		Msg string `json:"msg"`
		ResData Data `json:"data"`
	}
	var result Result
	json.Unmarshal(body, &result)
	// fmt.Println(result.ResData.Token)


	url := "http://dataworks.gw106.oneitfarm.com/"+result.ResData.RawUrl
	var bearer = "Bearer " + result.ResData.Token
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(responseJson))

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	respHttp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer respHttp.Body.Close()

	postBody, err := ioutil.ReadAll(respHttp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}		
	type PostResult struct {
		State string `json:"state"`
		Msg string `json:"msg"`
		PostResData string `json:"data"`
	}
	var postResult PostResult
	json.Unmarshal(postBody, &postResult)
	if postResult.Msg == "ok"{
		fmt.Println("Data has been sent to server")
	}else{
		fmt.Println("There is something wrong, Please try later!")
	}
	return nil
}