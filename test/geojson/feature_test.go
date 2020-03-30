package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Feature Test", func() {
	ginkgo.Describe("NewFeature()", func() {
		geom := geojson.NewGeometry(geojson.PointType, test.PointCoordinates, test.SrId)
		f := geojson.NewFeature(geom, nil)
		ginkgo.It("has Feature Type", func() {
			gomega.Expect(f.GetType()).To(gomega.Equal(geojson.FeatureType))
		})
		ginkgo.It("has Point Geometry", func() {
			gomega.Expect(f.GetGeometry()).NotTo(gomega.BeNil())
			pt, err := f.GetGeometry().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(pt.GetCoordinates()).To(gomega.Equal(test.PointCoordinates))
			gomega.Expect(pt.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Initialized Properties", func() {
			gomega.Expect(f.GetProperties()).NotTo(gomega.BeNil())
			gomega.Expect(len(*f.GetProperties())).To(gomega.Equal(0))
		})
		ginkgo.It("has Properties that was given into", func() {
			someProps := make(map[string]interface{})
			someProps["count"] = 42
			f := geojson.NewFeature(geom, &someProps)
			gomega.Expect(f.GetProperties()).NotTo(gomega.BeNil())
			gomega.Expect(len(*f.GetProperties())).To(gomega.Equal(1))
			gomega.Expect((*f.GetProperties())["count"]).To(gomega.Equal(42))
		})
	})
	ginkgo.Describe("Deserialize()", func() {
		feature, fErr := geojson.DeserializeFeature(test.FeatureGeoJSON)
		featureCrs, fCrsErr := geojson.DeserializeFeature(test.FeatureGeoJSONCrs)
		featureWithProps, fPropErr := geojson.DeserializeFeature(test.FeatureGeoJSONProps)
		ginkgo.It("has Feature Type", func() {
			gomega.Expect(fErr).To(gomega.BeNil())
			gomega.Expect(fCrsErr).To(gomega.BeNil())
			gomega.Expect(fPropErr).To(gomega.BeNil())
			gomega.Expect(feature.GetType()).To(gomega.Equal(geojson.FeatureType))
			gomega.Expect(featureWithProps.GetType()).To(gomega.Equal(geojson.FeatureType))
			gomega.Expect(featureCrs.GetType()).To(gomega.Equal(geojson.FeatureType))
		})
		ginkgo.It("has the Point as Geometry", func() {
			_, err := feature.GetGeometry().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			_, err = featureWithProps.GetGeometry().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			_, err = featureCrs.GetGeometry().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has 0 as SrId", func() {
			gomega.Expect(feature.GetGeometry().GetCRS().GetSrId()).To(gomega.Equal(0))
			gomega.Expect(featureWithProps.GetGeometry().GetCRS().GetSrId()).To(gomega.Equal(0))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(featureCrs.GetGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has empty Properties", func() {
			gomega.Expect(len(*feature.GetProperties())).To(gomega.Equal(0))
			gomega.Expect(len(*featureCrs.GetProperties())).To(gomega.Equal(0))
		})
		ginkgo.It("has Properties", func() {
			gomega.Expect(len(*featureWithProps.GetProperties())).To(gomega.Equal(1))
			gomega.Expect((*featureWithProps.GetProperties())["hello"]).To(gomega.Equal("world"))
		})
		ginkgo.It("deserialize array of features", func() {
			value, err := geojson.DeserializeFeatureList(test.FeatureGeoJSONArray)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(len(value)).To(gomega.Equal(1))
		})
	})
	ginkgo.Describe("Serialize()", func() {
		feature, fErr := geojson.DeserializeFeature(test.FeatureGeoJSON)
		featureCrs, fCrsErr := geojson.DeserializeFeature(test.FeatureGeoJSONCrs)
		featureWithProps, fPropErr := geojson.DeserializeFeature(test.FeatureGeoJSONProps)
		ginkgo.It("is Valid GeoJSON Feature", func() {
			gomega.Expect(fErr).To(gomega.BeNil())
			gomega.Expect(fCrsErr).To(gomega.BeNil())
			gomega.Expect(fPropErr).To(gomega.BeNil())
			gomega.Expect(feature.Serialize()).To(gomega.Equal(test.FeatureGeoJSON))
			gomega.Expect(featureCrs.Serialize()).To(gomega.Equal(test.FeatureGeoJSONCrs))
			gomega.Expect(featureWithProps.Serialize()).To(gomega.Equal(test.FeatureGeoJSONProps))
		})
	})
})
