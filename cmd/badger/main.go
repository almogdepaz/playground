package main

import (
	"context"
	"fmt"
	"log"
	"math"
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
	eth_amount := new(big.Int)
	uni_amount := new(big.Int)
	sushi_amount := new(big.Int)
	eth_amount.Mul(big.NewInt(1), big.NewInt(int64(math.Pow(10, decimals))))
	uni_amount.Mul(big.NewInt(1), big.NewInt(int64(math.Pow(10, decimals))))
	sushi_amount.Mul(big.NewInt(1), big.NewInt(int64(math.Pow(10, decimals))))

	bancor_eth := common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")      //bancor eth
	sdigg := common.HexToAddress("0x7e7e112a68d8d2e221e11047a72ffc1065c38e1a")           //sdigg
	uni_pool := common.HexToAddress("0xE86204c4eDDd2f70eE00EAd6805f917671F56c52")        //Uniswap WBTC/DIGG LP (UNI-V2)
	sushi_pool := common.HexToAddress("0x9a13867048e01c663ce8ce2fe0cdae69ff9f35e3")      //Sushiswap WBTC/DIGG LP (UNI-V2)
	badger_contract := common.HexToAddress("0x0F92Ca0fB07E420b2fED036A6bB023c6c9e49940") //badger contract
	bancor_contract := common.HexToAddress("0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0") //bancor contract
	digg := common.HexToAddress("0x798d1be841a82a273720ce31c822c61a67a601c3")            //digg
	wbtc := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")            //wbtc

	// uni
	uni_res := FetchPoolStatsUniswap(uni_pool, uni_amount, wbtc, digg)
	fmt.Printf("\nUniswap WBTC/DIGG wbtc amount in %v digg ammount out %v", uni_amount, uni_res)

	// sushi
	sushi_res := FetchPoolStatsUniswap(sushi_pool, sushi_amount, wbtc, digg)
	fmt.Printf("\nSushiswap WBTC/DIGG wbtc amount in %v digg ammount out %v", sushi_amount, sushi_res)

	// bancor
	bancor_res := FetchPoolStatsBancor(bancor_contract, bancor_eth, sdigg, eth_amount)
	fmt.Printf("\nBancor Eth amount in %v SDIGG ammount out %v", eth_amount, bancor_res)

	// bagder
	badger_caller, err := badger.NewBadgerCaller(badger_contract, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	price, err := badger_caller.GetPricePerFullShare(&bind.CallOpts{Context: context.TODO()})
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	fmt.Printf("\nBadger digg price %v", price)
}

func FetchPoolStatsBancor(bancor_contract common.Address, wbtc_contract common.Address, digg_contract common.Address, amount *big.Int) *big.Int {
	bancor_caller, err := bancor.NewBancorCaller(bancor_contract, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to create bancor caller: %v\n", err))
	}
	path, err := bancor_caller.ConversionPath(&bind.CallOpts{Context: context.TODO()}, wbtc_contract, digg_contract)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get conversion path: %v\n", err))
	}

	bancor_res, err := bancor_caller.RateByPath(&bind.CallOpts{Context: context.TODO()}, path, amount)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get rate for path: %v\n", err))
	}
	return bancor_res
}

// amount - the amount of token0 to send
// returns the recived amount of token1 given the input
func FetchPoolStatsUniswap(pool common.Address, amount_in *big.Int, in common.Address, out common.Address) *big.Float {
	pair, err := uniswap.NewUniswapv2pairCaller(pool, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	amount1, err := uniswap.GetExchangeAmount(pair, new(big.Float).SetInt(amount_in), in, out)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get exchange amount: %v\n", err))
	}
	return amount1
}
