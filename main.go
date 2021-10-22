package main

import (
	"fmt"
)

func main() {
	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	runScenario(scenarioNumber)
	// }

	var battery *Battery = NewBattery(1, 4, 60, 6, 5)
	columnC := battery.columnsList[2]
	columnC.elevatorsList[0].currentFloor = 1
	columnC.elevatorsList[0].direction = "up"
	columnC.elevatorsList[0].status = "stopped"
	columnC.elevatorsList[0].floorRequestsList = append(columnC.elevatorsList[0].floorRequestsList, 21)
	columnC.elevatorsList[1].currentFloor = 23
	columnC.elevatorsList[1].direction = "up"
	columnC.elevatorsList[1].status = "moving"
	columnC.elevatorsList[1].floorRequestsList = append(columnC.elevatorsList[1].floorRequestsList, 28)
	columnC.elevatorsList[2].currentFloor = 33
	columnC.elevatorsList[2].direction = "down"
	columnC.elevatorsList[2].status = "moving"
	columnC.elevatorsList[2].floorRequestsList = append(columnC.elevatorsList[2].floorRequestsList, 1)
	columnC.elevatorsList[3].currentFloor = 40
	columnC.elevatorsList[3].direction = "down"
	columnC.elevatorsList[3].status = "moving"
	columnC.elevatorsList[3].floorRequestsList = append(columnC.elevatorsList[3].floorRequestsList, 24)
	columnC.elevatorsList[4].currentFloor = 39
	columnC.elevatorsList[4].direction = "down"
	columnC.elevatorsList[4].status = "moving"
	columnC.elevatorsList[4].floorRequestsList = append(columnC.elevatorsList[4].floorRequestsList, 1)

	bestColumn, bestElevator := battery.assignElevator(36, "up")
	moveAllElevators(bestColumn)

	fmt.Printf("Best Elevator's floor is %d", bestElevator.currentFloor)
	bestColumn.requestElevator(59, "down")
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
