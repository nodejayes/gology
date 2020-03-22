package geometry

type Type string

const (
	PointType      Type = "Point"
	MultiPointType      = "MultiPoint"
	LineType            = "LineString"
	MultiLineType       = "MultiLineString"
	PolygonType         = "Polygon"
	MultiPolygon        = "MultiPolygon"
)
