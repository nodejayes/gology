package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Point Test", func() {
	ginkgo.Describe("NewPoint()", func() {
		geom, err := geojson.NewPoint(test.PointCoordinates, test.SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("SrId was set", func() {
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("Coordinates was set", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.PointCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		tmp, deserializeErr := geojson.DeserializeGeometry(test.PointGeoJSONCrs)
		ginkgo.It("deserialize successfully", func() {
			gomega.Expect(deserializeErr).To(gomega.BeNil())
		})
		pt, err := tmp.AsPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has Point Type", func() {
			gomega.Expect(pt.ToGeometry().GetType()).To(gomega.Equal(geojson.PointType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(pt.ToGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Point Coordinates", func() {
			coords, err := pt.ToGeometry().GetCoordinates().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(coords).To(gomega.Equal(test.PointCoordinates))
		})
	})
})
