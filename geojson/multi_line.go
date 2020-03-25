package geojson

type MultiLine struct {
	Coordinates [][][]float64
	SrId        int
}

func (ml *MultiLine) ToGeometry() *Geometry {
	return NewGeometry(MultiLineType, ml.Coordinates, ml.SrId)
}
