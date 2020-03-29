package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("FeatureCollection Test", func() {
	ginkgo.Describe("NewFeatureCollection()", func() {
		geom1 := geojson.NewGeometry(geojson.PointType, test.PointCoordinates, test.SrId)
		geom2 := geojson.NewGeometry(geojson.LineType, test.LineCoordinates, test.SrId)
		f1 := geojson.NewFeature(geom1, nil)
		f2 := geojson.NewFeature(geom2, nil)
		fc := geojson.NewFeatureCollection([]geojson.IFeature{f1, f2})

		ginkgo.It("has 2 Features", func() {
			gomega.Expect(len(fc.GetFeatures())).To(gomega.Equal(2))
		})
		ginkgo.It("has the SrId", func() {
			for _, f := range fc.GetFeatures() {
				gomega.Expect(f.GetGeometry().GetCRS().GetSrId()).To(gomega.Equal(fc.GetSrId()))
			}
		})
		ginkgo.It("has as first the Point", func() {
			p := fc.GetFeatures()[0]
			coords, err := p.GetGeometry().GetCoordinates().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(p.GetGeometry().GetType()).To(gomega.Equal(geojson.PointType))
			gomega.Expect(p.GetGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			gomega.Expect(coords).To(gomega.Equal(test.PointCoordinates))
		})
		ginkgo.It("has as second the Line", func() {
			l := fc.GetFeatures()[1]
			coords, err := l.GetGeometry().GetCoordinates().AsLine()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(l.GetGeometry().GetType()).To(gomega.Equal(geojson.LineType))
			gomega.Expect(l.GetGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			gomega.Expect(coords).To(gomega.Equal(test.LineCoordinates))
		})
		ginkgo.Describe("without CRS", func() {
			geom1 := geojson.NewGeometry(geojson.PointType, test.PointCoordinates, 0)
			geom2 := geojson.NewGeometry(geojson.LineType, test.LineCoordinates, 0)
			f1 := geojson.NewFeature(geom1, nil)
			f2 := geojson.NewFeature(geom2, nil)
			fc := geojson.NewFeatureCollection([]geojson.IFeature{f1, f2})
			ginkgo.It("has no SrId", func() {
				gomega.Expect(fc.GetSrId()).To(gomega.Equal(0))
			})
		})
	})
	ginkgo.Describe("Deserialize()", func() {
		featureCollection := geojson.DeserializeFeatureCollection(test.FeatureCollectionGeoJSON)
		featureCollectionCrs := geojson.DeserializeFeatureCollection(test.FeatureCollectionGeoJSONCrs)
		ginkgo.It("has Feature Type", func() {
			gomega.Expect(featureCollection.GetType()).To(gomega.Equal(geojson.FeatureCollectionType))
			gomega.Expect(featureCollectionCrs.GetType()).To(gomega.Equal(geojson.FeatureCollectionType))
		})
		ginkgo.It("has Features", func() {
			gomega.Expect(len(featureCollection.GetFeatures())).To(gomega.Equal(1))
			gomega.Expect(len(featureCollectionCrs.GetFeatures())).To(gomega.Equal(1))
		})
		ginkgo.It("has the Point as Geometry", func() {
			_, err := featureCollection.GetFeatures()[0].GetGeometry().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			_, err = featureCollectionCrs.GetFeatures()[0].GetGeometry().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has 0 as SrId", func() {
			firstFeature := featureCollection.GetFeatures()[0]
			geom := firstFeature.GetGeometry()
			geomCrs := geom.GetCRS()
			gomega.Expect(geomCrs.GetSrId()).To(gomega.Equal(0))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(featureCollectionCrs.GetFeatures()[0].GetGeometry().GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
	})
	ginkgo.Describe("Serialize()", func() {
		featureCollection := geojson.DeserializeFeatureCollection(test.FeatureCollectionGeoJSON)
		featureCollectionCrs := geojson.DeserializeFeatureCollection(test.FeatureCollectionGeoJSONCrs)
		ginkgo.It("is Valid GeoJSON FeatureCollection", func() {
			gomega.Expect(featureCollection.Serialize()).To(gomega.Equal(test.FeatureCollectionGeoJSON))
			gomega.Expect(featureCollectionCrs.Serialize()).To(gomega.Equal(test.FeatureCollectionGeoJSONCrs))
		})
	})
})
