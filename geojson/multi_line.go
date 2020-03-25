package geojson

// a Multi Line Geometry
type MultiLine struct {
	// the Coordinates of the Multi Line String
	Coordinates [][][]float64
	// the EPSG Code
	SrId int
}

// convert the Multi Line into a Geometry
func (ml *MultiLine) ToGeometry() *Geometry {
	return NewGeometry(MultiLineType, ml.Coordinates, ml.SrId)
}
