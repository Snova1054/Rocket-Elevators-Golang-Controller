package main

import (
	"fmt"
	"math"
)

var columnID rune = 'A'
var floorRequestButtonID int = 1

//Declares each Battery
type Battery struct {
	ID                      int
	status                  string
	columnsList             []Column
	floorRequestButtonsList []FloorRequestButton
}

//Function used to create new Batteries with the desired properties
func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *Battery {
	battery := Battery{ID: _id, status: "online", columnsList: []Column{}, floorRequestButtonsList: []FloorRequestButton{}}

	if _amountOfBasements > 0 {
		battery.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		battery.createBasementFloorRequestButtons(_amountOfBasements)
		_amountOfColumns--
	}

	battery.createFloorRequestButtons(float64(_amountOfFloors))
	battery.createColumns(float64(_amountOfColumns), float64(_amountOfFloors), _amountOfElevatorPerColumn)
	columnID = 'A'
	return &battery
}

//Method used by the Battery to create a basement Column
func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	var servedFloors []int
	floor := -1
	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	basementColumn := NewColumn(string(columnID), _amountOfElevatorPerColumn, servedFloors, true)
	b.columnsList = append(b.columnsList, *basementColumn)
	columnID++
}

//Method used to create Columns
func (b *Battery) createColumns(_amountOfColumns float64, _amountOfFloors float64, _amountOfElevatorPerColumn int) {
	amountOfFloorsPerColumn := math.Ceil(_amountOfFloors / float64(_amountOfColumns))
	floor := 1

	for i := 0; i < int(_amountOfColumns); i++ {
		var servedFloors []int

		for j := 0; j < int(amountOfFloorsPerColumn); j++ {
			if floor <= int(_amountOfFloors) {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := NewColumn(string(columnID), _amountOfElevatorPerColumn, servedFloors, false)
		b.columnsList = append(b.columnsList, *column)
		columnID++
	}
}

//Method used to create basement FloorRequestButtons
func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "down")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
}

//Method used to create FloorRequestButtons
func (b *Battery) createFloorRequestButtons(_amountOfFloors float64) {
	buttonFloor := 1
	for i := 0; i < int(_amountOfFloors); i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "up")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
}

//Method used to assign the best Elevator according to the best Column from the user's inputs
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	fmt.Printf("An elevator has been requested for the floor %d to go %s\n\n", _requestedFloor, _direction)
	bestColumn := b.findBestColumn(_requestedFloor)
	bestElevator := (bestColumn).findElevator(1, _direction)
	fmt.Printf("Elevator %q hon the floor #%d has been selected as the best elevator\n\n", bestElevator.ID, bestElevator.currentFloor)
	bestElevator.addNewRequest(1)
	bestElevator.move()

	bestElevator.addNewRequest(_requestedFloor)
	bestElevator.move()
	return bestColumn, bestElevator
}

//Method used to find the best Column
func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for i := 0; i < len(b.columnsList); i++ {
		var returnedColumn = b.columnsList[i]
		for j := 0; j < len(b.columnsList[i].servedFloors); j++ {
			if b.columnsList[i].servedFloors[j] == _requestedFloor {
				return &returnedColumn
			}
		}
	}
	return nil
}
