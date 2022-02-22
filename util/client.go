package util

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"playground/bindings/bancor"
	"playground/bindings/uniswap"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MyClient struct {
	*ethclient.Client
}

func GetClient(node string) *MyClient {
	cl, err := ethclient.Dial(node)
	if err != nil {
		log.Panic("cannot connect to node", err)
	}
	return &MyClient{cl}
}

func (cl *MyClient) GetBlock(blockId big.Int) (*types.Block, error) {
	return cl.BlockByNumber(context.Background(), &blockId)

}

func (cl *MyClient) GetAccountBalance(address string) (*big.Int, error) {
	return cl.BalanceAt(context.Background(), common.HexToAddress(address), nil)
}

func (cl *MyClient) CheckIsContract(address string) (bool, error) {
	bytecode, err := cl.CodeAt(context.Background(), common.HexToAddress(address), nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	return len(bytecode) > 0, nil
}

// amount - the amount of token0 to send
// returns the recived amount of token1 given the input
func (cl *MyClient) FetchPoolStatsBancor(bancor_contract common.Address, from common.Address, to common.Address, amount *big.Int, blockNumber *big.Int) (*big.Float, error) {

	opts := &bind.CallOpts{BlockNumber: blockNumber, Context: context.TODO()}
	bancor_caller, err := bancor.NewBancorCaller(bancor_contract, cl)
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
func (cl *MyClient) FetchPoolStatsUniswap(pool common.Address, amount_in *big.Int, in common.Address,
	out common.Address, token0Decimals, token1Decimals int, blockNumber *big.Int) (*big.Float, error) {

	decimalDiff := math.Abs(float64(token0Decimals - token1Decimals))
	if token0Decimals > token1Decimals {
		amount_in = new(big.Int).Mul(amount_in, big.NewInt(int64(math.Pow(10, decimalDiff))))
	} else {
		amount_in = new(big.Int).Quo(amount_in, big.NewInt(int64(math.Pow(10, decimalDiff))))

	}
	pair, err := uniswap.NewUniswapv2pairCaller(pool, cl)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("failed to instantiate pair caller:  %v", err)
	}
	amount1, err := GetExchangeAmount(pair, new(big.Float).SetInt(amount_in), in, out, blockNumber)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("failed to get exchange amount:  %v", err)
	}
	return amount1, nil
}
