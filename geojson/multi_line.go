package geojson

// a Multi Line Geometry
type IMultiLine interface {
	IGeometryConvertible
	ISrIdReadProperty
	GetCoordinates() [][][]float64
}

type multiLine struct {
	Coordinates [][][]float64
	SrId        int
}

// the EPSG Code
func (ml *multiLine) GetSrId() int {
	return ml.SrId
}

// the Coordinates of the Multi Line String
func (ml *multiLine) GetCoordinates() [][][]float64 {
	return ml.Coordinates
}

// convert the Multi Line into a Geometry
func (ml *multiLine) ToGeometry() IGeometry {
	return NewGeometry(MultiLineType, ml.Coordinates, ml.SrId)
}
