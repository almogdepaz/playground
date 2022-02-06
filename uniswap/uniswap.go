package uniswap

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"price_monitor/util"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const factory = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

type LP struct {
	Lp common.Address
	T0 common.Address
	T1 common.Address
}

func PullDataFromUniswap(client *ethclient.Client, index int, op *bind.CallOpts) ([]LP, error) {
	var err error

	lpAddress, err := GetAllPairs(client, index, op)
	if err != nil {
		return nil, err
	}

	lps, err := GetPairMeta(client, lpAddress, op)

	return lps, err
}

func GetPairMeta(client *ethclient.Client, pairs []common.Address, op *bind.CallOpts) ([]LP, error) {
	var err error

	lps := []LP{}
	for i, tk := range pairs {
		pair, err := NewUniswapv2pairCaller(tk, client)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to instantiate pair caller: %v\n", err))
			return nil, err
		}

		t0, err := pair.Token0(op)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to fetch %s t0 : %v\n", tk, err))
			return nil, err
		}

		t1, err := pair.Token1(op)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to fetch %s t1 : %v\n", tk, err))
			return nil, err
		}

		lps = append(lps, LP{tk, t0, t1})

		fmt.Println(fmt.Sprintf("GetPairMeta %d out of %d", i, len(pairs)))
	}
	return lps, err
}

func GetAllPairs(client *ethclient.Client, index int, op *bind.CallOpts) ([]common.Address, error) {
	var err error

	token, err := NewUniswapv2factoryCaller(common.HexToAddress(factory), client)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to instantiate a token contract: %v\n", err), false)
		return nil, err
	}

	ln, err := token.AllPairsLength(op)
	if err != nil {
		panic(err)
	}

	var lps []common.Address
	for i := int64(index); i < ln.Int64(); i++ {
		tk, err := token.AllPairs(op, big.NewInt(i))
		if err != nil {
			panic(err)
		}
		lps = append(lps, tk)
		fmt.Println(fmt.Sprintf("GetAllPairs %d out of %d", i, ln.Int64()))
	}
	return lps, nil
}

// GetReserves retursn the available reserves in a pair
func GetReserves(caller *Uniswapv2pairCaller, token0, token1 common.Address) (*struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	reserves, err := caller.GetReserves(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	// This is the tricky bit.
	// The reserve call returns the reserves for token0 and token1 in a sorted order.
	// This means we need to check if our token addresses are sorted or not and flip the reserves if they are not sorted.
	stoken0, _ := util.SortAddressess(token0, token1)
	if stoken0 != token0 {
		// We're not sorted, so the reserves need to be flipped to represent the actual reserves.
		reserves.Reserve0, reserves.Reserve1 = reserves.Reserve1, reserves.Reserve0
	}
	return &reserves, nil
}

// GetExchangeAmount returns the amount of tokens you'd receive when exchanging the given amount of token0 to token1.
func GetExchangeAmount(caller *Uniswapv2pairCaller, amount *big.Float, token0, token1 common.Address) (*big.Float, error) {
	reserves, err := GetReserves(caller, token0, token1)
	if err != nil {
		return nil, err
	}
	return Quote(amount, new(big.Float).SetInt(reserves.Reserve0), new(big.Float).SetInt(reserves.Reserve1)), nil
}

func Quote(amount, reserve0, reserve1 *big.Float) *big.Float {
	if reserve1.Cmp(big.NewFloat(0)) <= 0 ||
		reserve0.Cmp(big.NewFloat(0)) <= 0 ||
		amount.Cmp(big.NewFloat(0)) <= 0 {

		return new(big.Float)
	}

	// amountB = amountA.mul(reserveB) / reserveA;
	multiplied := new(big.Float).Mul(amount, reserve1)
	res := new(big.Float).Quo(multiplied, reserve0)
	return res
}
