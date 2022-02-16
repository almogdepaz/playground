package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"os"
	"price_monitor/badger"
	"price_monitor/bancor"
	uniswap "price_monitor/uniswap"
	"price_monitor/util"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const node = "https://mainnet.infura.io/v3/093f1d19defd46248d24aa7e734ea203"

const decimals = 18
const wbtc_decimals = 8
const diggDecimals = 9

const interval = 20 * time.Minute

var Client *ethclient.Client

func init() { Client = util.GetClient(node) }

func main() {
	LOG_FILE := "debug.log"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	mw := io.MultiWriter(os.Stdout, logFile)

	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(mw)
	for {
		ethAmount := new(big.Int)
		ethAmount.Mul(big.NewInt(1), big.NewInt(int64(math.Pow(10, decimals))))
		bancorEth := common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")         //null addr for bancor eth
		bDigg := common.HexToAddress("0x7e7e112a68d8d2e221e11047a72ffc1065c38e1a")             //bDigg
		uniWbtcDiggPool := common.HexToAddress("0xE86204c4eDDd2f70eE00EAd6805f917671F56c52")   //Uniswap WBTC/DIGG LP (UNI-V2)
		sushiWbtcDiggPool := common.HexToAddress("0x9a13867048e01c663ce8ce2fe0cdae69ff9f35e3") //Sushiswap WBTC/DIGG LP (UNI-V2)
		sushiEthWbtcPool := common.HexToAddress("0xceff51756c56ceffca006cd410b03ffc46dd3a58")
		badger_contract := common.HexToAddress("0x0F92Ca0fB07E420b2fED036A6bB023c6c9e49940") //badger contract
		bancor_contract := common.HexToAddress("0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0") //bancor contract
		digg := common.HexToAddress("0x798d1be841a82a273720ce31c822c61a67a601c3")            //digg
		wbtc := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")            //wbtc
		weth := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

		// eth to bdigg
		bancorRes := FetchPoolStatsBancor(bancor_contract, bancorEth, bDigg, ethAmount)
		log.Println(fmt.Sprintf("Bancor Eth amount in %v bDIGG ammount out %v", parseDecimalsFromInt(ethAmount, decimals), parseDecimalsFromInt(bancorRes, decimals)))
		// bdigg to digg
		badger_caller, err := badger.NewBadgerCaller(badger_contract, Client)
		if err != nil {
			log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
		}
		price, err := badger_caller.GetPricePerFullShare(&bind.CallOpts{Context: context.TODO()})
		if err != nil {
			log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
		}
		parsedSharePrice := big.NewInt(price.Int64() / int64(math.Pow(10, decimals)))
		log.Println(fmt.Sprintf("bDIGG share price %v", parsedSharePrice))
		diggOut := new(big.Int).Mul(bancorRes, parsedSharePrice)
		log.Println(fmt.Sprintf("\nBadger %v bDIGG in %v Digg out ", parseDecimalsFromInt(bancorRes, decimals), parseDecimalsFromInt(diggOut, decimals)))

		// digg to wbtc
		uniWbtcRes := FetchPoolStatsUniswap(uniWbtcDiggPool, big.NewInt(1), digg, wbtc)
		sushiWbtcRes := FetchPoolStatsUniswap(sushiWbtcDiggPool, diggOut, digg, wbtc)
		log.Println(fmt.Sprintf("Uniswap WBTC/DIGG digg amount in %v  wbtc ammount out %v", parseDecimalsFromInt(diggOut, decimals), parseDecimalsFromFloat(uniWbtcRes, decimals)))
		log.Println(fmt.Sprintf("Sushiswap WBTC/DIGG digg amount in %v wbtc ammount out %v", parseDecimalsFromInt(diggOut, decimals), parseDecimalsFromFloat(sushiWbtcRes, decimals)))
		wbtcRes := new(big.Int)

		if uniWbtcRes.Cmp(sushiWbtcRes) < 0 {
			sushiWbtcRes.Int(wbtcRes)
			log.Println("choose sushi")
		} else {
			uniWbtcRes.Int(wbtcRes)
			log.Println("choose uni")
		}

		// wbtc to eth
		sushiEthRes := FetchPoolStatsUniswap(sushiEthWbtcPool, wbtcRes, wbtc, weth)
		log.Println(fmt.Sprintf("%v wbtc in get %v eth out", parseDecimalsFromInt(wbtcRes, decimals), parseDecimalsFromFloat(sushiEthRes, decimals)))
		time.Sleep(interval)
		log.Println("------------------------------------------------------------------------------------------")
	}

}

func parseDecimalsFromInt(num *big.Int, decimals float64) float64 {
	parsed := new(big.Float).SetInt(num)
	return parseDecimalsFromFloat(parsed, decimals)
}

func parseDecimalsFromFloat(num *big.Float, decimals float64) float64 {
	parsed, _ := num.Float64()
	parsed = parsed / math.Pow(10, decimals)
	return parsed
}

func FetchPoolStatsBancor(bancor_contract common.Address, from common.Address, to common.Address, amount *big.Int) *big.Int {
	bancor_caller, err := bancor.NewBancorCaller(bancor_contract, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to create bancor caller: %v\n", err))
	}
	path, err := bancor_caller.ConversionPath(&bind.CallOpts{Context: context.TODO()}, from, to)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get conversion path: %v\n", err))
	}

	bancorRes, err := bancor_caller.RateByPath(&bind.CallOpts{Context: context.TODO()}, path, amount)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get rate for path: %v\n", err))
	}
	return bancorRes
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
