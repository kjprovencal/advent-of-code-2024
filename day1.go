package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// insert a number into a list in order
func insert(list []int, num int) []int {
	if len(list) == 0 {
		return []int{num}
	}

	for i := 0; i < len(list); i++ {
		if num < list[i] {
			list = append(list, 0)
			copy(list[i+1:], list[i:])
			list[i] = num
			return list
		}
	}

	return append(list, num)
}

// count the number of times a number appears in a sorted list
func count(list []int, num int) int {
	count := 0

	first := sort.SearchInts(list, num)
	for i := first; i < len(list) && list[i] == num; i++ {
		count++
	}

	return count
}

func day1() {
	filepath := "input/day1.txt"

	if len(os.Args) == 4 {
		filepath = os.Args[3]
	}

	file, ferr := os.Open(filepath)

	if ferr != nil {
		panic(ferr)
	}

	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		nums := regexp.MustCompile(`\s+`).Split(line, -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		list1 = insert(list1, num1)
		list2 = insert(list2, num2)
	}

	diffCount := 0
	simCount := 0

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diffCount -= diff
		} else {
			diffCount += diff
		}
		simCount += list1[i] * count(list2, list1[i])
	}

	println("The total distance is: ", diffCount)
	println("The total similarity is: ", simCount)
}
