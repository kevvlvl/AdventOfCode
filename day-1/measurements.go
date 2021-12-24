package day_1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Day1Measurements() {

	fmt.Println("Day1_measurements() Start")

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

	fmt.Println("Day1_measurements() End - Increases = ", inc)
}
