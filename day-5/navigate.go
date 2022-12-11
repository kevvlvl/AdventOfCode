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
	x1         int
	y1         int
	x2         int
	y2         int
	intersects int
}

type LineIntersect struct {
	l1 Line
	l2 Line
}

func FindOverlaps() {

	fmt.Println("FindLeastOverlap() Start")

	loadData()
	findIntersects()

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
		findStraightLines(line, r)
	}

	fmt.Println("loadData() End")
}

func findStraightLines(line string, r *regexp.Regexp) {

	if r.MatchString(line) {

		res := r.FindAllStringSubmatch(line, 1)

		fmt.Println("Line = ", line, " res = ", res)

		// Line = (x1  res[0][1], y1 res[0][2]) -> (x2 res[0][3], y2 res[0][4])
		// Straight lines where x1 = x2 or y1 = y2
		if (res[0][1] == res[0][3]) || (res[0][2] == res[0][4]) {

			var line Line
			line.x1, _ = strconv.Atoi(res[0][1])
			line.y1, _ = strconv.Atoi(res[0][2])
			line.x2, _ = strconv.Atoi(res[0][3])
			line.y2, _ = strconv.Atoi(res[0][4])
			line.intersects = 0

			lines = append(lines, line)
			fmt.Println("   Straight line found: ", line)
		}
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

	if l1.x1 == l1.x2 && l2.x1 == l2.x2 && l1.x1 == l2.x1 {
		fmt.Println("both lines vertical on the same X. verify if Y overlaps")

	} else if l1.y1 == l1.y2 && l2.y1 == l2.y2 && l1.y1 == l2.y1 {
		fmt.Println("both lines horizontal on the same Y. verify if X overlaps")

		//linesIntersects = append(linesIntersects, LineIntersect{l1: l1, l2: l2})
	} else {
		fmt.Println("Verify if both lines intersect")

		//let Line 1 = (l1x1,l1y1) -> (l1x2,l1y2) and Line 2 = (l2x1,l2y1) -> (l2x2,l2y2)

		denominator := (l1.x1-l1.x2)*(l2.y1-l2.y2) - (l1.y1-l1.y2)*(l2.x1-l2.x2)

		if denominator != 0 {

			xPoint := ((l1.x1*l1.y2-l1.y1*l1.x2)*(l2.x1-l2.x2) - (l1.x1-l1.x2)*(l2.x1*l2.y2-l2.y1*l2.x2)) / denominator
			yPoint := ((l1.x1*l1.y2-l1.y1*l1.x2)*(l2.y1-l2.y2) - (l1.y1-l1.y2)*(l2.x1*l2.y2-l2.y1*l2.x2)) / denominator

			fmt.Println("X,Y = ", xPoint, ",", yPoint)
		}

		//intersect := false
		//intersectX := 0
		//intersectY := 0
	}
}
