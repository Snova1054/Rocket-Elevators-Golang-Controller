package main

import (
	"fmt"
	"math"
)

var columnID rune = 'A'
var floorRequestButtonID int = 1

type Battery struct {
	ID                      int
	status                  string
	columnsList             []Column
	floorRequestButtonsList []FloorRequestButton
}

func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *Battery {
	battery := Battery{ID: _id, status: "online"}

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

func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "down")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
}

func (b *Battery) createFloorRequestButtons(_amountOfFloors float64) {
	buttonFloor := 1
	for i := 0; i < int(_amountOfFloors); i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, buttonFloor, "up")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
}

func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	bestColumn := b.findBestColumn(_requestedFloor)
	bestElevator := bestColumn.findElevator(1, _direction)
	fmt.Printf("Best Elevator's ID is %q and its current floor is %d \n", bestElevator.ID, bestElevator.currentFloor)
	bestElevator.addNewRequest(1)
	bestElevator.move()
	fmt.Printf("Requested floor is %d \n", _requestedFloor)
	bestElevator.addNewRequest(_requestedFloor)
	bestElevator.move()
	return bestColumn, bestElevator
}

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
