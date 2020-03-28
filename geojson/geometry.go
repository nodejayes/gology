package geojson

import (
	"encoding/json"
)

// GeometryTypes are the Type strings for the GeoJSON Geometry Types like Point, Line Polygon and so on.
type GeometryTypes string

const (
	// the GeoJSON Geometry Point Type
	PointType GeometryTypes = "Point"
	// the GeoJSON Geometry Line Type
	LineType GeometryTypes = "LineString"
	// the GeoJSON Geometry Polygon Type
	PolygonType GeometryTypes = "Polygon"
	// the GeoJSON Geometry MultiPoint Type
	MultiPointType GeometryTypes = "MultiPoint"
	// the GeoJSON Geometry MultiLine Type
	MultiLineType GeometryTypes = "MultiLineString"
	// the GeoJSON Geometry MultiPolygon Type
	MultiPolygonType GeometryTypes = "MultiPolygon"
)

type IGeometry interface {
	GetType() GeometryTypes
	GetCoordinates() ICoordinate
	GetCRS() IReferenceSystem
	Serialize() string
	AsPoint() (IPoint, error)
	AsLine() (ILine, error)
	AsPolygon() (IPolygon, error)
	AsMultiPoint() (IMultiPoint, error)
	AsMultiLine() (IMultiLine, error)
	AsMultiPolygon() (IMultiPolygon, error)
}

// the Definition of a abstract Geometry this can be a Point, Line , Polygon or MultiPoint, MultiLine, MultiPolygon
type geometry struct {
	// the Type of the Geometry
	Type GeometryTypes `json:"type"`
	// the Coordinates as float64 Array with N Dimensions represented as interface{}
	Coordinates ICoordinate `json:"coordinates"`
	// a optional Coordinate System Info was removed on Serialization when is nil
	CRS IReferenceSystem `json:"crs,omitempty"`
}

// create a new Geometry Object with a Type Coordinates and the EPSG Code
func NewGeometry(typ GeometryTypes, coordinates interface{}, srId int) IGeometry {
	return &geometry{
		Type:        typ,
		Coordinates: NewCoordinate(coordinates),
		CRS:         NewReferenceSystem(srId),
	}
}

