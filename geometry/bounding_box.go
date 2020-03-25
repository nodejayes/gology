package geometry

import "github.com/nodejayes/gology/geojson"

type BoundingBox struct {
	geometry *geojson.Polygon
}

func NewBoundingBox(coordinates [][][]float64, srid int) *BoundingBox {
	geometry, err := geojson.NewPolygon(coordinates, srid)
	if err != nil {
		return nil
	}
	return &BoundingBox{
		geometry: geometry,
	}
}

func (bb *BoundingBox) GetSrId() int {
	return bb.geometry.SrId
}
