package main

type Door struct {
	ID     int
	status string
}

func newDoor(_ID int) *Door {
	door := Door{
		ID:     _ID,
		status: "closed",
	}
	return &door
}
