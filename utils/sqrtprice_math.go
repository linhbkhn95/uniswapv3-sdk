package utils

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
)

var (
	ErrSqrtPriceLessThanZero = errors.New("sqrt price less than zero")
	ErrLiquidityLessThanZero = errors.New("liquidity less than zero")
	ErrInvariant             = errors.New("invariant violation")
)
var MaxUint160 = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil), constants.One)

func multiplyIn256(x, y *big.Int) *big.Int {
	product := new(big.Int).Mul(x, y)
	return new(big.Int).And(product, entities.MaxUint256)
}

func GetAmount0Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int, roundUp bool) *big.Int {
	// https://github.com/Uniswap/v3-core/blob/d8b1c635c275d2a9450bd6a78f3fa2484fef73eb/contracts/libraries/SqrtPriceMath.sol#L159
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	numerator1 := new(big.Int).Lsh(liquidity, 96)
	numerator2 := new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96)

	if roundUp {
		result := MulDivRoundingUpModified(numerator1, numerator2, sqrtRatioBX96)
		return DivRoundingUpModified(result, sqrtRatioAX96, numerator2)
	}
	return numerator1.Mul(numerator1, numerator2).Div(numerator1, sqrtRatioBX96).Div(numerator1, sqrtRatioAX96)
}

func GetAmount1Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int, roundUp bool) *big.Int {
	// https://github.com/Uniswap/v3-core/blob/d8b1c635c275d2a9450bd6a78f3fa2484fef73eb/contracts/libraries/SqrtPriceMath.sol#L188
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	result := new(big.Int)
	if roundUp {
		result.Sub(sqrtRatioBX96, sqrtRatioAX96).Mul(result, liquidity)
		return DivRoundingUpModified(result, constants.Q96, new(big.Int))
	}
	return result.Sub(sqrtRatioBX96, sqrtRatioAX96).Mul(result, liquidity).Div(result, constants.Q96)
}

func GetNextSqrtPriceFromInput(sqrtPX96, liquidity, amountIn *big.Int, zeroForOne bool) (*big.Int, error) {
	if sqrtPX96.Cmp(constants.Zero) <= 0 {
		return nil, ErrSqrtPriceLessThanZero
	}
	if liquidity.Cmp(constants.Zero) <= 0 {
		return nil, ErrLiquidityLessThanZero
	}
	if zeroForOne {
		return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountIn, true)
	}
	return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountIn, true)
}

func GetNextSqrtPriceFromOutput(sqrtPX96, liquidity, amountOut *big.Int, zeroForOne bool) (*big.Int, error) {
	if sqrtPX96.Cmp(constants.Zero) <= 0 {
		return nil, ErrSqrtPriceLessThanZero
	}
	if liquidity.Cmp(constants.Zero) <= 0 {
		return nil, ErrLiquidityLessThanZero
	}
	if zeroForOne {
		return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountOut, false)
	}
	return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountOut, false)
}

func getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amount *big.Int, add bool) (*big.Int, error) {
	if amount.Cmp(constants.Zero) == 0 {
		return sqrtPX96, nil
	}

	numerator1 := new(big.Int).Lsh(liquidity, 96)
	if add {
		product := multiplyIn256(amount, sqrtPX96)
		temp := new(big.Int)
		if temp.Div(product, amount).Cmp(sqrtPX96) == 0 {
			// addIn256
			temp = temp.Add(numerator1, product).And(temp, entities.MaxUint256)
			if temp.Cmp(numerator1) >= 0 {
				numerator1.Mul(numerator1, sqrtPX96)
				return DivRoundingUpModified(numerator1, temp, new(big.Int)), nil
			}
		}
		temp.Div(numerator1, sqrtPX96).Add(temp, amount)
		return DivRoundingUpModified(numerator1, temp, new(big.Int)), nil
	} else {
		product := multiplyIn256(amount, sqrtPX96)
		temp := new(big.Int)
		if temp.Div(product, amount).Cmp(sqrtPX96) != 0 {
			return nil, ErrInvariant
		}
		if numerator1.Cmp(product) <= 0 {
			return nil, ErrInvariant
		}
		temp.Sub(numerator1, product)
		numerator1.Mul(numerator1, sqrtPX96)
		return DivRoundingUpModified(numerator1, temp, new(big.Int)), nil
	}
}

func getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amount *big.Int, add bool) (*big.Int, error) {
	if add {
		quotient := new(big.Int)
		if amount.Cmp(MaxUint160) <= 0 {
			quotient = quotient.Set(amount).Lsh(quotient, 96).Div(quotient, liquidity)
		} else {
			quotient = quotient.Set(amount).Mul(quotient, constants.Q96).Div(quotient, liquidity)
		}
		return quotient.Add(sqrtPX96, quotient), nil
	}

	quotient := MulDivRoundingUp(amount, constants.Q96, liquidity)
	if sqrtPX96.Cmp(quotient) <= 0 {
		return nil, ErrInvariant
	}
	return quotient.Sub(sqrtPX96, quotient), nil
}
