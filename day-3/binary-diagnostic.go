package day_3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func MeasurePowerConsumption() {

	fmt.Println("MeasurePowerConsumption() Start")

	content, err := ioutil.ReadFile("day-3/puzzle-input.txt")
	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
		os.Exit(1)
	}

	stringData := bytes.Split(content, []byte{'\n'})
	fileLength := len(stringData)
	bitsLength := len(stringData[0])

	var bitsMatrix [][]string

	// Read file data into String matrix of bit values
	for i := 0; i < fileLength; i++ {

		lineStr := string(stringData[i])
		bitsArray := make([]string, bitsLength)

		for i, v := range lineStr {
			bitsArray[i] = string(v)
		}

		bitsMatrix = append(bitsMatrix, bitsArray)
	}

	// Matrix loaded with data
	zeroBitsCount := make([]int, bitsLength)

	var gammaRate string
	var epsilonRate string

	// Calculate sum of zero bit values at each index (and indirectly count of one bit values)
	for i := 0; i < len(bitsMatrix); i++ {
		for j := 0; j < len(bitsMatrix[i]); j++ {
			if bitsMatrix[i][j] == "0" {
				zeroBitsCount[j]++
			}
		}
	}

	// Calculate gamma and epsilon rates
	for i := 0; i < len(zeroBitsCount); i++ {

		// is the count of zero bits greater (or not) than the remainder of 1 bits (total length - count of zero bits = 1 bits)
		if zeroBitsCount[i] >= (fileLength - zeroBitsCount[i]) {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}

	// Calculate power
	gammaRateDec, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		fmt.Println("  FATAL: - Failed to parse gammaRate value ", gammaRate)
		return
	}

	epsilonRateDec, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		fmt.Println("  FATAL: - Failed to parse epsilonRate value ", epsilonRate)
		return
	}

	powerConsumption := gammaRateDec * epsilonRateDec

	fmt.Println("Bits count")
	fmt.Println("  - Zero bits: ", zeroBitsCount)
	fmt.Println("  - Gamma rate: ", gammaRate, ", Dec: ", gammaRateDec)
	fmt.Println("  - Epsilon rate: ", epsilonRate, ", Dec: ", epsilonRateDec)

	fmt.Println("MeasurePowerConsumption() End - Power = ", powerConsumption)
}
