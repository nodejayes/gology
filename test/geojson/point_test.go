package geojson

import (
	"git.agricon.de/map-factory/pkg/geojson"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Point Test", func() {
	ginkgo.Describe("NewPoint()", func() {
		geom, err := geojson.NewPoint(PointCoordinates, SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("SrId was set", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("Coordinates was set", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(PointCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		pt, err := geojson.DeserializeGeometry(PointGeoJSONCrs).AsPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has Point Type", func() {
			gomega.Expect(pt.ToGeometry().Type).To(gomega.Equal(geojson.PointType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(pt.ToGeometry().GetSrId()).To(gomega.Equal(SrId))
		})
		ginkgo.It("has Point Coordinates", func() {
			gomega.Expect(pt.ToGeometry().Coordinates).To(gomega.Equal(PointCoordinates))
		})
	})
})
