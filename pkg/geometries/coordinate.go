package geometries

import (
	"fmt"
	"math"
)

type ICoordinate interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	SetX(float64)
	SetY(float64)
	SetZ(float64)
	GetValues() []float64
	GetDimension() CoordinateDimension
}

type CoordinateDimension = int

const (
	DD CoordinateDimension = 2 << iota
	DDD
)

type Coordinate struct {
	x         float64
	y         float64
	z         float64
	dimension CoordinateDimension
}

func NewCoordinate(x, y, z float64) *Coordinate {
	d := DDD
	if math.IsInf(z, -1) || math.IsInf(z, 1) {
		d = DD
	}
	return &Coordinate{
		x:         x,
		y:         y,
		z:         z,
		dimension: d,
	}
}

func (c *Coordinate) GetX() float64 {
	return c.x
}

func (c *Coordinate) SetX(value float64) {
	c.x = value
}

func (c *Coordinate) GetY() float64 {
	return c.y
}

func (c *Coordinate) SetY(value float64) {
	c.y = value
}

func (c *Coordinate) GetZ() float64 {
	return c.z
}

func (c *Coordinate) SetZ(value float64) {
	c.z = value
	c.dimension = DDD
	if math.IsInf(value, -1) || math.IsInf(value, 1) {
		c.dimension = DD
	}
}

func (c *Coordinate) GetValues() []float64 {
	return []float64{c.x, c.y, c.z}
}

func (c *Coordinate) GetDimension() CoordinateDimension {
	return c.dimension
}

func (c *Coordinate) Copy() *Coordinate {
	return &Coordinate{
		x:         c.x,
		y:         c.y,
		z:         c.z,
		dimension: c.dimension,
	}
}

func (c *Coordinate) TakeOverCoordinate(value *Coordinate) {
	c.x = value.x
	c.y = value.y
	c.z = value.z
	c.dimension = value.dimension
}

func (c *Coordinate) Equals(other ICoordinate) bool {
	return c.GetX() == other.GetX() && c.GetY() == other.GetY() && c.GetZ() == other.GetZ()
}

func (c *Coordinate) EqualsTolerance(other ICoordinate, tolerance float64) bool {
	equalX := valueTolerance(c.GetX(), other.GetX(), tolerance)
	equalY := valueTolerance(c.GetY(), other.GetY(), tolerance)
	equalZ := valueTolerance(c.GetZ(), other.GetZ(), tolerance)
	return equalX && equalY && equalZ
}

func (c *Coordinate) Equals2D(other ICoordinate) bool {
	return c.GetX() == other.GetX() && c.GetY() == other.GetY()
}

func (c *Coordinate) Equals2DTolerance(other ICoordinate, tolerance float64) bool {
	equalX := valueTolerance(c.GetX(), other.GetX(), tolerance)
	equalY := valueTolerance(c.GetY(), other.GetY(), tolerance)
	return equalX && equalY
}

func (c *Coordinate) CompareTo(other ICoordinate) int {
	if c.GetX() < other.GetX() {
		return -1
	}
	if c.GetX() > other.GetX() {
		return 1
	}
	if c.GetY() < other.GetY() {
		return -1
	}
	if c.GetY() > other.GetY() {
		return 1
	}
	if c.GetZ() < other.GetZ() {
		return -1
	}
	if c.GetZ() > other.GetZ() {
		return 1
	}
	return 0
}

func (c *Coordinate) Distance(other ICoordinate) float64 {
	dx := c.x - other.GetX()
	dy := c.y - other.GetY()
	if c.dimension == DDD && other.GetDimension() == DDD {
		dz := c.z - other.GetZ()
		return math.Sqrt((dx * dx) + (dy * dy) + (dz * dz))
	}
	return math.Sqrt(dx*dx + dy*dy)
}

func (c *Coordinate) ToString() string {
	return fmt.Sprintf("(X: %v, Y: %v, Z: %v)", c.x, c.y, c.z)
}

func valueTolerance(v1, v2, tolerance float64) bool {
	if v1 < v2 {
		return v2-v1 <= tolerance
	}
	return v1-v2 <= tolerance
}
