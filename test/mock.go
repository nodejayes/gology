package test

import "fmt"

var SrId = 4326
var PointCoordinates = []float64{13.39185655117035, 51.30710261170077}
var PointGeoJSON = "{\"type\":\"Point\",\"coordinates\":[13.39185655117035,51.30710261170077]}"
var PointGeoJSONCrs = fmt.Sprintf("{\"type\":\"Point\",\"coordinates\":[13.39185655117035,51.30710261170077],\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"EPSG:%v\"}}}", SrId)
var LineCoordinates = [][]float64{
	{13.390472531318665, 51.307323945069115},
	{13.39232325553894, 51.30820591440422},
}
var LineGeoJSON = "{\"type\":\"LineString\",\"coordinates\":[[13.390472531318665,51.307323945069115],[13.39232325553894,51.30820591440422]]}"
var LineGeoJSONCrs = fmt.Sprintf("{\"type\":\"LineString\",\"coordinates\":[[13.390472531318665,51.307323945069115],[13.39232325553894,51.30820591440422]],\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"EPSG:%v\"}}}", SrId)
var PolygonCoordinates = [][][]float64{{
	{13.390740752220152, 51.30591208814888},
	{13.391990661621094, 51.30591208814888},
	{13.391990661621094, 51.30654256644033},
	{13.390740752220152, 51.30654256644033},
	{13.390740752220152, 51.30591208814888},
}}
var PolygonGeoJSON = "{\"type\":\"Polygon\",\"coordinates\":[[[13.390740752220152,51.30591208814888],[13.391990661621094,51.30591208814888],[13.391990661621094,51.30654256644033],[13.390740752220152,51.30654256644033],[13.390740752220152,51.30591208814888]]]}"
var PolygonGeoJSONCrs = fmt.Sprintf("{\"type\":\"Polygon\",\"coordinates\":[[[13.390740752220152,51.30591208814888],[13.391990661621094,51.30591208814888],[13.391990661621094,51.30654256644033],[13.390740752220152,51.30654256644033],[13.390740752220152,51.30591208814888]]],\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"EPSG:%v\"}}}", SrId)
var MultiPointCoordinates = [][]float64{
	{13.392763137817383, 51.30759222650863},
	{13.39185655117035, 51.30710261170077},
}
var MultiPointGeoJSON = "{\"type\":\"MultiPoint\",\"coordinates\":[[13.392763137817383,51.30759222650863],[13.39185655117035,51.30710261170077]]}"
var MultiPointGeoJSONCrs = fmt.Sprintf("{\"type\":\"MultiPoint\",\"coordinates\":[[13.392763137817383,51.30759222650863],[13.39185655117035,51.30710261170077]],\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"EPSG:%v\"}}}", SrId)
var MultiLineCoordinates = [][][]float64{
	{{13.392741680145264, 51.306150195330034}, {13.392945528030396, 51.30683097383521}},
	{{13.393229842185974, 51.30607306215311}, {13.39378237724304, 51.30697853040611}},
}
var MultiLineGeoJSON = "{\"type\":\"MultiLineString\",\"coordinates\":[[[13.392741680145264,51.306150195330034],[13.392945528030396,51.30683097383521]],[[13.393229842185974,51.30607306215311],[13.39378237724304,51.30697853040611]]]}"
var MultiLineGeoJSONCrs = fmt.Sprintf("{\"type\":\"MultiLineString\",\"coordinates\":[[[13.392741680145264,51.306150195330034],[13.392945528030396,51.30683097383521]],[[13.393229842185974,51.30607306215311],[13.39378237724304,51.30697853040611]]],\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"EPSG:%v\"}}}", SrId)
var MultiPolygonCoordinates = [][][][]float64{
	{{
		{13.387693762779236, 51.30805836177982}, {13.388970494270325, 51.30805836177982},
		{13.388970494270325, 51.30864521599963}, {13.387693762779236, 51.30864521599963},
		{13.387693762779236, 51.30805836177982},
	}},
	{{
		{13.386889100074768, 51.30635811890035}, {13.38933527469635, 51.30635811890035},
		{13.38933527469635, 51.307501681698135}, {13.386889100074768, 51.307501681698135},
		{13.386889100074768, 51.30635811890035},
	}},
}
var MultiPolygonGeoJSON = "{\"type\":\"MultiPolygon\",\"coordinates\":[[[[13.387693762779236,51.30805836177982],[13.388970494270325,51.30805836177982],[13.388970494270325,51.30864521599963],[13.387693762779236,51.30864521599963],[13.387693762779236,51.30805836177982]]],[[[13.386889100074768,51.30635811890035],[13.38933527469635,51.30635811890035],[13.38933527469635,51.307501681698135],[13.386889100074768,51.307501681698135],[13.386889100074768,51.30635811890035]]]]}"
var MultiPolygonGeoJSONCrs = fmt.Sprintf("{\"type\":\"MultiPolygon\",\"coordinates\":[[[[13.387693762779236,51.30805836177982],[13.388970494270325,51.30805836177982],[13.388970494270325,51.30864521599963],[13.387693762779236,51.30864521599963],[13.387693762779236,51.30805836177982]]],[[[13.386889100074768,51.30635811890035],[13.38933527469635,51.30635811890035],[13.38933527469635,51.307501681698135],[13.386889100074768,51.307501681698135],[13.386889100074768,51.30635811890035]]]],\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"EPSG:%v\"}}}", SrId)
var FeatureGeoJSON = fmt.Sprintf("{\"type\":\"Feature\",\"geometry\":%v,\"properties\":{}}", PointGeoJSON)
var FeatureGeoJSONProps = fmt.Sprintf("{\"type\":\"Feature\",\"geometry\":%v,\"properties\":{\"hello\":\"world\"}}", PointGeoJSON)
var FeatureGeoJSONCrs = fmt.Sprintf("{\"type\":\"Feature\",\"geometry\":%v,\"properties\":{}}", PointGeoJSONCrs)
var FeatureCollectionGeoJSON = fmt.Sprintf("{\"type\":\"FeatureCollection\",\"features\":[%v]}", FeatureGeoJSON)
var FeatureCollectionGeoJSONCrs = fmt.Sprintf("{\"type\":\"FeatureCollection\",\"features\":[%v],\"crs\":{\"type\":\"name\",\"properties\":{\"name\":\"EPSG:%v\"}}}", FeatureGeoJSONCrs, SrId)