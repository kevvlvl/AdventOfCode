package day_4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numbers []int
var lastNumberCalled int
var bingoMatrices = make([]BingoMatrix, 0)
var wonBoard BingoMatrix

type BingoMatrix struct {
	numbers [][]int
	found   [5][5]int
}

func FindBingoMatrix() {

	fmt.Println("FindBingoMatrix() Start")

	loadData()
	fmt.Println("Matrices loaded: ", bingoMatrices)

	findWinningMatrix()
	fmt.Println("The first board that won BINGO (horizontal or vertical) = ", wonBoard, " - Last number called = ", lastNumberCalled)

	sumOfUncalledNumbers := calculateScore()
	fmt.Println("The sum of all uncalled numbers = ", sumOfUncalledNumbers)

	puzzleAnswer := sumOfUncalledNumbers * lastNumberCalled
	fmt.Println("The Puzzle answer (Final score) = ", puzzleAnswer)

	fmt.Println("FindBingoMatrix() End")
}

func calculateScore() int {

	var sumOfUncalledNumbers int
	for i := 0; i < len(wonBoard.numbers); i++ {
		for j := 0; j < len(wonBoard.numbers[i]); j++ {

			if wonBoard.found[i][j] == 0 {
				sumOfUncalledNumbers += wonBoard.numbers[i][j]
			}
		}
	}

	return sumOfUncalledNumbers
}

func verifyAnyWonBoard() bool {

	for m := 0; m < len(bingoMatrices); m++ {

		for i := 0; i < len(bingoMatrices[m].found); i++ {

			horizontalSum := bingoMatrices[m].found[i][0] + bingoMatrices[m].found[i][1] + bingoMatrices[m].found[i][2] + bingoMatrices[m].found[i][3] + bingoMatrices[m].found[i][4]
			if horizontalSum == 5 {
				fmt.Println("Found a winning board!")
				wonBoard = bingoMatrices[m]
				return true
			}

			verticalSum := bingoMatrices[m].found[0][i] + bingoMatrices[m].found[1][i] + bingoMatrices[m].found[2][i] + bingoMatrices[m].found[3][i] + bingoMatrices[m].found[4][i]
			if verticalSum == 5 {
				fmt.Println("Found a winning board!")
				wonBoard = bingoMatrices[m]
				return true
			}
		}
	}

	return false
}

func findWinningMatrix() {

	for _, v := range numbers {

		fmt.Println("Current Number ", v)

		for m := 0; m < len(bingoMatrices); m++ {

			for i := 0; i < len(bingoMatrices[m].numbers); i++ {
				for j := 0; j < len(bingoMatrices[m].numbers[i]); j++ {
					if bingoMatrices[m].numbers[i][j] == v {
						bingoMatrices[m].found[i][j] = 1

						if verifyAnyWonBoard() {
							lastNumberCalled = v
							return
						}
					}
				}
			}
		}
	}
}

func loadData() {

	fmt.Println("loadData() Start")

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
	numbers = stringArrayToInt(scanner.Text(), ",")
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

		firstRowMatrix := stringArrayToIntWithWhitespace(scanner.Text())
		currentMatrix.numbers = append(currentMatrix.numbers, firstRowMatrix)

		scanner.Scan()
		secondRowMatrix := stringArrayToIntWithWhitespace(scanner.Text())
		currentMatrix.numbers = append(currentMatrix.numbers, secondRowMatrix)

		scanner.Scan()
		thirdRowMatrix := stringArrayToIntWithWhitespace(scanner.Text())
		currentMatrix.numbers = append(currentMatrix.numbers, thirdRowMatrix)

		scanner.Scan()
		fourthRowMatrix := stringArrayToIntWithWhitespace(scanner.Text())
		currentMatrix.numbers = append(currentMatrix.numbers, fourthRowMatrix)

		scanner.Scan()
		fifthRowMatrix := stringArrayToIntWithWhitespace(scanner.Text())
		currentMatrix.numbers = append(currentMatrix.numbers, fifthRowMatrix)

		bingoMatrices = append(bingoMatrices, currentMatrix)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR reading file")
	}

	fmt.Println("loadData() End")
}

func stringArrayToInt(stringNumbers string, separator string) []int {

	stringNumbersArr := strings.Split(stringNumbers, separator)
	intArray := make([]int, len(stringNumbersArr))

	for i, v := range stringNumbersArr {
		intArray[i], _ = strconv.Atoi(v)
	}

	return intArray
}

func stringArrayToIntWithWhitespace(stringNumbers string) []int {
	stringNumbersArr := regexp.MustCompile("[\\s]+").Split(strings.TrimSpace(stringNumbers), -1)
	intArray := make([]int, len(stringNumbersArr))

	for i, v := range stringNumbersArr {
		intArray[i], _ = strconv.Atoi(v)
	}

	return intArray
}
