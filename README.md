## NFT ANALYZE REPOT

To get started, make sure you have to clone this repository.

1. Clone this project:

   ```sh
   git clone https://github.com/saimunhossain/nft-analyze-api.git
   ```
2. Inside the folder `nft-analyze-api` and run the next command:

   ```sh
   go run main.go
   ```

### Head over to our browser

1. **To see latest block with transaction** `http://localhost:9000/api/v1/eth/get-latest-block`

2. **To see specific transaction** `http://localhost:9000/api/v1/eth/get-tx?hash={Press Your Transaction Hash}`

3. **To store specific transaction on server** `http://localhost:9000/api/v1/eth/getstore-tx?hash={Press Your Transaction Hash}`

4. **To see address balance** `http://localhost:9000/api/v1/eth/get-balance?address={Press Your Address Hash}`

4. **To Transfer Ethereum one to another** `http://localhost:9000/api/v1/eth/get-balance?privKey={Press Your private key}to={Press To Address}amount={Press Amount}`