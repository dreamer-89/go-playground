package number_utils

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %g\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    } else {
        z := float64(x)
        for i := 0; i < 100; i++ {
            z -= (z*z - x) / (2 * z)
        }
        return z, nil
    }
}
