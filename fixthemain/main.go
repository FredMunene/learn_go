package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

const (
	OPEN  = "open"
	CLOSE = "closed"
)

func main() {
	door := &Door{}

	OpenDoor(door)

	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) bool {
	if ptrDoor == nil {
		return false
	}
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	if ptrDoor == nil {
		return false
	}
	PrintStr("is the Door opened ?")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	if ptrDoor == nil {
		return false
	}
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func OpenDoor(ptrDoor *Door) bool {
	if ptrDoor == nil {
		return false
	}
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}
