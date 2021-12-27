package utils

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(X=%d, Y=%d)", p.X, p.Y)
}
