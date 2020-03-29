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

// a Geometry with some Properties
type IFeature interface {
	GetType() FeatureTypes
	GetGeometry() IGeometry
	GetProperties() *map[string]interface{}
	Serialize() string
}

type feature struct {
	Type       FeatureTypes            `json:"type"`
	Geometry   *geometry               `json:"geometry"`
	Properties *map[string]interface{} `json:"properties"`
}

// create a new Feature from a Geometry and some Properties you can pass nil when the Geometry has no Properties
// if so a empty Hash Map was created
func NewFeature(geom IGeometry, properties *map[string]interface{}) IFeature {
	g := &geometry{
		Type:        geom.GetType(),
		Coordinates: geom.GetCoordinates(),
		CRS: &referenceSystem{
			Type:       geom.GetCRS().GetType(),
			Properties: newReferenceSystemProperties(geom.GetCRS().GetSrId()),
		},
	}
	if properties == nil {
		emptyProperties := make(map[string]interface{})
		return &feature{
			Type:       FeatureType,
			Geometry:   g,
			Properties: &emptyProperties,
		}
	}
	return &feature{
		Type:       FeatureType,
		Geometry:   g,
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
