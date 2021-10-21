package main

import (
	"fmt"
	"math"
)

type Column struct {
	ID              string
	status          string
	servedFloors    []int
	isBasement      bool
	elevatorsList   []Elevator
	callButtonsList []CallButton
}

func NewColumn(_id string, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	column := Column{ID: _id, status: "online", servedFloors: _servedFloors, isBasement: _isBasement}
	column.createElevators(len(_servedFloors), _amountOfElevators)
	column.createCallButtons(len(_servedFloors), _isBasement)
	return &column
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	callButtonID := 1 // Look for global variable
	if _isBasement {
		buttonFloor := -1
		for i := 0; i < _amountOfFloors; i++ {
			var callButton = NewCallButton(callButtonID, buttonFloor, "up")
			c.callButtonsList = append(c.callButtonsList, *callButton)
			buttonFloor--
			callButtonID++
		}
	} else {
		buttonFloor := 1
		for i := 0; i < _amountOfFloors; i++ {
			var callButton = NewCallButton(callButtonID, buttonFloor, "down")
			c.callButtonsList = append(c.callButtonsList, *callButton)
			buttonFloor++
			callButtonID++
		}
	}
}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	elevatorID := 'A'
	for i := 0; i < _amountOfElevators; i++ {
		var elevator = NewElevator(string(elevatorID))
		c.elevatorsList = append(c.elevatorsList, *elevator)
		elevatorID++
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	var bestElevator = c.findElevator(_requestedFloor, _direction)
	bestElevator.addNewRequest(_requestedFloor)
	bestElevator.move()
	fmt.Print(bestElevator.currentFloor)
	bestElevator.addNewRequest(1)
	bestElevator.move()

	return &bestElevator
}

func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) Elevator {
	var bestElevator = c.elevatorsList[0]
	bestScore := 6
	var referenceGap float64 = float64(len(c.servedFloors)) * 2

	if _requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.currentFloor < 1 && elevator.direction == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.currentFloor > 1 && elevator.direction == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == _requestedFloor && elevator.status == "stopped" && elevator.direction == _requestedDirection {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.currentFloor < _requestedFloor && elevator.direction == "up" && _requestedDirection == "up" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.currentFloor > _requestedFloor && elevator.direction == "down" && _requestedDirection == "down" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			} else {
				bestElevator, bestScore, referenceGap = c.checkIfElevatorIsBetter(5, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
			}
		}
	}
	return bestElevator
}

func (c *Column) checkIfElevatorIsBetter(_scoreToCheck int, _newElevator Elevator, _bestScore int, _referenceGap float64, _bestElevator Elevator, _floor int) (Elevator, int, float64) {
	if _scoreToCheck < _bestScore {
		_bestScore = _scoreToCheck
		_bestElevator = _newElevator
		_referenceGap = math.Abs(float64(_newElevator.currentFloor) - +float64(_floor))
	} else if _bestScore == _scoreToCheck {
		gap := math.Abs(float64(_newElevator.currentFloor) - +float64(_floor))
		if _referenceGap > gap {
			_bestElevator = _newElevator
			_referenceGap = gap
		}
	}
	return _bestElevator, _bestScore, _referenceGap
}
