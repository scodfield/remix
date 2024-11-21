package goeth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ReqContractByHash() {
	conn, err := ethclient.Dial("https://sepolia.infura.io/v3/your_api_key")
	if err != nil {
		log.Fatalf("Dail infura went wrong: %v", err)
	}
	tx, pending, err := conn.TransactionByHash(context.Background(), common.HexToHash("0x30bccbd032dd9436e2a33ec49c5d135ab02d898d05c22c62c50774f4eb2db24d"))
	if err != nil {
		log.Fatalf("TransactionByHash went wrong: %v", err)
	}
	if !pending {
		fmt.Println(tx)
	}
}
