package main

import (
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"os"
	"os/signal"
	"price_monitor/badger"
	"price_monitor/bancor"
	uniswap "price_monitor/uniswap"
	"price_monitor/util"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

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
	ethAmount := new(big.Int).Mul(big.NewInt(1), big.NewInt(int64(1e18)))
	badger_contract := common.HexToAddress("0x7e7E112A68d8D2E221E11047a72fFC1065c38e1a") //badger contract
	bancor_contract := common.HexToAddress("0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0") //bancor contract

	bancorEth := common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	bDigg := common.HexToAddress("0x7e7e112a68d8d2e221e11047a72ffc1065c38e1a")
	digg := common.HexToAddress("0x798d1be841a82a273720ce31c822c61a67a601c3")
	wbtc := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")
	weth := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	// sushi pools
	sushiWbtcDiggPool, _ := uniswap.CalculatePoolAddressSushi(wbtc, digg)
	sushiEthWbtcPool, _ := uniswap.CalculatePoolAddressSushi(weth, wbtc)
	sushiEthBdiggPool, _ := uniswap.CalculatePoolAddressSushi(weth, bDigg)

	// uni pools
	uniWbtcDiggPool, _ := uniswap.CalculatePoolAddressUniV2(wbtc, digg)
	uniEthWbtcPool, _ := uniswap.CalculatePoolAddressUniV2(weth, wbtc)

	blockNumber := big.NewInt(0)
	log.SetOutput(mw)
	maxProfit := float64(0)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// print profit before exit
	go func() {
		<-c
		fmt.Printf("\n max profit %v", maxProfit)
		os.Exit(0)
	}()

	for {
		prev := blockNumber
		opts := bind.CallOpts{Context: context.TODO()}
		blockRes, err := Client.BlockNumber(context.TODO())
		blockNumber = new(big.Int).SetUint64(blockRes)
		if prev == blockNumber {
			log.Info(fmt.Sprintf("last block %v, waiting for next block", blockNumber))
			time.Sleep(interval)
			continue
		}

		if err != nil {
			log.Error(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
		}
		// eth to bdigg
		bancorRes, err := FetchPoolStatsBancor(bancor_contract, bancorEth, bDigg, ethAmount, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		sushiEthBdiggRes, err := FetchPoolStatsUniswap(sushiEthBdiggPool, ethAmount, weth, bDigg, std_decimals, std_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(fmt.Sprintf("bancor ETH/bDIGG %v eth in %v bdigg out", parseDecimalsFromInt(ethAmount, std_decimals), parseDecimalsFromFloat(bancorRes, std_decimals)))
		log.Info(fmt.Sprintf("Sushiswap ETH/bDIGG %v eth in %v bdigg out", parseDecimalsFromInt(ethAmount, std_decimals), parseDecimalsFromFloat(sushiEthBdiggRes, std_decimals)))
		res := sushiEthBdiggRes
		if res.Cmp(bancorRes) < 0 {
			res = bancorRes
		}
		// bdigg to digg
		badger_caller, err := badger.NewBadgerCaller(badger_contract, Client)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
			continue
		}
		supply, err := badger_caller.TotalSupply(&opts)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
			continue
		}
		balance, err := badger_caller.Balance(&opts)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
			continue
		}

		sharePrice := big.NewFloat(1 / (float64(new(big.Int).Div(supply, balance).Int64()) / 1e9))
		log.Info(fmt.Sprintf("bDIGG share price %v", sharePrice))
		diggOut, _ := new(big.Float).Mul(res, sharePrice).Int64()
		diggOutputInt := big.NewInt(diggOut)
		log.Info(fmt.Sprintf("Badger %v bDIGG in %v Digg out ", parseDecimalsFromFloat(res, std_decimals), parseDecimalsFromInt(diggOutputInt, std_decimals)))

		// digg to wbtc
		uniWbtcRes, err := FetchPoolStatsUniswap(uniWbtcDiggPool, diggOutputInt, digg, wbtc, diggDecimals, wbtc_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		sushiWbtcRes, err := FetchPoolStatsUniswap(sushiWbtcDiggPool, diggOutputInt, digg, wbtc, diggDecimals, wbtc_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(fmt.Sprintf("Uniswap WBTC/DIGG %v digg in %v  wbtc out", parseDecimalsFromInt(diggOutputInt, std_decimals), parseDecimalsFromFloat(uniWbtcRes, std_decimals)))
		log.Info(fmt.Sprintf("Sushiswap WBTC/DIGG %v digg in %v  wbtc out", parseDecimalsFromInt(diggOutputInt, std_decimals), parseDecimalsFromFloat(sushiWbtcRes, std_decimals)))
		wbtcRes := new(big.Int)
		if uniWbtcRes.Cmp(sushiWbtcRes) < 0 {
			sushiWbtcRes.Int(wbtcRes)
		} else {
			uniWbtcRes.Int(wbtcRes)
		}

		// wbtc to eth
		uniEthRes, err := FetchPoolStatsUniswap(uniEthWbtcPool, wbtcRes, wbtc, weth, wbtc_decimals, std_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		sushiEthRes, err := FetchPoolStatsUniswap(sushiEthWbtcPool, wbtcRes, wbtc, weth, wbtc_decimals, std_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(fmt.Sprintf("Uniswap ETH/WBTC %v wbtc in %v eth out", parseDecimalsFromInt(wbtcRes, std_decimals), parseDecimalsFromFloat(uniEthRes, std_decimals)))
		log.Info(fmt.Sprintf("Sushiswap ETH/WBTC  %v wbtc in %v eth out", parseDecimalsFromInt(wbtcRes, std_decimals), parseDecimalsFromFloat(sushiEthRes, std_decimals)))
		ethRes := new(big.Int)
		if uniEthRes.Cmp(sushiEthRes) < 0 {
			sushiEthRes.Int(ethRes)
		} else {
			uniEthRes.Int(ethRes)
		}

		log.Info(fmt.Sprintf("%v wbtc in get %v eth out", parseDecimalsFromInt(wbtcRes, std_decimals), parseDecimalsFromInt(ethRes, std_decimals)))
		parsedEthIn := parseDecimalsFromInt(ethAmount, std_decimals)
		parsedEthOut := parseDecimalsFromInt(ethRes, std_decimals)
		log.Info(fmt.Sprintf("results %v eth in get %v eth out", parsedEthIn, parsedEthOut))
		delta := parsedEthOut - parsedEthIn
		if delta > maxProfit {
			maxProfit = delta
			log.Info(fmt.Sprintf("found max profit %v ", maxProfit))
		}
		time.Sleep(interval)
		log.Info("------------------------------------------------------------------------------------------")
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

func FetchPoolStatsBancor(bancor_contract common.Address, from common.Address, to common.Address, amount *big.Int, blockNumber *big.Int) (*big.Float, error) {

	opts := &bind.CallOpts{BlockNumber: blockNumber, Context: context.TODO()}
	bancor_caller, err := bancor.NewBancorCaller(bancor_contract, Client)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("failed to create bancor caller: %v", err)

	}
	path, err := bancor_caller.ConversionPath(opts, from, to)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("Failed to get conversion path: %v", err)

	}

	bancorRes, err := bancor_caller.RateByPath(opts, path, amount)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("failed to get rate for path: %v", err)
	}

	return new(big.Float).SetInt64(bancorRes.Int64()), nil
}

// amount - the amount of token0 to send
// returns the recived amount of token1 given the input
func FetchPoolStatsUniswap(pool common.Address, amount_in *big.Int, in common.Address,
	out common.Address, token0Decimals, token1Decimals int, blockNumber *big.Int) (*big.Float, error) {

	decimalDiff := math.Abs(float64(token0Decimals - token1Decimals))
	if token0Decimals > token1Decimals {
		amount_in = new(big.Int).Mul(amount_in, big.NewInt(int64(math.Pow(10, decimalDiff))))
	} else {
		amount_in = new(big.Int).Quo(amount_in, big.NewInt(int64(math.Pow(10, decimalDiff))))

	}
	pair, err := uniswap.NewUniswapv2pairCaller(pool, Client)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("failed to instantiate pair caller:  %v", err)
	}
	amount1, err := uniswap.GetExchangeAmount(pair, new(big.Float).SetInt(amount_in), in, out, blockNumber)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("failed to get exchange amount:  %v", err)
	}
	return amount1, nil
}
