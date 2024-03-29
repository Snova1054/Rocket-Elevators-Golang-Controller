package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

//Function used to create new CallButtons with the desired properties
func NewCallButton(_ID int, _floor int, _direction string) *CallButton {
	callButton := CallButton{ID: _ID, status: "off", floor: _floor, direction: _direction}
	return &callButton
}
