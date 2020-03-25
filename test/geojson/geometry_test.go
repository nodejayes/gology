package geojson

import (
	"fmt"
	"github.com/nodejayes/topgology/geojson"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Geometry Test", func() {
	point := geojson.NewGeometry(geojson.PointType, PointCoordinates, SrId)
	line := geojson.NewGeometry(geojson.LineType, LineCoordinates, SrId)
	polygon := geojson.NewGeometry(geojson.PolygonType, PolygonCoordinates, SrId)
	multiPoint := geojson.NewGeometry(geojson.MultiPointType, MultiPointCoordinates, SrId)
	multiLine := geojson.NewGeometry(geojson.MultiLineType, MultiLineCoordinates, SrId)
	multiPolygon := geojson.NewGeometry(geojson.MultiPolygonType, MultiPolygonCoordinates, SrId)
	ginkgo.Describe("Getter", func() {
		tmp := geojson.NewGeometry(geojson.MultiPolygonType, MultiPolygonCoordinates, SrId)
		ginkgo.It("get the SrId", func() {
			gomega.Expect(tmp.GetSrId()).To(gomega.Equal(SrId))
		})
		ginkgo.It("invalid values return 0", func() {
			tmp.CRS.Properties.Name = ""
			gomega.Expect(tmp.GetSrId()).To(gomega.Equal(0))
			tmp.CRS.Properties.Name = "EPSG:ZZZZZZ"
			gomega.Expect(tmp.GetSrId()).To(gomega.Equal(0))
		})
	})
	ginkgo.Describe("GeometryTypes Check", func() {
		ginkgo.It("Point is a Point", func() {
			gomega.Expect(point.IsPoint()).To(gomega.BeTrue())
		})
		ginkgo.It("Point is not a Line", func() {
			gomega.Expect(point.IsLine()).To(gomega.BeFalse())
		})
		ginkgo.It("Point is not a Polygon", func() {
			gomega.Expect(point.IsPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("Point is not a MultiPoint", func() {
			gomega.Expect(point.IsMultiPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("Point is not a MultiLine", func() {
			gomega.Expect(point.IsMultiLine()).To(gomega.BeFalse())
		})
		ginkgo.It("Point is not a MultiPolygon", func() {
			gomega.Expect(point.IsMultiPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("Line is not a Point", func() {
			gomega.Expect(line.IsPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("Line is a Line", func() {
			gomega.Expect(line.IsLine()).To(gomega.BeTrue())
		})
		ginkgo.It("Line is not a Polygon", func() {
			gomega.Expect(line.IsPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("Line is not a MultiPoint", func() {
			gomega.Expect(line.IsMultiPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("Line is not a MultiLine", func() {
			gomega.Expect(line.IsMultiLine()).To(gomega.BeFalse())
		})
		ginkgo.It("Line is not a MultiPolygon", func() {
			gomega.Expect(line.IsMultiPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("Polygon is not a Point", func() {
			gomega.Expect(polygon.IsPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("Polygon is not a Line", func() {
			gomega.Expect(polygon.IsLine()).To(gomega.BeFalse())
		})
		ginkgo.It("Polygon is a Polygon", func() {
			gomega.Expect(polygon.IsPolygon()).To(gomega.BeTrue())
		})
		ginkgo.It("Polygon is not a MultiPoint", func() {
			gomega.Expect(polygon.IsMultiPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("Polygon is not a MultiLine", func() {
			gomega.Expect(polygon.IsMultiLine()).To(gomega.BeFalse())
		})
		ginkgo.It("Polygon is not a MultiPolygon", func() {
			gomega.Expect(polygon.IsMultiPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPoint is not a Point", func() {
			gomega.Expect(multiPoint.IsPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPoint is not a Line", func() {
			gomega.Expect(multiPoint.IsLine()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPoint is not a Polygon", func() {
			gomega.Expect(multiPoint.IsPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPoint is a MultiPoint", func() {
			gomega.Expect(multiPoint.IsMultiPoint()).To(gomega.BeTrue())
		})
		ginkgo.It("MultiPoint is not a MultiLine", func() {
			gomega.Expect(multiPoint.IsMultiLine()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPoint is not a MultiPolygon", func() {
			gomega.Expect(multiPoint.IsMultiPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiLine is not a Point", func() {
			gomega.Expect(multiLine.IsPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiLine is not a Line", func() {
			gomega.Expect(multiLine.IsLine()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiLine is not a Polygon", func() {
			gomega.Expect(multiLine.IsPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiLine is not a MultiPoint", func() {
			gomega.Expect(multiLine.IsMultiPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiLine is a MultiLine", func() {
			gomega.Expect(multiLine.IsMultiLine()).To(gomega.BeTrue())
		})
		ginkgo.It("MultiLine is not a MultiPolygon", func() {
			gomega.Expect(multiLine.IsMultiPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPolygon is not a Point", func() {
			gomega.Expect(multiPolygon.IsPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPolygon is not a Line", func() {
			gomega.Expect(multiPolygon.IsLine()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPolygon is not a Polygon", func() {
			gomega.Expect(multiPolygon.IsPolygon()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPolygon is not a MultiPoint", func() {
			gomega.Expect(multiPolygon.IsMultiPoint()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPolygon not is a MultiLine", func() {
			gomega.Expect(multiPolygon.IsMultiLine()).To(gomega.BeFalse())
		})
		ginkgo.It("MultiPolygon is a MultiPolygon", func() {
			gomega.Expect(multiPolygon.IsMultiPolygon()).To(gomega.BeTrue())
		})
	})
	ginkgo.Describe("NewGeometry()", func() {
		geom := geojson.NewGeometry(geojson.PointType, PointCoordinates, SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(geom.Type).To(gomega.Equal(geojson.PointType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.CRS.Properties.Name).To(gomega.Equal(fmt.Sprintf("EPSG:%v", SrId)))
		})
		ginkgo.It("has PointCoordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(PointCoordinates))
		})
	})
	ginkgo.Describe("AsPoint()", func() {
		geom, err := point.AsPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has Point Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(PointCoordinates))
		})
		ginkgo.Describe("can't be other Geometry", func() {
			ginkgo.It("AsLine() has a Error", func() {
				_, err = point.AsLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsPolygon() has a Error", func() {
				_, err = point.AsPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPoint() has a Error", func() {
				_, err = point.AsMultiPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiLine() has a Error", func() {
				_, err = point.AsMultiLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPolygon() has a Error", func() {
				_, err = point.AsMultiPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
	ginkgo.Describe("AsLine()", func() {
		geom, err := line.AsLine()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has Line Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(LineCoordinates))
		})
		ginkgo.Describe("can't be other Geometry", func() {
			ginkgo.It("AsPoint() has a Error", func() {
				_, err = line.AsPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsPolygon() has a Error", func() {
				_, err = line.AsPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPoint() has a Error", func() {
				_, err = line.AsMultiPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiLine() has a Error", func() {
				_, err = line.AsMultiLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPolygon() has a Error", func() {
				_, err = line.AsMultiPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
	ginkgo.Describe("AsPolygon()", func() {
		geom, err := polygon.AsPolygon()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has Polygon Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(PolygonCoordinates))
		})
		ginkgo.Describe("can't be other Geometry", func() {
			ginkgo.It("AsPoint() has a Error", func() {
				_, err = polygon.AsPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsLine() has a Error", func() {
				_, err = polygon.AsLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPoint() has a Error", func() {
				_, err = polygon.AsMultiPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiLine() has a Error", func() {
				_, err = polygon.AsMultiLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPolygon() has a Error", func() {
				_, err = polygon.AsMultiPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
	ginkgo.Describe("AsMultiPoint()", func() {
		geom, err := multiPoint.AsMultiPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(MultiPointCoordinates))
		})
		ginkgo.Describe("can't be other Geometry", func() {
			ginkgo.It("AsPoint() has a Error", func() {
				_, err = multiPoint.AsPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsLine() has a Error", func() {
				_, err = multiPoint.AsLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsPolygon() has a Error", func() {
				_, err = multiPoint.AsPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiLine() has a Error", func() {
				_, err = multiPoint.AsMultiLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPolygon() has a Error", func() {
				_, err = multiPoint.AsMultiPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
	ginkgo.Describe("AsMultiLine()", func() {
		geom, err := multiLine.AsMultiLine()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has MultiLine Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(MultiLineCoordinates))
		})
		ginkgo.Describe("can't be other Geometry", func() {
			ginkgo.It("AsPoint() has a Error", func() {
				_, err = multiLine.AsPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsLine() has a Error", func() {
				_, err = multiLine.AsLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsPolygon() has a Error", func() {
				_, err = multiLine.AsPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPoint() has a Error", func() {
				_, err = multiLine.AsMultiPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPolygon() has a Error", func() {
				_, err = multiLine.AsMultiPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
	ginkgo.Describe("AsMultiPolygon()", func() {
		geom, err := multiPolygon.AsMultiPolygon()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.SrId).To(gomega.Equal(SrId))
		})
		ginkgo.It("has MultiPolygon Coordinates", func() {
			gomega.Expect(geom.Coordinates).To(gomega.Equal(MultiPolygonCoordinates))
		})
		ginkgo.Describe("can't be other Geometry", func() {
			ginkgo.It("AsPoint() has a Error", func() {
				_, err = multiPolygon.AsPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsLine() has a Error", func() {
				_, err = multiPolygon.AsLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsPolygon() has a Error", func() {
				_, err = multiPolygon.AsPolygon()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiPoint() has a Error", func() {
				_, err = multiPolygon.AsMultiPoint()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
			ginkgo.It("AsMultiLine() has a Error", func() {
				_, err = multiPolygon.AsMultiLine()
				gomega.Expect(err).ToNot(gomega.BeNil())
			})
		})
	})
	ginkgo.Describe("Deserialize Test", func() {
		nothing := geojson.DeserializeGeometry("")
		ginkgo.It("invalid returns empty string", func() {
			gomega.Expect(nothing).To(gomega.BeNil())
		})
		ginkgo.Describe("without CRS", func() {
			point := geojson.DeserializeGeometry(PointGeoJSON)
			line := geojson.DeserializeGeometry(LineGeoJSON)
			polygon := geojson.DeserializeGeometry(PolygonGeoJSON)
			multiPoint := geojson.DeserializeGeometry(MultiPointGeoJSON)
			multiLine := geojson.DeserializeGeometry(MultiLineGeoJSON)
			multiPolygon := geojson.DeserializeGeometry(MultiPolygonGeoJSON)
			ginkgo.It("is a valid Point", func() {
				gomega.Expect(point).ToNot(gomega.BeNil())
				gomega.Expect(point.IsPoint()).To(gomega.BeTrue())
				gomega.Expect(point.GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid Line", func() {
				gomega.Expect(line).ToNot(gomega.BeNil())
				gomega.Expect(line.IsLine()).To(gomega.BeTrue())
				gomega.Expect(line.GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid Polygon", func() {
				gomega.Expect(polygon).ToNot(gomega.BeNil())
				gomega.Expect(polygon.IsPolygon()).To(gomega.BeTrue())
				gomega.Expect(polygon.GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid MultiPoint", func() {
				gomega.Expect(multiPoint).ToNot(gomega.BeNil())
				gomega.Expect(multiPoint.IsMultiPoint()).To(gomega.BeTrue())
				gomega.Expect(multiPoint.GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid MultiLine", func() {
				gomega.Expect(multiLine).ToNot(gomega.BeNil())
				gomega.Expect(multiLine.IsMultiLine()).To(gomega.BeTrue())
				gomega.Expect(multiLine.GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid MultiPolygon", func() {
				gomega.Expect(multiPolygon).ToNot(gomega.BeNil())
				gomega.Expect(multiPolygon.IsMultiPolygon()).To(gomega.BeTrue())
				gomega.Expect(multiPolygon.GetSrId()).To(gomega.Equal(0))
			})
		})
		ginkgo.Describe("with CRS", func() {
			point := geojson.DeserializeGeometry(PointGeoJSONCrs)
			line := geojson.DeserializeGeometry(LineGeoJSONCrs)
			polygon := geojson.DeserializeGeometry(PolygonGeoJSONCrs)
			multiPoint := geojson.DeserializeGeometry(MultiPointGeoJSONCrs)
			multiLine := geojson.DeserializeGeometry(MultiLineGeoJSONCrs)
			multiPolygon := geojson.DeserializeGeometry(MultiPolygonGeoJSONCrs)
			ginkgo.It("is a valid Point", func() {
				gomega.Expect(point).ToNot(gomega.BeNil())
				gomega.Expect(point.IsPoint()).To(gomega.BeTrue())
				gomega.Expect(point.GetSrId()).To(gomega.Equal(SrId))
			})
			ginkgo.It("is a valid Line", func() {
				gomega.Expect(line).ToNot(gomega.BeNil())
				gomega.Expect(line.IsLine()).To(gomega.BeTrue())
				gomega.Expect(line.GetSrId()).To(gomega.Equal(SrId))
			})
			ginkgo.It("is a valid Polygon", func() {
				gomega.Expect(polygon).ToNot(gomega.BeNil())
				gomega.Expect(polygon.IsPolygon()).To(gomega.BeTrue())
				gomega.Expect(polygon.GetSrId()).To(gomega.Equal(SrId))
			})
			ginkgo.It("is a valid MultiPoint", func() {
				gomega.Expect(multiPoint).ToNot(gomega.BeNil())
				gomega.Expect(multiPoint.IsMultiPoint()).To(gomega.BeTrue())
				gomega.Expect(multiPoint.GetSrId()).To(gomega.Equal(SrId))
			})
			ginkgo.It("is a valid MultiLine", func() {
				gomega.Expect(multiLine).ToNot(gomega.BeNil())
				gomega.Expect(multiLine.IsMultiLine()).To(gomega.BeTrue())
				gomega.Expect(multiLine.GetSrId()).To(gomega.Equal(SrId))
			})
			ginkgo.It("is a valid MultiPolygon", func() {
				gomega.Expect(multiPolygon).ToNot(gomega.BeNil())
				gomega.Expect(multiPolygon.IsMultiPolygon()).To(gomega.BeTrue())
				gomega.Expect(multiPolygon.GetSrId()).To(gomega.Equal(SrId))
			})
		})
	})
	ginkgo.Describe("Serialize Test", func() {
		ginkgo.Describe("without CRS", func() {
			point := geojson.DeserializeGeometry(PointGeoJSON).Serialize()
			line := geojson.DeserializeGeometry(LineGeoJSON).Serialize()
			polygon := geojson.DeserializeGeometry(PolygonGeoJSON).Serialize()
			multiPoint := geojson.DeserializeGeometry(MultiPointGeoJSON).Serialize()
			multiLine := geojson.DeserializeGeometry(MultiLineGeoJSON).Serialize()
			multiPolygon := geojson.DeserializeGeometry(MultiPolygonGeoJSON).Serialize()
			ginkgo.It("is a valid GeoJSON Point", func() {
				gomega.Expect(point).To(gomega.Equal(PointGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON Line", func() {
				gomega.Expect(line).To(gomega.Equal(LineGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON Polygon", func() {
				gomega.Expect(polygon).To(gomega.Equal(PolygonGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON MultiPoint", func() {
				gomega.Expect(multiPoint).To(gomega.Equal(MultiPointGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON MultiLine", func() {
				gomega.Expect(multiLine).To(gomega.Equal(MultiLineGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON MultiPolygon", func() {
				gomega.Expect(multiPolygon).To(gomega.Equal(MultiPolygonGeoJSON))
			})
		})
		ginkgo.Describe("with CRS", func() {
			point := geojson.DeserializeGeometry(PointGeoJSONCrs).Serialize()
			line := geojson.DeserializeGeometry(LineGeoJSONCrs).Serialize()
			polygon := geojson.DeserializeGeometry(PolygonGeoJSONCrs).Serialize()
			multiPoint := geojson.DeserializeGeometry(MultiPointGeoJSONCrs).Serialize()
			multiLine := geojson.DeserializeGeometry(MultiLineGeoJSONCrs).Serialize()
			multiPolygon := geojson.DeserializeGeometry(MultiPolygonGeoJSONCrs).Serialize()
			ginkgo.It("is a valid GeoJSON Point", func() {
				gomega.Expect(point).To(gomega.Equal(PointGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON Line", func() {
				gomega.Expect(line).To(gomega.Equal(LineGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON Polygon", func() {
				gomega.Expect(polygon).To(gomega.Equal(PolygonGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON MultiPoint", func() {
				gomega.Expect(multiPoint).To(gomega.Equal(MultiPointGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON MultiLine", func() {
				gomega.Expect(multiLine).To(gomega.Equal(MultiLineGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON MultiPolygon", func() {
				gomega.Expect(multiPolygon).To(gomega.Equal(MultiPolygonGeoJSONCrs))
			})
		})
	})
})
