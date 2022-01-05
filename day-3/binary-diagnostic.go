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
		var bitsArray []string

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

/**
Find if the bit is set to searchBit at the given (x) position
*/
func isBitSet(bits string, x int, searchBit byte) bool {
	return bits[x] == searchBit
}

/**
Return the common and least common bit for the given index for the list of bits values

1 is returned if the count matches the 0 (exact 50/50 split)
*/
func getCommonBit(bitsSlice []string, x int) (byte, byte) {

	oneBitCount := 0
	sliceLength := len(bitsSlice)

	for i := 0; i < sliceLength; i++ {
		if bitsSlice[i][x] == bytesValueOne[0] {
			oneBitCount++
		}
	}

	if oneBitCount >= (sliceLength - oneBitCount) {
		return bytesValueOne[0], bytesValueZero[0]
	} else {
		return bytesValueZero[0], bytesValueOne[0]
	}
}

func findRemainingBits(bitsSlice []string, pos int, isCommon bool) string {

	bitsLength := len(bitsSlice[0])
	commonBit, leastCommonBit := getCommonBit(bitsSlice, pos)
	var commonList []string

	if isCommon {
		commonList = compileListMatchingBits(bitsSlice, pos, commonBit)
	} else {
		commonList = compileListMatchingBits(bitsSlice, pos, leastCommonBit)
	}

	if len(commonList) == 1 && pos <= (bitsLength-1) {
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
