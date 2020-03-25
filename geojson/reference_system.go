package geojson

import (
	"fmt"
	"strconv"
	"strings"
)

// the Properties of the Coordinate Reference System
//
// only EPSG Codes are supported like 'EPSG:4326'
type ReferenceSystemProperties struct {
	// EPSG Code definition
	Name string `json:"name"`
}

// the Coordinate Reference System Information
type ReferenceSystem struct {
	// the Name of the Reference System Type
	//
	// only 'name' are supported
	Type string `json:"type"`
	// the Coordinate Reference System Properties
	Properties *ReferenceSystemProperties `json:"properties"`
}

// create a new Reference System from EPSG Code
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
