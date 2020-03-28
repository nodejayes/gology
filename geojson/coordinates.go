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
	res, ok := c.Data.([]float64)
	if !ok {
		return nil, WrongCoordinateTypeError
	}
	return res, nil
}

func (c coordinate) AsLine() ([][]float64, error) {
	res, ok := c.Data.([][]float64)
	if !ok {
		return nil, WrongCoordinateTypeError
	}
	return res, nil
}

func (c coordinate) AsPolygon() ([][][]float64, error) {
	res, ok := c.Data.([][][]float64)
	if !ok {
		return nil, WrongCoordinateTypeError
	}
	return res, nil
}

func (c coordinate) AsMultiPoint() ([][]float64, error) {
	res, ok := c.Data.([][]float64)
	if !ok {
		return nil, WrongCoordinateTypeError
	}
	return res, nil
}

func (c coordinate) AsMultiLine() ([][][]float64, error) {
	res, ok := c.Data.([][][]float64)
	if !ok {
		return nil, WrongCoordinateTypeError
	}
	return res, nil
}

func (c coordinate) AsMultiPolygon() ([][][][]float64, error) {
	res, ok := c.Data.([][][][]float64)
	if !ok {
		return nil, WrongCoordinateTypeError
	}
	return res, nil
}
