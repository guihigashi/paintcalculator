package domain

import (
	"fmt"
	"math"

	"github.com/samber/lo"
)

type RealCans struct{}

func (rc RealCans) Cans() []float64 {
	return Latas()
}

type Canner interface {
	Cans() []float64
}

func CansNeeded(canner Canner) func(float64) (map[string]int, float64) {
	return func(paintNeeded float64) (map[string]int, float64) {
		availableCans := canner.Cans()

		availableCansDL := lo.Map(availableCans, func(c float64, _ int) int {
			return int(math.Ceil(10 * c))
		})
		paintNeededDL := int(math.Ceil(10 * paintNeeded))

		var selectedCansDL []int
		ok := false
		for !ok {
			selectedCansDL, ok = lo.TryOr(func() ([]int, error) {
				return change(availableCansDL, paintNeededDL), nil
			}, []int{})
			if !ok {
				paintNeededDL++
			}
		}

		// format
		finalTarget := float64(paintNeededDL) / 10
		cansNeeded := lo.Reduce(availableCansDL,
			func(acc map[string]int, cur int, _ int) map[string]int {
				key := fmt.Sprintf("%.1fL", float64(cur)/10)

				n := lo.Count(selectedCansDL, cur)
				if n > 0 {
					acc[key] = n
				}

				return acc
			}, map[string]int{})

		return cansNeeded, finalTarget
	}
}

// https://github.com/AndreaGhizzoni/coins-change/blob/master/change_dynamic.go
// This function takes a slice of coins cuts T and the amount of change N.
// Returns a slice of choices of T.
// Complexity: O( N*len(T) ) ~ O( N^2 )
func change(T []int, N int) []int {
	// coins choices. N+1 because we will use S[1_N]
	S := make([]int, N+1)
	// auxiliary slice. N+1 because we will use C[1_N]
	C := make([]int, N+1)
	C[0] = 0

	// solve the problem N bottom-up. m is the current change.
	for m := 1; m <= N; m++ {
		C[m] = math.MaxInt32 // set current change to max

		// for every coins cut, check if I can give back that cut
		for j := 0; j < len(T); j++ {
			if m >= T[j] && C[m-T[j]]+1 < C[m] {
				C[m] = C[m-T[j]] + 1
				// save that j+1 can be the change for the problem m
				S[m] = j + 1
			}
		}
	}

	// func printChange(T []int, N int) {
	result := []int{}
	for N > 0 {
		result = append(result, T[S[N]-1])
		N = N - T[S[N]-1]
	}

	return result
}
