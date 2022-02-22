package util

import (
	"context"
	"encoding/hex"
	"math/big"
	"playground/bindings/uniswap"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	FactoryAddrV3  = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	FactoryAddrV2  = "0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f"
	FactoryAddress = "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
)

var (
	Address, _       = abi.NewType("address", "", nil)
	Uint24, _        = abi.NewType("uint24", "", nil)
	saltAbiArguments = abi.Arguments{
		abi.Argument{
			Name:    "token0",
			Type:    Address,
			Indexed: false,
		},
		abi.Argument{
			Name:    "token1",
			Type:    Address,
			Indexed: false,
		},
		abi.Argument{
			Name:    "fee",
			Type:    Uint24,
			Indexed: false,
		},
	}
	PoolInitCodeV3, _ = hex.DecodeString("e34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54")
	PoolInitCodeV2, _ = hex.DecodeString("96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f")
	SUSHIINITCODE, _  = hex.DecodeString("e18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303")
)

func CalculatePoolAddressSushi(token0, token1 common.Address) (pairAddress common.Address, err error) {
	return CalculatePoolAddress(token0, token1, common.HexToAddress(FactoryAddress), SUSHIINITCODE)
}

func CalculatePoolAddressUniV2(token0, token1 common.Address) (pairAddress common.Address, err error) {
	return CalculatePoolAddress(token0, token1, common.HexToAddress(FactoryAddrV2), PoolInitCodeV2)
}

func CalculatePoolAddress(token0, token1, factoryAddr common.Address, poolInitCode []byte) (pairAddress common.Address, err error) {
	tkn0, tkn1 := SortAddressess(token0, token1)

	msg := []byte{255}
	msg = append(msg, factoryAddr.Bytes()...)
	addrBytes := tkn0.Bytes()
	addrBytes = append(addrBytes, tkn1.Bytes()...)
	msg = append(msg, crypto.Keccak256(addrBytes)...)

	msg = append(msg, poolInitCode...)
	hash := crypto.Keccak256(msg)
	pairAddressBytes := big.NewInt(0).SetBytes(hash)
	pairAddressBytes = pairAddressBytes.Abs(pairAddressBytes)
	return common.BytesToAddress(pairAddressBytes.Bytes()), nil
}

// CalculatePoolAddressV3 calculate uniswapV3 pool address offline from pool tokens and fee
func CalculatePoolAddressV3(tokenA, tokenB string, fee *big.Int) (poolAddress common.Address, err error) {
	tkn0, tkn1 := SortAddressess(common.HexToAddress(tokenA), common.HexToAddress(tokenB))
	paramsPacked, err := saltAbiArguments.Pack(tkn0, tkn1, fee)
	if err != nil {
		err = errors.Wrap(err, "pack arguments")
		return
	}

	salt := crypto.Keccak256(paramsPacked)
	// "0xff"
	msg := []byte{255}
	msg = append(msg, common.HexToAddress(FactoryAddrV3).Bytes()...)
	msg = append(msg, salt...)
	msg = append(msg, PoolInitCodeV3...)

	hash := crypto.Keccak256(msg)
	return common.BytesToAddress(hash[12:]), nil
}

// GetReserves retursn the available reserves in a pair
func GetReserves(caller *uniswap.Uniswapv2pairCaller, token0, token1 common.Address, blockNumber *big.Int) (*struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	reserves, err := caller.GetReserves(&bind.CallOpts{BlockNumber: blockNumber, Context: ctx})
	if err != nil {
		return nil, err
	}
	// This is the tricky bit.
	// The reserve call returns the reserves for token0 and token1 in a sorted order.
	// This means we need to check if our token addresses are sorted or not and flip the reserves if they are not sorted.
	stoken0, _ := SortAddressess(token0, token1)
	if stoken0 != token0 {
		// We're not sorted, so the reserves need to be flipped to represent the actual reserves.
		reserves.Reserve0, reserves.Reserve1 = reserves.Reserve1, reserves.Reserve0
	}
	return &reserves, nil
}

// GetExchangeAmount returns the amount of tokens you'd receive when exchanging the given amount of token0 to token1.
func GetExchangeAmount(caller *uniswap.Uniswapv2pairCaller, amount *big.Float, token0, token1 common.Address, blockNumber *big.Int) (*big.Float, error) {
	reserves, err := GetReserves(caller, token0, token1, blockNumber)
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
