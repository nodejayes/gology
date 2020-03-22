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
