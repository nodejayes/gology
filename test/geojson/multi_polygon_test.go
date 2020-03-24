package geojson

import (
	"github.com/nodejayes/topgology/pkg/geojson"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("MultiPolygon Test", func() {
	ginkgo.Describe("NewMultiPolygon()", func() {
		geom, err := geojson.NewMultiPolygon(MultiPolygonCoordinates, SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has MultiPolygon Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(MultiPolygonCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		l, err := geojson.DeserializeGeometry(MultiPolygonGeoJSONCrs).AsMultiPolygon()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has MultiPolygon Type", func() {
			gomega.Expect(l.ToGeometry().Type).To(gomega.Equal(geojson.MultiPolygonType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetSrId()).To(gomega.Equal(SrId))
		})
		ginkgo.It("has MultiPolygon Coordinates", func() {
			t := l.ToGeometry().Coordinates
			println(t)
			gomega.Expect(t).To(gomega.Equal(MultiPolygonCoordinates))
		})
	})
})
