package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type MusicServicesService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewMusicServicesService(deviceUrl *url.URL) *MusicServicesService {
	c, _ := url.Parse("/MusicServices/Control")
	e, _ := url.Parse("/MusicServices/Event")
	return &MusicServicesService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *MusicServicesService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *MusicServicesService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type MusicServicesEnvelope struct {
	XMLName       xml.Name          `xml:"s:Envelope"`
	XMLNameSpace  string            `xml:"xmlns:s,attr"`
	EncodingStyle string            `xml:"s:encodingStyle,attr"`
	Body          MusicServicesBody `xml:"s:Body"`
}
type MusicServicesBody struct {
	XMLName                 xml.Name                                  `xml:"s:Body"`
	GetSessionId            *MusicServicesGetSessionIdArgs            `xml:"u:GetSessionId,omitempty"`
	ListAvailableServices   *MusicServicesListAvailableServicesArgs   `xml:"u:ListAvailableServices,omitempty"`
	UpdateAvailableServices *MusicServicesUpdateAvailableServicesArgs `xml:"u:UpdateAvailableServices,omitempty"`
}
type MusicServicesEnvelopeResponse struct {
	XMLName       xml.Name                  `xml:"Envelope"`
	XMLNameSpace  string                    `xml:"xmlns:s,attr"`
	EncodingStyle string                    `xml:"encodingStyle,attr"`
	Body          MusicServicesBodyResponse `xml:"Body"`
}
type MusicServicesBodyResponse struct {
	XMLName                 xml.Name                                      `xml:"Body"`
	GetSessionId            *MusicServicesGetSessionIdResponse            `xml:"GetSessionIdResponse"`
	ListAvailableServices   *MusicServicesListAvailableServicesResponse   `xml:"ListAvailableServicesResponse"`
	UpdateAvailableServices *MusicServicesUpdateAvailableServicesResponse `xml:"UpdateAvailableServicesResponse"`
}

func (s *MusicServicesService) _MusicServicesExec(soapAction string, httpClient *http.Client, envelope *MusicServicesEnvelope) (*MusicServicesEnvelopeResponse, error) {
	postBody, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("soapAction %s: postBody %v\n", soapAction, string(postBody))
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
	// fmt.Printf("responseBody %v\n", string(responseBody))
	var envelopeResponse MusicServicesEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type MusicServicesGetSessionIdArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ServiceId    uint32 `xml:"ServiceId"`
	Username     string `xml:"Username"`
}
type MusicServicesGetSessionIdResponse struct {
	SessionId string `xml:"SessionId"`
}

func (s *MusicServicesService) GetSessionId(httpClient *http.Client, args *MusicServicesGetSessionIdArgs) (*MusicServicesGetSessionIdResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:MusicServices:1"
	r, err := s._MusicServicesExec("urn:schemas-upnp-org:service:MusicServices:1#GetSessionId", httpClient,
		&MusicServicesEnvelope{
			Body:          MusicServicesBody{GetSessionId: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSessionId == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetSessionId, nil
}

type MusicServicesListAvailableServicesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type MusicServicesListAvailableServicesResponse struct {
	AvailableServiceDescriptorList string `xml:"AvailableServiceDescriptorList"`
	AvailableServiceTypeList       string `xml:"AvailableServiceTypeList"`
	AvailableServiceListVersion    string `xml:"AvailableServiceListVersion"`
}

func (s *MusicServicesService) ListAvailableServices(httpClient *http.Client, args *MusicServicesListAvailableServicesArgs) (*MusicServicesListAvailableServicesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:MusicServices:1"
	r, err := s._MusicServicesExec("urn:schemas-upnp-org:service:MusicServices:1#ListAvailableServices", httpClient,
		&MusicServicesEnvelope{
			Body:          MusicServicesBody{ListAvailableServices: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ListAvailableServices == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ListAvailableServices, nil
}

type MusicServicesUpdateAvailableServicesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type MusicServicesUpdateAvailableServicesResponse struct {
}

func (s *MusicServicesService) UpdateAvailableServices(httpClient *http.Client, args *MusicServicesUpdateAvailableServicesArgs) (*MusicServicesUpdateAvailableServicesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:MusicServices:1"
	r, err := s._MusicServicesExec("urn:schemas-upnp-org:service:MusicServices:1#UpdateAvailableServices", httpClient,
		&MusicServicesEnvelope{
			Body:          MusicServicesBody{UpdateAvailableServices: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.UpdateAvailableServices == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.UpdateAvailableServices, nil
}

type MusicServicesUpnpEvent struct {
	XMLName      xml.Name                `xml:"propertyset"`
	XMLNameSpace string                  `xml:"xmlns:e,attr"`
	Properties   []MusicServicesProperty `xml:"property"`
}
type MusicServicesProperty struct {
	XMLName            xml.Name `xml:"property"`
	ServiceListVersion *string  `xml:"ServiceListVersion"`
}

func MusicServicesDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt MusicServicesUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.ServiceListVersion != nil:
			dispatchMusicServicesServiceListVersion(*prop.ServiceListVersion) // string
		}
	}
}
