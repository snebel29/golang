package main

import (
	"fmt"
	. "github.com/snebel29/golang/mars-rover/rover"
)

func main() {
	fmt.Println("Starting rover mission!")

	mars := Map{Height: 100, Width: 100}
	r1 := mars.NewRover("Rover one", Position{0, Coordinates{X: 0, Y: 0}})

	commands := "ffrfflff"
	r1.RunCommands(commands)
}


