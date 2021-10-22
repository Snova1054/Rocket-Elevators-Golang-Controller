package main

//Declares each Door
type Door struct {
	ID     int
	status string
}

//Function used to create new Doors with the desired properties
func NewDoor(_ID int) *Door {
	door := Door{ID: _ID, status: "closed"}
	return &door
}
