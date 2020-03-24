package geojson

import (
	"fmt"
	"strconv"
	"strings"
)

type ReferenceSystemProperties struct {
	Name string `json:"name"`
}

type ReferenceSystem struct {
	Type       string                     `json:"type"`
	Properties *ReferenceSystemProperties `json:"properties"`
}

func NewReferenceSystem(srid int) *ReferenceSystem {
	return &ReferenceSystem{
		Type: "name",
		Properties: &ReferenceSystemProperties{
			Name: fmt.Sprintf("EPSG:%v", srid),
		},
	}
}

func (rf *ReferenceSystem) GetSrId() int {
	tmp := strings.Split(rf.Properties.Name, ":")
	if len(tmp) < 2 {
		return 0
	}
	res, err := strconv.ParseInt(tmp[1], 10, 32)
	if err != nil {
		return 0
	}
	return int(res)
}
