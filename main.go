package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello")

	// var elevator *Elevator = NewElevator("1")
	// fmt.Println("Elevator's floor is", elevator.currentFloor)
	// elevator.addNewRequest(12390812)
	// elevator.move()
	// fmt.Println("Elevator's floor is", elevator.currentFloor)
	var yes []int

	var column *Column = NewColumn("A", 5, yes, false)
	for i := 0; i < len(column.elevatorsList); i++ {
		fmt.Println("Elevator ", i, "current floor is ", column.elevatorsList[i].currentFloor)
	}
	var superevator = column.requestElevator(20, "up")
	fmt.Print("Chosen elevator's floor is ", superevator.currentFloor)
	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	runScenario(scenarioNumber)
	// }

	//########### ALL WORKING TESTS
	//
	// f := new(elevator1.Elevator)
	// var door1 Door
	// p := &door1.ID
	// *p = 123
	// callButton1 := newCallButton(1, 1, "up")
	// floorRequestButton1 := newFloorRequestButton(1, 1, "down")
	// fmt.Print(door1.ID, callButton1.direction, floorRequestButton1.direction)
	// fmt.Print(f)

	//########### ALL WORKING TESTS

}
