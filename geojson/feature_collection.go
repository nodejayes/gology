package geojson

import (
	"encoding/json"
)

// a Collection of Features
type IFeatureCollection interface {
	GetType() FeatureTypes
	GetFeatures() []IFeature
	GetCRS() IReferenceSystem
	GetSrId() int
	Serialize() string
}

type featureCollection struct {
	Type     FeatureTypes     `json:"type"`
	Features []*feature       `json:"features"`
	CRS      *referenceSystem `json:"crs,omitempty"`
}

// get the GeoJSON Type of the FeatureCollection
func (fc *featureCollection) GetType() FeatureTypes {
	return fc.Type
}

// get the Feature List of the Feature Collection
func (fc *featureCollection) GetFeatures() []IFeature {
	var res []IFeature
	for _, f := range fc.Features {
		res = append(res, f)
	}
	return res
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
	var feats []*feature
	for _, f := range features {
		feats = append(feats, &feature{
			Type: f.GetType(),
			Geometry: &geometry{
				Type:        f.GetGeometry().GetType(),
				Coordinates: f.GetGeometry().GetCoordinates(),
				CRS: &referenceSystem{
					Type:       f.GetGeometry().GetCRS().GetType(),
					Properties: newReferenceSystemProperties(f.GetGeometry().GetCRS().GetSrId()),
				},
			},
			Properties: f.GetProperties(),
		})
	}
	return &featureCollection{
		Type:     FeatureCollectionType,
		Features: feats,
		CRS: &referenceSystem{
			Type:       "name",
			Properties: newReferenceSystemProperties(srId),
		},
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
func DeserializeFeatureCollection(input string) (IFeatureCollection, error) {
	var res *featureCollection
	err := json.Unmarshal([]byte(input), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeserializeFeatureCollectionList(input string) ([]IFeatureCollection, error) {
	var tmp []*featureCollection
	err := json.Unmarshal([]byte(input), &tmp)
	if err != nil {
		return nil, err
	}
	var res []IFeatureCollection
	for _, fc := range tmp {
		res = append(res, fc)
	}
	return res, nil
}
