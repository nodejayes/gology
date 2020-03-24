package geojson

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
)

type GeometryTypes string

const (
	PointType        GeometryTypes = "Point"
	MultiPointType   GeometryTypes = "MultiPoint"
	LineType         GeometryTypes = "LineString"
	MultiLineType    GeometryTypes = "MultiLineString"
	PolygonType      GeometryTypes = "Polygon"
	MultiPolygonType GeometryTypes = "MultiPolygon"
)

type IGeometryConverter interface {
	ToGeometry() *Geometry
}

type Geometry struct {
	Type        GeometryTypes    `json:"type"`
	Coordinates interface{}      `json:"coordinates"`
	CRS         *ReferenceSystem `json:"crs,omitempty"`
}

func NewGeometry(typ GeometryTypes, coordinates interface{}, srId int) *Geometry {
	return &Geometry{
		Type:        typ,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
}

func NewPoint(coordinate []float64, srId int) (*Point, error) {
	geom := &Geometry{
		Type:        PointType,
		Coordinates: coordinate,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsPoint()
}

func NewLine(coordinates [][]float64, srId int) (*Line, error) {
	geom := &Geometry{
		Type:        LineType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsLine()
}

func NewPolygon(coordinates [][][]float64, srId int) (*Polygon, error) {
	geom := &Geometry{
		Type:        PolygonType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsPolygon()
}

func NewMultiPoint(coordinates [][]float64, srId int) (*MultiPoint, error) {
	geom := &Geometry{
		Type:        MultiPointType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiPoint()
}

func NewMultiLine(coordinates [][][]float64, srId int) (*MultiLine, error) {
	geom := &Geometry{
		Type:        MultiLineType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiLine()
}

func NewMultiPolygon(coordinates [][][][]float64, srId int) (*MultiPolygon, error) {
	geom := &Geometry{
		Type:        MultiPolygonType,
		Coordinates: coordinates,
		CRS:         NewReferenceSystem(srId),
	}
	return geom.AsMultiPolygon()
}

func DeserializeGeometry(str string) *Geometry {
	var geom *Geometry
	err := jsoniter.Unmarshal([]byte(str), &geom)
	if err != nil {
		return nil
	}
	return geom
}

func (g *Geometry) Serialize() string {
	stream, err := jsoniter.Marshal(g)
	if err != nil {
		return ""
	}
	return string(stream)
}

func (g *Geometry) GetSrId() int {
	if g.CRS == nil {
		return 0
	}
	return g.CRS.GetSrId()
}

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

func (g *Geometry) IsPoint() bool {
	return g.Type == PointType
}

func (g *Geometry) IsLine() bool {
	return g.Type == LineType
}

func (g *Geometry) IsPolygon() bool {
	return g.Type == PolygonType
}

func (g *Geometry) IsMultiPoint() bool {
	return g.Type == MultiPointType
}

func (g *Geometry) IsMultiLine() bool {
	return g.Type == MultiLineType
}

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
