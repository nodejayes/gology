package geojson

// a Polygon Geometry
type IPolygon interface {
	IGeometryConvertible
	ISrIdReadProperty
	GetCoordinates() [][][]float64
}

type polygon struct {
	Coordinates [][][]float64
	SrId        int
}

// the EPSG Code
func (p *polygon) GetSrId() int {
	return p.SrId
}

// the Coordinates of the Polygon Geometry
func (p *polygon) GetCoordinates() [][][]float64 {
	return p.Coordinates
}

// convert the Polygon into a Geometry
func (p *polygon) ToGeometry() IGeometry {
	return NewGeometry(PolygonType, p.Coordinates, p.SrId)
}
