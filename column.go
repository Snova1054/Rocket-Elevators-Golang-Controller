package main

import (
	"fmt"
	"math"
)

var elevatorID rune = 'A'

//Declares each Column
type Column struct {
	ID              string
	status          string
	servedFloors    []int
	isBasement      bool
	elevatorsList   []Elevator
	callButtonsList []CallButton
}

//Declares a BestElevatorInformations before using it in the method : findElevator
type BestElevatorInformations struct {
	bestElevator *Elevator
	bestScore    int
	referenceGap float64
}

//Function used to create new Columns with the desired properties
func NewColumn(_id string, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	column := new(Column)
	column.ID = _id
	column.status = "online"
	column.servedFloors = _servedFloors
	column.isBasement = _isBasement
	column.elevatorsList = []Elevator{}
	column.callButtonsList = []CallButton{}
	column.createElevators(len(_servedFloors), _amountOfElevators)
	column.createCallButtons(len(_servedFloors), _isBasement)
	return column
}

//Method used by the Column to create Call Buttons
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

//Method used to create Elevators
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
	fmt.Printf("An elevator has been requested for the floor %d to go %s\n\n", _requestedFloor, _direction)
	bestElevator := c.findElevator(_requestedFloor, _direction)
	fmt.Printf("Elevator %q on the floor #%d has been selected as the best elevator\n\n", bestElevator.ID, bestElevator.currentFloor)
	bestElevator.addNewRequest(_requestedFloor)
	bestElevator.move()

	fmt.Print(bestElevator.currentFloor)
	bestElevator.addNewRequest(1)
	bestElevator.move()

	return bestElevator
}

//Method used to find the best Elevator possible
func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) *Elevator {
	bestElevatorInformations := BestElevatorInformations{&c.elevatorsList[0], 6, float64(len(c.servedFloors) * 2)}

	if _requestedFloor == 1 {
		for i := 0; i < len(c.elevatorsList); i++ {
			if c.elevatorsList[i].currentFloor == 1 && c.elevatorsList[i].status == "stopped" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(1, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else if c.elevatorsList[i].currentFloor == 1 && c.elevatorsList[i].status == "idle" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(2, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else if c.elevatorsList[i].currentFloor < 1 && c.elevatorsList[i].direction == "up" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(3, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else if c.elevatorsList[i].currentFloor > 1 && c.elevatorsList[i].direction == "down" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(3, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else if c.elevatorsList[i].status == "idle" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(4, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(5, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			}

		}
	} else {
		for i := 0; i < len(c.elevatorsList); i++ {
			if c.elevatorsList[i].currentFloor == _requestedFloor && c.elevatorsList[i].status == "stopped" && c.elevatorsList[i].direction == _requestedDirection {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(1, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else if c.elevatorsList[i].currentFloor < _requestedFloor && c.elevatorsList[i].direction == "up" && _requestedDirection == "up" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(2, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else if c.elevatorsList[i].currentFloor > _requestedFloor && c.elevatorsList[i].direction == "down" && _requestedDirection == "down" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(2, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else if c.elevatorsList[i].status == "idle" {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(4, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			} else {
				bestElevatorInformations = *c.checkIfElevatorIsBetter(5, &c.elevatorsList[i], bestElevatorInformations, _requestedFloor)
			}
		}
	}
	return bestElevatorInformations.bestElevator
}

//Method used to compared a new Elevator's information with the bestElevator's
func (c *Column) checkIfElevatorIsBetter(_scoreToCheck int, _newElevator *Elevator, bestElevatorInformations BestElevatorInformations, _floor int) *BestElevatorInformations {
	if _scoreToCheck < bestElevatorInformations.bestScore {
		bestElevatorInformations.bestScore = _scoreToCheck
		bestElevatorInformations.bestElevator = _newElevator
		bestElevatorInformations.referenceGap = math.Abs(float64(_newElevator.currentFloor) - float64(_floor))
	} else if bestElevatorInformations.bestScore == _scoreToCheck {
		gap := math.Abs(float64(_newElevator.currentFloor) - float64(_floor))
		if bestElevatorInformations.referenceGap > gap {
			bestElevatorInformations.bestElevator = _newElevator
			bestElevatorInformations.referenceGap = gap
		}
	}
	return &bestElevatorInformations
}
