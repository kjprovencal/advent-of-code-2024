package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// 1 for increasing, -1 for decreasing, 0 for no change (bad)
func levelDirection(a string, b string) int {
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)
	diff := abs(aInt - bInt)

	if diff < 1 || diff > 3 {
		return 0
	}

	return (aInt - bInt) / diff
}

// recursive func, 1 if chunk is valid/increasing, -1 if valid/decreasing, 0 if invalid
func checkReport(report []string, withDampener bool, dir int) int {
	// assume each report has at least 2 levels
	for i := 0; i < len(report)-1; i++ {
		newDir := levelDirection(report[i], report[i+1])

		if dir == 0 {
			dir = newDir
		}

		if (dir != 0 && dir != newDir) || newDir == 0 {
			// if not using dampener (part one) or already dampened, report fails
			if !withDampener {
				return 0
			}

			// remove report[i]
			reportA := make([]string, 0, len(report)-1)
			reportA = append(reportA, report[:i]...)
			reportA = append(reportA, report[i+1:]...)
			fmt.Println("Report A:", reportA)

			// remove report[i+1]
			reportB := make([]string, 0, len(report)-1)
			reportB = append(reportB, report[:i+1]...)
			if i+1 < len(report) {
				reportB = append(reportB, report[i+2:]...)
			}
			fmt.Println("Report B:", reportB)

			reportAStatus := checkReport(reportA, false, dir)
			if reportAStatus != 0 {
				fmt.Println("Dampened report A:", reportA)
				return reportAStatus
			}

			reportBStatus := checkReport(reportB, false, dir)
			if reportBStatus != 0 {
				fmt.Println("Dampened report B:", reportB)
				return reportBStatus
			}

			return 0
		}
	}
	return dir
}

func day2() {
	filepath := "input/day2.txt"

	if len(os.Args) == 4 {
		filepath = os.Args[3]
	}

	file, ferr := os.Open(filepath)

	if ferr != nil {
		panic(ferr)
	}

	safeCount := 0
	dampenedCount := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		nums := regexp.MustCompile(`\s+`).Split(line, -1)

		if checkReport(nums, false, 0) != 0 {
			safeCount++
			dampenedCount++
		} else if checkReport(nums, true, 1) != 0 {
			dampenedCount++
		} else if checkReport(nums, true, -1) != 0 {
			dampenedCount++
		}
	}

	println("Safe count:", safeCount)
	println("Dampened count:", dampenedCount)
}
