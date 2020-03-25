package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("MultiLine Test", func() {
	ginkgo.Describe("NewMultiLine()", func() {
		geom, err := geojson.NewMultiLine(test.MultiLineCoordinates, test.SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiLine Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(test.MultiLineCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		l, err := geojson.DeserializeGeometry(test.MultiLineGeoJSONCrs).AsMultiLine()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has MultiLine Type", func() {
			gomega.Expect(l.ToGeometry().Type).To(gomega.Equal(geojson.MultiLineType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiLine Coordinates", func() {
			gomega.Expect(l.ToGeometry().Coordinates).To(gomega.Equal(test.MultiLineCoordinates))
		})
	})
})
