package geometries

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestNewPointFromLatLong(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromLatLong(12.571105957031248, 50.999928855859636, 0)

	g.Expect(pt.coordinate.x).To(gomega.Equal(12.571105957031248))
	g.Expect(pt.coordinate.y).To(gomega.Equal(50.999928855859636))
	g.Expect(pt.coordinate.z).To(gomega.Equal(0.0))
	g.Expect(pt.coordinate.dimension).To(gomega.Equal(DDD))
	g.Expect(pt.srid).To(gomega.Equal(4326))
}

func TestNewPointFromXYZ(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)

	g.Expect(pt.coordinate.x).To(gomega.Equal(1.5))
	g.Expect(pt.coordinate.y).To(gomega.Equal(2.5))
	g.Expect(pt.coordinate.z).To(gomega.Equal(3.5))
	g.Expect(pt.coordinate.dimension).To(gomega.Equal(DDD))
	g.Expect(pt.srid).To(gomega.Equal(4326))
}

func TestPoint_GetCoordinate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	g.Expect(pt.GetCoordinate()).To(gomega.Equal([]float64{1.5, 2.5, 3.5}))
}

func TestPoint_GetType(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	g.Expect(pt.GetType()).To(gomega.Equal(PointType))
}

func TestPoint_GetX(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	g.Expect(pt.GetX()).To(gomega.Equal(1.5))
}

func TestPoint_GetY(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	g.Expect(pt.GetY()).To(gomega.Equal(2.5))
}

func TestPoint_GetZ(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	g.Expect(pt.GetZ()).To(gomega.Equal(3.5))
}

func TestPoint_SetX(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	pt.SetX(15.5)
	g.Expect(pt.coordinate.x).To(gomega.Equal(15.5))
}

func TestPoint_SetY(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	pt.SetY(15.5)
	g.Expect(pt.coordinate.y).To(gomega.Equal(15.5))
}

func TestPoint_SetZ(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	pt := NewPointFromXYZ(1.5, 2.5, 3.5, 4326)
	pt.SetZ(15.5)
	g.Expect(pt.coordinate.z).To(gomega.Equal(15.5))
}
