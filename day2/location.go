package main

import "fmt"

type Location struct {
	Horizontal int
	Depth      int
}

func (location *Location) goUp(steps int) {
	location.Depth = location.Depth - steps
}

func (location *Location) goDown(steps int) {
	location.Depth = location.Depth + steps
}

func (location *Location) goForward(steps int) {
	location.Horizontal = location.Horizontal + steps
}

func main() {
	loc := Location{0, 0}
	for _, dir := range Input {
		switch dir.Direction {
		case "forward":
			loc.goForward(dir.Speed)
		case "up":
			loc.goUp(dir.Speed)
		case "down":
			loc.goDown(dir.Speed)
		default:
			fmt.Println("Incorrect direction: ", dir.Direction)
		}
	}
	fmt.Println(loc)
	fmt.Println("Final location ", (loc.Horizontal * loc.Depth))
}
