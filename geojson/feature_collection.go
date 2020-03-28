package geojson

import (
	"encoding/json"
)

type IFeatureCollection interface {
	GetType() FeatureTypes
	GetFeatures() []IFeature
	GetCRS() IReferenceSystem
	GetSrId() int
	Serialize() string
}

// a Collection of Features
type featureCollection struct {
	// the GeoJSON FeatureCollection Type
	Type FeatureTypes `json:"type"`
	// the List of Features
	Features []IFeature `json:"features"`
	// a optional Coordinate Reference System for all Features
	CRS IReferenceSystem `json:"crs,omitempty"`
}

// get the GeoJSON Type of the FeatureCollection
func (fc *featureCollection) GetType() FeatureTypes {
	return fc.Type
}

// get the Feature List of the Feature Collection
func (fc *featureCollection) GetFeatures() []IFeature {
	return fc.Features
}

// get the ReferenceSystem of the Feature Collection
func (fc *featureCollection) GetCRS() IReferenceSystem {
	return fc.CRS
}

// create a new Feature Collection from a List of Features
// the EPSG Code was taken from the First Feature that have one otherwise 0 was set
//
// It is recommended to use only features in the same coordinate system!
func NewFeatureCollection(features []IFeature) IFeatureCollection {
	srId := 0
	for _, feat := range features {
		srId = feat.GetGeometry().GetCRS().GetSrId()
		if srId > 0 {
			break
		}
	}
	return &featureCollection{
		Type:     FeatureCollectionType,
		Features: features,
		CRS:      NewReferenceSystem(srId),
	}
}

// returns the EPSG Code of the Coordinate Reference System
func (fc *featureCollection) GetSrId() int {
	return fc.CRS.GetSrId()
}

// serialize the current Feature Collection into a valid GeoJSON String
func (fc *featureCollection) Serialize() string {
	stream, err := json.Marshal(fc)
	if err != nil {
		return ""
	}
	return string(stream)
}

// Deserialize a GeoJSON String into a Feature Collection if the String is invalid nil was returned
func DeserializeFeatureCollection(input string) IFeatureCollection {
	var res *featureCollection
	err := json.Unmarshal([]byte(input), &res)
	if err != nil {
		return nil
	}
	return res
}
