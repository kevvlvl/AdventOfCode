package main

import (
	day1 "AdventOfCode/day-1"
	day2 "AdventOfCode/day-2"
	day3 "AdventOfCode/day-3"
	day4 "AdventOfCode/day-4"
)

func main() {

	//puzzleDay1()
	//puzzleDay2()
	//puzzleDay3()
	puzzleDay4()
}

func puzzleDay1() {
	day1.Measurements()
	day1.ThreeMeasurements()
}

func puzzleDay2() {
	day2.MeasureDive()
	day2.MeasureDiveWithAim()
}

func puzzleDay3() {
	day3.MeasurePowerConsumption()
	day3.MeasureLifeSupportingRate()
}

func puzzleDay4() {
	//day4.FindBingoMatrix()
	day4.FindLastWinningBingoMatrix()
}
