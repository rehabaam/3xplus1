///////////////////////////////////////////////////////////////////////////////////////
//                                                                                   //
//                          Collatz conjecture (known as 3x+1)                       //
//                                                                                   //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"math/big"
	"time"
)

// Break the conjecture
func main() {

	// Choose a number
	var num = new(big.Int)
	num.Exp(big.NewInt(300), big.NewInt(5), nil)

	// Starts the fun!!
	check_conjecture(num)
}

// Starts the fun!!
func check_conjecture(num *big.Int) {

	// Start the timer
	t1 := time.Now()

	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	three := big.NewInt(3)

	// Print the initial number
	fmt.Printf("First number %v\n", num)

	for {

		// Handle even cases
		if isEven(num, zero, two) {

			// Print the input number
			fmt.Printf("Is Even %v\n", num)

			// Divide by 2
			num = num.Div(num, two)

			// Print the new number
			fmt.Printf("New half %v\n", num)
			continue
		}

		// If it reaches 1, then stop, no need to get stuck in a forever loop
		if num.Cmp(one) == 0 {

			// Print the time taken for reaching 1
			t2 := time.Now()
			fmt.Printf("Time taken to reach '1': %v", t2.Sub(t1))
			break
		}

		// Print the input number
		fmt.Printf("Is Odd %v\n", num)

		// Do 3x+1 magic
		num = num.Add(num.Mul(num, three), one)

		// Print the new number
		fmt.Printf("new Even %v\n", num)
	}

}

// Just returns even or odd
func isEven(x, zero, two *big.Int) bool {
	mod := new(big.Int)
	return zero.Cmp(mod.Mod(x, two)) == 0
}
