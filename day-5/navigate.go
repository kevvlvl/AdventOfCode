package day_5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lines = make([]Line, 0)
var linesIntersects = make([]LineIntersect, 0)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

/*
Horizontal line = same Y value
*/
func (line Line) isHorizontal() bool {
	return line.y1 == line.y2
}

/*
Vertical line = same X value
*/
func (line Line) isVertical() bool {
	return line.x1 == line.x2
}

type LineIntersect struct {
	l1 Line
	l2 Line
}

func FindOverlaps() {

	fmt.Println("FindLeastOverlap() Start")

	loadData()
	findIntersects()

	fmt.Println("Number of intersects of two lines: ", len(linesIntersects))

	fmt.Println("FindLeastOverlap() End")
}

func loadData() {

	fmt.Println("loadData() Start")

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

		line := scanner.Text()
		strToLine(line, r)
	}

	fmt.Println("Parsed ", len(lines), " of strings as Line structs")
	fmt.Println("loadData() End")
}

func strToLine(line string, r *regexp.Regexp) {

	if r.MatchString(line) {

		res := r.FindAllStringSubmatch(line, 1)

		// Line = (x1  res[0][1], y1 res[0][2]) -> (x2 res[0][3], y2 res[0][4])

		var line Line
		line.x1, _ = strconv.Atoi(res[0][1])
		line.y1, _ = strconv.Atoi(res[0][2])
		line.x2, _ = strconv.Atoi(res[0][3])
		line.y2, _ = strconv.Atoi(res[0][4])

		lines = append(lines, line)
	} else {
		fmt.Println("ERROR - Line does not match regex: ", line)
	}
}

func findIntersects() {

	for i, v1 := range lines {

		for j, v2 := range lines {

			// line = (x1,y1)->(x2,y2)
			// intersects happen where two lines whose values x1 and x2 are within range or y1 and y2 are within range

			// skip if current lines refer to each-other
			if i == j {
				continue
			}

			calc(v1, v2)
		}
	}
}

func calc(l1 Line, l2 Line) {

	// let Line 1 = (l1x1,l1y1) -> (l1x2,l1y2) and Line 2 = (l2x1,l2y1) -> (l2x2,l2y2)

	if l1.isHorizontal() {

		if l2.isHorizontal() {
			parallelHorizontalLinesOverlap(l1, l2)
		} else if l2.isVertical() {
			perpendicularLinesOverlap(l1, l2)
		}
	} else if l1.isVertical() {
		if l2.isHorizontal() {
			perpendicularLinesOverlap(l2, l1)
		} else if l2.isVertical() {
			parallelVerticalLinesOverlap(l1, l2)
		}
	}
}

func parallelHorizontalLinesOverlap(l1 Line, l2 Line) {

	if l1.y1 == l2.y1 {

		line1X1OverlapLine2 := (l1.x1 >= l2.x1 && l1.x1 <= l2.x2) || (l1.x1 >= l2.x2 && l1.x1 <= l2.x1)
		line1X2OverlapLine2 := (l1.x2 >= l2.x1 && l1.x2 <= l2.x2) || (l1.x2 >= l2.x2 && l1.x2 <= l2.x1)

		line2X1OverlapLine1 := (l2.x1 >= l1.x1 && l2.x1 <= l1.x2) || (l2.x1 >= l1.x2 && l2.x1 <= l1.x1)
		line2X2OverlapLine1 := (l2.x2 >= l1.x1 && l2.x2 <= l1.x2) || (l2.x2 >= l1.x2 && l2.x2 <= l1.x1)

		overlap := line1X1OverlapLine2 ||
			line1X2OverlapLine2 ||
			line2X1OverlapLine1 ||
			line2X2OverlapLine1

		if overlap {
			addIntersectToList(l1, l2)
		}
	}
}

func parallelVerticalLinesOverlap(l1 Line, l2 Line) {

	if l1.x1 == l2.x1 {

		line1Y1OverlapLine2 := (l1.y1 >= l2.y1 && l1.y1 <= l2.y2) || (l1.y1 >= l2.y2 && l1.y1 <= l2.y1)
		line1Y2OverlapLine2 := (l1.y2 >= l2.y1 && l1.y2 <= l2.y2) || (l1.y2 >= l2.y2 && l1.y2 <= l2.y1)

		line2Y1OverlapLine1 := (l2.y1 >= l1.y1 && l2.y1 <= l1.y2) || (l2.y1 >= l1.y2 && l2.y1 <= l1.y1)
		line2Y2OverlapLine1 := (l2.y2 >= l1.y1 && l2.y2 <= l1.y2) || (l2.y2 >= l1.y2 && l2.y2 <= l1.y1)

		overlap := line1Y1OverlapLine2 ||
			line1Y2OverlapLine2 ||
			line2Y1OverlapLine1 ||
			line2Y2OverlapLine1

		if overlap {
			addIntersectToList(l1, l2)
		}
	}
}

func perpendicularLinesOverlap(horizontalLine Line, verticalLine Line) {

	verticalLowerBoundary := min(verticalLine.y1, verticalLine.y2)
	verticalUpperBoundary := max(verticalLine.y1, verticalLine.y2)

	if horizontalLine.x1 <= horizontalLine.x2 {

		for i := horizontalLine.x1; i <= horizontalLine.x2; i++ {

			if i == verticalLine.x1 && horizontalLine.y1 >= verticalLowerBoundary && horizontalLine.y1 <= verticalUpperBoundary {
				addIntersectToList(horizontalLine, verticalLine)
				break
			}
		}
	} else {

		for i := horizontalLine.x2; i <= horizontalLine.x1; i++ {

			if i == verticalLine.x1 && horizontalLine.y1 >= verticalLowerBoundary && horizontalLine.y1 <= verticalUpperBoundary {
				addIntersectToList(horizontalLine, verticalLine)
				break
			}
		}
	}
}

/*
Append the intersect of the two lines only if it ha snot already been saved.
*/
func addIntersectToList(l1 Line, l2 Line) {

	//fmt.Println("Both lines (l1 ", l1, ", l2 ", l2, ") intersect")

	if !isIntersectInList(l1, l2) {
		linesIntersects = append(linesIntersects, LineIntersect{l1: l1, l2: l2})
	}
}

func isIntersectInList(l1 Line, l2 Line) bool {

	var found = false

	for i := range linesIntersects {
		if (linesIntersects[i].l1 == l1 && linesIntersects[i].l2 == l2) ||
			linesIntersects[i].l1 == l2 && linesIntersects[i].l2 == l1 {
			found = true
			break
		}
	}

	return found
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
