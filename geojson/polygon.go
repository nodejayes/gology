package geojson

// a Polygon Geometry
type Polygon struct {
	// the Coordinates of the Polygon Geometry
	Coordinates [][][]float64
	// the EPSG Code
	SrId int
}

// convert the Polygon into a Geometry
func (p *Polygon) ToGeometry() *Geometry {
	return NewGeometry(PolygonType, p.Coordinates, p.SrId)
}
