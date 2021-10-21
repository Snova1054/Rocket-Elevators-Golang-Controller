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
	var theNewDoor *Door = NewDoor(1)
	elevator := Elevator{ID: _elevatorID, status: "idle", currentFloor: 1, direction: "null", door: *theNewDoor}
	return &elevator
}

//moving method, need to test it
func (e *Elevator) move() {
	for {
		if len(e.floorRequestsList) == 0 {
			break
		} else {
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
			//operateDoors()
			e.floorRequestsList = e.floorRequestsList[1:1]
		}
	}
	e.status = "idle"
}

func (e *Elevator) addNewRequest(_requestedFloor int) {
	e.completedRequestsList = append(e.completedRequestsList, _requestedFloor)
}
