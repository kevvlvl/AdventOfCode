package day_1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Day1_measurements() {

	fmt.Println("day1_measurements() Start")

	var content, err = ioutil.ReadFile("day-1/puzzle-input.txt")

	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
	}

	var stringData = bytes.Split(content, []byte{'\n'})

	inc := 0
	dec := 0

	for i := 1; i < len(stringData); i++ {

		intCurrent, _ := strconv.Atoi(string(stringData[i]))
		intPrev, _ := strconv.Atoi(string(stringData[i-1]))

		if intCurrent > intPrev {
			inc++
		} else {
			dec++
		}
	}

	fmt.Println("day1_measurements() End - Increases = ", inc, ", Decreases = ", dec)
}
