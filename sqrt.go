// Package floats provides the implementation of a few additional operations for the
// standard library big.Float type.
package floats

import (
	"math"
	"math/big"
)

// Sqrt returns a big.Float representation of the square root of z. Precision is
// the same as the one of the argument. The function panics if z is negative, returns ±0
// when z = ±0, and +Inf when z = +Inf.
func Sqrt(z *big.Float) *big.Float {

	// panic on negative z
	if z.Sign() == -1 {
		panic("Sqrt: argument is negative")
	}

	// Sqrt(±0) = ±0
	if z.Sign() == 0 {
		return big.NewFloat(float64(z.Sign()))
	}

	// Sqrt(+Inf) = +Inf
	if z.IsInf() {
		return big.NewFloat(math.Inf(+1))
	}

	prec := z.Prec() + 64

	one := new(big.Float).SetPrec(prec).SetInt64(1)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// Halve exponent for the initial estimate
	x := new(big.Float).SetPrec(prec).SetMantExp(one, z.MantExp(nil)/2)

	t := new(big.Float)

	// Newton:
	//     x_{n+1} = 1/2 * ( x_n + (S / x_n) )

	// we need at least log_2(prec) iterations
	steps := int(math.Log2(float64(prec)))

	for i := 0; i < steps; i++ {
		t.Quo(z, x)    // t = S / x_n
		t.Add(t, x)    // t = x_n + (S / x_n)
		x.Mul(t, half) // x = t / 2
	}

	return x.SetPrec(z.Prec())

}
