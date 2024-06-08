package internal

type Camera struct {
	x         int
	y         int
	targetX   int
	targetY   int
	xDrag     int
	yDrag     int
	smoothing float64
}

func (c *Camera) Update() {

}
