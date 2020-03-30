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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Polygon Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.PolygonCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		tmp, deserializeErr := geojson.DeserializeGeometry(test.PolygonGeoJSONCrs)
		ginkgo.It("deserialize successfully", func() {
			gomega.Expect(deserializeErr).To(gomega.BeNil())
		})
		l, err := tmp.AsPolygon()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has Polygon Type", func() {
			gomega.Expect(l.ToGeometry().GetType()).To(gomega.Equal(geojson.PolygonType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Polygon Coordinates", func() {
			coords, err := l.ToGeometry().GetCoordinates().AsPolygon()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(coords).To(gomega.Equal(test.PolygonCoordinates))
		})
		ginkgo.It("deserialize polygon array", func() {
			value, err := geojson.DeserializeGeometryList(test.PolygonGeoJSONArray)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(len(value)).To(gomega.Equal(1))
		})
	})
})
