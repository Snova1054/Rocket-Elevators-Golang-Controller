package main

import (
	"fmt"
	// "os"
	// "strconv"
)

func main() {
	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	runScenario(scenarioNumber)
	// }

	battery := NewBattery(1, 4, 60, 6, 5) // watch at battery line 101 for error.
	battery.columnsList[1].elevatorsList[0].currentFloor = 19
	battery.columnsList[1].elevatorsList[1].currentFloor = 15
	battery.columnsList[1].elevatorsList[2].currentFloor = 14
	battery.columnsList[1].elevatorsList[3].currentFloor = 12
	battery.columnsList[1].elevatorsList[4].currentFloor = 10
	bestColumn, bestElevator := battery.assignElevator(15, "up")
	fmt.Printf("Best Elevator's floor is %d", bestElevator.currentFloor)
	bestColumn.requestElevator(15, "down")
	fmt.Println("Hello")

	// var elevator *Elevator = NewElevator("1")
	// fmt.Println("Elevator's floor is", elevator.currentFloor)
	// elevator.addNewRequest(12390812)
	// elevator.move()
	// fmt.Println("Elevator's floor is", elevator.currentFloor)
	// var yes []int

	// var column *Column = NewColumn("A", 5, yes, false)
	// for i := 0; i < len(column.elevatorsList); i++ {
	// 	fmt.Println("Elevator ", i, "current floor is ", column.elevatorsList[i].currentFloor)
	// }
	// superevator := column.requestElevator(20, "up")
	// fmt.Print("Chosen elevator's floor is ", superevator.currentFloor)
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
