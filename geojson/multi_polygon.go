package geojson

// a Multi Polygon Geometry
type IMultiPolygon interface {
	IGeometryConvertible
	ISrIdReadProperty
	GetCoordinates() [][][][]float64
}

type multiPolygon struct {
	Coordinates [][][][]float64
	SrId        int
}

// the EPSG Code
func (mp *multiPolygon) GetSrId() int {
	return mp.SrId
}

// the Coordinates of the Multi Polygon Geometry
func (mp *multiPolygon) GetCoordinates() [][][][]float64 {
	return mp.Coordinates
}

// convert the Multi Polygon into a Geometry
func (mp *multiPolygon) ToGeometry() IGeometry {
	return NewGeometry(MultiPolygonType, mp.Coordinates, mp.SrId)
}
