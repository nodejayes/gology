package geojson

// a Multi Point Geometry
type IMultiPoint interface {
	IGeometryConvertible
	ISrIdReadProperty
	GetCoordinates() [][]float64
}

type multiPoint struct {
	Coordinates [][]float64
	SrId        int
}

// the EPSG Code
func (mp *multiPoint) GetSrId() int {
	return mp.SrId
}

// the Coordinates of the Multi Point
func (mp *multiPoint) GetCoordinates() [][]float64 {
	return mp.Coordinates
}

// convert the Multi Line into a Geometry
func (mp *multiPoint) ToGeometry() IGeometry {
	return NewGeometry(MultiPointType, mp.Coordinates, mp.SrId)
}
