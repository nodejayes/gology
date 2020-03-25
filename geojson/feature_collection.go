package geojson

import jsoniter "github.com/json-iterator/go"

type FeatureCollection struct {
	Type     FeatureTypes     `json:"type"`
	Features []*Feature       `json:"features"`
	CRS      *ReferenceSystem `json:"crs,omitempty"`
}

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

func (fc *FeatureCollection) GetSrId() int {
	return fc.CRS.GetSrId()
}

func (fc *FeatureCollection) Serialize() string {
	stream, err := jsoniter.Marshal(fc)
	if err != nil {
		return ""
	}
	return string(stream)
}

func DeserializeFeatureCollection(input string) *FeatureCollection {
	var res *FeatureCollection
	err := jsoniter.Unmarshal([]byte(input), &res)
	if err != nil {
		return nil
	}
	return res
}
