package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/S-Amine/go-uni/contracts/uniswapv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	defaultSleepDuration = 5 * time.Second
	maxSleepDuration     = 15 * time.Second
	lastBlocksToCheck    = 1
)

func main() {
	// Handle graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Connect to an Ethereum client
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	// Connect to a websocket client
	wsClient, err := ethclient.Dial("wss://mainnet.gateway.tenderly.co")
	if err != nil {
		log.Fatalf("Failed to connect to the WebSocket client: %v", err)
	}
	defer wsClient.Close()

	// Instantiate the Uniswapv2 contract
	contractAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	instance, err := uniswapv2.NewUniswapv2(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Uniswapv2 contract: %v", err)
	}

	// Create a channel to receive new block headers
	headers := make(chan *types.Header)
	subscription, err := wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Failed to subscribe to new block headers: %v", err)
	}
	defer subscription.Unsubscribe()

	// Start event processing loop
	for {
		select {
		case <-interrupt:
			log.Println("Received termination signal. Shutting down...")
			return
		case header := <-headers:
			go processNewBlock(instance, header)
		}
	}
}

func processNewBlock(instance *uniswapv2.Uniswapv2, header *types.Header) {
	// FilterOpts for the PairCreated events in the last block
	filterOpts := &bind.FilterOpts{
		Context: context.Background(),
		Start:   header.Number.Uint64() - lastBlocksToCheck,
		End:     nil,
	}

	// Fetch PairCreated events
	pairCreatedEvents, err := instance.FilterPairCreated(filterOpts, nil, nil)
	if err != nil {
		log.Printf("Error fetching PairCreated events: %++v\n", err)
		return
	}
	defer pairCreatedEvents.Close()

	// Process PairCreated events concurrently using goroutines
	for pairCreatedEvents.Next() {
		event := pairCreatedEvents.Event
		go func(event *uniswapv2.Uniswapv2PairCreated) {
			fmt.Printf(`{"pair": "%v", "token0: "%v", "token1": "%v", "arg3": "%v"}`+"\n",
				event.Token0.Hex(), event.Token1.Hex(), event.Pair.Hex(), event.Arg3.String())
		}(event)
	}
}
