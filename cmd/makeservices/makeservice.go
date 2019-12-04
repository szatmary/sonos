package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
		return ""
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
	Xmlns          string          `xml:"type,attr"`
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

func MakeServiceApi(ServiceName, serviceControlEndpoint, serviceEventEndpoint string, scdp []byte) ([]byte, error) {
	var s Scpd
	err := xml.Unmarshal(scdp, &s)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "package %s\n\n", strings.ToLower(ServiceName))
	fmt.Fprint(buf, "import (\n\"net/url\"\n\"errors\"\n\"io/ioutil\"\n\"encoding/xml\"\n\"bytes\"\n\"net/http\"\n)\n")

	fmt.Fprint(buf, "const (\n")
	fmt.Fprintf(buf, "_ServiceURN = \"urn:schemas-upnp-org:service:%s:1\"\n", ServiceName)
	fmt.Fprintf(buf, "_EncodingSchema = \"http://schemas.xmlsoap.org/soap/encoding/\"\n")
	fmt.Fprintf(buf, "_EnvelopeSchema = \"http://schemas.xmlsoap.org/soap/envelope/\"\n")
	fmt.Fprint(buf, ")\n")

	// Service object
	fmt.Fprintf(buf, "type Service struct {\nControlEndpoint *url.URL\nEventEndpoint *url.URL\n}\n")
	fmt.Fprintf(buf, "func NewService(deviceUrl *url.URL) *Service {\n")
	fmt.Fprintf(buf, "c, err := url.Parse(`%s`)\nif nil != err { panic(err) }\n", serviceControlEndpoint)
	fmt.Fprintf(buf, "e, err := url.Parse(`%s`)\nif nil != err { panic(err) }\n", serviceEventEndpoint)
	fmt.Fprintf(buf, "return &Service{\nControlEndpoint: deviceUrl.ResolveReference(c),\nEventEndpoint: deviceUrl.ResolveReference(e),\n}\n}\n")

	// Martial structs
	fmt.Fprintf(buf, "type Envelope struct {\n")
	fmt.Fprintf(buf, "XMLName xml.Name `xml:\"s:Envelope\"`\n")
	fmt.Fprintf(buf, "Xmlns string `xml:\"xmlns:s,attr\"`\n")
	fmt.Fprintf(buf, "EncodingStyle string `xml:\"s:encodingStyle,attr\"`\n")
	fmt.Fprintf(buf, "Body Body `xml:\"s:Body\"`\n")
	fmt.Fprintf(buf, "}\n")

	fmt.Fprintf(buf, "type Body struct {\n")
	fmt.Fprint(buf, "XMLName xml.Name `xml:\"s:Body\"`\n")
	for _, action := range s.Actions {
		fmt.Fprintf(buf, "%s *%sArgs `xml:\"u:%s,omitempty\"`\n", action.Name, action.Name, action.Name)
	}
	fmt.Fprint(buf, "}\n")

	// Unmartial structs
	fmt.Fprintf(buf, "type EnvelopeResponse struct {\n")
	fmt.Fprintf(buf, "XMLName xml.Name `xml:\"Envelope\"`\n")
	fmt.Fprintf(buf, "Xmlns string `xml:\"xmlns:s,attr\"`\n")
	fmt.Fprintf(buf, "EncodingStyle string `xml:\"encodingStyle,attr\"`\n")
	fmt.Fprintf(buf, "Body BodyResponse `xml:\"Body\"`\n")
	fmt.Fprintf(buf, "}\n")

	fmt.Fprintf(buf, "type BodyResponse struct {\n")
	fmt.Fprint(buf, "XMLName xml.Name `xml:\"Body\"`\n")
	for _, action := range s.Actions {
		fmt.Fprintf(buf, "%s *%sResponse `xml:\"%sResponse,omitempty\"`\n", action.Name, action.Name, action.Name)
	}
	fmt.Fprintf(buf, "}\n")

	// exec function
	fmt.Fprintf(buf, "func (s *Service) exec(actionName string, httpClient *http.Client, envelope *Envelope) (*EnvelopeResponse, error) {\n")
	fmt.Fprintf(buf, "marshaled, err := xml.Marshal(envelope)\n")
	fmt.Fprintf(buf, "if err != nil { return nil, err\n}\n")
	fmt.Fprintf(buf, "postBody := []byte(`<?xml version=\"1.0\"?>`)\n")
	fmt.Fprintf(buf, "postBody = append(postBody, marshaled...)\n")
	fmt.Fprintf(buf, "req, err := http.NewRequest(`POST`, s.ControlEndpoint.String(), bytes.NewBuffer(postBody))\n")
	fmt.Fprintf(buf, "if err != nil { return nil, err\n}\n")
	fmt.Fprintf(buf, "req.Header.Set(`Content-Type`, `text/xml; charset=\"utf-8\"`)\n")
	fmt.Fprintf(buf, "req.Header.Set(`SOAPAction`, _ServiceURN+`#`+actionName)\n")
	fmt.Fprintf(buf, "res, err := httpClient.Do(req)\n")
	fmt.Fprintf(buf, "if err != nil { return nil, err\n}\n")
	fmt.Fprintf(buf, "defer res.Body.Close()\n")
	fmt.Fprintf(buf, "responseBody, err := ioutil.ReadAll(res.Body)\n")
	fmt.Fprintf(buf, "if err != nil { return nil, err\n}\n")
	fmt.Fprintf(buf, "var envelopeResponse EnvelopeResponse\n")
	fmt.Fprintf(buf, "err = xml.Unmarshal(responseBody,&envelopeResponse)\n")
	fmt.Fprintf(buf, "if err != nil { return nil, err\n}\n")
	fmt.Fprintf(buf, "return &envelopeResponse, nil\n}\n")

	for _, action := range s.Actions {
		var inArguments, outArguments []Argument
		for _, argument := range action.Arguments {
			switch argument.Direction {
			case "in":
				inArguments = append(inArguments, argument)
			case "out":
				outArguments = append(outArguments, argument)
			default:
				return []byte{}, errors.New("unexpected action direction")
			}
		}

		fmt.Fprintf(buf, "type %sArgs struct {\n", action.Name)
		fmt.Fprintf(buf, "Xmlns string `xml:\"xmlns:u,attr\"`\n")
		for _, argument := range inArguments {
			sv := s.GetStateVariable(argument.RelatedStateVariable)
			if sv == nil {
				return []byte{}, fmt.Errorf("unexpected state variable %s", argument.RelatedStateVariable)
			}
			if sv.AllowedValueRange != nil {
				fmt.Fprintf(buf, "// Allowed Range: %s -> %s step: %s\n", sv.AllowedValueRange.Minimum, sv.AllowedValueRange.Maximum, sv.AllowedValueRange.Step)
			}
			for _, allowedValue := range sv.AllowedValues {
				fmt.Fprintf(buf, "// Allowed Value: %s\n", allowedValue)
			}
			fmt.Fprintf(buf, "%s %s `xml:\"%s\"`\n", argument.Name, sv.GoDataType(), argument.Name)
		}
		fmt.Fprintf(buf, "}\n")

		fmt.Fprintf(buf, "type %sResponse struct {\n", action.Name)
		for _, argument := range outArguments {
			sv := s.GetStateVariable(argument.RelatedStateVariable)
			if sv == nil {
				return []byte{}, fmt.Errorf("unexpected state variable %s", argument.RelatedStateVariable)
			}
			fmt.Fprintf(buf, "%s %s\t`xml:\"%s\"`\n", argument.Name, sv.GoDataType(), argument.Name)
		}
		fmt.Fprintf(buf, "}\n")

		// TODO Validate, inputs
		fmt.Fprintf(buf, "func (s *Service) %s(httpClient *http.Client, args *%sArgs) (*%sResponse, error) {\n", action.Name, action.Name, action.Name)
		fmt.Fprintf(buf, "args.Xmlns = _ServiceURN\n")
		fmt.Fprintf(buf, "r, err := s.exec(`%s`, httpClient, \n&Envelope{\n", action.Name)
		fmt.Fprintf(buf, "EncodingStyle: _EncodingSchema,\n")
		fmt.Fprintf(buf, "Xmlns: _EnvelopeSchema,\n")
		fmt.Fprintf(buf, "Body: Body{%s: args},\n", action.Name)
		fmt.Fprintf(buf, "})\n")
		fmt.Fprintf(buf, "if err != nil { return nil, err }\n")
		fmt.Fprintf(buf, "if r.Body.%s == nil { return nil, errors.New(`unexpected respose from service calling %s.%s()`) }\n",
			action.Name, strings.ToLower(ServiceName), action.Name)
		fmt.Fprintf(buf, "\nreturn r.Body.%s, nil }\n", action.Name)
	}
	return buf.Bytes(), nil
}

func main() {
	// "http://192.168.131.242:1400/xml/RenderingControl1.xml"
	// "RenderingControl"
	serviceName := os.Args[1]
	serviceEndpoint := os.Args[2]
	controlEndpoint := os.Args[3]
	serviceXml := os.Args[4]
	body, err := ioutil.ReadFile(serviceXml)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	dotgo, err := MakeServiceApi(serviceName, serviceEndpoint, controlEndpoint, body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("%s\n", string(dotgo))
}
