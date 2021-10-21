package main

type Battery struct {
	ID     int
	status string
	// columnsList             []Column
	// floorRequestButtonsList []FloorRequestButton
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	battery := Battery{ID: _id, status: "online"}
	return &battery
}

// func (b *Battery) findBestColumn(_requestedFloor int) *Column {

// }

// //Simulate when a user press a button at the lobby
// func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {

// }
