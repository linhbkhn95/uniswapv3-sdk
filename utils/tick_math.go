package utils

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
)

const (
	MinTick = -887272  // The minimum tick that can be used on any pool.
	MaxTick = -MinTick // The maximum tick that can be used on any pool.
)

var (
	Q32             = big.NewInt(1 << 32)
	MinSqrtRatio    = big.NewInt(4295128739)                                                          // The sqrt ratio corresponding to the minimum tick that could be used on any pool.
	MaxSqrtRatio, _ = new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10) // The sqrt ratio corresponding to the maximum tick that could be used on any pool.
)

var (
	ErrInvalidTick      = errors.New("invalid tick")
	ErrInvalidSqrtRatio = errors.New("invalid sqrt ratio")
)

func mulShiftX(inout *big.Int, mulBy *big.Int) {
	inout.Mul(inout, mulBy).Rsh(inout, 128)
}

var (
	sqrtConst1, _  = new(big.Int).SetString("fffcb933bd6fad37aa2d162d1a594001", 16)
	sqrtConst2, _  = new(big.Int).SetString("100000000000000000000000000000000", 16)
	sqrtConst3, _  = new(big.Int).SetString("fff97272373d413259a46990580e213a", 16)
	sqrtConst4, _  = new(big.Int).SetString("fff2e50f5f656932ef12357cf3c7fdcc", 16)
	sqrtConst5, _  = new(big.Int).SetString("ffe5caca7e10e4e61c3624eaa0941cd0", 16)
	sqrtConst6, _  = new(big.Int).SetString("ffcb9843d60f6159c9db58835c926644", 16)
	sqrtConst7, _  = new(big.Int).SetString("ff973b41fa98c081472e6896dfb254c0", 16)
	sqrtConst8, _  = new(big.Int).SetString("ff2ea16466c96a3843ec78b326b52861", 16)
	sqrtConst9, _  = new(big.Int).SetString("fe5dee046a99a2a811c461f1969c3053", 16)
	sqrtConst10, _ = new(big.Int).SetString("fcbe86c7900a88aedcffc83b479aa3a4", 16)
	sqrtConst11, _ = new(big.Int).SetString("f987a7253ac413176f2b074cf7815e54", 16)
	sqrtConst12, _ = new(big.Int).SetString("f3392b0822b70005940c7a398e4b70f3", 16)
	sqrtConst13, _ = new(big.Int).SetString("e7159475a2c29b7443b29c7fa6e889d9", 16)
	sqrtConst14, _ = new(big.Int).SetString("d097f3bdfd2022b8845ad8f792aa5825", 16)
	sqrtConst15, _ = new(big.Int).SetString("a9f746462d870fdf8a65dc1f90e061e5", 16)
	sqrtConst16, _ = new(big.Int).SetString("70d869a156d2a1b890bb3df62baf32f7", 16)
	sqrtConst17, _ = new(big.Int).SetString("31be135f97d08fd981231505542fcfa6", 16)
	sqrtConst18, _ = new(big.Int).SetString("9aa508b5b7a84e1c677de54f3e99bc9", 16)
	sqrtConst19, _ = new(big.Int).SetString("5d6af8dedb81196699c329225ee604", 16)
	sqrtConst20, _ = new(big.Int).SetString("2216e584f5fa1ea926041bedfe98", 16)
	sqrtConst21, _ = new(big.Int).SetString("48a170391f7dc42444e8fa2", 16)
)

/**
 * Returns the sqrt ratio as a Q64.96 for the given tick. The sqrt ratio is computed as sqrt(1.0001)^tick
 * @param tick the tick for which to compute the sqrt ratio
 */
