package main

import "fmt"

func main() {
	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	runScenario(scenarioNumber)
	// }
	door1 := newDoor(1)
	callButton1 := newCallButton(1, 1, "up")
	floorRequestButton1 := newFloorRequestButton(1, 1, "down")
	fmt.Print(door1.status, callButton1.direction, floorRequestButton1.ID)

}
