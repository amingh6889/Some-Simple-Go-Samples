/*
Introduction
Your body is made up of cells that contain DNA. Those cells regularly wear out and need replacing, which they achieve by dividing into daughter cells. In fact, the average human body experiences about 10 quadrillion cell divisions in a lifetime!

When cells divide, their DNA replicates too. Sometimes during this process mistakes happen and single pieces of DNA get encoded with the incorrect information. If we compare two strands of DNA and count the differences between them, we can see how many mistakes occurred. This is known as the "Hamming distance".

The Hamming distance is useful in many areas of science, not just biology, so it's a nice phrase to be familiar with :)

Instructions
Calculate the Hamming distance between two DNA strands.

We read DNA using the letters C, A, G and T. Two strands might look like this:

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^
They have 7 differences, and therefore the Hamming distance is 7.

Implementation notes
The Hamming distance is only defined for sequences of equal length, so an attempt to calculate it between sequences of different lengths should not work.
*/

package main

import "fmt"

func main() {

	var FirstString string
	var SecondString string

	fmt.Println("Enter first string:")
	fmt.Scan(&FirstString)

	fmt.Println("Enter second string:")
	fmt.Scan(&SecondString)

	FirstStringLength := len(FirstString)
	SecondStringLength := len(SecondString)

	var OutPut []int

	if FirstStringLength != SecondStringLength {
		fmt.Println("Your entered strings are not equal in length!")
	} else {
		for i := 0; i < FirstStringLength; i++ {
			if FirstString[i] != SecondString[i] {
				OutPut = append(OutPut, i+1)
			}

		}
	}

	fmt.Println(OutPut)

}

