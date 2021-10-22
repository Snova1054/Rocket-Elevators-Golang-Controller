package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	scenarioNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		runScenario(scenarioNumber)
	}
	fmt.Println("Hello")
}

func contains(_listToCheck []int, _intToCheck int) bool {

	for i := 0; i < len(_listToCheck); i++ {
		if _listToCheck[i] == _intToCheck {

			return true
		}
	}
	return false
}