func GetSqrtRatioAtTick(tick int) (*big.Int, error) {
	if tick < MinTick || tick > MaxTick {
		return nil, ErrInvalidTick
	}
	absTick := tick
	if tick < 0 {
		absTick = -tick
	}
	ratio := new(big.Int)
	if absTick&0x1 != 0 {
		ratio.Set(sqrtConst1)
	} else {
		ratio.Set(sqrtConst2)
	}
	if (absTick & 0x2) != 0 {
		mulShiftX(ratio, sqrtConst3)
	}
	if (absTick & 0x4) != 0 {
		mulShiftX(ratio, sqrtConst4)
	}
	if (absTick & 0x8) != 0 {
		mulShiftX(ratio, sqrtConst5)
	}
	if (absTick & 0x10) != 0 {
		mulShiftX(ratio, sqrtConst6)
	}
	if (absTick & 0x20) != 0 {
		mulShiftX(ratio, sqrtConst7)
	}
	if (absTick & 0x40) != 0 {
		mulShiftX(ratio, sqrtConst8)
	}
	if (absTick & 0x80) != 0 {
		mulShiftX(ratio, sqrtConst9)
	}
	if (absTick & 0x100) != 0 {
		mulShiftX(ratio, sqrtConst10)
	}
	if (absTick & 0x200) != 0 {
		mulShiftX(ratio, sqrtConst11)
	}
	if (absTick & 0x400) != 0 {
		mulShiftX(ratio, sqrtConst12)
	}
	if (absTick & 0x800) != 0 {
		mulShiftX(ratio, sqrtConst13)
	}
	if (absTick & 0x1000) != 0 {
		mulShiftX(ratio, sqrtConst14)
	}
	if (absTick & 0x2000) != 0 {
		mulShiftX(ratio, sqrtConst15)
	}
	if (absTick & 0x4000) != 0 {
		mulShiftX(ratio, sqrtConst16)
	}
	if (absTick & 0x8000) != 0 {
		mulShiftX(ratio, sqrtConst17)
	}
	if (absTick & 0x10000) != 0 {
		mulShiftX(ratio, sqrtConst18)
	}
	if (absTick & 0x20000) != 0 {
		mulShiftX(ratio, sqrtConst19)
	}
	if (absTick & 0x40000) != 0 {
		mulShiftX(ratio, sqrtConst20)
	}
	if (absTick & 0x80000) != 0 {
		mulShiftX(ratio, sqrtConst21)
	}
	if tick > 0 {
		ratio.Div(entities.MaxUint256, ratio)
	}

	// back to Q96
	if new(big.Int).Rem(ratio, Q32).Cmp(constants.Zero) > 0 {
		return ratio.Div(ratio, Q32).Add(ratio, constants.One), nil
	} else {
		return ratio.Div(ratio, Q32), nil
	}
}

var (
	magicSqrt10001, _ = new(big.Int).SetString("255738958999603826347141", 10)
	magicTickLow, _   = new(big.Int).SetString("3402992956809132418596140100660247210", 10)
	magicTickHigh, _  = new(big.Int).SetString("291339464771989622907027621153398088495", 10)
)

/**
 * Returns the tick corresponding to a given sqrt ratio, s.t. #getSqrtRatioAtTick(tick) <= sqrtRatioX96
 * and #getSqrtRatioAtTick(tick + 1) > sqrtRatioX96
 * @param sqrtRatioX96 the sqrt ratio as a Q64.96 for which to compute the tick
 */
func GetTickAtSqrtRatio(sqrtRatioX96 *big.Int) (int, error) {
	if sqrtRatioX96.Cmp(MinSqrtRatio) < 0 || sqrtRatioX96.Cmp(MaxSqrtRatio) >= 0 {
		return 0, ErrInvalidSqrtRatio
	}
	sqrtRatioX128 := new(big.Int).Lsh(sqrtRatioX96, 32)
	msb, err := MostSignificantBit(sqrtRatioX128)
	if err != nil {
		return 0, err
	}
	var r *big.Int
	if msb >= 128 {
		r = new(big.Int).Rsh(sqrtRatioX128, uint(msb-127))
	} else {
		r = new(big.Int).Lsh(sqrtRatioX128, uint(127-msb))
	}

	log2 := big.NewInt(msb - 128)
	log2.Lsh(log2, 64)

	f := new(big.Int)
	for i := 0; i < 14; i++ {
		r.Mul(r, r).Rsh(r, 127)
		f.Rsh(r, 128)
		rshf := f.Int64()
		log2.Or(log2, f.Lsh(f, uint(63-i)))
		r.Rsh(r, uint(rshf))
	}

	logSqrt10001 := log2.Mul(log2, magicSqrt10001)
	bigTick := new(big.Int)
	tickLow := bigTick.Sub(logSqrt10001, magicTickLow).Rsh(bigTick, 128).Int64()
	tickHigh := bigTick.Add(logSqrt10001, magicTickHigh).Rsh(bigTick, 128).Int64()

	if tickLow == tickHigh {
		return int(tickLow), nil
	}

	sqrtRatio, err := GetSqrtRatioAtTick(int(tickHigh))
	if err != nil {
		return 0, err
	}
	if sqrtRatio.Cmp(sqrtRatioX96) <= 0 {
		return int(tickHigh), nil
	} else {
		return int(tickLow), nil
	}
}
