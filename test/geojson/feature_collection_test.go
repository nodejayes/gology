package geojson

import (
	"github.com/nodejayes/topgology/geojson"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("FeatureCollection Test", func() {
	ginkgo.Describe("NewFeatureCollection()", func() {
		geom1 := geojson.NewGeometry(geojson.PointType, PointCoordinates, SrId)
		geom2 := geojson.NewGeometry(geojson.LineType, LineCoordinates, SrId)
		f1 := geojson.NewFeature(geom1, nil)
		f2 := geojson.NewFeature(geom2, nil)
		fc := geojson.NewFeatureCollection([]*geojson.Feature{f1, f2})

		ginkgo.It("has 2 Features", func() {
			gomega.Expect(len(fc.Features)).To(gomega.Equal(2))
		})
		ginkgo.It("has the SrId", func() {
			for _, f := range fc.Features {
				gomega.Expect(f.Geometry.GetSrId()).To(gomega.Equal(fc.GetSrId()))
			}
		})
		ginkgo.It("has as first the Point", func() {
			p := fc.Features[0]
			gomega.Expect(p.Geometry.Type).To(gomega.Equal(geojson.PointType))
			gomega.Expect(p.Geometry.CRS.GetSrId()).To(gomega.Equal(SrId))
			gomega.Expect(p.Geometry.Coordinates).To(gomega.Equal(PointCoordinates))
		})
		ginkgo.It("has as second the Line", func() {
			l := fc.Features[1]
			gomega.Expect(l.Geometry.Type).To(gomega.Equal(geojson.LineType))
			gomega.Expect(l.Geometry.CRS.GetSrId()).To(gomega.Equal(SrId))
			gomega.Expect(l.Geometry.Coordinates).To(gomega.Equal(LineCoordinates))
		})
		ginkgo.Describe("without CRS", func() {
			geom1 := geojson.NewGeometry(geojson.PointType, PointCoordinates, 0)
			geom2 := geojson.NewGeometry(geojson.LineType, LineCoordinates, 0)
			f1 := geojson.NewFeature(geom1, nil)
			f2 := geojson.NewFeature(geom2, nil)
			fc := geojson.NewFeatureCollection([]*geojson.Feature{f1, f2})
			ginkgo.It("has no SrId", func() {
				gomega.Expect(fc.GetSrId()).To(gomega.Equal(0))
			})
		})
	})
	ginkgo.Describe("Deserialize()", func() {
		featureCollection := geojson.DeserializeFeatureCollection(FeatureCollectionGeoJSON)
		featureCollectionCrs := geojson.DeserializeFeatureCollection(FeatureCollectionGeoJSONCrs)
		ginkgo.It("has Feature Type", func() {
			gomega.Expect(featureCollection.Type).To(gomega.Equal(geojson.FeatureCollectionType))
			gomega.Expect(featureCollectionCrs.Type).To(gomega.Equal(geojson.FeatureCollectionType))
		})
		ginkgo.It("has Features", func() {
			gomega.Expect(len(featureCollection.Features)).To(gomega.Equal(1))
			gomega.Expect(len(featureCollectionCrs.Features)).To(gomega.Equal(1))
		})
		ginkgo.It("has the Point as Geometry", func() {
			gomega.Expect(featureCollection.Features[0].Geometry.IsPoint()).To(gomega.BeTrue())
			gomega.Expect(featureCollectionCrs.Features[0].Geometry.IsPoint()).To(gomega.BeTrue())
		})
		ginkgo.It("has 0 as SrId", func() {
			gomega.Expect(featureCollection.Features[0].Geometry.GetSrId()).To(gomega.Equal(0))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(featureCollectionCrs.Features[0].Geometry.GetSrId()).To(gomega.Equal(SrId))
		})
	})
	ginkgo.Describe("Serialize()", func() {
		featureCollection := geojson.DeserializeFeatureCollection(FeatureCollectionGeoJSON)
		featureCollectionCrs := geojson.DeserializeFeatureCollection(FeatureCollectionGeoJSONCrs)
		ginkgo.It("is Valid GeoJSON FeatureCollection", func() {
			gomega.Expect(featureCollection.Serialize()).To(gomega.Equal(FeatureCollectionGeoJSON))
			gomega.Expect(featureCollectionCrs.Serialize()).To(gomega.Equal(FeatureCollectionGeoJSONCrs))
		})
	})
})
