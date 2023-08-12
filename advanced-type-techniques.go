package main

import "fmt"

type Position struct {
	x, y int
}

func (p Position) Move(val int) {
	fmt.Println("the player is moved by:", val)
}

func (p *Position) Move(val int) {
	fmt.Println("the player is moved by:", val)
}

type Player struct {
	Position
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	Entity
	specialField int
}

type Color int

// String makes Color satisfy the Stringer interface.
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "Blue"
	case ColorBlack:
		return "Black"
	case ColorYellow:
		return "Yellow"
	case ColorPink:
		return "Pink"
	default:
		panic("invalid colour given")
	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	e := SpecialEntity{
		Entity: Entity{
			name:    "entity",
			id:      "3",
			version: "1.2",
			Position: Position{
				x: 100,
				y: 200,
			},
		},
		specialField: 10,
	}
	fmt.Printf("%+v\n", e)

	fmt.Println("the color is: ", ColorBlack)

	p := Player{}
	p.Move(10)
}
