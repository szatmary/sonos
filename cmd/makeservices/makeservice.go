package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type AllowedValueRange struct {
	XMLName xml.Name `xml:"allowedValueRange"`
	Minimum string   `xml:"minimum"`
	Maximum string   `xml:"maximum"`
	Step    string   `xml:"step"`
}

type StateVariable struct {
	XMLName           xml.Name           `xml:"stateVariable"`
	SendEvents        string             `xml:"sendEvents,attr"`
	Multicast         string             `xml:"multicast,attr"`
	Name              string             `xml:"name"`
	DataType          string             `xml:"dataType"`
	DefaultValue      string             `xml:"defaultValue"`
	AllowedValueRange *AllowedValueRange `xml:"allowedValueRange"`
	AllowedValues     []string           `xml:"allowedValueList>allowedValue"`
}

func (s *StateVariable) GoDataType() string {
	switch s.DataType {
	case "ui1":
		return "uint8"
	case "ui2":
		return "uint16"
	case "ui4":
		return "uint32"
	case "i1":
		return "int8"
	case "i2":
		return "int16"
	case "i4":
		return "int32"
	case "int":
		return "int64"
	case "r4":
		return "flaot32"
	case "number", "r8":
		return "float"
	case "float", "float64":
		return "float"
	// case "fixed.14.4": // TODO fixed
	case "char":
		return "rune"
	case "string":
		return "string"
		// TODO data/time
	case "date", "dateTime", " dateTime.tz", "time", "time.tz":
		return "string"
	case "boolean":
		return "bool"
		// TODO
	// case "bin.base64", "bin.hex":
	// 	return "string"
	case "uri":
		return "*url.URL"
	case "uuid":
		return "string"
	default:
		panic(s.DataType)
	}
}

type Argument struct {
	XMLName              xml.Name `xml:"argument"`
	Name                 string   `xml:"name"`
	Direction            string   `xml:"direction"`
	RelatedStateVariable string   `xml:"relatedStateVariable"`
}

type Action struct {
	XMLName   xml.Name   `xml:"action"`
	Name      string     `xml:"name"`
	Arguments []Argument `xml:"argumentList>argument"`
}

type SpecVersion struct {
	XMLName xml.Name `xml:"specVersion"`
	Major   int      `xml:"major"`
	Minor   int      `xml:"minor"`
}

type Scpd struct {
	XMLName        xml.Name        `xml:"scpd"`
	XMLNameSpace   string          `xml:"type,attr"`
	SpecVersion    SpecVersion     `xml:"specVersion"`
	StateVariables []StateVariable `xml:"serviceStateTable>stateVariable"`
	Actions        []Action        `xml:"actionList>action"`
}

func (s *Scpd) GetStateVariable(name string) *StateVariable {
	for _, sv := range s.StateVariables {
		if sv.Name == name {
			return &sv
		}
	}
	return nil
}

