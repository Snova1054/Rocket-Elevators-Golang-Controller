package main

import (
	"os"
	"strconv"
)

func main() {
	scenarioNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		runScenario(scenarioNumber)
	}
}

//Function used to search if a specific array/slice contains a value, returns a boolean
func contains(_listToCheck []int, _intToCheck int) bool {
	for i := 0; i < len(_listToCheck); i++ {
		if _listToCheck[i] == _intToCheck {
			return true
		}
	}
	return false
}
