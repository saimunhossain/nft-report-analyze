## NFT REPORT ANALYZE

To get started, make sure you have to clone this repository.

1. Clone this project:

   ```sh
   git clone https://github.com/saimunhossain/nft-report-analyze.git
   ```
2. Inside the folder `nft-report-analyze` and run the next command:

   ```sh
   go run main.go
   ```

### Head over to our browser

1. **To see latest block with transaction** `http://localhost:7000/api/v1/get-latest-block`

2. **To see specific transaction** `http://localhost:7000/api/v1/get-tx?hash={Press Your Transaction Hash}`

3. **To store specific transaction on server** `http://localhost:7000/api/v1/store-tx?hash={Press Your Transaction Hash}`

4. **To see address balance** `http://localhost:7000/api/v1/get-balance?address={Press Your Address Hash}`

To generate ABI and bin:

   ```sh
   solc --bin --abi erc721.sol -o build
   ```
   ```sh
   abigen --abi=./build/ERC721.abi --pkg=erc721 --out=erc721.go
   ```

5. **To Transfer Ethereum one to another** `http://localhost:7000/api/v1/send-eth?privKey={Press Your private key}to={Press To Address}amount={Press Amount}`