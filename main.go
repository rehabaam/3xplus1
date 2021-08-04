///////////////////////////////////////////////////////////////////////////////////////
//                                                                                   //
//                          Collatz conjecture (known as 3x+1)                       //
//                                                                                   //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"math"
	"time"
)

// Starts the fun!!
func main() {

	// Choose a number and then start the timer
	var num float64 = 3112412422142425353344343434424347435435435436346346443634627527527575230953665899508087094643634634431
	t1 := time.Now()

	// Print the initial number
	fmt.Printf("First number %f\n", num)

	for {

		// Handle even cases
		if isEven(num) {

			// Print the input number
			fmt.Printf("Is Even %f\n", num)

			// Divide by 2
			num = num / 2

			// Print the new number
			fmt.Printf("New half %f\n", num)
			continue
		}

		// If it reaches 1, then stop, no need to get stuck in a forever loop
		if num == 1 {

			// Print the time taken for reaching 1
			t2 := time.Now()
			fmt.Printf("Time taken to reach '1': %v", t2.Sub(t1))
			break
		}

		// Print the input number
		fmt.Printf("Is Odd %f\n", num)

		// Do 3x+1 magic
		num = 3*num + 1

		// Print the new number
		fmt.Printf("new Even %f\n", num)
	}

}

// Just returns even or odd
func isEven(x float64) bool {
	return math.Mod(x, 2) == 0
}
