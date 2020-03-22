package geometries

import (
	"bytes"
	"math"
	"strconv"
)

type CoordinateDimension = int

const (
	DD CoordinateDimension = iota
	DDD
)

func equalsWithTolerance(v1, v2, tolerance float64) bool {
	return math.Abs(v1-v2) <= tolerance
}

type Coordinate struct {
	x         float64
	y         float64
	z         float64
	dimension CoordinateDimension
}

func NewCoordinateFromXY(x, y float64) *Coordinate {
	return &Coordinate{
		x:         x,
		y:         y,
		z:         0,
		dimension: DD,
	}
}

func NewCoordinateFromXYZ(x, y, z float64) *Coordinate {
	return &Coordinate{
		x:         x,
		y:         y,
		z:         z,
		dimension: DDD,
	}
}

func NewCoordinateFromCoordinate(c *Coordinate) *Coordinate {
	return &Coordinate{
		x: c.x,
		y: c.y,
		z: c.z,
	}
}

func (c *Coordinate) GetCoordinateValue() *Coordinate {
	return c
}

func (c *Coordinate) GetCoordinate() []float64 {
	return []float64{c.x, c.y, c.z}
}

func (c *Coordinate) SetCoordinateValue(value *Coordinate) {
	c.x = value.x
	c.y = value.y
}

func (c *Coordinate) Equals2D(other *Coordinate) bool {
	return c.x == other.x && c.y == other.y
}

func (c *Coordinate) Equals2DTolerance(other *Coordinate, tolerance float64) bool {
	if !equalsWithTolerance(c.x, other.x, tolerance) {
		return false
	}
	if !equalsWithTolerance(c.y, other.y, tolerance) {
		return false
	}
	return true
}

func (c *Coordinate) EqualsCoordinate(other *Coordinate) bool {
	return c.Equals2D(other)
}

func (c *Coordinate) CompareToCoordinate(other *Coordinate) int {
	if other == nil {
		panic("Coordinate.CompareToCoordinate => other is null")
	}

	if c.x < other.x {
		return -1
	}
	if c.x > other.x {
		return 1
	}
	if c.y < other.y {
		return -1
	}
	if c.y > other.y {
		return 1
	}
	return 0
}

func (c *Coordinate) CompareTo(other interface{}) int {
	switch coord := other.(type) {
	case Coordinate:
		return c.CompareToCoordinate(&coord)
	case *Coordinate:
		return c.CompareToCoordinate(coord)
	}
	panic("invalid type expected Coordinate or Pointer of Coordinate")
}

func (c *Coordinate) Copy() *Coordinate {
	return NewCoordinateFromCoordinate(c)
}

func (c *Coordinate) Distance(other *Coordinate) float64 {
	dx := c.x - other.x
	dy := c.y - other.y
	if c.dimension == DDD {
		dz := c.z - other.z
		return math.Sqrt((dx * dx) + (dy * dy) + (dz * dz))
	}
	return math.Sqrt(dx*dx + dy*dy)
}

func (c *Coordinate) Equals(other interface{}) bool {
	switch coord := other.(type) {
	case Coordinate:
		return c.EqualsCoordinate(&coord)
	case *Coordinate:
		return c.EqualsCoordinate(coord)
	}
	return false
}

func (c *Coordinate) ToString() string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("(")
	buf.WriteString(strconv.FormatFloat(c.x, 'f', -1, 64))
	buf.WriteString(",")
	buf.WriteString(strconv.FormatFloat(c.y, 'f', -1, 64))
	if c.dimension == DDD {
		buf.WriteString(",")
		buf.WriteString(strconv.FormatFloat(c.z, 'f', -1, 64))
	}
	buf.WriteString(")")
	return buf.String()
}
