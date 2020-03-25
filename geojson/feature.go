package geojson

import jsoniter "github.com/json-iterator/go"

// the GeoJSON Feature Types like 'Feature' and 'FeatureCollection'
type FeatureTypes string

const (
	// the GeoJSON Feature Type
	FeatureType FeatureTypes = "Feature"
	// the GeoJSON FeatureCollection Type
	FeatureCollectionType FeatureTypes = "FeatureCollection"
)

// a Geometry with some Properties
type Feature struct {
	// the GeoJSON Feature Type
	Type FeatureTypes `json:"type"`
	// the Geometry
	Geometry *Geometry `json:"geometry"`
	// the Properties of the Geometry Object
	Properties *map[string]interface{} `json:"properties"`
}

// create a new Feature from a Geometry and some Properties you can pass nil when the Geometry has no Properties
// if so a empty Hash Map was created
func NewFeature(geometry *Geometry, properties *map[string]interface{}) *Feature {
	if properties == nil {
		emptyProperties := make(map[string]interface{})
		return &Feature{
			Type:       FeatureType,
			Geometry:   geometry,
			Properties: &emptyProperties,
		}
	}
	return &Feature{
		Type:       FeatureType,
		Geometry:   geometry,
		Properties: properties,
	}
}

// Deserialize a GeoJSON String into a Feature if a invalid GeoJSON String was given nil was returned
func DeserializeFeature(input string) *Feature {
	var res *Feature
	err := jsoniter.Unmarshal([]byte(input), &res)
	if err != nil {
		return nil
	}
	if res.Properties == nil {
		emptyProps := make(map[string]interface{})
		res.Properties = &emptyProps
	}
	return res
}

// Serialize the current Feature to a GeoJSON String
func (f *Feature) Serialize() string {
	stream, err := jsoniter.Marshal(f)
	if err != nil {
		return ""
	}
	return string(stream)
}
