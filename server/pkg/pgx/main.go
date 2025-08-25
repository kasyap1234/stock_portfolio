package pgxhelpers

import (
	"fmt"
	"math"
	"math/big"
"errors"
	"github.com/jackc/pgx/v5/pgtype"
	"math/big"
	
)

// FloatToNumeric2Decimal converts a float64 to pgtype.Numeric with 2 decimal places
func FloatToNumeric2Decimal(f float64) (pgtype.Numeric, error) {
	if math.IsNaN(f) {
		return pgtype.Numeric{NaN: true, Valid: true}, nil
	}
	if math.IsInf(f, 1) {
		return pgtype.Numeric{InfinityModifier: pgtype.Infinity, Valid: true}, nil
	}
	if math.IsInf(f, -1) {
		return pgtype.Numeric{InfinityModifier: pgtype.NegativeInfinity, Valid: true}, nil
	}

	scale := 2
	factor := math.Pow10(scale)
	scaled := f * factor

	// Round to nearest integer to preserve 2 decimal places
	rounded := math.Round(scaled)

	// Check overflow
	if rounded > math.MaxInt64 || rounded < math.MinInt64 {
		return pgtype.Numeric{}, fmt.Errorf("value out of range for pgtype.Numeric")
	}

	return pgtype.Numeric{
		Int:              big.NewInt(int64(rounded)),
		Exp:              int32(-scale), // -2 for 2 decimal places
		NaN:              false,
		InfinityModifier: pgtype.Finite,
		Valid:            true,
	}, nil
}

func NumericToIntFast(num pgtype.Numeric) int64 {
	i, err := num.Int64Value()
	if err != nil {
		return 0
	}
	if !i.Valid {
		return 0
	}
	return i.Int64
}


// NumericToFloat converts pgtype.Numeric to float64
func NumericToFloat(num pgtype.Numeric) (float64, error) {
	if !num.Valid {
		return 0, errors.New("numeric is null")
	}

	// Use Int and Exp to calculate the float value
	bf := new(big.Float).SetInt(num.Int)
	exp := big.NewFloat(1)

	if num.Exp != 0 {
		exp.SetFloat64(pow10(int(num.Exp)))
	}

	bf.Mul(bf, exp)

	f, _ := bf.Float64()
	return f, nil
}

// pow10 handles negative exponents
func pow10(exp int) float64 {
	if exp == 0 {
		return 1
	}

	res := 1.0
	if exp > 0 {
		for i := 0; i < exp; i++ {
			res *= 10
		}
	} else {
		for i := 0; i < -exp; i++ {
			res /= 10
		}
	}
	return res
}
