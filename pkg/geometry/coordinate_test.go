package geometry

import (
	"github.com/onsi/gomega"
	"math"
	"testing"
)

func TestNewCoordinate(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	coord2d := NewCoordinate(1.0, 2.0, math.Inf(-1))
	g.Expect(coord2d.x).To(gomega.Equal(1.0))
	g.Expect(coord2d.y).To(gomega.Equal(2.0))
	g.Expect(coord2d.z).To(gomega.Equal(math.Inf(-1)))
	g.Expect(coord2d.dimension).To(gomega.Equal(DD))

	coord3d := NewCoordinate(1.0, 2.0, 3.0)
	g.Expect(coord3d.x).To(gomega.Equal(1.0))
	g.Expect(coord3d.y).To(gomega.Equal(2.0))
	g.Expect(coord3d.z).To(gomega.Equal(3.0))
	g.Expect(coord3d.dimension).To(gomega.Equal(DDD))
}

func TestCoordinate_CompareTo(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	coord1 := NewCoordinate(1.5, 2.5, 3.5)
	coord2 := NewCoordinate(1.5, 2.5, 3.5)
	coord3 := NewCoordinate(1.0, 2.0, 3.5)
	coord4 := NewCoordinate(2.0, 3.0, 4.5)
	coord5 := NewCoordinate(1.5, 1.5, 4.5)
	coord6 := NewCoordinate(1.5, 2.5, 4.5)
	coord7 := NewCoordinate(1.5, 3.0, 4.5)

	g.Expect(coord1.CompareTo(coord2)).To(gomega.Equal(0))
	g.Expect(coord1.CompareTo(coord3)).To(gomega.Equal(1))
	g.Expect(coord1.CompareTo(coord4)).To(gomega.Equal(-1))
	g.Expect(coord1.CompareTo(coord5)).To(gomega.Equal(1))
	g.Expect(coord1.CompareTo(coord6)).To(gomega.Equal(-1))
	g.Expect(coord6.CompareTo(coord1)).To(gomega.Equal(1))
	g.Expect(coord1.CompareTo(coord7)).To(gomega.Equal(-1))

	g.Expect(func() {
		coord1.CompareTo(nil)
	}).To(gomega.Panic())
}

func TestCoordinate_GetterSetter(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinate(1.5, 2.5, 3.5)
	g.Expect(coord.GetX()).To(gomega.Equal(1.5))
	g.Expect(coord.GetY()).To(gomega.Equal(2.5))
	g.Expect(coord.GetZ()).To(gomega.Equal(3.5))
	g.Expect(coord.GetDimension()).To(gomega.Equal(DDD))

	coord.SetX(15.5)
	coord.SetY(15.5)
	coord.SetZ(15.5)
	g.Expect(coord.GetX()).To(gomega.Equal(15.5))
	g.Expect(coord.GetY()).To(gomega.Equal(15.5))
	g.Expect(coord.GetZ()).To(gomega.Equal(15.5))
	g.Expect(coord.GetDimension()).To(gomega.Equal(DDD))

	coord.SetZ(math.Inf(-1))
	g.Expect(coord.GetDimension()).To(gomega.Equal(DD))
}

func TestCoordinate_TakeOverCoordinate(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinate(1.5, 2.5, 3.5)
	g.Expect(coord.GetX()).To(gomega.Equal(1.5))
	g.Expect(coord.GetY()).To(gomega.Equal(2.5))
	g.Expect(coord.GetZ()).To(gomega.Equal(3.5))
	g.Expect(coord.GetDimension()).To(gomega.Equal(DDD))

	coord.TakeOverCoordinate(NewCoordinate(15.5, 15.5, math.Inf(-1)))
	g.Expect(coord.GetX()).To(gomega.Equal(15.5))
	g.Expect(coord.GetY()).To(gomega.Equal(15.5))
	g.Expect(coord.GetDimension()).To(gomega.Equal(DD))
}

func TestCoordinate_Copy(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)

	coord1 := NewCoordinate(1.5, 2.5, 3.5)
	coord2 := coord1.Copy()
	coord2.x = 15.5
	coord2.y = 15.5
	coord2.z = 15.5

	g.Expect(coord1.x).To(gomega.Equal(1.5))
	g.Expect(coord1.y).To(gomega.Equal(2.5))
	g.Expect(coord1.z).To(gomega.Equal(3.5))
}

