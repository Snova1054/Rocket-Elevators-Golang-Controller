package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

//Function used to create new NewFloorRequestButtons with the desired properties
func NewFloorRequestButton(_ID int, _floor int, _direction string) *FloorRequestButton {
	floorRequestButton := FloorRequestButton{ID: _ID, status: "off", floor: _floor, direction: _direction}
	return &floorRequestButton
}
