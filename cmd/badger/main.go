package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"price_monitor/badger"
	uniswap "price_monitor/uniswap"
	"price_monitor/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const node = "https://mainnet.infura.io/v3/093f1d19defd46248d24aa7e734ea203"

const decimals = 18

var Client *ethclient.Client

func init() { Client = util.GetClient(node) }

func main() {
	addr1 := common.HexToAddress("0xE86204c4eDDd2f70eE00EAd6805f917671F56c52") //Uniswap WBTC/DIGG LP (UNI-V2)
	pair, err := uniswap.NewUniswapv2pairCaller(addr1, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}

	for i := 0; i < 10; i++ {
		res := FetchPoolStatsUniswap(pair, i)
		fmt.Printf("\nWBTC/DIGG amount in %v ammount out %v", i, res)
	}
	addr2 := common.HexToAddress("0x0F92Ca0fB07E420b2fED036A6bB023c6c9e49940") //badger contract

	badger_caller, err := badger.NewBadgerCaller(addr2, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	price, err := badger_caller.GetPricePerFullShare(&bind.CallOpts{Context: context.TODO()})
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	fmt.Printf("Badger digg price %v", price)
}

// amount - the amount of token0 to send
// returns the recived amount of token1 given the input
func FetchPoolStatsUniswap(pair *uniswap.Uniswapv2pairCaller, amount_in int) *big.Float {
	token0, _ := pair.Token0(&bind.CallOpts{Context: context.TODO()})
	token1, _ := pair.Token1(&bind.CallOpts{Context: context.TODO()})
	amount1, err := uniswap.GetExchangeAmount(pair, big.NewFloat(float64(amount_in)), token0, token1)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get exchange amount: %v\n", err))
	}
	return amount1
}
