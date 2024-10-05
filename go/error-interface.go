package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %v\n", float64(e))
	// ↓のように変更すると、
	// ErrNegativeSqrt.Error()でfmt.Sprint(e)を呼び出すと、
	// fmt.Sprint()処理内で引数eのe.Error()を呼び出す。
	// このe.Error()はErrNegativeSqrt.Error()です。
	// そのため再びfmt.Sprint(e)が評価されます。
	// このようにErrNegativeSqrt.Error()の呼び出しが無限に続いてしまいます。
	// return fmt.Sprint(e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
