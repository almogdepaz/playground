package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"regexp"
	"roi/erc20"
	"roi/uniswap"
	"roi/util"
)

type MyClient struct {
	*ethclient.Client
}

const node = "https://mainnet.infura.io/v3/093f1d19defd46248d24aa7e734ea203"

const account = "0xf4721f8cc66436456f2230764d91782f7c09be8r"

const decimals = 18

var reg = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

var Client MyClient

func init() {
	cl, err := ethclient.Dial(node)
	if err != nil {
		log.Panic("cannot connect to node", err)
	}
	Client = MyClient{cl}
}

func (cl MyClient) getBlock(blockId big.Int) (*types.Block, error) {
	return cl.BlockByNumber(context.Background(), &blockId)

}

func (cl MyClient) getAccountBalance(address string) (*big.Int, error) {
	return cl.BalanceAt(context.Background(), common.HexToAddress(address), nil)
}

func (cl MyClient) checkIsContract(address string) (bool, error) {
	bytecode, err := Client.CodeAt(context.Background(), common.HexToAddress(address), nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	return len(bytecode) > 0, nil
}

func main() {
	fmt.Printf("\n--------------start--------------\n")
	if !util.IsValidAddress(account) {
		log.Panic("invalid address")
	}

	var lps []uniswap.LP

	err := util.ReadGob("./uni_listings.gob", &lps)
	if err != nil {
		fmt.Println("err reading uniswap data from disk,", err)
	}

	fetchAndPersist(len(lps))

	stats := []*erc20.TokenBalance{}
	visited := map[common.Address]struct{}{}
	// print all token balances
	for _, t := range lps {

		if _, ok := visited[t.Lp]; !ok {
			st := tokenStats(t.Lp)
			if util.ToDecimal(st.Balance, decimals).GreaterThan(decimal.Zero) {
				stats = append(stats, st)
				printStats(st)
			}
			visited[t.Lp] = struct{}{}
		}
		if _, ok := visited[t.T0]; !ok {
			st1 := tokenStats(t.T0)
			if util.ToDecimal(st1.Balance, decimals).GreaterThan(decimal.Zero) {
				stats = append(stats, st1)
				printStats(st1)
			}
			visited[t.T0] = struct{}{}
		}
		if _, ok := visited[t.T1]; !ok {
			st2 := tokenStats(t.T1)
			if util.ToDecimal(st2.Balance, decimals).GreaterThan(decimal.Zero) {
				stats = append(stats, st2)
				printStats(st2)
			}
			visited[t.T1] = struct{}{}
		}
	}

}

func fetchAndPersist(index int) {
	fmt.Println("fetching data from uniswap ......")

	if index > 0 {
		fmt.Println(fmt.Sprintf("starting from index %d", index))
	}
	lps, err := uniswap.PullDataFromUniswap(Client.Client, index, &bind.CallOpts{Context: context.TODO()})
	if err != nil {
		panic("failed pulling token data from uniswap contract")
	}

	err = util.WriteGob("./uni_listings.gob", lps)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("persisted uniswap data to disk")
	}
}

func tokenStats(tokenAddress common.Address) *erc20.TokenBalance {
	instance, err := erc20.New(Client.Client, tokenAddress.String(), account)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

//todo get token balances and relevant pairs through tx history
func PrintAccountTXHistory(tokenAddress common.Address) {

}

func printStats(instance *erc20.TokenBalance) {
	fmt.Printf("name: %s\n", instance.Name)
	fmt.Printf("symbol: %s\n", instance.Symbol)
	fmt.Printf("decimals: %v\n", instance.Decimals)
	fmt.Printf("token balance is: %s\n", util.ToDecimal(instance.Balance, decimals).String())
}

func ethBalance() (*big.Int, error) {
	// print eth balance
	bal, err := Client.getAccountBalance(account)
	if err != nil {
		log.Panic("failed fetching balance\n", err)
	}

	fmt.Printf("\n%s eth balance is %v", account, util.ToDecimal(bal, decimals))
	return bal, err
}
