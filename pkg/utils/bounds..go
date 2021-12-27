package utils

type Bounds struct {
	TopLeft     Point
	BottomRight Point
}

func (b Bounds) Contains(p Point) bool {
	return p.X >= b.TopLeft.X && p.X <= b.BottomRight.X && p.Y <= b.TopLeft.Y && p.Y >= b.BottomRight.Y
}
