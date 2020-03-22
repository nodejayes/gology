package geometries

type GeometryType string

const (
	PointType      GeometryType = "Point"
	MultiPointType              = "MultiPoint"
	LineType                    = "LineString"
	MultiLineType               = "MultiLineString"
	PolygonType                 = "Polygon"
	MultiPolygon                = "MultiPolygon"
)
