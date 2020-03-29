package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("MultiPolygon Test", func() {
	ginkgo.Describe("NewMultiPolygon()", func() {
		geom, err := geojson.NewMultiPolygon(test.MultiPolygonCoordinates, test.SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPolygon Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.MultiPolygonCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		tmp, deserializeErr := geojson.DeserializeGeometry(test.MultiPolygonGeoJSONCrs)
		ginkgo.It("deserialize successfully", func() {
			gomega.Expect(deserializeErr).To(gomega.BeNil())
		})
		l, err := tmp.AsMultiPolygon()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has MultiPolygon Type", func() {
			gomega.Expect(l.ToGeometry().GetType()).To(gomega.Equal(geojson.MultiPolygonType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPolygon Coordinates", func() {
			coords, err := l.ToGeometry().GetCoordinates().AsMultiPolygon()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(coords).To(gomega.Equal(test.MultiPolygonCoordinates))
		})
	})
})
