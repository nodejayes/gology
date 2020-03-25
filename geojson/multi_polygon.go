package geojson

// a Multi Polygon Geometry
type MultiPolygon struct {
	// the Coordinates of the Multi Polygon Geometry
	Coordinates [][][][]float64
	// the EPSG Code
	SrId int
}

// convert the Multi Polygon into a Geometry
func (mp *MultiPolygon) ToGeometry() *Geometry {
	return NewGeometry(MultiPolygonType, mp.Coordinates, mp.SrId)
}
