package day_2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func MeasureDive() {

	fmt.Println("MeasureDive() Start")

	var content, err = ioutil.ReadFile("day-2/puzzle-input.txt")
	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
		os.Exit(1)
	}

	// horizontal position and depth
	var positions [1][2]int

	var stringData = bytes.Split(content, []byte{'\n'})
	for i := 0; i < len(stringData); i++ {

		re := regexp.MustCompile(`^(forward|up|down) ([0-9]+)$`)
		matches := re.FindStringSubmatch(string(stringData[i]))

		if len(matches) == 3 {
			cmd := matches[1]
			value, _ := strconv.Atoi(matches[2])

			switch cmd {
			case "forward":
				positions[0][0] += value
			case "up":
				positions[0][1] -= value
			case "down":
				positions[0][1] += value
			}
		}
	}

	fmt.Println("MeasureDive() End - Horizontal pos = ", positions[0][0], " Depth = ", positions[0][1], " - Multiplied = ", positions[0][0]*positions[0][1])
}

func MeasureDiveWithAim() {

	fmt.Println("MeasureDiveWithAim() Start")

	var content, err = ioutil.ReadFile("day-2/puzzle-input.txt")
	if err != nil {
		fmt.Println("  FATAL: - Failed to read file")
		os.Exit(1)
	}

	// horizontal position, depth, and aim
	var positions [1][3]int

	var stringData = bytes.Split(content, []byte{'\n'})
	for i := 0; i < len(stringData); i++ {

		re := regexp.MustCompile(`^(forward|up|down) ([0-9]+)$`)
		matches := re.FindStringSubmatch(string(stringData[i]))

		if len(matches) == 3 {
			cmd := matches[1]
			value, _ := strconv.Atoi(matches[2])

			switch cmd {
			case "forward":
				positions[0][0] += value
				positions[0][1] += positions[0][2] * value
			case "up":
				positions[0][2] -= value
			case "down":
				positions[0][2] += value
			}
		}
	}

	fmt.Println("MeasureDiveWithAim() End - Horizontal pos = ", positions[0][0], " Depth = ", positions[0][1], " - Multiplied = ", positions[0][0]*positions[0][1])
}
