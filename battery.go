package main

import (
	"fmt"
	"math"
	"sort"
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
	fmt.Printf("Best Elevator's ID is %s and its current floor is %d", bestElevator.ID, bestElevator.currentFloor)
	bestElevator.addNewRequest(1)
	bestElevator.move()

	bestElevator.addNewRequest(_requestedFloor)
	bestElevator.move()

	return &bestColumn, &bestElevator
}

func (b *Battery) findBestColumn(_requestedFloor int) Column {
	returnedObject := b.columnsList[1]
	for _, column := range b.columnsList {
		indexFound := sort.SearchInts(column.servedFloors, _requestedFloor)
		if column.servedFloors[indexFound-1] == _requestedFloor {
			returnedObject = column
		}
	}
	return returnedObject
}
