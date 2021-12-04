package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://testnet-dex.binance.org/api/ws")
	fmt.Printf("client%+v\n", client)
	if err != nil {
		log.Fatal(err)
	}

	var contractAddress = common.HexToAddress("0x98b3f2219a2b7a047B6234c19926673ad4aac83A")
	var topics = common.HexToHash("0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topics}},
	}
	fmt.Printf("query%+v\n", query)

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("logs%+v\n", logs)
	fmt.Printf("sub%+v\n", sub)
	fmt.Printf("err%+v\n", sub)

	for {
		fmt.Println("_........................_")
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
