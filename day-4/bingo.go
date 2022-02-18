package day_4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers []string
var bingoMatrices = make([]BingoMatrix, 0)

type BingoMatrix struct {
	numbers [][]int
	found   [][]int
}

func FindBingoMatrix() {

	fmt.Println("FindBingoMatrix() Start")

	loadData()
	fmt.Println("Matrices loaded: ", bingoMatrices)

	fmt.Println("FindBingoMatrix() End")
}

func loadData() {

	file, err := os.Open("day-4/puzzle-input.txt")
	if err != nil {
		fmt.Println("ERROR opening file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("ERROR closing the file ", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	// first line are the bingo numbers
	scanner.Scan()
	numbers = strings.Split(scanner.Text(), ",")
	fmt.Println(numbers)

	// next lines will be loaded into matrices

	for {

		scanned := scanner.Scan()
		if err := scanner.Err(); err != nil || !scanned {
			break
		}

		scanned = scanner.Scan()
		if err := scanner.Err(); err != nil || !scanned {
			break
		}

		var currentMatrix BingoMatrix

		currentLine := scanner.Text()
		firstRowMatrix := stringArrayToInt(currentLine)
		currentMatrix.numbers = append(currentMatrix.numbers, firstRowMatrix)

		scanner.Scan()
		currentLine = scanner.Text()
		secondRowMatrix := stringArrayToInt(currentLine)
		currentMatrix.numbers = append(currentMatrix.numbers, secondRowMatrix)

		scanner.Scan()
		currentLine = scanner.Text()
		thirdRowMatrix := stringArrayToInt(currentLine)
		currentMatrix.numbers = append(currentMatrix.numbers, thirdRowMatrix)

		scanner.Scan()
		currentLine = scanner.Text()
		fourthRowMatrix := stringArrayToInt(currentLine)
		currentMatrix.numbers = append(currentMatrix.numbers, fourthRowMatrix)

		scanner.Scan()
		currentLine = scanner.Text()
		fifthRowMatrix := stringArrayToInt(currentLine)
		currentMatrix.numbers = append(currentMatrix.numbers, fifthRowMatrix)

		bingoMatrices = append(bingoMatrices, currentMatrix)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR reading file")
	}
}

func stringArrayToInt(stringNumbers string) []int {

	stringNumbersArr := strings.Split(stringNumbers, " ")
	intArray := make([]int, len(stringNumbersArr))

	for i, v := range stringNumbersArr {
		intArray[i], _ = strconv.Atoi(v)
	}

	return intArray
}
