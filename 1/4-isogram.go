/*
Instructions
Determine if a word or phrase is an isogram.

An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

Examples of isograms:

lumberjacks
background
downstream
six-year-old
The word isograms, however, is not an isogram, because the s repeats.
*/

package main

import "fmt"

func main() {
	fmt.Println("Enter your string:")
	var InputString string
	fmt.Scan(&InputString)
	var InputStringLength = len(InputString)
	var IsIsogram = true

	for i := 0; i < InputStringLength-1; i++ {
		for j := i + 1; j < InputStringLength; j++ {
			if InputString[i] == InputString[j] && InputString[i] != '-' && InputString[i] != ' ' {
				IsIsogram = false
				break
			}
		}

		if IsIsogram == false {
			break
		}
	}

	if IsIsogram == false {
		fmt.Println("It is not Isogram")
	} else {
		fmt.Println("It is Isogram")
	}

}

