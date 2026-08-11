/*
Instructions
Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.

The square of the sum of the first ten natural numbers is (1 + 2 + ... + 10)² = 55² = 3025.

The sum of the squares of the first ten natural numbers is 1² + 2² + ... + 10² = 385.

Hence the difference between the square of the sum of the first ten natural numbers and the sum of the squares of the first ten natural numbers is 3025 - 385 = 2640.

You are not expected to discover an efficient solution to this yourself from first principles; research is allowed, indeed, encouraged. Finding the best algorithm for the problem is a key skill in software engineering.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Enter natural number:")

	var InputString string
	fmt.Scan(&InputString)

	InputNumber, error := strconv.Atoi(InputString)

	if error != nil {
		fmt.Println("Input numbers not string!")
		return
	}
	if InputNumber < 0 {
		fmt.Println("Input a positive number!")
		return
	}

	var SumOfSquare int = 0
	var SquareOfSum int = 0
	for i := 1; i <= InputNumber; i++ {
		SumOfSquare += i * i
		SquareOfSum += i
	}

	SquareOfSum *= SquareOfSum
	Result := SquareOfSum - SumOfSquare

	fmt.Println(Result)

}

