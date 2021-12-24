package day_1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Day1Measurements() {

	fmt.Println("Day1Measurements() Start")

	inc := 0

	var content, err = ioutil.ReadFile("day-1/puzzle-input.txt")

	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
	}

	var stringData = bytes.Split(content, []byte{'\n'})
	for i := 1; i < len(stringData); i++ {

		intCurrent, _ := strconv.Atoi(string(stringData[i]))
		intPrev, _ := strconv.Atoi(string(stringData[i-1]))

		if intCurrent > intPrev {
			inc++
		}
	}

	fmt.Println("Day1Measurements() End - Increases = ", inc)
}

func Day1ThreeMeasurements() {

	fmt.Println("Day1ThreeMeasurements() Start")

	inc := 0
	prev := 0

	var content, err = ioutil.ReadFile("day-1/puzzle-input.txt")

	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
	}

	var strData = bytes.Split(content, []byte{'\n'})
	var dataLength = len(strData)

	for i := 0; i < dataLength; i++ {

		if i+2 <= dataLength-1 {

			first, _ := strconv.Atoi(string(strData[i]))
			second, _ := strconv.Atoi(string(strData[i+1]))
			third, _ := strconv.Atoi(string(strData[i+2]))

			sum := first + second + third

			if prev > 0 && sum > prev {
				inc++
			}

			prev = sum
		}
	}

	fmt.Println("Day1ThreeMeasurements() End - Increases = ", inc)
}
