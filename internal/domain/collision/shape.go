package collision

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}
