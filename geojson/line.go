package geojson

type Line struct {
	Coordinates [][]float64
	SrId        int
}

func (l *Line) ToGeometry() *Geometry {
	return NewGeometry(LineType, l.Coordinates, l.SrId)
}
