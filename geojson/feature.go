package geojson

import (
	"encoding/json"
)

// the GeoJSON Feature Types like 'Feature' and 'FeatureCollection'
type FeatureTypes string

const (
	// the GeoJSON Feature Type
	FeatureType FeatureTypes = "Feature"
	// the GeoJSON FeatureCollection Type
	FeatureCollectionType FeatureTypes = "FeatureCollection"
)

type IFeature interface {
	GetType() FeatureTypes
	GetGeometry() IGeometry
	GetProperties() *map[string]interface{}
	Serialize() string
}

// a Geometry with some Properties
type feature struct {
	// the GeoJSON Feature Type
	Type FeatureTypes `json:"type"`
	// the Geometry
	Geometry IGeometry `json:"geometry"`
	// the Properties of the Geometry Object
	Properties *map[string]interface{} `json:"properties"`
}

// create a new Feature from a Geometry and some Properties you can pass nil when the Geometry has no Properties
// if so a empty Hash Map was created
func NewFeature(geometry IGeometry, properties *map[string]interface{}) IFeature {
	if properties == nil {
		emptyProperties := make(map[string]interface{})
		return &feature{
			Type:       FeatureType,
			Geometry:   geometry,
			Properties: &emptyProperties,
		}
	}
	return &feature{
		Type:       FeatureType,
		Geometry:   geometry,
		Properties: properties,
	}
}

// Deserialize a GeoJSON String into a Feature if a invalid GeoJSON String was given nil was returned
func DeserializeFeature(input string) IFeature {
	var res *feature
	err := json.Unmarshal([]byte(input), &res)
	if err != nil {
		return nil
	}
	if res.Properties == nil {
		emptyProps := make(map[string]interface{})
		res.Properties = &emptyProps
	}
	return res
}

// get the GeoJSON Type of the Feature
func (f *feature) GetType() FeatureTypes {
	return f.Type
}

// get the Geometry of the Feature
func (f *feature) GetGeometry() IGeometry {
	return f.Geometry
}

// get the Properties of the Feature
func (f *feature) GetProperties() *map[string]interface{} {
	return f.Properties
}

// Serialize the current Feature to a GeoJSON String
func (f *feature) Serialize() string {
	stream, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(stream)
}
