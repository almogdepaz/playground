package main

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/signal"
	"playground/bindings/badger"
	"playground/util"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const node = "https://mainnet.infura.io/v3/093f1d19defd46248d24aa7e734ea203"

const std_decimals = 18
const wbtc_decimals = 8
const diggDecimals = 9

const interval = 10 * time.Second

var cl *util.MyClient

func init() {
	cl = util.GetClient(node)
}

func main() {
	// addr := common.HexToAddress("0x424A70C78b1242B19DC1e62Ee14cB2Fc2ae57998")
	badger_contract := common.HexToAddress("0x7e7E112A68d8D2E221E11047a72fFC1065c38e1a")
	bancor_contract := common.HexToAddress("0x2F9EC37d6CcFFf1caB21733BdaDEdE11c823cCB0")

	bancorEth := common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	bDigg := common.HexToAddress("0x7e7e112a68d8d2e221e11047a72ffc1065c38e1a")
	digg := common.HexToAddress("0x798d1be841a82a273720ce31c822c61a67a601c3")
	wbtc := common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")
	weth := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	// sushi pools
	sushiWbtcDiggPool, _ := util.CalculatePoolAddressSushi(wbtc, digg)
	sushiEthWbtcPool, _ := util.CalculatePoolAddressSushi(weth, wbtc)
	sushiEthBdiggPool, _ := util.CalculatePoolAddressSushi(weth, bDigg)

	// uni pools
	uniWbtcDiggPool, _ := util.CalculatePoolAddressUniV2(wbtc, digg)
	uniEthWbtcPool, _ := util.CalculatePoolAddressUniV2(weth, wbtc)

	LOG_FILE := "debug.log"
	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	mw := io.MultiWriter(os.Stdout, logFile)

	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	ethAmount := new(big.Int).Mul(big.NewInt(1), big.NewInt(int64(1e18)))

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
		blockRes, err := cl.BlockNumber(context.TODO())
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
		bancorRes, err := cl.FetchPoolStatsBancor(bancor_contract, bancorEth, bDigg, ethAmount, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		sushiEthBdiggRes, err := cl.FetchPoolStatsUniswap(sushiEthBdiggPool, ethAmount, weth, bDigg, std_decimals, std_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(fmt.Sprintf("bancor ETH/bDIGG %v eth in %v bdigg out", util.ParseDecimalsFromInt(ethAmount, std_decimals), util.ParseDecimalsFromFloat(bancorRes, std_decimals)))
		log.Info(fmt.Sprintf("Sushiswap ETH/bDIGG %v eth in %v bdigg out", util.ParseDecimalsFromInt(ethAmount, std_decimals), util.ParseDecimalsFromFloat(sushiEthBdiggRes, std_decimals)))
		res := sushiEthBdiggRes
		if res.Cmp(bancorRes) < 0 {
			res = bancorRes
		}
		// bdigg to digg
		badger_caller, err := badger.NewBadgerCaller(badger_contract, cl)
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
		log.Info(fmt.Sprintf("Badger %v bDIGG in %v Digg out ", util.ParseDecimalsFromFloat(res, std_decimals), util.ParseDecimalsFromInt(diggOutputInt, std_decimals)))

		// digg to wbtc
		uniWbtcRes, err := cl.FetchPoolStatsUniswap(uniWbtcDiggPool, diggOutputInt, digg, wbtc, diggDecimals, wbtc_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		sushiWbtcRes, err := cl.FetchPoolStatsUniswap(sushiWbtcDiggPool, diggOutputInt, digg, wbtc, diggDecimals, wbtc_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(fmt.Sprintf("Uniswap WBTC/DIGG %v digg in %v  wbtc out", util.ParseDecimalsFromInt(diggOutputInt, std_decimals), util.ParseDecimalsFromFloat(uniWbtcRes, std_decimals)))
		log.Info(fmt.Sprintf("Sushiswap WBTC/DIGG %v digg in %v  wbtc out", util.ParseDecimalsFromInt(diggOutputInt, std_decimals), util.ParseDecimalsFromFloat(sushiWbtcRes, std_decimals)))
		wbtcRes := new(big.Int)
		if uniWbtcRes.Cmp(sushiWbtcRes) < 0 {
			sushiWbtcRes.Int(wbtcRes)
		} else {
			uniWbtcRes.Int(wbtcRes)
		}

		// wbtc to eth
		uniEthRes, err := cl.FetchPoolStatsUniswap(uniEthWbtcPool, wbtcRes, wbtc, weth, wbtc_decimals, std_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		sushiEthRes, err := cl.FetchPoolStatsUniswap(sushiEthWbtcPool, wbtcRes, wbtc, weth, wbtc_decimals, std_decimals, blockNumber)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info(fmt.Sprintf("Uniswap ETH/WBTC %v wbtc in %v eth out", util.ParseDecimalsFromInt(wbtcRes, std_decimals), util.ParseDecimalsFromFloat(uniEthRes, std_decimals)))
		log.Info(fmt.Sprintf("Sushiswap ETH/WBTC  %v wbtc in %v eth out", util.ParseDecimalsFromInt(wbtcRes, std_decimals), util.ParseDecimalsFromFloat(sushiEthRes, std_decimals)))
		ethRes := new(big.Int)
		if uniEthRes.Cmp(sushiEthRes) < 0 {
			sushiEthRes.Int(ethRes)
		} else {
			uniEthRes.Int(ethRes)
		}

		log.Info(fmt.Sprintf("%v wbtc in get %v eth out", util.ParseDecimalsFromInt(wbtcRes, std_decimals), util.ParseDecimalsFromInt(ethRes, std_decimals)))
		parsedEthIn := util.ParseDecimalsFromInt(ethAmount, std_decimals)
		parsedEthOut := util.ParseDecimalsFromInt(ethRes, std_decimals)
		log.Info(fmt.Sprintf("results %v eth in get %v eth out", parsedEthIn, parsedEthOut))
		delta := parsedEthOut - parsedEthIn
		if delta > maxProfit {
			maxProfit = delta
			log.Info(fmt.Sprintf("found max profit %v ", maxProfit))
		}
		time.Sleep(interval)
		// build call and estimate gas
		// call:=&CallMsg{
		// 	From      common.Address  // the sender of the 'transaction'
		// 	To        *common.Address // the destination contract (nil for contract creation)
		// 	Gas       uint64          // if 0, the call executes with near-infinite gas
		// 	GasPrice  *big.Int        // wei <-> gas exchange ratio
		// 	GasFeeCap *big.Int        // EIP-1559 fee cap per gas.
		// 	GasTipCap *big.Int        // EIP-1559 tip per gas.
		// 	Value     *big.Int        // amount of wei sent along with the call
		// 	Data      []byte          // input data, usually an ABI-encoded contract method invocation

		// 	AccessList types.AccessList // EIP-2930 access list.
		// }

		// gas := Client.EstimateGas(context.TODO(),call)
		// // txData := &types.DynamicFeeTx{
		// // 	Nonce:      tx.Nonce,
		// // 	To:         to,
		// // 	Data:       data,
		// // 	Gas:        Gas,
		// // 	AccessList: AccessList,
		// // 	Value:      value,
		// // 	ChainID:    chainid,
		// // 	GasTipCap:  tip,
		// // 	GasFeeCap:  fee,
		// // 	V:          v,
		// // 	R:          r,
		// // 	S:          s,
		// // }
		// // tx := types.NewTx(txData)
		// // Client.SendTransaction(context.TODO(), txData)
		// log.Info("------------------------------------------------------------------------------------------")
	}

}
