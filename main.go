package main

import (
	"fmt"
)

func main() {
	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	runScenario(scenarioNumber)
	// }

	//scenario 1 elevator

	var battery *Battery = NewBattery(1, 4, 60, 6, 5)
	column := battery.columnsList[1]

	column.elevatorsList[0].currentFloor = 20
	column.elevatorsList[0].direction = "down"
	column.elevatorsList[0].status = "moving"
	column.elevatorsList[0].floorRequestsList = append(column.elevatorsList[0].floorRequestsList, 5)

	column.elevatorsList[1].currentFloor = 3
	column.elevatorsList[1].direction = "up"
	column.elevatorsList[1].status = "moving"
	column.elevatorsList[1].floorRequestsList = append(column.elevatorsList[1].floorRequestsList, 15)

	column.elevatorsList[2].currentFloor = 13
	column.elevatorsList[2].direction = "down"
	column.elevatorsList[2].status = "moving"
	column.elevatorsList[2].floorRequestsList = append(column.elevatorsList[2].floorRequestsList, 1)

	column.elevatorsList[3].currentFloor = 15
	column.elevatorsList[3].direction = "down"
	column.elevatorsList[3].status = "moving"
	column.elevatorsList[3].floorRequestsList = append(column.elevatorsList[3].floorRequestsList, 2)

	column.elevatorsList[4].currentFloor = 6
	column.elevatorsList[4].direction = "down"
	column.elevatorsList[4].status = "moving"
	column.elevatorsList[4].floorRequestsList = append(column.elevatorsList[4].floorRequestsList, 2)

	chosenColumn, chosenElevator := battery.assignElevator(20, "up")
	moveAllElevators(chosenColumn)

	fmt.Printf("Best Elevator's floor is %d", chosenElevator.currentFloor)
	// chosenColumn.requestElevator(59, "down")
	fmt.Println("Hello")
}

func moveAllElevators(c *Column) {
	for i := 0; i < len(c.elevatorsList); i++ {
		for len(c.elevatorsList[i].floorRequestsList) != 0 {
			c.elevatorsList[i].move()
			fmt.Println("Elevator's ", c.elevatorsList[i].ID, " last position is floor ", c.elevatorsList[i].currentFloor)
		}
	}
}
