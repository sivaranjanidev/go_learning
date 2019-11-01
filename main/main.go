package main

import (
	"fmt"
	"./rover"
)

func main() {
	rover1 := rover.Rover{}
	var instructions string

	rover1.SetRoverPosition()
	fmt.Println("Current position of rover", rover1)

	fmt.Println("Enter the instructions:")
	fmt.Scanf("%s", &instructions)

	rover1.MoveRover(instructions)
	fmt.Println("After Rover Instructions its position is  ", rover1)
}
