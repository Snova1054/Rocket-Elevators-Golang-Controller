package main

type Door struct {
	ID     int
	status string
}

func (d Door) newDoor(_ID int) *Door {
	d.ID = _ID
	d.status = "idle"
	return &d
}
