package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var prnt = fmt.Println

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// Returns 2 and 3 occurances existing in a map
func getOccurances(m map[string]int) (bool, bool) {
	hasTwo := false
	hasThree := false
	for _, val := range m {
		if val == 2 {
			hasTwo = true
		} else if val == 3 {
			hasThree = true
		}
	}
	return hasTwo, hasThree
}

func main() {
	fmt.Println("Advent of Code - Day 2")

	// Read the whole file into memory
	var inputFilename = "./input.txt"
	fmt.Println("Reading input file: " + inputFilename)
	data, err := ioutil.ReadFile(inputFilename)
	checkError(err)

	textData := string(data)

	// Remove all "\r" chars and covert to string
	textData = strings.Replace(textData, "\r", "", -1)

	splitText := strings.Split(textData, "\n")

	// Declare a map of string->int
	charMap := make(map[string]int)
	numberOfTwos, numberOfThrees := 0, 0

	// Iterate through each ID
	for i := 0; i < len(splitText); i++ {
		// Iterate through each char in the ID
		charMap = make(map[string]int)
		for _, char := range splitText[i] {

			_, exists := charMap[string(char)]

			if exists {
				currentVal := charMap[string(char)]
				currentVal++
				charMap[string(char)] = currentVal
			} else {
				// Doesn't exist
				charMap[string(char)] = 1
			}
		}

		hasTwo, hasThree := getOccurances(charMap)
		if hasTwo {
			numberOfTwos++
		}
		if hasThree {
			numberOfThrees++
		}
		// prnt(splitText[i], " -- ", hasTwo, ", ", hasThree)
	}

	checksum := numberOfTwos * numberOfThrees

	prnt("Answer: ", checksum)

}
