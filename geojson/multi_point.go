package geojson

// a Multi Point Geometry
type MultiPoint struct {
	// the Coordinates of the Multi Point
	Coordinates [][]float64
	// the EPSG Code
	SrId int
}

// convert the Multi Line into a Geometry
func (mp *MultiPoint) ToGeometry() *Geometry {
	return NewGeometry(MultiPointType, mp.Coordinates, mp.SrId)
}
