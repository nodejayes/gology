package geojson

import "errors"

var WrongCoordinateTypeError = errors.New("Wrong Coordinate Type")

type ICoordinate interface {
	AsPoint() ([]float64, error)
	AsLine() ([][]float64, error)
	AsPolygon() ([][][]float64, error)
	AsMultiPoint() ([][]float64, error)
	AsMultiLine() ([][][]float64, error)
	AsMultiPolygon() ([][][][]float64, error)
}

type coordinate struct {
	Data interface{}
}

func NewCoordinate(coords interface{}) ICoordinate {
	return &coordinate{Data: coords}
}

func (c coordinate) AsPoint() ([]float64, error) {
	switch res := c.Data.(type) {
	case []float64:
		return res, nil
	case *[]float64:
		return *res, nil
	case *coordinate:
		return res.AsPoint()
	case coordinate:
		return res.AsPoint()
	case []interface{}:
		return forceFloat64Array1D(res), nil
	case *[]interface{}:
		return forceFloat64Array1D(*res), nil
	}
	return nil, WrongCoordinateTypeError
}

func (c coordinate) AsLine() ([][]float64, error) {
	switch res := c.Data.(type) {
	case [][]float64:
		return res, nil
	case *[][]float64:
		return *res, nil
	case *coordinate:
		return res.AsLine()
	case coordinate:
		return res.AsLine()
	case []interface{}:
		return forceFloat64Array2D(res), nil
	case *[]interface{}:
		return forceFloat64Array2D(*res), nil
	}
	return nil, WrongCoordinateTypeError
}

func (c coordinate) AsPolygon() ([][][]float64, error) {
	switch res := c.Data.(type) {
	case [][][]float64:
		return res, nil
	case *[][][]float64:
		return *res, nil
	case *coordinate:
		return res.AsPolygon()
	case coordinate:
		return res.AsPolygon()
	case []interface{}:
		return forceFloat64Array3D(res), nil
	case *[]interface{}:
		return forceFloat64Array3D(*res), nil
	}
	return nil, WrongCoordinateTypeError
}

func (c coordinate) AsMultiPoint() ([][]float64, error) {
	switch res := c.Data.(type) {
	case [][]float64:
		return res, nil
	case *[][]float64:
		return *res, nil
	case *coordinate:
		return res.AsMultiPoint()
	case coordinate:
		return res.AsMultiPoint()
	case []interface{}:
		return forceFloat64Array2D(res), nil
	case *[]interface{}:
		return forceFloat64Array2D(*res), nil
	}
	return nil, WrongCoordinateTypeError
}

func (c coordinate) AsMultiLine() ([][][]float64, error) {
	switch res := c.Data.(type) {
	case [][][]float64:
		return res, nil
	case *[][][]float64:
		return *res, nil
	case *coordinate:
		return res.AsMultiLine()
	case coordinate:
		return res.AsMultiLine()
	case []interface{}:
		return forceFloat64Array3D(res), nil
	case *[]interface{}:
		return forceFloat64Array3D(*res), nil
	}
	return nil, WrongCoordinateTypeError
}

func (c coordinate) AsMultiPolygon() ([][][][]float64, error) {
	switch res := c.Data.(type) {
	case [][][][]float64:
		return res, nil
	case *[][][][]float64:
		return *res, nil
	case *coordinate:
		return res.AsMultiPolygon()
	case coordinate:
		return res.AsMultiPolygon()
	case []interface{}:
		return forceFloat64Array4D(res), nil
	case *[]interface{}:
		return forceFloat64Array4D(*res), nil
	}
	return nil, WrongCoordinateTypeError
}
