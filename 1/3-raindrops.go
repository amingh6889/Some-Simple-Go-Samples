/*
Introduction
Raindrops is a slightly more complex version of the FizzBuzz challenge, a classic interview question.

Instructions
Your task is to convert a number into its corresponding raindrop sounds.

If a given number:

is divisible by 3, add "Pling" to the result.
is divisible by 5, add "Plang" to the result.
is divisible by 7, add "Plong" to the result.
is not divisible by 3, 5, or 7, the result should be the number as a string.
Examples
28 is divisible by 7, but not 3 or 5, so the result would be "Plong".
30 is divisible by 3 and 5, but not 7, so the result would be "PlingPlang".
34 is not divisible by 3, 5, or 7, so the result would be "34".
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var InputString string
	fmt.Println("Enter your number:")
	fmt.Scan(&InputString)

	InputNumber, error := strconv.Atoi(InputString)
	if error != nil {
		fmt.Println("Enter number only!")
	} else {
		var DivisionNumber = [3]int{3, 5, 7}
		var IsDividable = false
		var Result string

		for _, value := range DivisionNumber {
			if InputNumber%value == 0 && value == 3 {
				Result += "Pling"
				IsDividable = true
			} else if InputNumber%value == 0 && value == 5 {
				Result += "Plang"
				IsDividable = true
			} else if InputNumber%value == 0 && value == 7 {
				Result += "Plong"
				IsDividable = true
			}
		}

		if IsDividable == false {
			Result = strconv.Itoa(InputNumber)
		}
		fmt.Println(Result)

	}

}

