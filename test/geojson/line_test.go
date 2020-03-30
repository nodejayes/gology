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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Line Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.LineCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		tmp, deserializeErr := geojson.DeserializeGeometry(test.LineGeoJSONCrs)
		ginkgo.It("deserialize successfully", func() {
			gomega.Expect(deserializeErr).To(gomega.BeNil())
		})
		l, err := tmp.AsLine()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has Line Type", func() {
			gomega.Expect(l.ToGeometry().GetType()).To(gomega.Equal(geojson.LineType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Line Coordinates", func() {
			coords, err := l.ToGeometry().GetCoordinates().AsLine()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(coords).To(gomega.Equal(test.LineCoordinates))
		})
		ginkgo.It("deserialize line array", func() {
			value, err := geojson.DeserializeGeometryList(test.LineGeoJSONArray)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(len(value)).To(gomega.Equal(1))
		})
	})
})
