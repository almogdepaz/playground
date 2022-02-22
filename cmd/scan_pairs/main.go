package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"playground/bindings/erc20"
	"playground/bindings/uniswap"

	"playground/util"
	"regexp"
	"runtime"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/shopspring/decimal"
)

const node = "https://mainnet.infura.io/v3/093f1d19defd46248d24aa7e734ea203"
const uniswapFactory = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

const decimals = 18

var reg = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

var cl *util.MyClient

func init() {
	cl = util.GetClient(node)
}

func main() {
	fmt.Printf("\n--------------start--------------\n")

	// fmt.Print("Enter text: ")
	// reader := bufio.NewReader(os.Stdin)
	// // ReadString will block until the delimiter is entered
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An error occured while reading input. Please try again", err)
	// 	return
	// }

	// // remove the delimeter from the string
	// account := strings.TrimSuffix(input, "\n")
	// if !util.IsValidAddress(account) {
	// 	log.Panic("invalid address")
	// }

	var lps []LP
	account := ""
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
			st := tokenStats(t.Lp, account)
			if util.ToDecimal(st.Balance, decimals).GreaterThan(decimal.Zero) {
				stats = append(stats, st)
				printStats(st)
			}
			visited[t.Lp] = struct{}{}
		}
		if _, ok := visited[t.T0]; !ok {
			st1 := tokenStats(t.T0, account)
			if util.ToDecimal(st1.Balance, decimals).GreaterThan(decimal.Zero) {
				stats = append(stats, st1)
				printStats(st1)
			}
			visited[t.T0] = struct{}{}
		}
		if _, ok := visited[t.T1]; !ok {
			st2 := tokenStats(t.T1, account)
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
	lps, err := PullDataFromUniswap(cl, index, &bind.CallOpts{Context: context.TODO()}, uniswapFactory, node)
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

func tokenStats(tokenAddress common.Address, account string) *erc20.TokenBalance {
	instance, err := erc20.New(cl.Client, tokenAddress.String(), account)
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

func ethBalance(account string) (*big.Int, error) {
	// print eth balance
	bal, err := cl.GetAccountBalance(account)
	if err != nil {
		log.Panic("failed fetching balance\n", err)
	}

	fmt.Printf("\n%s eth balance is %v", account, util.ToDecimal(bal, decimals))
	return bal, err
}

type LP struct {
	Lp common.Address
	T0 common.Address
	T1 common.Address
}

func PullDataFromUniswap(cl *util.MyClient, index int, op *bind.CallOpts, factory string, node string) ([]LP, error) {
	var err error

	lpAddress, err := GetAllPairs(cl, index, op, factory, node)
	if err != nil {
		return nil, err
	}

	lps, err := GetPairMeta(cl, lpAddress, op)

	return lps, err
}

func GetPairMeta(cl *util.MyClient, pairs []common.Address, op *bind.CallOpts) ([]LP, error) {
	var err error

	lps := []LP{}
	for i, tk := range pairs {
		pair, err := uniswap.NewUniswapv2pairCaller(tk, cl)
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

		fmt.Printf("\nGetPairMeta %d out of %d", i, len(pairs))
	}
	return lps, err
}

func GetAllPairs(cl *util.MyClient, index int, op *bind.CallOpts, factory string, node string) ([]common.Address, error) {
	var err error
	caller, err := uniswap.NewUniswapv2factoryCaller(common.HexToAddress(factory), cl)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to instantiate a token contract: %v\n", err), false)
		return nil, err
	}
	ln, err := caller.AllPairsLength(op)
	if err != nil {
		panic(err)
	}

	c := make(chan int64)
	resChan := make(chan common.Address)
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go pairWorker(cl, c, resChan, ln, op, factory, node)
	}

	go func() {
		for i := int64(index); i < ln.Int64(); i++ {
			c <- i
		}
	}()

	var lps []common.Address
	for addr := range resChan {
		lps = append(lps, addr)
		fmt.Printf("\nGetAllPairs %d out of %d", len(lps), ln.Int64())
	}

	return lps, nil
}

func pairWorker(cl *util.MyClient, c chan int64, resChan chan common.Address, ln *big.Int, op *bind.CallOpts, factory string, node string) {
	caller, err := uniswap.NewUniswapv2factoryCaller(common.HexToAddress(factory), cl)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to instantiate a token contract: %v\n", err), false)
	}
	for i := range c {
		tk, err := caller.AllPairs(op, big.NewInt(i))
		if err != nil {
			panic(err)
		}
		resChan <- tk
	}
}
