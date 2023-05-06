package exosx

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
)

// Response : Typed representation of any XML API response
type Response struct {
	Version    string   `xml:"VERSION,attr"`
	Objects    []Object `xml:"OBJECT"`
	ObjectsMap map[string]*Object
}

// Object : Typed representation of any XML API object
type Object struct {
	Typ           string     `xml:"basetype,attr"`
	Name          string     `xml:"name,attr"`
	OID           int32      `xml:"oid,attr"`
	Format        string     `xml:"format,attr,omitempty"`
	Objects       []Object   `xml:"OBJECT"`
	Properties    []Property `xml:"PROPERTY"`
	ObjectsMap    map[string]*Object
	PropertiesMap map[string]*Property
}

// Property : Typed representation of any XML API property
type Property struct {
	Name        string `xml:"name,attr"`
	Typ         string `xml:"type,attr"`
	Size        int32  `xml:"size,attr"`
	Draw        bool   `xml:"draw,attr"`
	Sort        string `xml:"sort,attr"`
	DisplayName string `xml:"display-name,attr"`
	Data        string `xml:",chardata"`
}

// ResponseStatus : Final representation of the "status" object in every API response
type ResponseStatus struct {
	ResponseType        string
	ResponseTypeNumeric int
	Response            string
	ReturnCode          int
	Time                time.Time
}

// NewResponse : Unmarshals the raw data into a typed response object
// and generate a hash map from fields for optimization
func NewResponse(data []byte) (*Response, error) {
	res := &Response{}
	err := xml.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	res.ObjectsMap = make(map[string]*Object)
	for idx := range res.Objects {
		obj := &res.Objects[idx]
		fillObjectMap(obj)
		res.ObjectsMap[obj.Name] = obj
	}

	return res, nil
}

// NewErrorStatus : Creates an error status when response is not available
func NewErrorStatus(err string) *common.ResponseStatus {
	return &common.ResponseStatus{
		ResponseType: "Error",
		Response:     err,
		Time:         time.Now(),
	}
}

// GetStatus : Creates and returns the final ResponseStatus struct
// from the raw status object in response
func (res *Response) GetStatus() *common.ResponseStatus {
	statusObject := res.ObjectsMap["status"]
	responseTypeNumeric, _ := strconv.Atoi(statusObject.PropertiesMap["response-type-numeric"].Data)
	returnCode, _ := strconv.Atoi(statusObject.PropertiesMap["return-code"].Data)
	timestampNumeric, _ := strconv.Atoi(statusObject.PropertiesMap["time-stamp-numeric"].Data)

	return &common.ResponseStatus{
		ResponseType:        statusObject.PropertiesMap["response-type"].Data,
		ResponseTypeNumeric: responseTypeNumeric,
		Response:            statusObject.PropertiesMap["response"].Data,
		ReturnCode:          returnCode,
		Time:                time.Unix(int64(timestampNumeric), 0),
	}
}

func (object *Object) GetProperties(names ...string) ([]*Property, error) {
	var properties []*Property

	for _, name := range names {
		if property, ok := object.PropertiesMap[name]; ok {
			properties = append(properties, property)
		} else {
			return nil, fmt.Errorf("missing property %q", name)
		}
	}

	return properties, nil
}

func fillObjectMap(obj *Object) {
	obj.PropertiesMap = make(map[string]*Property)

	for idx2 := range obj.Properties {
		prop := &obj.Properties[idx2]
		prop.Data = strings.TrimSpace(prop.Data)
		obj.PropertiesMap[prop.Name] = prop
	}

	obj.ObjectsMap = make(map[string]*Object)
	for idx2 := range obj.Objects {
		subObject := &obj.Objects[idx2]
		fillObjectMap(subObject)
		obj.ObjectsMap[subObject.Name] = subObject
	}
}
