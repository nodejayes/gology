package geojson

type MultiPolygon struct {
	Coordinates [][][][]float64
	SrId        int
}

func (mp *MultiPolygon) ToGeometry() *Geometry {
	return NewGeometry(MultiPolygonType, mp.Coordinates, mp.SrId)
}
