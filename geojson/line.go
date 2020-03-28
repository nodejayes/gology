package geojson

// a Line Geometry
type ILine interface {
	IGeometryConvertible
	ISrIdReadProperty
	GetCoordinates() [][]float64
}

type line struct {
	Coordinates [][]float64
	SrId        int
}

// the EPSG Code
func (l *line) GetSrId() int {
	return l.SrId
}

// the Coordinates of the Line
func (l *line) GetCoordinates() [][]float64 {
	return l.Coordinates
}

// convert the Line to a Geometry
func (l *line) ToGeometry() IGeometry {
	return NewGeometry(LineType, l.Coordinates, l.SrId)
}
