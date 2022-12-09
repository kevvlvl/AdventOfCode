package day_5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lines = make([]Line, 0)

type Line struct {
	x1         int
	y1         int
	x2         int
	y2         int
	intersects int
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

			// skip if current lines refer to each-other (same line!)
			if i == j {
				continue
			}

			showLineDirection(v1)
			showLineDirection(v2)

			// TODO: calculate intersects
			//  X intersection:
			//    where (v1.x1 >= v2.x1 && v1.x2 <= v2.x2) or (v2.x1 >= v1.x1 && v2.x2 <= v1.x2)
			//  Y intersection:
			//    where (v1.y1 >= v2.y1 && v1.y2 <= v2.y2) or (v2.y1 >= v1.y1 && v2.y2 <= v1.y2)

			// TODO increment intersects counter
			//  keep in memory lines intersecting each other to not duplicate counts
		}
	}
}

func showLineDirection(line Line) {

	if line.x1 == line.x2 {
		fmt.Println("is Vertical line: ", line)
	} else {
		fmt.Println("is Horizontal line: ", line)
	}
}
