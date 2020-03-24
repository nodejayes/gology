package geojson

type Point struct {
	Coordinates []float64
	SrId        int
}

func (p *Point) ToGeometry() *Geometry {
	return NewGeometry(PointType, p.Coordinates, p.SrId)
}
