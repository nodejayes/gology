package constants

type datum struct {
	ToWgs84   string
	Ellipse   string
	DatumName string
}

var WGS84Date = datum{
	ToWgs84:   "0,0,0",
	Ellipse:   "WGS84",
	DatumName: "WGS84",
}

var CH1903 = datum{
	ToWgs84:   "674.374,15.056,405.346",
	Ellipse:   "bessel",
	DatumName: "swiss",
}

var GGRS87 = datum{
	ToWgs84:   "-199.87,74.79,246.62",
	Ellipse:   "GRS80",
	DatumName: "Greek_Geodetic_Reference_System_1987",
}

var NAD83 = datum{
	ToWgs84:   "0,0,0",
	Ellipse:   "GRS80",
	DatumName: "North_American_Datum_1983",
}

var NAD27 = datum{
	ToWgs84:   "@conus,@alaska,@ntv2_0.gsb,@ntv1_can.dat",
	Ellipse:   "clrk66",
	DatumName: "North_American_Datum_1927",
}

var Potsdam = datum{
	ToWgs84:   "606.0,23.0,413.0",
	Ellipse:   "bessel",
	DatumName: "Potsdam Rauenberg 1950 DHDN",
}

var Carthage = datum{
	ToWgs84:   "-263.0,6.0,431.0",
	Ellipse:   "clark80",
	DatumName: "Carthage 1934 Tunisia",
}

var Hermannskogel = datum{
	ToWgs84:   "653.0,-212.0,449.0",
	Ellipse:   "bessel",
	DatumName: "Hermannskogel",
}

var OSNI52 = datum{
	ToWgs84:   "482.530,-130.596,564.557,-1.042,-0.214,-0.631,8.15",
	Ellipse:   "airy",
	DatumName: "Irish National",
}

var IRE65 = datum{
	ToWgs84:   "482.530,-130.596,564.557,-1.042,-0.214,-0.631,8.15",
	Ellipse:   "mod_airy",
	DatumName: "Ireland 1965",
}

var Rassadiran = datum{
	ToWgs84:   "-133.63,-157.5,-158.62",
	Ellipse:   "intl",
	DatumName: "Rassadiran",
}

var NZGD49 = datum{
	ToWgs84:   "59.47,-5.04,187.44,0.47,-0.1,1.024,-4.5993",
	Ellipse:   "intl",
	DatumName: "New Zealand Geodetic Datum 1949",
}

var OSGB36 = datum{
	ToWgs84:   "446.448,-125.157,542.060,0.1502,0.2470,0.8421,-20.4894",
	Ellipse:   "airy",
	DatumName: "Airy 1830",
}

var SJTSK = datum{
	ToWgs84:   "589,76,480",
	Ellipse:   "bessel",
	DatumName: "S-JTSK (Ferro)",
}

var Beduaram = datum{
	ToWgs84:   "-106,-87,188",
	Ellipse:   "clrk80",
	DatumName: "Beduaram",
}

var GunungSegara = datum{
	ToWgs84:   "-403,684,41",
	Ellipse:   "bessel",
	DatumName: "Gunung Segara Jakarta",
}

var RNB72 = datum{
	ToWgs84:   "106.869,-52.2978,103.724,-0.33657,0.456955,-1.84218,1",
	Ellipse:   "intl",
	DatumName: "Reseau National Belge 1972",
}
