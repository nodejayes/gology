package geojson

import jsoniter "github.com/json-iterator/go"

type FeatureTypes string

const (
	FeatureType           FeatureTypes = "Feature"
	FeatureCollectionType FeatureTypes = "FeatureCollection"
)

type Feature struct {
	Type       FeatureTypes            `json:"type"`
	Geometry   *Geometry               `json:"geometry"`
	Properties *map[string]interface{} `json:"properties"`
}

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

func (f *Feature) Serialize() string {
	stream, err := jsoniter.Marshal(f)
	if err != nil {
		return ""
	}
	return string(stream)
}
