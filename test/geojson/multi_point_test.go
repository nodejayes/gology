package geojson

import (
	"github.com/nodejayes/topgology/pkg/geojson"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("MultiPoint Test", func() {
	ginkgo.Describe("NewMultiPoint()", func() {
		geom, err := geojson.NewMultiPoint(MultiPointCoordinates, SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(MultiPointCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		l, err := geojson.DeserializeGeometry(MultiPointGeoJSONCrs).AsMultiPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has MultiPoint Type", func() {
			gomega.Expect(l.ToGeometry().Type).To(gomega.Equal(geojson.MultiPointType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetSrId()).To(gomega.Equal(SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			gomega.Expect(l.ToGeometry().Coordinates).To(gomega.Equal(MultiPointCoordinates))
		})
	})
})
