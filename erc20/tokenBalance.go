package erc20

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"price_monitor/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenBalance struct {
	Contract common.Address
	Wallet   common.Address
	Name     string
	Symbol   string
	Balance  *big.Int
	ETH      *big.Int
	Decimals int64
	Block    int64
	ctx      context.Context
}

type tokenBalanceJson struct {
	Contract string `json:"token"`
	Wallet   string `json:"wallet"`
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Balance  string `json:"balance"`
	ETH      string `json:"eth_balance"`
	Decimals int64  `json:"decimals"`
	Block    int64  `json:"block"`
}

func New(client *ethclient.Client, contract, wallet string) (*TokenBalance, error) {
	var err error
	tb := &TokenBalance{
		Contract: common.HexToAddress(contract),
		Wallet:   common.HexToAddress(wallet),
		Decimals: 0,
		Balance:  big.NewInt(0),
		ctx:      context.TODO(),
	}
	err = tb.query(client)
	return tb, err
}

func (tb *TokenBalance) ETHString() string {
	return util.BigIntString(tb.ETH, 18)
}

func (tb *TokenBalance) BalanceString() string {
	if tb.Decimals == 0 {
		return tb.Balance.String()
	}
	return util.BigIntString(tb.Balance, tb.Decimals)
}

func (tb *TokenBalance) query(client *ethclient.Client) error {
	var err error

	token, err := NewUniswapv2erc20Caller(tb.Contract, client)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to instantiate a token contract: %v\n", err), false)
		return err
	}

	block, err := client.BlockByNumber(tb.ctx, nil)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to get current block number: %v\n", err), false)
	}

	tb.Block = block.Number().Int64()

	decimals, err := token.Decimals(nil)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to get decimals from contract: %v \n", tb.Contract.String()), false)
		return err
	}
	tb.Decimals = int64(decimals)

	tb.ETH, err = client.BalanceAt(tb.ctx, tb.Wallet, nil)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to get ethereum balance from address: %v \n", tb.Wallet.String()), false)
	}

	tb.Balance, err = token.BalanceOf(nil, tb.Wallet)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to get balance from contract: %v %v\n", tb.Contract.String(), err), false)
		tb.Balance = big.NewInt(0)
	}

	tb.Symbol, err = token.Symbol(nil)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to get symbol from contract: %v \n", tb.Contract.String()), false)
		tb.Symbol = symbolFix(tb.Contract.String())
	}

	tb.Name, err = token.Name(nil)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to retrieve token name from contract: %v | %v\n", tb.Contract.String(), err), false)
		tb.Name = "MISSING"
	}

	return err
}

func symbolFix(contract string) string {
	switch common.HexToAddress(contract).String() {
	case "0x86Fa049857E0209aa7D9e616F7eb3b3B78ECfdb0":
		return "EOS"
	}
	return "MISSING"
}

func (tb *TokenBalance) ToJSON() string {
	jsonData := tokenBalanceJson{
		Contract: tb.Contract.String(),
		Wallet:   tb.Wallet.String(),
		Name:     tb.Name,
		Symbol:   tb.Symbol,
		Balance:  tb.BalanceString(),
		ETH:      tb.ETHString(),
		Decimals: tb.Decimals,
		Block:    tb.Block,
	}
	d, _ := json.Marshal(jsonData)
	return string(d)
}
