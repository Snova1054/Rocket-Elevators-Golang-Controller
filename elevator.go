package main

import (
	"sort"
)

type Elevator struct {
	ID                    string
	status                string
	currentFloor          int
	direction             string
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
}

func NewElevator(_elevatorID string) *Elevator {

	elevator := new(Elevator)
	elevator.ID = _elevatorID
	elevator.status = "idle"
	elevator.currentFloor = 1
	elevator.direction = ""
	elevator.door = *NewDoor(1)
	return elevator
}

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
						e.currentFloor++
					}
				} else if e.currentFloor > destination {
					e.direction = "down"
					sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
					for e.currentFloor > destination {
						e.currentFloor--
					}
				}
				e.status = "stopped"
				e.operateDoors()
				e.floorRequestsList = e.floorRequestsList[1:i]
			}
		}
	}
	e.status = "idle"
}

func (e *Elevator) operateDoors() {
	e.door.status = "opened"
	//Wait 5 seconds
	e.door.status = "closed"
}

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
