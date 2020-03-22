// Package geometry contains geometry objects such as points, lines and polygons
package geometry

import (
	"fmt"
	"math"
)

// ICoordinate represent a Coordinate that can be handled
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

// the type of the Coordinate Dimension DD stands for 2D and DDD for 3D
type CoordinateDimension = int

const (
	DD CoordinateDimension = 2 << iota
	DDD
)

// Coordinate struct
type Coordinate struct {
	x         float64
	y         float64
	z         float64
	dimension CoordinateDimension
}

// create a new Coordinate and identify the Dimension on the z value
// if you need a 2D Coordinate you can pass mathInf(-1) or mathInf(1)
// // coord3D := NewCoordinate(1,2,3)
// // coord2D := NewCoordinate(1,2,math.Inf(-1))
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

// getter for the Value of the X
func (c *Coordinate) GetX() float64 {
	return c.x
}

// setter for the Value of the X
func (c *Coordinate) SetX(value float64) {
	c.x = value
}

// getter for the Value of the Y
func (c *Coordinate) GetY() float64 {
	return c.y
}

// setter for the Value of the Y
func (c *Coordinate) SetY(value float64) {
	c.y = value
}

// getter for the Value of the Z
func (c *Coordinate) GetZ() float64 {
	return c.z
}

// setter for the Value of the Z
func (c *Coordinate) SetZ(value float64) {
	c.z = value
	c.dimension = DDD
	if math.IsInf(value, -1) || math.IsInf(value, 1) {
		c.dimension = DD
	}
}

// get all 3 Dimensions in a Slice x,y,z
func (c *Coordinate) GetValues() []float64 {
	return []float64{c.x, c.y, c.z}
}

// get the Coordinate Dimension
func (c *Coordinate) GetDimension() CoordinateDimension {
	return c.dimension
}

// create a new Instance of the current Coordinate
func (c *Coordinate) Copy() *Coordinate {
	return &Coordinate{
		x:         c.x,
		y:         c.y,
		z:         c.z,
		dimension: c.dimension,
	}
}

// take the Values of the given Coordinate and writes them into this Coordinate
func (c *Coordinate) TakeOverCoordinate(value *Coordinate) {
	c.x = value.x
	c.y = value.y
	c.z = value.z
	c.dimension = value.dimension
}

// compare this Coordinate with the given Coordinate on all 3 Dimensions
func (c *Coordinate) Equals(other ICoordinate) bool {
	return c.GetX() == other.GetX() && c.GetY() == other.GetY() && c.GetZ() == other.GetZ()
}

// see Equals but with a accepted Tolerance
func (c *Coordinate) EqualsTolerance(other ICoordinate, tolerance float64) bool {
	equalX := valueTolerance(c.GetX(), other.GetX(), tolerance)
	equalY := valueTolerance(c.GetY(), other.GetY(), tolerance)
	equalZ := valueTolerance(c.GetZ(), other.GetZ(), tolerance)
	return equalX && equalY && equalZ
}

// compare this Coordinate with the given Coordinate on 2 Dimensions x,y
func (c *Coordinate) Equals2D(other ICoordinate) bool {
	return c.GetX() == other.GetX() && c.GetY() == other.GetY()
}

// see Equals2D but with a accepted Tolerance
func (c *Coordinate) Equals2DTolerance(other ICoordinate, tolerance float64) bool {
	equalX := valueTolerance(c.GetX(), other.GetX(), tolerance)
	equalY := valueTolerance(c.GetY(), other.GetY(), tolerance)
	return equalX && equalY
}

// give the Sort Order of this Coordinate again the given Coordinate
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

// calculates the flat Distance from this Coordinate to the given Coordinate
func (c *Coordinate) Distance(other ICoordinate) float64 {
	dx := c.x - other.GetX()
	dy := c.y - other.GetY()
	if c.dimension == DDD && other.GetDimension() == DDD {
		dz := c.z - other.GetZ()
		return math.Sqrt((dx * dx) + (dy * dy) + (dz * dz))
	}
	return math.Sqrt(dx*dx + dy*dy)
}

// converts the Coordinate to a readable String
func (c *Coordinate) ToString() string {
	dimensionString := "2D"
	if c.dimension == DDD {
		dimensionString = "3D"
	}
	return fmt.Sprintf("(X: %v, Y: %v, Z: %v) %v", c.x, c.y, c.z, dimensionString)
}

func valueTolerance(v1, v2, tolerance float64) bool {
	if v1 < v2 {
		return v2-v1 <= tolerance
	}
	return v1-v2 <= tolerance
}
