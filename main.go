package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Please provide a day number as an argument.")
		return
	}

	funcMap := map[string]func(){
		"1": day1,
	}

	dayArg := os.Args[1]

	if fn, exists := funcMap[dayArg]; exists {
		fn()
	} else {
		fmt.Printf("Day %s is not implemented yet.\n", dayArg)
	}
}
