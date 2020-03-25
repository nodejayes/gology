package geojson

type Polygon struct {
	Coordinates [][][]float64
	SrId        int
}

func (p *Polygon) ToGeometry() *Geometry {
	return NewGeometry(PolygonType, p.Coordinates, p.SrId)
}
