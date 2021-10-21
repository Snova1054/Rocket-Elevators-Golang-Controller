package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_ID int) *Door {
	door := Door{ID: _ID, status: "closed"}
	return &door
}
