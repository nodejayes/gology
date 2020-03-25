package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Line Test", func() {
	ginkgo.Describe("NewLine()", func() {
		geom, err := geojson.NewLine(test.LineCoordinates, test.SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Line Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(test.LineCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		l, err := geojson.DeserializeGeometry(test.LineGeoJSONCrs).AsLine()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has Line Type", func() {
			gomega.Expect(l.ToGeometry().Type).To(gomega.Equal(geojson.LineType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Line Coordinates", func() {
			gomega.Expect(l.ToGeometry().Coordinates).To(gomega.Equal(test.LineCoordinates))
		})
	})
})
