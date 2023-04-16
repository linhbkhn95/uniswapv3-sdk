package utils

import (
	"math/big"

	"github.com/daoleno/uniswapv3-sdk/constants"
)

func MulDivRoundingUp(a, b, denominator *big.Int) *big.Int {
	product := new(big.Int).Mul(a, b)
	if new(big.Int).Rem(product, denominator).Cmp(constants.Zero) != 0 {
		return product.Div(product, denominator).Add(product, constants.One)
	} else {
		return product.Div(product, denominator)
	}
}

// MulDivRoundingUpModified avoid malloc by reusing inputs a and b, a will be modified and used as result
func MulDivRoundingUpModified(a, b, denominator *big.Int) *big.Int {
	a.Mul(a, b)
	if b.Rem(a, denominator).Cmp(constants.Zero) != 0 {
		return a.Div(a, denominator).Add(a, constants.One)
	} else {
		return a.Div(a, denominator)
	}
}

// MulDivRoundingUpModified avoid malloc by reusing inputs a and helper, current value of helper doesn't matter
func DivRoundingUpModified(a, denominator, helper *big.Int) *big.Int {
	if helper.Rem(a, denominator).Cmp(constants.Zero) != 0 {
		return a.Div(a, denominator).Add(a, constants.One)
	} else {
		return a.Div(a, denominator)
	}
}
