package geojson

import jsoniter "github.com/json-iterator/go"

// a Collection of Features
type FeatureCollection struct {
	// the GeoJSON FeatureCollection Type
	Type FeatureTypes `json:"type"`
	// the List of Features
	Features []*Feature `json:"features"`
	// a optional Coordinate Reference System for all Features
	CRS *ReferenceSystem `json:"crs,omitempty"`
}

// create a new Feature Collection from a List of Features
// the EPSG Code was taken from the First Feature that have one otherwise 0 was set
//
// It is recommended to use only features in the same coordinate system!
func NewFeatureCollection(features []*Feature) *FeatureCollection {
	srId := 0
	for _, feat := range features {
		srId = feat.Geometry.GetSrId()
		if srId > 0 {
			break
		}
	}
	return &FeatureCollection{
		Type:     FeatureCollectionType,
		Features: features,
		CRS:      NewReferenceSystem(srId),
	}
}

// returns the EPSG Code of the Coordinate Reference System
func (fc *FeatureCollection) GetSrId() int {
	return fc.CRS.GetSrId()
}

// serialize the current Feature Collection into a valid GeoJSON String
func (fc *FeatureCollection) Serialize() string {
	stream, err := jsoniter.Marshal(fc)
	if err != nil {
		return ""
	}
	return string(stream)
}

// Deserialize a GeoJSON String into a Feature Collection if the String is invalid nil was returned
func DeserializeFeatureCollection(input string) *FeatureCollection {
	var res *FeatureCollection
	err := jsoniter.Unmarshal([]byte(input), &res)
	if err != nil {
		return nil
	}
	return res
}
