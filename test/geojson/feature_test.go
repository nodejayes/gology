package geojson

import (
	"github.com/nodejayes/topgology/pkg/geojson"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Feature Test", func() {
	ginkgo.Describe("NewFeature()", func() {
		geom := geojson.NewGeometry(geojson.PointType, PointCoordinates, SrId)
		f := geojson.NewFeature(geom, nil)
		ginkgo.It("has Feature Type", func() {
			gomega.Expect(f.Type).To(gomega.Equal(geojson.FeatureType))
		})
		ginkgo.It("has Point Geometry", func() {
			gomega.Expect(f.Geometry).NotTo(gomega.BeNil())
			pt, err := f.Geometry.AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(pt.Coordinates).To(gomega.Equal(PointCoordinates))
			gomega.Expect(pt.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has Initialized Properties", func() {
			gomega.Expect(f.Properties).NotTo(gomega.BeNil())
			gomega.Expect(len(*f.Properties)).To(gomega.Equal(0))
		})
		ginkgo.It("has Properties that was given into", func() {
			someProps := make(map[string]interface{})
			someProps["count"] = 42
			f := geojson.NewFeature(geom, &someProps)
			gomega.Expect(f.Properties).NotTo(gomega.BeNil())
			gomega.Expect(len(*f.Properties)).To(gomega.Equal(1))
			gomega.Expect((*f.Properties)["count"]).To(gomega.Equal(42))
		})
	})
})
