package geometries

import "strconv"

func checkForSrid(points []*Point) int {
	srid := 0
	for _, pt := range points {
		if srid == 0 {
			srid = pt.srid
		}
		if srid != pt.srid {
			panic("point list has multiple srids expect all points to have same srid " + strconv.FormatInt(int64(srid), 10))
		}
	}
	return srid
}

func checkMultiPointLength(mp *MultiPoint, pos int) {
	if len(mp.points)-1 < pos {
		panic("MultiPoint has not enough Points request point " +
			strconv.FormatInt(int64(pos), 10) + " has " +
			strconv.FormatInt(int64(len(mp.points)-1), 10))
	}
}

func calculateBoundingBox(points *MultiPoint) *BoundingBox {
	var minX float64
	var maxX float64
	var minY float64
	var maxY float64
	var minZ float64
	var maxZ float64
	for idx, pt := range points.points {
		if idx == 0 || pt.coordinate.x < minX {
			minX = pt.coordinate.x
		}
		if idx == 0 || pt.coordinate.x > maxX {
			maxX = pt.coordinate.x
		}
		if idx == 0 || pt.coordinate.y < minY {
			minY = pt.coordinate.y
		}
		if idx == 0 || pt.coordinate.y > maxY {
			maxY = pt.coordinate.y
		}
		if idx == 0 || pt.coordinate.z < minZ {
			minZ = pt.coordinate.z
		}
		if idx == 0 || pt.coordinate.z > maxZ {
			maxZ = pt.coordinate.z
		}
	}
	topLeft := NewPointFromXYZ(minX, maxY, minZ, points.srid)
	topRight := NewPointFromXYZ(maxX, maxY, minZ, points.srid)
	bottomRight := NewPointFromXYZ(maxX, minY, minZ, points.srid)
	bottomLeft := NewPointFromXYZ(minX, minY, minZ, points.srid)
	return NewBoundingBoxFromPoints(topLeft, topRight, bottomRight, bottomLeft)
}

type MultiPoint struct {
	points      []*Point
	srid        int
	boundingBox *BoundingBox
}

func NewMultiPointFromPoints(points []*Point) *MultiPoint {
	if points == nil {
		return &MultiPoint{
			points:      []*Point{},
			srid:        0,
			boundingBox: nil,
		}
	}
	srid := checkForSrid(points)
	mp := &MultiPoint{
		points:      points,
		srid:        srid,
		boundingBox: nil,
	}
	mp.boundingBox = calculateBoundingBox(mp)
	return mp
}

func (mp *MultiPoint) GetCoordinates() []*Coordinate {
	var tmp []*Coordinate
	for _, pt := range mp.points {
		tmp = append(tmp, pt.coordinate)
	}
	return tmp
}

func (mp *MultiPoint) GetCoordinateValues() [][]float64 {
	var tmp [][]float64
	for _, pt := range mp.points {
		tmp = append(tmp, pt.coordinate.GetCoordinateValue())
	}
	return tmp
}

func (mp *MultiPoint) GetCoordinate(pos int) *Coordinate {
	checkMultiPointLength(mp, pos)
	return mp.points[pos].coordinate.GetCoordinate()
}

func (mp *MultiPoint) GetCoordinateValue(pos int) []float64 {
	checkMultiPointLength(mp, pos)
	return mp.points[pos].coordinate.GetCoordinateValue()
}
