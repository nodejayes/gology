package geojson

// a Point Geometry
type IPoint interface {
	IGeometryConvertible
	ISrIdReadProperty
	GetCoordinates() []float64
}

type point struct {
	Coordinates []float64
	SrId        int
}

// the EPSG Code
func (p *point) GetSrId() int {
	return p.SrId
}

// the Coordinate of the Point Geometry
func (p *point) GetCoordinates() []float64 {
	return p.Coordinates
}

// convert the Point into a Geometry
func (p *point) ToGeometry() IGeometry {
	return NewGeometry(PointType, p.Coordinates, p.SrId)
}
