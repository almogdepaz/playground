package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"price_monitor/badger"
	"price_monitor/bancor"
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
	amount := 10
	digg_contract := common.HexToAddress("0x798d1be841a82a273720ce31c822c61a67a601c3") //digg
	wbtc_contract := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599") //wbtc

	// uni
	uni_pool := common.HexToAddress("0xE86204c4eDDd2f70eE00EAd6805f917671F56c52") //Uniswap WBTC/DIGG LP (UNI-V2)
	uni_res := FetchPoolStatsUniswap(uni_pool, amount)
	fmt.Printf("\nUniswap WBTC/DIGG amount in %v ammount out %v", amount, uni_res)

	// sushi
	sushi_pool := common.HexToAddress("0x9a13867048e01c663ce8ce2fe0cdae69ff9f35e3") //Sushiswap WBTC/DIGG LP (UNI-V2)
	sushi_res := FetchPoolStatsUniswap(sushi_pool, amount)
	fmt.Printf("\nWBTC/DIGG amount in %v ammount out %v", amount, sushi_res)

	// bancor
	bancor_contract := common.HexToAddress("0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0") //bancor contract
	bancor_res := FetchPoolStatsBancor(bancor_contract, wbtc_contract, digg_contract, amount)
	fmt.Printf("\nWBTC/DIGG amount in %v ammount out %v", amount, bancor_res)
	addr2 := common.HexToAddress("0x0F92Ca0fB07E420b2fED036A6bB023c6c9e49940") //badger contract
	badger_caller, err := badger.NewBadgerCaller(addr2, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	price, err := badger_caller.GetPricePerFullShare(&bind.CallOpts{Context: context.TODO()})
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	fmt.Printf("\nBadger digg price %v", price)
}

func FetchPoolStatsBancor(bancor_contract common.Address, wbtc_contract common.Address, digg_contract common.Address, amount int) *big.Int {
	bancor_caller, err := bancor.NewBancorCaller(bancor_contract, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to create bancor caller: %v\n", err))
	}
	path, err := bancor_caller.ConversionPath(&bind.CallOpts{Context: context.TODO()}, wbtc_contract, digg_contract)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get conversion path: %v\n", err))
	}
	bancor_res, err := bancor_caller.RateByPath(&bind.CallOpts{Context: context.TODO()}, path, big.NewInt(int64(amount)))
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get rate for path: %v\n", err))
	}
	return bancor_res
}

// amount - the amount of token0 to send
// returns the recived amount of token1 given the input
func FetchPoolStatsUniswap(pool common.Address, amount_in int) *big.Float {
	pair, err := uniswap.NewUniswapv2pairCaller(pool, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	token0, _ := pair.Token0(&bind.CallOpts{Context: context.TODO()})
	token1, _ := pair.Token1(&bind.CallOpts{Context: context.TODO()})
	amount1, err := uniswap.GetExchangeAmount(pair, big.NewFloat(float64(amount_in)), token0, token1)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get exchange amount: %v\n", err))
	}
	return amount1
}
