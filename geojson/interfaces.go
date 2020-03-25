package geojson

type ILocatable interface {
	GetSrId() int
}

type IGeometryConvertible interface {
	ToGeometry() *Geometry
}

type ISerializable interface {
	Serialize() string
}
