package geojson

type MultiPoint struct {
	Coordinates [][]float64
	SrId        int
}

func (mp *MultiPoint) ToGeometry() *Geometry {
	return NewGeometry(MultiPointType, mp.Coordinates, mp.SrId)
}
