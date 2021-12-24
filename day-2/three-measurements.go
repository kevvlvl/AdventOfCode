package day_2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Day2ThreeMeasurements() {

	fmt.Println("Day2_three_measurements() Start")

	inc := 0
	prev := 0

	var content, err = ioutil.ReadFile("day-2/puzzle-input.txt")

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

	fmt.Println("Day2_three_measurements() End - Increases = ", inc)
}
