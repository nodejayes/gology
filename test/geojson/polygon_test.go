package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Polygon Test", func() {
	ginkgo.Describe("NewPolygon()", func() {
		geom, err := geojson.NewPolygon(test.PolygonCoordinates, test.SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Polygon Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(test.PolygonCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		l, err := geojson.DeserializeGeometry(test.PolygonGeoJSONCrs).AsPolygon()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has Polygon Type", func() {
			gomega.Expect(l.ToGeometry().Type).To(gomega.Equal(geojson.PolygonType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Polygon Coordinates", func() {
			gomega.Expect(l.ToGeometry().Coordinates).To(gomega.Equal(test.PolygonCoordinates))
		})
	})
})
