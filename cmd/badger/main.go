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

const std_decimals = 18
const wbtc_decimals = 8
const diggDecimals = 9

const interval = 10 * time.Second

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
	ethAmount := new(big.Int).Mul(big.NewInt(1), big.NewInt(int64(math.Pow(10, std_decimals))))
	badger_contract := common.HexToAddress("0x0F92Ca0fB07E420b2fED036A6bB023c6c9e49940") //badger contract
	bancor_contract := common.HexToAddress("0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0") //bancor contract

	bancorEth := common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") //null addr for bancor eth
	bDigg := common.HexToAddress("0x7e7e112a68d8d2e221e11047a72ffc1065c38e1a")     //bDigg
	digg := common.HexToAddress("0x798d1be841a82a273720ce31c822c61a67a601c3")      //digg
	wbtc := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")      //wbtc
	weth := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	uniWbtcDiggPool, _ := uniswap.CalculatePoolAddressUniV2(wbtc, digg)   //Uniswap WBTC/DIGG LP (UNI-V2)
	sushiWbtcDiggPool, _ := uniswap.CalculatePoolAddressSushi(wbtc, digg) //Sushiswap WBTC/DIGG LP (UNI-V2)
	sushiEthWbtcPool, _ := uniswap.CalculatePoolAddressSushi(weth, wbtc)
	uniEthWbtcPool, _ := uniswap.CalculatePoolAddressUniV2(weth, wbtc)

	blockNumber := uint64(0)
	log.SetOutput(mw)
	for {
		prev := blockNumber
		blockNumber, err = Client.BlockNumber(context.TODO())
		if prev == blockNumber {
			log.Println(fmt.Sprintf("last block %v, waiting for next block", blockNumber))
			time.Sleep(interval)
			continue
		}

		if err != nil {
			log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
		}
		// eth to bdigg
		bancorRes := FetchPoolStatsBancor(bancor_contract, bancorEth, bDigg, ethAmount, blockNumber)
		log.Println(fmt.Sprintf("Bancor %v Eth in %v bDIGG  out", parseDecimalsFromInt(ethAmount, std_decimals), parseDecimalsFromInt(bancorRes, std_decimals)))
		// bdigg to digg
		badger_caller, err := badger.NewBadgerCaller(badger_contract, Client)
		if err != nil {
			log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
		}
		price, err := badger_caller.GetPricePerFullShare(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(blockNumber), Context: context.TODO()})
		if err != nil {
			log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
		}

		// parsedSharePrice := big.NewInt(price.Int64() / int64(math.Pow(10, std_decimals)))
		parsedSharePrice := big.NewInt(price.Int64() / int64(math.Pow(10, std_decimals)))
		log.Println(fmt.Sprintf("bDIGG share price %v", parsedSharePrice))
		diggOut := new(big.Int).Mul(bancorRes, parsedSharePrice)
		log.Println(fmt.Sprintf("Badger %v bDIGG in %v Digg out ", parseDecimalsFromInt(bancorRes, std_decimals), parseDecimalsFromInt(diggOut, std_decimals)))

		// digg to wbtc
		uniWbtcRes := FetchPoolStatsUniswap(uniWbtcDiggPool, diggOut, digg, wbtc, diggDecimals, wbtc_decimals, blockNumber)
		sushiWbtcRes := FetchPoolStatsUniswap(sushiWbtcDiggPool, diggOut, digg, wbtc, diggDecimals, wbtc_decimals, blockNumber)
		log.Println(fmt.Sprintf("Uniswap WBTC/DIGG %v digg in %v  wbtc out", parseDecimalsFromInt(diggOut, std_decimals), parseDecimalsFromFloat(uniWbtcRes, std_decimals)))
		log.Println(fmt.Sprintf("Sushiswap WBTC/DIGG %v digg in %v  wbtc out", parseDecimalsFromInt(diggOut, std_decimals), parseDecimalsFromFloat(sushiWbtcRes, std_decimals)))
		wbtcRes := new(big.Int)
		if uniWbtcRes.Cmp(sushiWbtcRes) < 0 {
			sushiWbtcRes.Int(wbtcRes)
			log.Println("choose sushi")
		} else {
			uniWbtcRes.Int(wbtcRes)
			log.Println("choose uni")
		}

		// wbtc to eth
		uniEthRes := FetchPoolStatsUniswap(uniEthWbtcPool, wbtcRes, wbtc, weth, wbtc_decimals, std_decimals, blockNumber)
		sushiEthRes := FetchPoolStatsUniswap(sushiEthWbtcPool, wbtcRes, wbtc, weth, wbtc_decimals, std_decimals, blockNumber)
		log.Println(fmt.Sprintf("Uniswap ETH/WBTC %v wbtc in %v eth out", parseDecimalsFromInt(wbtcRes, std_decimals), parseDecimalsFromFloat(uniEthRes, std_decimals)))
		log.Println(fmt.Sprintf("Sushiswap ETH/WBTC  %v wbtc in %v eth out %v", parseDecimalsFromInt(wbtcRes, std_decimals), parseDecimalsFromFloat(sushiEthRes, std_decimals)))
		ethRes := new(big.Int)
		if uniEthRes.Cmp(sushiEthRes) < 0 {
			sushiWbtcRes.Int(ethRes)
			log.Println("choose sushi")
		} else {
			uniEthRes.Int(ethRes)
			log.Println("choose uni")
		}
		log.Println(fmt.Sprintf("%v wbtc in get %v eth out", parseDecimalsFromInt(wbtcRes, std_decimals), parseDecimalsFromFloat(sushiEthRes, std_decimals)))
		time.Sleep(interval)
		log.Println("------------------------------------------------------------------------------------------")
	}

}

func parseDecimalsFromInt(num *big.Int, decimals float64) float64 {
	return parseDecimalsFromFloat(new(big.Float).SetInt(num), decimals)
}

func parseDecimalsFromFloat(num *big.Float, decimals float64) float64 {
	parsed, _ := num.Float64()
	parsed = parsed / math.Pow(10, decimals)
	return parsed
}

func FetchPoolStatsBancor(bancor_contract common.Address, from common.Address, to common.Address, amount *big.Int, blockNumber uint64) *big.Int {

	opts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(blockNumber), Context: context.TODO()}
	bancor_caller, err := bancor.NewBancorCaller(bancor_contract, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to create bancor caller: %v\n", err))
	}
	path, err := bancor_caller.ConversionPath(opts, from, to)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get conversion path: %v\n", err))
	}

	bancorRes, err := bancor_caller.RateByPath(opts, path, amount)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get rate for path: %v\n", err))
	}
	return bancorRes
}

// amount - the amount of token0 to send
// returns the recived amount of token1 given the input
func FetchPoolStatsUniswap(pool common.Address, amount_in *big.Int, in common.Address, out common.Address, token0Decimals, token1Decimals int, blockNumber uint64) *big.Float {

	decimalDiff := math.Abs(float64(token0Decimals - token1Decimals))
	if token0Decimals > token1Decimals {
		amount_in = new(big.Int).Mul(amount_in, big.NewInt(int64(math.Pow(10, decimalDiff))))
	} else {
		amount_in = new(big.Int).Quo(amount_in, big.NewInt(int64(math.Pow(10, decimalDiff))))

	}

	pair, err := uniswap.NewUniswapv2pairCaller(pool, Client)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
	}
	amount1, err := uniswap.GetExchangeAmount(pair, new(big.Float).SetInt(amount_in), in, out, blockNumber)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to get exchange amount: %v\n", err))
	}
	return amount1
}
