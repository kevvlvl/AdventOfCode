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

var bingoMatrices = make([]BingoMatrix, 0)
var firstBoardWon BingoMatrix
var lastBoardWon BingoMatrix
var wonBoardLastNumberCalled int

type BingoMatrix struct {
	numbers [][]int
	found   [5][5]int
	won     bool
}

func FindBingoMatrix() {

	fmt.Println("FindBingoMatrix() Start")

	loadData()
	fmt.Println("Matrices loaded: ", bingoMatrices)

	findWinningMatrix()
	fmt.Println("The first board that won BINGO (horizontal or vertical) = ", firstBoardWon, " - Last number called = ", wonBoardLastNumberCalled)

	sumOfUncalledNumbers := calculateScore(firstBoardWon)
	fmt.Println("The sum of all uncalled numbers = ", sumOfUncalledNumbers)

	puzzleAnswer := sumOfUncalledNumbers * wonBoardLastNumberCalled
	fmt.Println("The Puzzle answer (Final score) = ", puzzleAnswer)

	fmt.Println("FindBingoMatrix() End")
}

func FindLastWinningBingoMatrix() {

	fmt.Println("FindLastWinningBingoMatrix() Start")

	loadData()
	fmt.Println("Matrices loaded: ", bingoMatrices)

	findLastWinningMatrix()
	fmt.Println("The last board that won BINGO (horizontal or vertical) = ", lastBoardWon, " - Last number called = ", wonBoardLastNumberCalled)

	sumOfUncalledNumbers := calculateScore(lastBoardWon)
	fmt.Println("The sum of all uncalled numbers = ", sumOfUncalledNumbers)

	puzzleAnswer := sumOfUncalledNumbers * wonBoardLastNumberCalled
	fmt.Println("The Puzzle answer (Final score) = ", puzzleAnswer)

	fmt.Println("FindLastWinningBingoMatrix() End")
}

func calculateScore(board BingoMatrix) int {

	var sumOfUncalledNumbers int
	for i := 0; i < len(board.numbers); i++ {
		for j := 0; j < len(board.numbers[i]); j++ {

			if board.found[i][j] == 0 {
				sumOfUncalledNumbers += board.numbers[i][j]
			}
		}
	}

	return sumOfUncalledNumbers
}

func verifyBoard(board BingoMatrix) bool {

	for i := 0; i < len(board.found); i++ {

		horizontalSum := board.found[i][0] + board.found[i][1] + board.found[i][2] + board.found[i][3] + board.found[i][4]
		verticalSum := board.found[0][i] + board.found[1][i] + board.found[2][i] + board.found[3][i] + board.found[4][i]

		if horizontalSum == 5 || verticalSum == 5 {
			fmt.Println("This board ", board, " it's a bingo!")
			return true
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

						fmt.Println("Number", v, " found!")
						bingoMatrices[m].found[i][j] = 1

						won := verifyBoard(bingoMatrices[m])
						if won {

							bingoMatrices[m].won = true
							firstBoardWon = bingoMatrices[m]
							wonBoardLastNumberCalled = v
							return
						}
					}
				}
			}
		}
	}
}

func findLastWinningMatrix() {
	for _, v := range numbers {

		fmt.Println("Current Number ", v)

		for m := 0; m < len(bingoMatrices); m++ {

			for i := 0; i < len(bingoMatrices[m].numbers); i++ {

				if bingoMatrices[m].won {
					break
				}

				for j := 0; j < len(bingoMatrices[m].numbers[i]); j++ {
					if bingoMatrices[m].numbers[i][j] == v {

						fmt.Println("Number", v, " found!")
						bingoMatrices[m].found[i][j] = 1

						won := verifyBoard(bingoMatrices[m])

						if won {
							bingoMatrices[m].won = true
							lastBoardWon = bingoMatrices[m]
							wonBoardLastNumberCalled = v
							break
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

	bingoMatrices = make([]BingoMatrix, 0)
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
	stringNumbersArr := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(stringNumbers), -1)
	intArray := make([]int, len(stringNumbersArr))

	for i, v := range stringNumbersArr {
		intArray[i], _ = strconv.Atoi(v)
	}

	return intArray
}
