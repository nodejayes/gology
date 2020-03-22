package geometries

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestNewCoordinateFromXY(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinateFromXY(1.0, 2.0)

	g.Expect(coord.x).To(gomega.Equal(1.0))
	g.Expect(coord.y).To(gomega.Equal(2.0))
	g.Expect(coord.z).To(gomega.Equal(0.0))
	g.Expect(coord.dimension).To(gomega.Equal(DD))
}

func TestNewCoordinateFromXYZ(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinateFromXYZ(1.0, 2.0, 3.0)

	g.Expect(coord.x).To(gomega.Equal(1.0))
	g.Expect(coord.y).To(gomega.Equal(2.0))
	g.Expect(coord.z).To(gomega.Equal(3.0))
	g.Expect(coord.dimension).To(gomega.Equal(DDD))
}

func TestNewCoordinateFromCoordinate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	testCoord := NewCoordinateFromCoordinate(coord)

	g.Expect(coord.x).To(gomega.Equal(testCoord.x))
	g.Expect(coord.y).To(gomega.Equal(testCoord.y))
	g.Expect(coord.z).To(gomega.Equal(testCoord.z))
	g.Expect(coord.dimension).To(gomega.Equal(testCoord.dimension))
}

func TestCoordinate_CompareTo(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	coord1 := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	coord2 := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	coord3 := NewCoordinateFromXYZ(1.0, 2.0, 3.5)
	coord4 := NewCoordinateFromXYZ(2.0, 3.0, 4.5)
	coord5 := NewCoordinateFromXYZ(1.5, 1.5, 4.5)
	coord6 := NewCoordinateFromXYZ(1.5, 2.5, 4.5)
	coord7 := NewCoordinateFromXYZ(1.5, 3.0, 4.5)

	g.Expect(coord1.CompareTo(coord2)).To(gomega.Equal(0))
	g.Expect(coord1.CompareTo(coord3)).To(gomega.Equal(1))
	g.Expect(coord1.CompareTo(coord4)).To(gomega.Equal(-1))
	g.Expect(coord1.CompareTo(coord5)).To(gomega.Equal(1))
	g.Expect(coord1.CompareTo(coord6)).To(gomega.Equal(0))
	g.Expect(coord1.CompareTo(coord7)).To(gomega.Equal(-1))

	g.Expect(coord1.CompareTo(*coord2)).To(gomega.Equal(0))
	g.Expect(coord1.CompareTo(*coord3)).To(gomega.Equal(1))
	g.Expect(coord1.CompareTo(*coord4)).To(gomega.Equal(-1))
	g.Expect(coord1.CompareTo(*coord5)).To(gomega.Equal(1))
	g.Expect(coord1.CompareTo(*coord6)).To(gomega.Equal(0))
	g.Expect(coord1.CompareTo(*coord7)).To(gomega.Equal(-1))

	g.Expect(func() {
		coord1.CompareTo(1)
	}).To(gomega.Panic())

	g.Expect(func() {
		coord1.CompareTo(nil)
	}).To(gomega.Panic())
}

func TestCoordinate_Copy(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	coord1 := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	coord2 := coord1.Copy()
	coord2.x = 15.5
	coord2.y = 15.5
	coord2.z = 15.5

	g.Expect(coord1.x).To(gomega.Equal(1.5))
	g.Expect(coord1.y).To(gomega.Equal(2.5))
	g.Expect(coord1.z).To(gomega.Equal(3.5))
}

func TestCoordinate_Distance(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinateFromXYZ(0.0, 0.0, 0.0)
	coord2 := NewCoordinateFromXYZ(1.0, 0.0, 0.0)
	coord3 := NewCoordinateFromXYZ(1.0, 1.0, 0.0)

	coord4 := NewCoordinateFromXY(5.5, 4.8)
	coord5 := NewCoordinateFromXY(15.5, 2.789)

	g.Expect(coord1.Distance(coord2)).To(gomega.Equal(1.0))
	g.Expect(coord1.Distance(coord3)).To(gomega.Equal(1.4142135623730951))
	g.Expect(coord1.Distance(coord4)).To(gomega.Equal(7.3))
	g.Expect(coord1.Distance(coord5)).To(gomega.Equal(15.748921264645398))
	g.Expect(coord4.Distance(coord5)).To(gomega.Equal(10.200202007803572))
}

func TestCoordinate_Equals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	coord2 := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	coord3 := NewCoordinateFromXYZ(1.4, 2.5, 3.5)

	g.Expect(coord1.Equals(coord2)).To(gomega.BeTrue())
	g.Expect(coord1.Equals(coord3)).To(gomega.BeFalse())
	g.Expect(coord1.Equals(*coord2)).To(gomega.BeTrue())
	g.Expect(coord1.Equals(*coord3)).To(gomega.BeFalse())
	g.Expect(coord1.Equals(1)).To(gomega.BeFalse())
}

func TestCoordinate_Equals2DTolerance(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	coord2 := NewCoordinateFromXYZ(1.4, 2.5, 3.5)
	coord3 := NewCoordinateFromXYZ(1.41, 2.5, 3.5)

	g.Expect(coord1.Equals2DTolerance(coord2, 0.0)).To(gomega.BeFalse())
	g.Expect(coord1.Equals2DTolerance(coord2, 0.1)).To(gomega.BeFalse())
	g.Expect(coord1.Equals2DTolerance(coord3, 0.1)).To(gomega.BeTrue())

	g.Expect(coord2.Equals2DTolerance(coord1, 0.0)).To(gomega.BeFalse())
	g.Expect(coord2.Equals2DTolerance(coord1, 0.1)).To(gomega.BeFalse())
	g.Expect(coord3.Equals2DTolerance(coord1, 0.1)).To(gomega.BeTrue())
}

func TestCoordinate_GetCoordinateValue(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	g.Expect(coord.GetCoordinateValue()).To(gomega.Equal([]float64{1.5, 2.5, 3.5}))
}

func TestCoordinate_GetCoordinate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	g.Expect(coord.GetCoordinate().x).To(gomega.Equal(coord.x))
	g.Expect(coord.GetCoordinate().y).To(gomega.Equal(coord.y))
	g.Expect(coord.GetCoordinate().z).To(gomega.Equal(coord.z))
	g.Expect(coord.GetCoordinate().dimension).To(gomega.Equal(coord.dimension))
}

func TestCoordinate_SetCoordinateValue(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	coord.SetCoordinateValue(NewCoordinateFromXYZ(2.5, 3.5, 4.5))
	g.Expect(coord.x).To(gomega.Equal(2.5))
	g.Expect(coord.y).To(gomega.Equal(3.5))
	g.Expect(coord.z).To(gomega.Equal(4.5))
	g.Expect(coord.dimension).To(gomega.Equal(DDD))

	coord.SetCoordinateValue(NewCoordinateFromXY(15.5, 15.5))
	g.Expect(coord.x).To(gomega.Equal(15.5))
	g.Expect(coord.y).To(gomega.Equal(15.5))
	g.Expect(coord.z).To(gomega.Equal(0.0))
	g.Expect(coord.dimension).To(gomega.Equal(DD))
}

func TestCoordinate_ToString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinateFromXYZ(1.5, 2.5, 3.5)
	g.Expect(coord.ToString()).To(gomega.Equal("(1.5, 2.5, 3.5)"))
}
