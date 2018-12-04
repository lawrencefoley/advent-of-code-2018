package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var prnt = fmt.Println

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// Claim represents a rectangle of fabric
type Claim struct {
	leftDistance, topDistance, width, height int
}

func main() {
	prnt("Advent of Code - Day 3")

	// Read the whole file into memory
	var inputFilename = "./input.txt"
	fmt.Println("Reading input file: " + inputFilename)
	data, err := ioutil.ReadFile(inputFilename)
	checkError(err)

	textData := string(data)

	// Remove all "\r" chars and covert to string
	textData = strings.Replace(textData, "\r", "", -1)

	claims := strings.Split(textData, "\n")
	claimsSlice := make([]Claim, 0)

	var fabricMatrix [1000][1000]int

	for _, claim := range claims {
		claimValues := strings.Split(claim, " ")
		claimValuesText := strings.Replace(claimValues[2], ":", "", -1)
		distanceValues := strings.Split(claimValuesText, ",")

		sizeValuesText := claimValues[3]
		sizeValues := strings.Split(sizeValuesText, "x")

		leftDistance, _ := strconv.Atoi(distanceValues[0])
		topDistance, _ := strconv.Atoi(distanceValues[1])
		width, _ := strconv.Atoi(sizeValues[0])
		height, _ := strconv.Atoi(sizeValues[1])

		currentClaim := Claim{leftDistance: leftDistance, topDistance: topDistance,
			width: width, height: height}

		claimsSlice = append(claimsSlice, currentClaim)

	}

	for _, claim := range claimsSlice {
		for i := claim.topDistance; i < claim.height+claim.topDistance; i++ {
			for j := claim.leftDistance; j < claim.leftDistance+claim.width; j++ {
				fabricMatrix[i][j] = fabricMatrix[i][j] + 1
			}
		}
	}

	numberOfOverlapping := 0
	for row := 0; row < len(fabricMatrix); row++ {
		for col := 0; col < len(fabricMatrix); col++ {
			if fabricMatrix[row][col] >= 2 {
				numberOfOverlapping++
			}
		}
	}
	prnt("Answer: ", numberOfOverlapping)
}
