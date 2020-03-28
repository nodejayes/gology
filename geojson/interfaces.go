package geojson

type INameReadProperty interface {
	GetName() string
}

type ITypeReadProperty interface {
	GetType() string
}

type ISrIdReadProperty interface {
	GetSrId() int
}

// a Struct that can be serialize into a string
type ISerializable interface {
	Serialize() string
}

type IReferenceSystemReadProperty interface {
	GetCRS() IReferenceSystem
}

// a Struct that can be Converted into a Geometry struct
type IGeometryConvertible interface {
	ToGeometry() IGeometry
}