func MakeServiceApi(ServiceName, servicecontrolEndpoint, serviceeventEndpoint, xmlFile string) []byte {
	var s Scpd
	scdp, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(scdp, &s)
	if err != nil {
		panic(err)
	}

	eventCount := 0
	for _, sv := range s.StateVariables {
		if sv.SendEvents != "yes" {
			continue
		}
		eventCount++
	}

	// State

	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, `package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// State Variables
`)

	for _, sv := range s.StateVariables {
		if sv.SendEvents != "yes" {
			continue
		}
		fmt.Fprintf(buf, "type %s_%s %s\n", ServiceName, sv.Name, sv.GoDataType())
	}

	fmt.Fprintf(buf, `

type %sService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
	// State
`, ServiceName)

	for _, sv := range s.StateVariables {
		if sv.SendEvents != "yes" {
			continue
		}
		fmt.Fprintf(buf, "%s *%s_%s\n", sv.Name, ServiceName, sv.Name)
	}

	fmt.Fprintf(buf, `
}
func New%sService(deviceUrl *url.URL) *%sService {
	c, _ := url.Parse("%s")
	e, _ := url.Parse("%s")
	return &%sService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *%sService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *%sService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

`, ServiceName, ServiceName, servicecontrolEndpoint, serviceeventEndpoint, ServiceName, ServiceName, ServiceName)

	// Martial structs
	fmt.Fprintf(buf, "type %sEnvelope struct {\n", ServiceName)
	fmt.Fprint(buf, "XMLName xml.Name `xml:\"s:Envelope\"`\n")
	fmt.Fprint(buf, "XMLNameSpace string `xml:\"xmlns:s,attr\"`\n")
	fmt.Fprint(buf, "EncodingStyle string `xml:\"s:encodingStyle,attr\"`\n")
	fmt.Fprintf(buf, "Body %sBody `xml:\"s:Body\"`\n", ServiceName)
	fmt.Fprint(buf, "}\n")

	fmt.Fprintf(buf, "type %sBody struct {\n", ServiceName)
	fmt.Fprint(buf, "XMLName xml.Name `xml:\"s:Body\"`\n")
	for _, action := range s.Actions {
		fmt.Fprintf(buf, "%s *%s%sArgs `xml:\"u:%s,omitempty\"`\n", action.Name, ServiceName, action.Name, action.Name)
	}
	fmt.Fprint(buf, "}\n")

	// Unmartial structs
	fmt.Fprintf(buf, "type %sEnvelopeResponse struct {\n", ServiceName)
	fmt.Fprint(buf, "XMLName xml.Name `xml:\"Envelope\"`\n")
	fmt.Fprint(buf, "XMLNameSpace string `xml:\"xmlns:s,attr\"`\n")
	fmt.Fprint(buf, "EncodingStyle string `xml:\"encodingStyle,attr\"`\n")
	fmt.Fprintf(buf, "Body %sBodyResponse `xml:\"Body\"`\n", ServiceName)
	fmt.Fprint(buf, "}\n")

	fmt.Fprintf(buf, "type %sBodyResponse struct {\n", ServiceName)
	fmt.Fprint(buf, "XMLName xml.Name `xml:\"Body\"`\n")
	for _, action := range s.Actions {
		fmt.Fprintf(buf, "%s *%s%sResponse `xml:\"%sResponse\"`\n", action.Name, ServiceName, action.Name, action.Name)
	}
	fmt.Fprint(buf, "}\n")

	// exec function
	fmt.Fprintf(buf, `func (s *%sService) _%sExec(soapAction string, httpClient *http.Client, envelope *%sEnvelope) (*%sEnvelopeResponse, error) {
	postBody, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("soapAction %%s: postBody %%v\n", soapAction, string(postBody))
	req, err := http.NewRequest("POST", s.controlEndpoint.String(), bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("SOAPAction", soapAction)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("responseBody %%v\n", string(responseBody))
	var envelopeResponse %sEnvelopeResponse
	err = xml.Unmarshal(responseBody,&envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}
`, ServiceName, ServiceName, ServiceName, ServiceName, ServiceName)

	for _, action := range s.Actions {
		fmt.Fprintf(buf, "type %s%sArgs struct {\n", ServiceName, action.Name)
		fmt.Fprint(buf, "XMLNameSpace string `xml:\"xmlns:u,attr\"`\n")
		for _, argument := range action.Arguments {
			if argument.Direction != "in" {
				continue
			}
			sv := s.GetStateVariable(argument.RelatedStateVariable)
			if sv == nil {
				panic("unexpected state variable " + argument.RelatedStateVariable)
			}
			if sv.AllowedValueRange != nil {
				fmt.Fprintf(buf, "// Allowed Range: %s -> %s step: %s\n", sv.AllowedValueRange.Minimum, sv.AllowedValueRange.Maximum, sv.AllowedValueRange.Step)
			}
			for _, allowedValue := range sv.AllowedValues {
				fmt.Fprintf(buf, "// Allowed Value: %s\n", allowedValue)
			}
			fmt.Fprintf(buf, "%s %s `xml:\"%s\"`\n", argument.Name, sv.GoDataType(), argument.Name)
		}
		fmt.Fprint(buf, "}\n")

		fmt.Fprintf(buf, "type %s%sResponse struct {\n", ServiceName, action.Name)
		for _, argument := range action.Arguments {
			if argument.Direction != "out" {
				continue
			}
			sv := s.GetStateVariable(argument.RelatedStateVariable)
			if sv == nil {
				panic("unexpected state variable " + argument.RelatedStateVariable)
			}
			fmt.Fprintf(buf, "%s %s\t`xml:\"%s\"`\n", argument.Name, sv.GoDataType(), argument.Name)
		}
		fmt.Fprint(buf, "}\n")

		// TODO Validate, inputs
		fmt.Fprintf(buf, `func (s *%sService) %s(httpClient *http.Client, args *%s%sArgs) (*%s%sResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:%s:1"
	r, err := s._%sExec("urn:schemas-upnp-org:service:%s:1#%s", httpClient,
		&%sEnvelope{
			Body: %sBody{%s: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
			})
	if err != nil {
		return nil, err
	}
	if r.Body.%s == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.%s, nil
}
`, ServiceName, action.Name, ServiceName, action.Name, ServiceName, action.Name,
			ServiceName, ServiceName, ServiceName, action.Name, ServiceName,
			ServiceName, action.Name, action.Name, action.Name)
	}

	if eventCount == 0 {
		fmt.Fprintf(buf, `func (zp *%sService) ParseEvent([]byte) []interface{} {
			return []interface{}{}
		}`, ServiceName)
		return buf.Bytes()
	}

	fmt.Fprintf(buf, "type %sUpnpEvent struct {\nXMLName xml.Name `xml:\"propertyset\"`\nXMLNameSpace string `xml:\"xmlns:e,attr\"`\nProperties []%sProperty `xml:\"property\"`\n}\n", ServiceName, ServiceName)
	fmt.Fprintf(buf, "type %sProperty struct {\nXMLName xml.Name `xml:\"property\"`\n", ServiceName)
	for _, sv := range s.StateVariables {
		if sv.SendEvents != "yes" {
			continue
		}
		// fmt.Fprintf(buf, "%s *%s%s `xml:\"%s\"`\n", sv.Name, ServiceName, sv.Name, sv.Name)
		fmt.Fprintf(buf, "%s *%s_%s `xml:\"%s\"`\n", sv.Name, ServiceName, sv.Name, sv.Name)
	}

	fmt.Fprint(buf, "}\n")
	fmt.Fprintf(buf, `func (zp *%sService) ParseEvent(body []byte) []interface{} {
	var evt %sUpnpEvent
	var events []interface{}
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return events
	}
	for _, prop := range evt.Properties {
	switch {
`, ServiceName, ServiceName)
	for _, sv := range s.StateVariables {
		if sv.SendEvents != "yes" {
			continue
		}
		// fmt.Fprintf(buf, "case prop.%s != nil:\n zp.EventCallback(*prop.%s)\n", sv.Name, sv.Name)
		fmt.Fprintf(buf, "case prop.%s != nil:\n", sv.Name)
		fmt.Fprintf(buf, "zp.%s = prop.%s\n", sv.Name, sv.Name)
		fmt.Fprintf(buf, "events = append(events, *prop.%s)\n", sv.Name)
	}
	fmt.Fprintf(buf, "}\n}\nreturn events\n}")
	return buf.Bytes()
}

func main() {
	serviceName := os.Args[1]
	serviceEndpoint := os.Args[2]
	controlEndpoint := os.Args[3]
	serviceXml := os.Args[4]
	dotgo := MakeServiceApi(serviceName, serviceEndpoint, controlEndpoint, serviceXml)
	fmt.Printf("%s\n", string(dotgo))
}
