package main

type Elevator struct {
	ID                    string
	status                string
	currentFloor          int
	direction             string
	door                  Door
	floorRequestsList     int
	completedRequestsList int
}

func NewElevator(_elevatorID string) *Elevator {
	elevator := Elevator{
		ID:           _elevatorID,
		status:       "idle",
		currentFloor: 1,
		direction:    "null",
		door:         newDoor(1),
	}
	return &elevator
}

// func (e *Elevator) move() {

// }
