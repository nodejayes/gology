package geojson

import (
	"github.com/nodejayes/gology/geojson"
	"github.com/nodejayes/gology/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Geometry Test", func() {
	point := geojson.NewGeometry(geojson.PointType, test.PointCoordinates, test.SrId)
	line := geojson.NewGeometry(geojson.LineType, test.LineCoordinates, test.SrId)
	polygon := geojson.NewGeometry(geojson.PolygonType, test.PolygonCoordinates, test.SrId)
	multiPoint := geojson.NewGeometry(geojson.MultiPointType, test.MultiPointCoordinates, test.SrId)
	multiLine := geojson.NewGeometry(geojson.MultiLineType, test.MultiLineCoordinates, test.SrId)
	multiPolygon := geojson.NewGeometry(geojson.MultiPolygonType, test.MultiPolygonCoordinates, test.SrId)
	ginkgo.Describe("Getter", func() {
		tmp := geojson.NewGeometry(geojson.MultiPolygonType, test.MultiPolygonCoordinates, test.SrId)
		ginkgo.It("get the SrId", func() {
			gomega.Expect(tmp.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("invalid values return 0", func() {
			pt, err := geojson.NewPoint([]float64{1.5, 2.5}, -1)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(pt.GetSrId()).To(gomega.Equal(0))
		})
	})
	ginkgo.Describe("GeometryTypes Check", func() {
		ginkgo.It("Point is a Point", func() {
			_, err := point.AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("Point is not a Line", func() {
			_, err := point.AsLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Point is not a Polygon", func() {
			_, err := point.AsPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Point is not a MultiPoint", func() {
			_, err := point.AsMultiPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Point is not a MultiLine", func() {
			_, err := point.AsMultiLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Point is not a MultiPolygon", func() {
			_, err := point.AsMultiPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Line is not a Point", func() {
			_, err := line.AsPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Line is a Line", func() {
			_, err := line.AsLine()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("Line is not a Polygon", func() {
			_, err := line.AsPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Line is a MultiPoint", func() {
			_, err := line.AsMultiPoint()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("Line is not a MultiLine", func() {
			_, err := line.AsMultiLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Line is not a MultiPolygon", func() {
			_, err := line.AsPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Polygon is not a Point", func() {
			_, err := polygon.AsPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Polygon is not a Line", func() {
			_, err := polygon.AsLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Polygon is a Polygon", func() {
			_, err := polygon.AsPolygon()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("Polygon is not a MultiPoint", func() {
			_, err := polygon.AsMultiPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("Polygon is a MultiLine", func() {
			_, err := polygon.AsMultiLine()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("Polygon is not a MultiPolygon", func() {
			_, err := polygon.AsMultiPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPoint is not a Point", func() {
			_, err := multiPoint.AsPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPoint is a Line", func() {
			_, err := multiPoint.AsLine()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("MultiPoint is not a Polygon", func() {
			_, err := multiPoint.AsPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPoint is a MultiPoint", func() {
			_, err := multiPoint.AsMultiPoint()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("MultiPoint is not a MultiLine", func() {
			_, err := multiPoint.AsMultiLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPoint is not a MultiPolygon", func() {
			_, err := multiPoint.AsMultiPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiLine is not a Point", func() {
			_, err := multiLine.AsPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiLine is not a Line", func() {
			_, err := multiLine.AsLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiLine is a Polygon", func() {
			_, err := multiLine.AsPolygon()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("MultiLine is not a MultiPoint", func() {
			_, err := multiLine.AsMultiPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiLine is a MultiLine", func() {
			_, err := multiLine.AsMultiLine()
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("MultiLine is not a MultiPolygon", func() {
			_, err := multiLine.AsMultiPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPolygon is not a Point", func() {
			_, err := multiPolygon.AsPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPolygon is not a Line", func() {
			_, err := multiPolygon.AsLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPolygon is not a Polygon", func() {
			_, err := multiPolygon.AsPolygon()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPolygon is not a MultiPoint", func() {
			_, err := multiPolygon.AsMultiPoint()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPolygon not is a MultiLine", func() {
			_, err := multiPolygon.AsMultiLine()
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.It("MultiPolygon is a MultiPolygon", func() {
			_, err := multiPolygon.AsMultiPolygon()
			gomega.Expect(err).To(gomega.BeNil())
		})
	})
	ginkgo.Describe("NewGeometry()", func() {
		geom := geojson.NewGeometry(geojson.PointType, test.PointCoordinates, test.SrId)
		ginkgo.It("has no Error", func() {
			gomega.Expect(geom.GetType()).To(gomega.Equal(geojson.PointType))
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has PointCoordinates", func() {
			coords, err := geom.GetCoordinates().AsPoint()
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(coords).To(gomega.Equal(test.PointCoordinates))
		})
	})
	ginkgo.Describe("AsPoint()", func() {
		geom, err := point.AsPoint()
		ginkgo.It("has no Error", func() {
			gomega.Expect(err).To(gomega.BeNil())
		})
		ginkgo.It("has SrId", func() {
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Point Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.PointCoordinates))
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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Line Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.LineCoordinates))
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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has Polygon Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.PolygonCoordinates))
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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPoint Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.MultiPointCoordinates))
		})
		ginkgo.Describe("can't be other Geometry", func() {
			ginkgo.It("AsPoint() has a Error", func() {
				_, err = multiPoint.AsPoint()
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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiLine Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.MultiLineCoordinates))
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
			gomega.Expect(geom.GetSrId()).To(gomega.Equal(test.SrId))
		})
		ginkgo.It("has MultiPolygon Coordinates", func() {
			gomega.Expect(geom.GetCoordinates()).To(gomega.Equal(test.MultiPolygonCoordinates))
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
		nothing, err := geojson.DeserializeGeometry("")
		ginkgo.It("invalid returns empty string", func() {
			gomega.Expect(nothing).To(gomega.BeNil())
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
		ginkgo.Describe("without CRS", func() {
			point, pointErr := geojson.DeserializeGeometry(test.PointGeoJSON)
			line, lineErr := geojson.DeserializeGeometry(test.LineGeoJSON)
			polygon, polygonErr := geojson.DeserializeGeometry(test.PolygonGeoJSON)
			multiPoint, multiPointErr := geojson.DeserializeGeometry(test.MultiPointGeoJSON)
			multiLine, multiLineErr := geojson.DeserializeGeometry(test.MultiLineGeoJSON)
			multiPolygon, multiPolygonErr := geojson.DeserializeGeometry(test.MultiPolygonGeoJSON)
			ginkgo.It("is a valid Point", func() {
				_, err := point.AsPoint()
				gomega.Expect(point).NotTo(gomega.BeNil())
				gomega.Expect(pointErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(point.GetCRS().GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid Line", func() {
				_, err := line.AsLine()
				gomega.Expect(line).ToNot(gomega.BeNil())
				gomega.Expect(lineErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(line.GetCRS().GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid Polygon", func() {
				_, err := polygon.AsPolygon()
				gomega.Expect(polygon).NotTo(gomega.BeNil())
				gomega.Expect(polygonErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(polygon.GetCRS().GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid MultiPoint", func() {
				_, err := multiPoint.AsMultiPoint()
				gomega.Expect(multiPoint).NotTo(gomega.BeNil())
				gomega.Expect(multiPointErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(multiPoint.GetCRS().GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid MultiLine", func() {
				_, err := multiLine.AsLine()
				gomega.Expect(multiLine).NotTo(gomega.BeNil())
				gomega.Expect(multiLineErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(multiLine.GetCRS().GetSrId()).To(gomega.Equal(0))
			})
			ginkgo.It("is a valid MultiPolygon", func() {
				_, err := multiPolygon.AsMultiPolygon()
				gomega.Expect(multiPolygon).NotTo(gomega.BeNil())
				gomega.Expect(multiPolygonErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(multiPolygon.GetCRS().GetSrId()).To(gomega.Equal(0))
			})
		})
		ginkgo.Describe("with CRS", func() {
			point, pointErr := geojson.DeserializeGeometry(test.PointGeoJSONCrs)
			line, lineErr := geojson.DeserializeGeometry(test.LineGeoJSONCrs)
			polygon, polygonErr := geojson.DeserializeGeometry(test.PolygonGeoJSONCrs)
			multiPoint, multiPointErr := geojson.DeserializeGeometry(test.MultiPointGeoJSONCrs)
			multiLine, multiLineErr := geojson.DeserializeGeometry(test.MultiLineGeoJSONCrs)
			multiPolygon, multiPolygonErr := geojson.DeserializeGeometry(test.MultiPolygonGeoJSONCrs)
			ginkgo.It("is a valid Point", func() {
				_, err := point.AsPoint()
				gomega.Expect(point).NotTo(gomega.BeNil())
				gomega.Expect(pointErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(point.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			})
			ginkgo.It("is a valid Line", func() {
				_, err := line.AsLine()
				gomega.Expect(line).NotTo(gomega.BeNil())
				gomega.Expect(lineErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(line.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			})
			ginkgo.It("is a valid Polygon", func() {
				_, err := polygon.AsPolygon()
				gomega.Expect(polygon).NotTo(gomega.BeNil())
				gomega.Expect(polygonErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(polygon.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			})
			ginkgo.It("is a valid MultiPoint", func() {
				_, err := multiPoint.AsMultiPoint()
				gomega.Expect(multiPoint).NotTo(gomega.BeNil())
				gomega.Expect(multiPointErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(multiPoint.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			})
			ginkgo.It("is a valid MultiLine", func() {
				_, err := multiLine.AsMultiLine()
				gomega.Expect(multiLine).NotTo(gomega.BeNil())
				gomega.Expect(multiLineErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(multiLine.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			})
			ginkgo.It("is a valid MultiPolygon", func() {
				_, err := multiPolygon.AsMultiPolygon()
				gomega.Expect(multiPolygon).NotTo(gomega.BeNil())
				gomega.Expect(multiPolygonErr).To(gomega.BeNil())
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(multiPolygon.GetCRS().GetSrId()).To(gomega.Equal(test.SrId))
			})
		})
	})
	ginkgo.Describe("Serialize Test", func() {
		ginkgo.Describe("without CRS", func() {
			point, _ := geojson.DeserializeGeometry(test.PointGeoJSON)
			line, _ := geojson.DeserializeGeometry(test.LineGeoJSON)
			polygon, _ := geojson.DeserializeGeometry(test.PolygonGeoJSON)
			multiPoint, _ := geojson.DeserializeGeometry(test.MultiPointGeoJSON)
			multiLine, _ := geojson.DeserializeGeometry(test.MultiLineGeoJSON)
			multiPolygon, _ := geojson.DeserializeGeometry(test.MultiPolygonGeoJSON)
			ginkgo.It("is a valid GeoJSON Point", func() {
				gomega.Expect(point.Serialize()).To(gomega.Equal(test.PointGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON Line", func() {
				gomega.Expect(line.Serialize()).To(gomega.Equal(test.LineGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON Polygon", func() {
				gomega.Expect(polygon.Serialize()).To(gomega.Equal(test.PolygonGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON MultiPoint", func() {
				gomega.Expect(multiPoint.Serialize()).To(gomega.Equal(test.MultiPointGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON MultiLine", func() {
				gomega.Expect(multiLine.Serialize()).To(gomega.Equal(test.MultiLineGeoJSON))
			})
			ginkgo.It("is a valid GeoJSON MultiPolygon", func() {
				t := multiPolygon.Serialize()
				gomega.Expect(t).To(gomega.Equal(test.MultiPolygonGeoJSON))
			})
		})
		ginkgo.Describe("with CRS", func() {
			point, _ := geojson.DeserializeGeometry(test.PointGeoJSONCrs)
			line, _ := geojson.DeserializeGeometry(test.LineGeoJSONCrs)
			polygon, _ := geojson.DeserializeGeometry(test.PolygonGeoJSONCrs)
			multiPoint, _ := geojson.DeserializeGeometry(test.MultiPointGeoJSONCrs)
			multiLine, _ := geojson.DeserializeGeometry(test.MultiLineGeoJSONCrs)
			multiPolygon, _ := geojson.DeserializeGeometry(test.MultiPolygonGeoJSONCrs)
			ginkgo.It("is a valid GeoJSON Point", func() {
				gomega.Expect(point.Serialize()).To(gomega.Equal(test.PointGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON Line", func() {
				gomega.Expect(line.Serialize()).To(gomega.Equal(test.LineGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON Polygon", func() {
				gomega.Expect(polygon.Serialize()).To(gomega.Equal(test.PolygonGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON MultiPoint", func() {
				gomega.Expect(multiPoint.Serialize()).To(gomega.Equal(test.MultiPointGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON MultiLine", func() {
				gomega.Expect(multiLine.Serialize()).To(gomega.Equal(test.MultiLineGeoJSONCrs))
			})
			ginkgo.It("is a valid GeoJSON MultiPolygon", func() {
				gomega.Expect(multiPolygon.Serialize()).To(gomega.Equal(test.MultiPolygonGeoJSONCrs))
			})
		})
	})
})