// create a new Point from Coordinates and EPSG Code
func NewPoint(coordinate []float64, srId int) (IPoint, error) {
	geom := &geometry{
		Type:        PointType,
		Coordinates: NewCoordinate(coordinate),
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsPoint()
}

// create a new Line from Coordinates and EPSG Code
func NewLine(coordinates [][]float64, srId int) (ILine, error) {
	geom := &geometry{
		Type:        LineType,
		Coordinates: NewCoordinate(coordinates),
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsLine()
}

// create a new Polygon from Coordinates and EPSG Code
func NewPolygon(coordinates [][][]float64, srId int) (IPolygon, error) {
	geom := &geometry{
		Type:        PolygonType,
		Coordinates: NewCoordinate(coordinates),
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsPolygon()
}

// create a new MultiPoint from Coordinates and EPSG Code
func NewMultiPoint(coordinates [][]float64, srId int) (IMultiPoint, error) {
	geom := &geometry{
		Type:        MultiPointType,
		Coordinates: NewCoordinate(coordinates),
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiPoint()
}

// create a new MultiLine from Coordinates and EPSG Code
func NewMultiLine(coordinates [][][]float64, srId int) (IMultiLine, error) {
	geom := &geometry{
		Type:        MultiLineType,
		Coordinates: NewCoordinate(coordinates),
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiLine()
}

// create a new MultiPolygon from Coordinates and EPSG Code
func NewMultiPolygon(coordinates [][][][]float64, srId int) (IMultiPolygon, error) {
	geom := &geometry{
		Type:        MultiPolygonType,
		Coordinates: NewCoordinate(coordinates),
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiPolygon()
}

// try to Deserialize a string into a Geometry when the string is not a valid GeoJSON String nil was returned
func DeserializeGeometry(str string) IGeometry {
	var geom *geometry
	err := json.Unmarshal([]byte(str), &geom)
	if err != nil {
		return nil
	}
	return geom
}

// get the Type of the Geometry
func (g *geometry) GetType() GeometryTypes {
	return g.Type
}

func (g *geometry) GetCoordinates() ICoordinate {
	return g.Coordinates
}

func (g *geometry) GetCRS() IReferenceSystem {
	return g.CRS
}

// serialize the current Geometry into a GeoJSON String
func (g *geometry) Serialize() string {
	stream, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(stream)
}

// get the EPSG Code from the Coordinate Reference System when the Geometry has some. If not 0 was returned
func (g *geometry) GetSrId() int {
	if g.CRS == nil {
		return 0
	}
	return g.CRS.GetSrId()
}

// try to convert the abstract Geometry into a Point
func (g *geometry) AsPoint() (IPoint, error) {
	coords, err := g.Coordinates.AsPoint()
	if err != nil {
		return nil, err
	}
	return &point{
		Coordinates: coords,
		SrId:        g.GetSrId(),
	}, nil
}

// try to convert the abstract Geometry into a Line
func (g *geometry) AsLine() (ILine, error) {
	coords, err := g.Coordinates.AsLine()
	if err != nil {
		return nil, err
	}
	return &line{
		Coordinates: coords,
		SrId:        g.GetSrId(),
	}, nil
}

// try to convert the abstract Geometry into a Polygon
func (g *geometry) AsPolygon() (IPolygon, error) {
	coords, err := g.Coordinates.AsPolygon()
	if err != nil {
		return nil, err
	}
	return &polygon{
		Coordinates: coords,
		SrId:        g.GetSrId(),
	}, nil
}

// try to convert the abstract Geometry into a MultiPoint
func (g *geometry) AsMultiPoint() (IMultiPoint, error) {
	coords, err := g.Coordinates.AsMultiPoint()
	if err != nil {
		return nil, err
	}
	return &multiPoint{
		Coordinates: coords,
		SrId:        g.GetSrId(),
	}, nil
}

// try to convert the abstract Geometry into a MultiLine
func (g *geometry) AsMultiLine() (IMultiLine, error) {
	coords, err := g.Coordinates.AsMultiLine()
	if err != nil {
		return nil, err
	}
	return &multiLine{
		Coordinates: coords,
		SrId:        g.GetSrId(),
	}, nil
}

// try to convert the abstract Geometry into a MultiPolygon
func (g *geometry) AsMultiPolygon() (IMultiPolygon, error) {
	coords, err := g.Coordinates.AsMultiPolygon()
	if err != nil {
		return nil, err
	}
	return &multiPolygon{
		Coordinates: coords,
		SrId:        g.GetSrId(),
	}, nil
}

// check the current Geometry Type for the Point Type
func (g *geometry) IsPoint() bool {
	return g.Type == PointType
}

// check the current Geometry Type for the Line Type
func (g *geometry) IsLine() bool {
	return g.Type == LineType
}

// check the current Geometry Type for the Polygon Type
func (g *geometry) IsPolygon() bool {
	return g.Type == PolygonType
}

// check the current Geometry Type for the MultiPoint Type
func (g *geometry) IsMultiPoint() bool {
	return g.Type == MultiPointType
}

// check the current Geometry Type for the MultiLine Type
func (g *geometry) IsMultiLine() bool {
	return g.Type == MultiLineType
}

// check the current Geometry Type for the MultiPolygon Type
func (g *geometry) IsMultiPolygon() bool {
	return g.Type == MultiPolygonType
}

func forceFloat64Array1D(input []interface{}) []float64 {
	var tmp []float64
	for _, element := range input {
		switch c := element.(type) {
		case float64:
			tmp = append(tmp, c)
		}
	}
	return tmp
}

func forceFloat64Array2D(input []interface{}) [][]float64 {
	var tmp [][]float64
	for _, element := range input {
		switch c := element.(type) {
		case []interface{}:
			tmp = append(tmp, forceFloat64Array1D(c))
		}
	}
	return tmp
}

func forceFloat64Array3D(input []interface{}) [][][]float64 {
	var tmp [][][]float64
	for _, element := range input {
		switch c := element.(type) {
		case []interface{}:
			tmp = append(tmp, forceFloat64Array2D(c))
		}
	}
	return tmp
}

func forceFloat64Array4D(input []interface{}) [][][][]float64 {
	var tmp [][][][]float64
	for _, element := range input {
		switch c := element.(type) {
		case []interface{}:
			tmp = append(tmp, forceFloat64Array3D(c))
		}
	}
	return tmp
}
