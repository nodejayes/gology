package geojson

// a Struct that have a Coordinate Reference System with a EPSG Code
type ILocatable interface {
	GetSrId() int
}

// a Struct that can be Converted into a Geometry struct
type IGeometryConvertible interface {
	ToGeometry() *Geometry
}

// a Struct that can be serialize into a string
type ISerializable interface {
	Serialize() string
}
