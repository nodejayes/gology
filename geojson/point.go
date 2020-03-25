package geojson

// a Point Geometry
type Point struct {
	// the Coordinate of the Point Geometry
	Coordinates []float64
	// the EPSG Code
	SrId int
}

// convert the Point into a Geometry
func (p *Point) ToGeometry() *Geometry {
	return NewGeometry(PointType, p.Coordinates, p.SrId)
}
