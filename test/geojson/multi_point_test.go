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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.MultiPointCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		tmp, deserializeErr := geojson.DeserializeGeometry(test.MultiPointGeoJSONCrs)
		ginkgo.It("deserialize successfully", func() {
			gomega.Expect(deserializeErr).To(gomega.BeNil())
		})
		l, err := tmp.AsMultiPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has MultiPoint Type", func() {
			gomega.Expect(l.ToGeometry().GetType()).To(gomega.Equal(geojson.MultiPointType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			coords, err := l.ToGeometry().GetCoordinates().AsMultiPoint()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(coords).To(gomega.Equal(test.MultiPointCoordinates))
		})
		ginkgo.It("deserialize multi point array", func() {
			value, err := geojson.DeserializeGeometryList(test.MultiPointGeoJSONArray)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(len(value)).To(gomega.Equal(1))
		})
	})
})
