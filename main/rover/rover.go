package rover

import (
	"fmt"
	"strings"
)

type Rover struct {
	xPosition        int
	yPosition        int
	compassDirection string
}

func (rover *Rover) SetRoverPosition() {
	fmt.Println("Enter X Coordinate")
	fmt.Scan(&(*rover).xPosition)
	fmt.Println("Enter Y Coordinate")
	fmt.Scan(&(*rover).yPosition)
	fmt.Println("Enter direction")
	fmt.Scan(&(*rover).compassDirection)
}

func (rover *Rover) rotateLeft() {
	switch rover.compassDirection {
	case "N":
		rover.compassDirection = "W"
	case "W":
		rover.compassDirection = "S"
	case "S":
		rover.compassDirection = "E"
	case "E":
		rover.compassDirection = "N"
	}
}

func (rover *Rover) rotateRight() {
	switch rover.compassDirection {
	case "N":
		rover.compassDirection = "E"
	case "W":
		rover.compassDirection = "N"
	case "S":
		rover.compassDirection = "W"
	case "E":
		rover.compassDirection = "S"
	}
}

func (rover *Rover) moveRoverForward() {
	switch rover.compassDirection {
	case "N":
		rover.yPosition++
	case "W":
		rover.xPosition--
	case "S":
		rover.yPosition--
	case "E":
		rover.xPosition++
	}
}

func (rover *Rover) MoveRover(instructions string) {
	for _, value := range instructions {
		movement := strings.ToUpper(string(value))
		switch string(movement) {
		case "L":
			rover.rotateLeft()
		case "R":
			rover.rotateRight()
		case "M":
			rover.moveRoverForward()
		}
	}
}
