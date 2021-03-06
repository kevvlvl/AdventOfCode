package day_3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var bytesValueZero = []byte("0")
var bytesValueOne = []byte("1")

func MeasurePowerConsumption() {

	fmt.Println("MeasurePowerConsumption() Start")

	content, err := ioutil.ReadFile("day-3/puzzle-input.txt")
	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
		os.Exit(1)
	}

	stringData := bytes.Split(content, []byte{'\n'})
	bitsMatrix, matrixLength, bitsLength := getBitsMatrix(stringData)

	// Matrix loaded with data
	zeroBitsCount := make([]int, bitsLength)

	var gammaRate string
	var epsilonRate string

	// Calculate sum of zero bit values at each index (and indirectly count of one bit values)
	for i := 0; i < len(bitsMatrix); i++ {
		for j := 0; j < bitsLength; j++ {
			if bitsMatrix[i][j] == "0" {
				zeroBitsCount[j]++
			}
		}
	}

	// Calculate gamma and epsilon rates
	for i := 0; i < len(zeroBitsCount); i++ {

		// is the count of zero bits greater (or not) than the remainder of 1 bits (total length - count of zero bits = 1 bits)
		if zeroBitsCount[i] >= (matrixLength - zeroBitsCount[i]) {
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

func MeasureLifeSupportingRate() {

	fmt.Println("MeasureLifeSupportingRate() Start")

	content, err := ioutil.ReadFile("day-3/puzzle-input.txt")
	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
		os.Exit(1)
	}

	stringData := bytes.Split(content, []byte{'\n'})
	bitsSlice, _, _ := getBitsSlice(stringData)

	// Slice loaded with data
	oxygenGeneratorRating := findRemainingBits(bitsSlice, 0, true)
	co2ScrubberRating := findRemainingBits(bitsSlice, 0, false)

	oxygenGeneratorRatingDec, err := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	if err != nil {
		fmt.Println("  FATAL: - Failed to parse oxygenGeneratorRating value ", oxygenGeneratorRating)
		return
	}

	co2ScrubberRatingDec, err := strconv.ParseInt(co2ScrubberRating, 2, 64)
	if err != nil {
		fmt.Println("  FATAL: - Failed to parse co2ScrubberRating value ", co2ScrubberRating)
		return
	}

	fmt.Println("Oxygen (based on common bit): ", oxygenGeneratorRating, " - Dec = ", oxygenGeneratorRatingDec)
	fmt.Println("CO2 (based on least common bit): ", co2ScrubberRating, " - Dec = ", co2ScrubberRatingDec)

	lifeSupportRating := oxygenGeneratorRatingDec * co2ScrubberRatingDec

	fmt.Println("MeasureLifeSupportingRate() End - Life Support Rating = ", lifeSupportRating)
}

func getBitsMatrix(stringData [][]byte) ([][]string, int, int) {

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

	return bitsMatrix, fileLength, bitsLength
}

func getBitsSlice(stringData [][]byte) ([]string, int, int) {

	fileLength := len(stringData)
	bitsLength := len(stringData[0])

	var bitsMatrix []string

	// Read file data into String matrix of bit values
	for i := 0; i < fileLength; i++ {

		lineStr := string(stringData[i])
		bitsMatrix = append(bitsMatrix, lineStr)
	}

	return bitsMatrix, fileLength, bitsLength
}

func findRemainingBits(bitsSlice []string, pos int, isCommon bool) string {

	// Find common bit at current index

	var commonBit byte
	var leastCommonBit byte
	oneBitCount := 0
	sliceLength := len(bitsSlice)

	for i := 0; i < sliceLength; i++ {
		if bitsSlice[i][pos] == bytesValueOne[0] {
			oneBitCount++
		}
	}

	if oneBitCount >= (sliceLength - oneBitCount) {
		commonBit = bytesValueOne[0]
		leastCommonBit = bytesValueZero[0]
	} else {
		commonBit = bytesValueZero[0]
		leastCommonBit = bytesValueOne[0]
	}

	// Compile list of matching bits for common bit

	var commonList []string

	if isCommon {
		commonList = compileListMatchingBits(bitsSlice, pos, commonBit)
	} else {
		commonList = compileListMatchingBits(bitsSlice, pos, leastCommonBit)
	}

	if len(commonList) == 1 {
		return commonList[0]
	} else {
		return findRemainingBits(commonList, pos+1, isCommon)
	}
}

func compileListMatchingBits(bitsSlice []string, pos int, matchingBit byte) []string {

	var list []string

	for i := 0; i < len(bitsSlice); i++ {
		if bitsSlice[i][pos] == matchingBit {
			list = append(list, bitsSlice[i])
		}
	}

	return list
}
