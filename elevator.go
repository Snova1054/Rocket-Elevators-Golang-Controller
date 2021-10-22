package main

import (
	"fmt"
	"sort"
)

//Declares each Elevator
type Elevator struct {
	ID                    string
	status                string
	currentFloor          int
	direction             string
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
}

//Function used to create new Elevators with the desired properties
func NewElevator(_elevatorID string) *Elevator {

	elevator := new(Elevator)
	elevator.ID = _elevatorID
	elevator.status = "idle"
	elevator.currentFloor = 1
	elevator.direction = ""
	elevator.door = *NewDoor(1)
	return elevator
}

//Method used by the Column or the Battery to move the Elevator
func (e *Elevator) move() {
	for {
		if len(e.floorRequestsList) == 0 {
			break
		} else {
			for i := len(e.floorRequestsList); i >= 1; i-- {
				var destination int = e.floorRequestsList[0]
				e.status = "moving"
				if e.currentFloor < destination {
					e.direction = "up"
					sort.Ints(e.floorRequestsList)
					for e.currentFloor < destination {
						fmt.Printf("Elevator %q is on the floor #%d\n", e.ID, e.currentFloor)
						e.currentFloor++
					}
				} else if e.currentFloor > destination {
					e.direction = "down"
					sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
					for e.currentFloor > destination {
						fmt.Printf("Elevator %q is on the floor #%d\n", e.ID, e.currentFloor)
						e.currentFloor--
					}
				}
				fmt.Printf("Elevator %q has arrived at the floor %d\n", e.ID, e.currentFloor)
				e.status = "stopped"
				e.operateDoors()
				e.floorRequestsList = e.floorRequestsList[1:i]
			}
		}
	}
	e.status = "idle"
}

//Method used by the Elevator to operate its Doors
func (e *Elevator) operateDoors() {
	e.door.status = "opened"
	fmt.Println("Elevator's doors have opened")
	//Wait 5 seconds
	fmt.Printf("Elevator's doors have closed\n\n")
	e.door.status = "closed"
}

//Method used to add new floor requests
func (e *Elevator) addNewRequest(_requestedFloor int) {
	e.completedRequestsList = append(e.completedRequestsList, _requestedFloor)
	if !contains(e.floorRequestsList, _requestedFloor) {
		e.floorRequestsList = append(e.floorRequestsList, _requestedFloor)
	}
	if e.currentFloor < _requestedFloor {
		e.direction = "up"
	} else if e.currentFloor > _requestedFloor {
		e.direction = "down"
	}
}
