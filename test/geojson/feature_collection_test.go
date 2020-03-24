package geojson

import (
	"git.agricon.de/map-factory/pkg/geojson"
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
})
