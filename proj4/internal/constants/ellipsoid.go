package constants

type ellipsoid struct {
	a           float64
	b           float64
	ellipseName string
}

var MERIT = ellipsoid{
	a:           6378137.0,
	b:           298.257,
	ellipseName: "MERIT 1983",
}

var SGS85 = ellipsoid{
	a:           6378136.0,
	b:           298.257,
	ellipseName: "Soviet Geodetic System 85",
}

var GRS80 = ellipsoid{
	a:           6378137.0,
	b:           298.257222101,
	ellipseName: "GRS 1980(IUGG, 1980)",
}

var IAU76 = ellipsoid{
	a:           6378140.0,
	b:           298.257,
	ellipseName: "IAU 1976",
}

var airy = ellipsoid{
	a:           6377563.396,
	b:           6356256.910,
	ellipseName: "Airy 1830",
}

var APL4 = ellipsoid{
	a:           6378137,
	b:           298.25,
	ellipseName: "Appl. Physics. 1965",
}

var NWL9D = ellipsoid{
	a:           6378145.0,
	b:           298.25,
	ellipseName: "Naval Weapons Lab., 1965",
}

var mod_airy = ellipsoid{
	a:           6377340.189,
	b:           6356034.446,
	ellipseName: "Modified Airy",
}

var andrae = ellipsoid{
	a:           6377104.43,
	b:           300.0,
	ellipseName: "Andrae 1876 (Den., Iclnd.)",
}

var aust_SA = ellipsoid{
	a:           6378160.0,
	b:           298.25,
	ellipseName: "Australian Natl & S. Amer. 1969",
}

var GRS67 = ellipsoid{
	a:           6378160.0,
	b:           298.2471674270,
	ellipseName: "GRS 67(IUGG 1967)",
}

var bessel = ellipsoid{
	a:           6377397.155,
	b:           299.1528128,
	ellipseName: "Bessel 1841",
}

var bess_nam = ellipsoid{
	a:           6377483.865,
	b:           299.1528128,
	ellipseName: "Bessel 1841 (Namibia)",
}

var clrk66 = ellipsoid{
	a:           6378206.4,
	b:           6356583.8,
	ellipseName: "Clarke 1866",
}

var clrk80 = ellipsoid{
	a:           6378249.145,
	b:           293.4663,
	ellipseName: "Clarke 1880 mod.",
}

var clrk58 = ellipsoid{
	a:           6378293.645208759,
	b:           294.2606763692654,
	ellipseName: "Clarke 1858",
}

var CPM = ellipsoid{
	a:           6375738.7,
	b:           334.29,
	ellipseName: "Comm. des Poids et Mesures 1799",
}

var delmbr = ellipsoid{
	a:           6376428.0,
	b:           311.5,
	ellipseName: "Delambre 1810 (Belgium)",
}

var engelis = ellipsoid{
	a:           6378136.05,
	b:           298.2566,
	ellipseName: "Engelis 1985",
}

var evrst30 = ellipsoid{
	a:           6377276.345,
	b:           300.8017,
	ellipseName: "Everest 1830",
}

var evrst48 = ellipsoid{
	a:           6377304.063,
	b:           300.8017,
	ellipseName: "Everest 1948",
}

var evrst56 = ellipsoid{
	a:           6377301.243,
	b:           300.8017,
	ellipseName: "Everest 1956",
}

var evrst69 = ellipsoid{
	a:           6377295.664,
	b:           300.8017,
	ellipseName: "Everest 1969",
}

var evrstSS = ellipsoid{
	a:           6377298.556,
	b:           300.8017,
	ellipseName: "Everest (Sabah & Sarawak)",
}

var fschr60 = ellipsoid{
	a:           6378166.0,
	b:           298.3,
	ellipseName: "Fischer (Mercury Datum) 1960",
}

var fschr60m = ellipsoid{
	a:           6378155.0,
	b:           298.3,
	ellipseName: "Fischer 1960",
}

var fschr68 = ellipsoid{
	a:           6378150.0,
	b:           298.3,
	ellipseName: "Fischer 1968",
}

var helmert = ellipsoid{
	a:           6378200.0,
	b:           298.3,
	ellipseName: "Helmert 1906",
}

var hough = ellipsoid{
	a:           6378270.0,
	b:           297.0,
	ellipseName: "Hough",
}

var intl = ellipsoid{
	a:           6378388.0,
	b:           297.0,
	ellipseName: "International 1909 (Hayford)",
}

var kaula = ellipsoid{
	a:           6378163.0,
	b:           298.24,
	ellipseName: "Kaula 1961",
}

var lerch = ellipsoid{
	a:           6378139.0,
	b:           298.257,
	ellipseName: "Lerch 1979",
}

var mprts = ellipsoid{
	a:           6397300.0,
	b:           191.0,
	ellipseName: "Maupertius 1738",
}

var new_intl = ellipsoid{
	a:           6378157.5,
	b:           6356772.2,
	ellipseName: "New International 1967",
}

var plessis = ellipsoid{
	a:           6376523.0,
	b:           6355863.0,
	ellipseName: "Plessis 1817 (France)",
}

var krass = ellipsoid{
	a:           6378245.0,
	b:           298.3,
	ellipseName: "Krassovsky, 1942",
}

var SEasia = ellipsoid{
	a:           6378155.0,
	b:           6356773.3205,
	ellipseName: "Southeast Asia",
}

var walbeck = ellipsoid{
	a:           6376896.0,
	b:           6355834.8467,
	ellipseName: "Walbeck",
}

var WGS60 = ellipsoid{
	a:           6378165.0,
	b:           298.3,
	ellipseName: "WGS 60",
}

var WGS66 = ellipsoid{
	a:           6378145.0,
	b:           298.25,
	ellipseName: "WGS 66",
}

var WGS7 = ellipsoid{
	a:           6378135.0,
	b:           298.26,
	ellipseName: "WGS 72",
}

var WGS84 = ellipsoid{
	a:           6378137.0,
	b:           298.257223563,
	ellipseName: "WGS 84",
}

var sphere = ellipsoid{
	a:           6370997.0,
	b:           6370997.0,
	ellipseName: "Normal Sphere (r=6370997)",
}
