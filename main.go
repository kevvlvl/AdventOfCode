package main

import (
	day1 "AdventOfCode/day-1"
	day2 "AdventOfCode/day-2"
	day3 "AdventOfCode/day-3"
	day4 "AdventOfCode/day-4"
	day5 "AdventOfCode/day-5"
	"fmt"
)

func main() {

	puzzleDay1()
	puzzleDay2()
	puzzleDay3()
	puzzleDay4()
	puzzleDay5()
}

func puzzleDay1() {

	fmt.Print(`***************
DAY 1
***************
`)

	day1.Measurements()
	day1.ThreeMeasurements()
}

func puzzleDay2() {

	fmt.Print(`***************
DAY 2
***************
`)

	day2.MeasureDive()
	day2.MeasureDiveWithAim()
}

func puzzleDay3() {

	fmt.Print(`***************
DAY 3
***************
`)

	day3.MeasurePowerConsumption()
	day3.MeasureLifeSupportingRate()
}

func puzzleDay4() {

	fmt.Print(`***************
DAY 4
***************
`)

	day4.FindBingoMatrix()
	day4.FindLastWinningBingoMatrix()
}

func puzzleDay5() {

	fmt.Print(`***************
DAY 5
***************
`)

	day5.FindOverlaps()
	day5.FindOverLapsWithDiagnonalLines()
}
