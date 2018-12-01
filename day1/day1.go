package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Keep track of the fequency value
	var frequency = 0

	fmt.Println("Advent of Code - Day 1")

	// Read the whole file into memory
	var inputFilename = "./input.txt"
	var outputFilename = "./output.txt"
	fmt.Println("Reading input file: " + inputFilename)
	data, err := ioutil.ReadFile(inputFilename)
	checkError(err)

	// Remove all "\r" chars and covert to string
	text := strings.Replace(string(data), "\r", "", -1)

	splitText := strings.Split(string(text), "\n")

	for i := 0; i < len(splitText); i++ {
		// Convert the string to an int
		parsedInt, err := strconv.Atoi(splitText[i])
		checkError(err)
		frequency += parsedInt
	}

	err = ioutil.WriteFile(outputFilename, []byte(strconv.Itoa(frequency)), 0777)

	fmt.Println("Frequency: " + strconv.Itoa(frequency))
}
