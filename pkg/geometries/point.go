package geometries

type Point struct {
	coordinate *Coordinate
	srid       int
}

func NewPointFromLatLong(latitude, longitude, altitude float64) *Point {
	return &Point{
		coordinate: NewCoordinateFromXYZ(latitude, longitude, altitude),
		srid:       4326,
	}
}

func NewPointFromXYZ(x, y, z float64, srid int) *Point {
	return &Point{
		coordinate: NewCoordinateFromXYZ(x, y, z),
		srid:       srid,
	}
}

func (p *Point) GetX() float64 {
	return p.coordinate.x
}

func (p *Point) GetY() float64 {
	return p.coordinate.y
}

func (p *Point) GetZ() float64 {
	return p.coordinate.z
}

func (p *Point) SetX(value float64) {
	p.coordinate.x = value
}

func (p *Point) SetY(value float64) {
	p.coordinate.y = value
}

func (p *Point) SetZ(value float64) {
	p.coordinate.z = value
}

func (p *Point) GetCoordinate() []float64 {
	return p.coordinate.GetCoordinate()
}
