package uniswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const factory = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

//uniswap = w3.eth.contract('0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95', abi=abi)
//events = uniswap.events.NewExchange.createFilter(fromBlock=6627917).get_all_entries()
//token_exchange = {e.args.token: e.args.exchange for e in events}
//
//for token, exchange in token_exchange.items():
//print(token, exchange)


type LP struct {
	Lp common.Address
	T0 common.Address
	T1 common.Address
}


func PullDataFromUniswap(client *ethclient.Client,index int,op *bind.CallOpts) ([]LP, error) {
	var err error

	lpAddress, err := GetAllPairs(client,index, op)
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
		pair, err := NewPairCaller(tk, client)
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

func GetAllPairs(client *ethclient.Client, index int,op *bind.CallOpts) ([]common.Address, error) {
	var err error

	token, err := NewFactoryCaller(common.HexToAddress(factory), client)
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
