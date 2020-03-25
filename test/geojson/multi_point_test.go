package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("MultiPoint Test", func() {
	ginkgo.Describe("NewMultiPoint()", func() {
		geom, err := geojson.NewMultiPoint(test.MultiPointCoordinates, test.SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(test.MultiPointCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		l, err := geojson.DeserializeGeometry(test.MultiPointGeoJSONCrs).AsMultiPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has MultiPoint Type", func() {
			gomega.Expect(l.ToGeometry().Type).To(gomega.Equal(geojson.MultiPointType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			gomega.Expect(l.ToGeometry().Coordinates).To(gomega.Equal(test.MultiPointCoordinates))
		})
	})
})
