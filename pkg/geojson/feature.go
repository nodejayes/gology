package geojson

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
