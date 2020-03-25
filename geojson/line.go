package geojson

// a Line Geometry
type Line struct {
	// the Coordinates of the Line
	Coordinates [][]float64
	// the EPSG Code
	SrId int
}

// convert the Line to a Geometry
func (l *Line) ToGeometry() *Geometry {
	return NewGeometry(LineType, l.Coordinates, l.SrId)
}
