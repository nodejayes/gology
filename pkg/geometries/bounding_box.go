package geometries

type BoundingBox struct {
	TopLeft     *Point
	TopRight    *Point
	BottomRight *Point
	BottomLeft  *Point
}

func NewBoundingBoxFromPoints(topLeft *Point, topRight *Point, bottomRight *Point, bottomLeft *Point) *BoundingBox {
	return &BoundingBox{
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomRight: bottomRight,
		BottomLeft:  bottomLeft,
	}
}
