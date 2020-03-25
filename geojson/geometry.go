package geojson

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
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

// the Definition of a abstract Geometry this can be a Point, Line , Polygon or MultiPoint, MultiLine, MultiPolygon
type Geometry struct {
	// the Type of the Geometry
	Type GeometryTypes `json:"type"`
	// the Coordinates as float64 Array with N Dimensions represented as interface{}
	Coordinates interface{} `json:"coordinates"`
	// a optional Coordinate System Info was removed on Serialization when is nil
	CRS *ReferenceSystem `json:"crs,omitempty"`
}

// create a new Geometry Object with a Type Coordinates and the EPSG Code
func NewGeometry(typ GeometryTypes, coordinates interface{}, srId int) *Geometry {
	return &Geometry{
		Type:        typ,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
}

// create a new Point from Coordinates and EPSG Code
func NewPoint(coordinate []float64, srId int) (*Point, error) {
	geom := &Geometry{
		Type:        PointType,
		Coordinates: coordinate,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsPoint()
}

// create a new Line from Coordinates and EPSG Code
func NewLine(coordinates [][]float64, srId int) (*Line, error) {
	geom := &Geometry{
		Type:        LineType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsLine()
}

// create a new Polygon from Coordinates and EPSG Code
func NewPolygon(coordinates [][][]float64, srId int) (*Polygon, error) {
	geom := &Geometry{
		Type:        PolygonType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsPolygon()
}

// create a new MultiPoint from Coordinates and EPSG Code
func NewMultiPoint(coordinates [][]float64, srId int) (*MultiPoint, error) {
	geom := &Geometry{
		Type:        MultiPointType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiPoint()
}

// create a new MultiLine from Coordinates and EPSG Code
func NewMultiLine(coordinates [][][]float64, srId int) (*MultiLine, error) {
	geom := &Geometry{
		Type:        MultiLineType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiLine()
}

// create a new MultiPolygon from Coordinates and EPSG Code
func NewMultiPolygon(coordinates [][][][]float64, srId int) (*MultiPolygon, error) {
	geom := &Geometry{
		Type:        MultiPolygonType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiPolygon()
}

// try to Deserialize a string into a Geometry when the string is not a valid GeoJSON String nil was returned
func DeserializeGeometry(str string) *Geometry {
	var geom *Geometry
	err := jsoniter.Unmarshal([]byte(str), &geom)
	if err != nil {
		return nil
	}
	return geom
}

// serialize the current Geometry into a GeoJSON String
func (g *Geometry) Serialize() string {
	stream, err := jsoniter.Marshal(g)
	if err != nil {
		return ""
	}
	return string(stream)
}

// get the EPSG Code from the Coordinate Reference System when the Geometry has some. If not 0 was returned
func (g *Geometry) GetSrId() int {
	if g.CRS == nil {
		return 0
	}
	return g.CRS.GetSrId()
}

// try to convert the abstract Geometry into a Point
func (g *Geometry) AsPoint() (*Point, error) {
	if g.IsPoint() {
		switch coords := g.Coordinates.(type) {
		case []float64:
			return &Point{
				Coordinates: coords,
				SrId:        g.GetSrId(),
			}, nil
		case []interface{}:
			return &Point{
				Coordinates: forceFloat64Array1D(coords),
				SrId:        g.GetSrId(),
			}, nil
		}
	}
	return nil, errors.New("can't convert to Point")
}

// try to convert the abstract Geometry into a Line
func (g *Geometry) AsLine() (*Line, error) {
	if g.IsLine() {
		switch coords := g.Coordinates.(type) {
		case [][]float64:
			return &Line{
				Coordinates: coords,
				SrId:        g.GetSrId(),
			}, nil
		case []interface{}:
			return &Line{
				Coordinates: forceFloat64Array2D(coords),
				SrId:        g.GetSrId(),
			}, nil
		}
	}
	return nil, errors.New("can't convert to Line")
}

// try to convert the abstract Geometry into a Polygon
func (g *Geometry) AsPolygon() (*Polygon, error) {
	if g.IsPolygon() {
		switch coords := g.Coordinates.(type) {
		case [][][]float64:
			return &Polygon{
				Coordinates: coords,
				SrId:        g.GetSrId(),
			}, nil
		case []interface{}:
			return &Polygon{
				Coordinates: forceFloat64Array3D(coords),
				SrId:        g.GetSrId(),
			}, nil
		}
	}
	return nil, errors.New("can't convert to Polygon")
}

// try to convert the abstract Geometry into a MultiPoint
func (g *Geometry) AsMultiPoint() (*MultiPoint, error) {
	if g.IsMultiPoint() {
		switch coords := g.Coordinates.(type) {
		case [][]float64:
			return &MultiPoint{
				Coordinates: coords,
				SrId:        g.GetSrId(),
			}, nil
		case []interface{}:
			return &MultiPoint{
				Coordinates: forceFloat64Array2D(coords),
				SrId:        g.GetSrId(),
			}, nil
		}
	}
	return nil, errors.New("can't convert to MultiPoint")
}

// try to convert the abstract Geometry into a MultiLine
func (g *Geometry) AsMultiLine() (*MultiLine, error) {
	if g.IsMultiLine() {
		switch coords := g.Coordinates.(type) {
		case [][][]float64:
			return &MultiLine{
				Coordinates: coords,
				SrId:        g.GetSrId(),
			}, nil
		case []interface{}:
			return &MultiLine{
				Coordinates: forceFloat64Array3D(coords),
				SrId:        g.GetSrId(),
			}, nil
		}
	}
	return nil, errors.New("can't convert to MultiLine")
}

// try to convert the abstract Geometry into a MultiPolygon
func (g *Geometry) AsMultiPolygon() (*MultiPolygon, error) {
	if g.IsMultiPolygon() {
		switch coords := g.Coordinates.(type) {
		case [][][][]float64:
			return &MultiPolygon{
				Coordinates: coords,
				SrId:        g.GetSrId(),
			}, nil
		case []interface{}:
			return &MultiPolygon{
				Coordinates: forceFloat64Array4D(coords),
				SrId:        g.GetSrId(),
			}, nil
		}
	}
	return nil, errors.New("can't convert to MultiPolygon")
}

// check the current Geometry Type for the Point Type
func (g *Geometry) IsPoint() bool {
	return g.Type == PointType
}

// check the current Geometry Type for the Line Type
func (g *Geometry) IsLine() bool {
	return g.Type == LineType
}

// check the current Geometry Type for the Polygon Type
func (g *Geometry) IsPolygon() bool {
	return g.Type == PolygonType
}

// check the current Geometry Type for the MultiPoint Type
func (g *Geometry) IsMultiPoint() bool {
	return g.Type == MultiPointType
}

// check the current Geometry Type for the MultiLine Type
func (g *Geometry) IsMultiLine() bool {
	return g.Type == MultiLineType
}

// check the current Geometry Type for the MultiPolygon Type
func (g *Geometry) IsMultiPolygon() bool {
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
