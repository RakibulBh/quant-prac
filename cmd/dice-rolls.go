package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func (app *application) SimulateDiceRolls(count int64) {

	counts := map[int64]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	for range count {
		n, _ := rand.Int(rand.Reader, big.NewInt(6))
		counts[n.Int64()+1] += 1
	}

	fmt.Printf("Results after %d dice rolls:\n\n", count)
	for i := int64(1); i <= 6; i++ {
		probability := float64(counts[i]) / float64(count) * 100
		fmt.Printf("Number %d: Rolled %d times (%.2f%% probability)\n", i, counts[i], probability)
	}
	fmt.Printf("\nNote: Expected probability for each number is approximately 16.67%%\n")

}
