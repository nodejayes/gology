package geojson

import (
	"fmt"
	"strconv"
	"strings"
)

// the Properties of the Coordinate Reference System
//
// only EPSG Codes are supported like 'EPSG:4326'
type IReferenceSystemProperties interface {
	INameReadProperty
}

type referenceSystemProperties struct {
	Name string `json:"name"`
}

func newReferenceSystemProperties(srId int) *referenceSystemProperties {
	if srId < 1 {
		srId = 0
	}
	return &referenceSystemProperties{
		Name: fmt.Sprintf("EPSG:%v", srId),
	}
}

// EPSG Code definition
func (rp *referenceSystemProperties) GetName() string {
	return rp.Name
}

// the Coordinate Reference System Information
type IReferenceSystem interface {
	// the Name of the Reference System Type
	//
	// only 'name' are supported
	ITypeReadProperty
	// the Coordinate Reference System Properties
	ISrIdReadProperty
}

type referenceSystem struct {
	Type       string                     `json:"type"`
	Properties *referenceSystemProperties `json:"properties"`
}

func newReferenceSystem(srId int) *referenceSystem {
	return &referenceSystem{
		Type:       "name",
		Properties: newReferenceSystemProperties(srId),
	}
}

// create a new Reference System from EPSG Code
func NewReferenceSystem(srId int) IReferenceSystem {
	return newReferenceSystem(srId)
}

func (rf *referenceSystem) GetType() string {
	return rf.Type
}

func (rf *referenceSystem) GetSrId() int {
	tmp := strings.Split(rf.Properties.GetName(), ":")
	if len(tmp) < 2 {
		return 0
	}
	res, err := strconv.ParseInt(tmp[1], 10, 32)
	if err != nil {
		return 0
	}
	return int(res)
}
