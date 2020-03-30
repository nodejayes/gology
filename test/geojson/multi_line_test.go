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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiLine Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.MultiLineCoordinates))
		})
	})
	ginkgo.Describe("ToGeometry()", func() {
		tmp, deserializeErr := geojson.DeserializeGeometry(test.MultiLineGeoJSONCrs)
		ginkgo.It("deserialize successfully", func() {
			gomega.Expect(deserializeErr).To(gomega.BeNil())
		})
		l, err := tmp.AsMultiLine()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has MultiLine Type", func() {
			gomega.Expect(l.ToGeometry().GetType()).To(gomega.Equal(geojson.MultiLineType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(l.ToGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiLine Coordinates", func() {
			coords, err := l.ToGeometry().GetCoordinates().AsMultiLine()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(coords).To(gomega.Equal(test.MultiLineCoordinates))
		})
		ginkgo.It("deserialize multi line array", func() {
			value, err := geojson.DeserializeGeometryList(test.MultiLineGeoJSONArray)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(len(value)).To(gomega.Equal(1))
		})
	})
})
