package main

import (
	"fmt"
	"math"
)

var elevatorID rune = 'A'

type Column struct {
	ID                       string
	status                   string
	servedFloors             []int
	isBasement               bool
	elevatorsList            []Elevator
	callButtonsList          []CallButton
	bestElevatorInformations struct {
		bestElevator Elevator
		bestScore    int
		referenceGap float64
	}
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
	elevatorID = 'A'
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

	return bestElevator
}

func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) *Elevator {
	c.bestElevatorInformations.bestElevator = c.elevatorsList[0]
	c.bestElevatorInformations.bestScore = 6
	c.bestElevatorInformations.referenceGap = float64(len(c.servedFloors)) * 2

	if _requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				c.checkIfElevatorIsBetter(1, elevator, _requestedFloor)
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				c.checkIfElevatorIsBetter(2, elevator, _requestedFloor)
			} else if elevator.currentFloor < 1 && elevator.direction == "up" {
				c.checkIfElevatorIsBetter(3, elevator, _requestedFloor)
			} else if elevator.currentFloor > 1 && elevator.direction == "down" {
				c.checkIfElevatorIsBetter(3, elevator, _requestedFloor)
			} else if elevator.status == "idle" {
				c.checkIfElevatorIsBetter(4, elevator, _requestedFloor)
			} else {
				c.checkIfElevatorIsBetter(5, elevator, _requestedFloor)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == _requestedFloor && elevator.status == "stopped" && elevator.direction == _requestedDirection {
				c.checkIfElevatorIsBetter(1, elevator, _requestedFloor)
			} else if elevator.currentFloor < _requestedFloor && elevator.direction == "up" && _requestedDirection == "up" {
				c.checkIfElevatorIsBetter(2, elevator, _requestedFloor)
			} else if elevator.currentFloor > _requestedFloor && elevator.direction == "down" && _requestedDirection == "down" {
				c.checkIfElevatorIsBetter(2, elevator, _requestedFloor)
			} else if elevator.status == "idle" {
				c.checkIfElevatorIsBetter(4, elevator, _requestedFloor)
			} else {
				c.checkIfElevatorIsBetter(5, elevator, _requestedFloor)
			}
		}
	}
	return &c.bestElevatorInformations.bestElevator
}

func (c *Column) checkIfElevatorIsBetter(_scoreToCheck int, _newElevator Elevator, _floor int) {
	if _scoreToCheck < c.bestElevatorInformations.bestScore {
		c.bestElevatorInformations.bestScore = _scoreToCheck
		c.bestElevatorInformations.bestElevator = _newElevator
		c.bestElevatorInformations.referenceGap = math.Abs(float64(_newElevator.currentFloor) - float64(_floor))
	} else if c.bestElevatorInformations.bestScore == _scoreToCheck {
		gap := math.Abs(float64(_newElevator.currentFloor) - float64(_floor))
		if c.bestElevatorInformations.referenceGap > gap {
			c.bestElevatorInformations.bestElevator = _newElevator
			c.bestElevatorInformations.referenceGap = gap
		}
	}

}
