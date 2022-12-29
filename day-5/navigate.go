package day_5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lines = make([]Line, 0)
var m = make(map[Coordinate]int)

type Line struct {
	p1 Coordinate
	p2 Coordinate
}

type Coordinate struct {
	x int
	y int
}

/*
Horizontal line = same Y value
*/
func (line Line) isHorizontal() bool {
	return line.p1.y == line.p2.y
}

/*
Vertical line = same X value
*/
func (line Line) isVertical() bool {
	return line.p1.x == line.p2.x
}

func FindOverlaps() {

	fmt.Println("FindOverlaps() Start")

	readFile(false)
	loadMap()
	countIntersectsGreaterThan(2)

	fmt.Println("FindOverlaps() End")
}

func FindOverLapsWithDiagnonalLines() {

	// empty data already in-memory from previous execution
	lines = make([]Line, 0)
	m = make(map[Coordinate]int)

	fmt.Println("FindOverLapsWithDiagnonalLines() Start")

	readFile(true)
	loadMap()

	fmt.Println("FindOverLapsWithDiagnonalLines() End")

}

func readFile(loadDiagonal bool) {

	fmt.Println("readFile() Start")

	file, err := os.Open("day-5/day5-input.txt")
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
	r := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	for scanner.Scan() {

		// Line = (x1  res[0][1], y1 res[0][2]) -> (x2 res[0][3], y2 res[0][4])
		line := scanner.Text()

		if r.MatchString(line) {

			res := r.FindAllStringSubmatch(line, 1)

			x1, _ := strconv.Atoi(res[0][1])
			y1, _ := strconv.Atoi(res[0][2])
			x2, _ := strconv.Atoi(res[0][3])
			y2, _ := strconv.Atoi(res[0][4])

			if x1 == x2 || y1 == y2 {
				lines = append(lines, Line{p1: Coordinate{x1, y1}, p2: Coordinate{x2, y2}})
			} else if loadDiagonal {

				// TODO: load diagonal 45 degrees lines where value of slope (y2 - y1 / x2 - x1) = 1 or -1
			}

		} else {
			fmt.Println("ERROR - Line (", line, ") does not match regex")
		}
	}

	fmt.Println("Parsed ", len(lines), " of strings as Line structs")
	fmt.Println("readFile() End")
}

func loadMap() {

	fmt.Println("loadMap() Start loading coordinates into a Map")

	for _, v := range lines {

		if v.isHorizontal() {

			lowerX := min(v.p1.x, v.p2.x)
			upperX := max(v.p1.x, v.p2.x)

			for i := lowerX; i <= upperX; i++ {
				m[Coordinate{x: i, y: v.p1.y}]++
			}

		} else if v.isVertical() {

			lowerY := min(v.p1.y, v.p2.y)
			upperY := max(v.p1.y, v.p2.y)

			for i := lowerY; i <= upperY; i++ {
				m[Coordinate{x: v.p1.x, y: i}]++
			}
		}
	}

	fmt.Println("loadMap() End")
}

func countIntersectsGreaterThan(minIntersects int) {

	sumOfIntersects := 0

	for _, v := range m {
		if v >= minIntersects {
			sumOfIntersects++
		}
	}

	fmt.Println("The number of coordinates with intersects greater than ", minIntersects, " is ", sumOfIntersects)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