func TestCoordinate_Distance(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinate(0.0, 0.0, 0.0)
	coord2 := NewCoordinate(1.0, 0.0, 0.0)
	coord3 := NewCoordinate(1.0, 1.0, 0.0)

	coord4 := NewCoordinate(5.5, 4.8, math.Inf(-1))
	coord5 := NewCoordinate(15.5, 2.789, math.Inf(-1))

	g.Expect(coord1.Distance(coord2)).To(gomega.Equal(1.0))
	g.Expect(coord1.Distance(coord3)).To(gomega.Equal(1.4142135623730951))
	g.Expect(coord1.Distance(coord4)).To(gomega.Equal(7.3))
	g.Expect(coord1.Distance(coord5)).To(gomega.Equal(15.748921264645398))
	g.Expect(coord4.Distance(coord5)).To(gomega.Equal(10.200202007803572))
}

func TestCoordinate_Equals(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinate(1.5, 2.5, 3.5)
	coord2 := NewCoordinate(1.5, 2.5, 3.5)
	coord3 := NewCoordinate(1.4, 2.5, 3.5)

	g.Expect(coord1.Equals(coord2)).To(gomega.BeTrue())
	g.Expect(coord1.Equals(coord3)).To(gomega.BeFalse())
}

func TestCoordinate_EqualsTolerance(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinate(1.5, 2.5, 3.5)
	coord2 := NewCoordinate(1.5, 2.5, 3.5)
	coord3 := NewCoordinate(1.4, 2.5, 3.5)
	coord4 := NewCoordinate(1.41, 2.5, 3.5)

	g.Expect(coord1.EqualsTolerance(coord2, 0.1)).To(gomega.BeTrue())
	g.Expect(coord1.EqualsTolerance(coord3, 0.1)).To(gomega.BeFalse())
	g.Expect(coord1.EqualsTolerance(coord4, 0.1)).To(gomega.BeTrue())
}

func TestCoordinate_Equals2D(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinate(1.5, 2.5, 3.5)
	coord2 := NewCoordinate(1.5, 2.5, 3.5)
	coord3 := NewCoordinate(1.5, 2.5, 4.5)
	coord4 := NewCoordinate(1.4, 2.5, 4.5)

	g.Expect(coord1.Equals2D(coord2)).To(gomega.BeTrue())
	g.Expect(coord1.Equals2D(coord3)).To(gomega.BeTrue())
	g.Expect(coord1.Equals2D(coord4)).To(gomega.BeFalse())
}

func TestCoordinate_Equals2DTolerance(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord1 := NewCoordinate(1.5, 2.5, 3.5)
	coord2 := NewCoordinate(1.4, 2.5, 3.5)
	coord3 := NewCoordinate(1.41, 2.5, 3.5)

	g.Expect(coord1.Equals2DTolerance(coord2, 0.0)).To(gomega.BeFalse())
	g.Expect(coord1.Equals2DTolerance(coord2, 0.1)).To(gomega.BeFalse())
	g.Expect(coord1.Equals2DTolerance(coord3, 0.1)).To(gomega.BeTrue())

	g.Expect(coord2.Equals2DTolerance(coord1, 0.0)).To(gomega.BeFalse())
	g.Expect(coord2.Equals2DTolerance(coord1, 0.1)).To(gomega.BeFalse())
	g.Expect(coord3.Equals2DTolerance(coord1, 0.1)).To(gomega.BeTrue())
}

func TestCoordinate_GetValues(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinate(1.5, 2.5, 3.5)
	g.Expect(coord.GetValues()).To(gomega.Equal([]float64{1.5, 2.5, 3.5}))
}

func TestCoordinate_ToString(t *testing.T) {
	t.Parallel()
	g := gomega.NewGomegaWithT(t)
	coord := NewCoordinate(1.5, 2.5, 3.5)
	g.Expect(coord.ToString()).To(gomega.Equal("(X: 1.5, Y: 2.5, Z: 3.5) 3D"))

	coordInf := NewCoordinate(1.5, 2.5, math.Inf(-1))
	g.Expect(coordInf.ToString()).To(gomega.Equal("(X: 1.5, Y: 2.5, Z: -Inf) 2D"))
}
