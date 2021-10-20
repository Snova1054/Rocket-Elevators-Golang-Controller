package main

import (
	elevator1 "Rocket-Elevators-Golang-Controller/elevators"
	"fmt"
)

func main() {
	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	runScenario(scenarioNumber)
	// }

	//########### ALL WORKING TESTS

	f := new(elevator1.Elevator)
	var door1 Door
	p := &door1.ID
	*p = 123
	callButton1 := newCallButton(1, 1, "up")
	floorRequestButton1 := newFloorRequestButton(1, 1, "down")
	fmt.Print(door1.ID, callButton1.direction, floorRequestButton1.direction)
	fmt.Print(f)

	//########### ALL WORKING TESTS

}
