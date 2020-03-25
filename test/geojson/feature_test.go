package geojson

import (
	"github.com/nodejayes/topgology/geojson"
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
	ginkgo.Describe("Deserialize()", func() {
		feature := geojson.DeserializeFeature(FeatureGeoJSON)
		featureCrs := geojson.DeserializeFeature(FeatureGeoJSONCrs)
		featureWithProps := geojson.DeserializeFeature(FeatureGeoJSONProps)
		ginkgo.It("has Feature Type", func() {
			gomega.Expect(feature.Type).To(gomega.Equal(geojson.FeatureType))
			gomega.Expect(featureWithProps.Type).To(gomega.Equal(geojson.FeatureType))
			gomega.Expect(featureCrs.Type).To(gomega.Equal(geojson.FeatureType))
		})
		ginkgo.It("has the Point as Geometry", func() {
			gomega.Expect(feature.Geometry.IsPoint()).To(gomega.BeTrue())
			gomega.Expect(featureWithProps.Geometry.IsPoint()).To(gomega.BeTrue())
			gomega.Expect(featureCrs.Geometry.IsPoint()).To(gomega.BeTrue())
		})
		ginkgo.It("has 0 as SrId", func() {
			gomega.Expect(feature.Geometry.GetSrId()).To(gomega.Equal(0))
			gomega.Expect(featureWithProps.Geometry.GetSrId()).To(gomega.Equal(0))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(featureCrs.Geometry.GetSrId()).To(gomega.Equal(SrId))
		})
		ginkgo.It("has empty Properties", func() {
			gomega.Expect(len(*feature.Properties)).To(gomega.Equal(0))
			gomega.Expect(len(*featureCrs.Properties)).To(gomega.Equal(0))
		})
		ginkgo.It("has Properties", func() {
			gomega.Expect(len(*featureWithProps.Properties)).To(gomega.Equal(1))
			gomega.Expect((*featureWithProps.Properties)["hello"]).To(gomega.Equal("world"))
		})
	})
	ginkgo.Describe("Serialize()", func() {
		feature := geojson.DeserializeFeature(FeatureGeoJSON)
		featureCrs := geojson.DeserializeFeature(FeatureGeoJSONCrs)
		featureWithProps := geojson.DeserializeFeature(FeatureGeoJSONProps)
		ginkgo.It("is Valid GeoJSON Feature", func() {
			gomega.Expect(feature.Serialize()).To(gomega.Equal(FeatureGeoJSON))
			gomega.Expect(featureCrs.Serialize()).To(gomega.Equal(FeatureGeoJSONCrs))
			gomega.Expect(featureWithProps.Serialize()).To(gomega.Equal(FeatureGeoJSONProps))
		})
	})
})
