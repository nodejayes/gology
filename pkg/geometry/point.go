package geometry

type IPoint interface {
	GetCoordinate() ICoordinate
	SetCoordinate(ICoordinate)
}

type Point struct {
	coordinate ICoordinate
}

func NewPoint(coord ICoordinate) *Point {
	return &Point{coordinate: coord}
}

func (p *Point) GetCoordinate() ICoordinate {
	return p.coordinate
}

func (p *Point) SetCoordinate(value ICoordinate) {
	p.coordinate = value
}
