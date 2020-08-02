package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ConnectionManagerService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewConnectionManagerService(deviceUrl *url.URL) *ConnectionManagerService {
	c, _ := url.Parse("/MediaServer/ConnectionManager/Control")
	e, _ := url.Parse("/MediaServer/ConnectionManager/Event")
	return &ConnectionManagerService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *ConnectionManagerService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *ConnectionManagerService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type ConnectionManagerEnvelope struct {
	XMLName       xml.Name              `xml:"s:Envelope"`
	XMLNameSpace  string                `xml:"xmlns:s,attr"`
	EncodingStyle string                `xml:"s:encodingStyle,attr"`
	Body          ConnectionManagerBody `xml:"s:Body"`
}
type ConnectionManagerBody struct {
	XMLName                  xml.Name                                       `xml:"s:Body"`
	GetProtocolInfo          *ConnectionManagerGetProtocolInfoArgs          `xml:"u:GetProtocolInfo,omitempty"`
	GetCurrentConnectionIDs  *ConnectionManagerGetCurrentConnectionIDsArgs  `xml:"u:GetCurrentConnectionIDs,omitempty"`
	GetCurrentConnectionInfo *ConnectionManagerGetCurrentConnectionInfoArgs `xml:"u:GetCurrentConnectionInfo,omitempty"`
}
type ConnectionManagerEnvelopeResponse struct {
	XMLName       xml.Name                      `xml:"Envelope"`
	XMLNameSpace  string                        `xml:"xmlns:s,attr"`
	EncodingStyle string                        `xml:"encodingStyle,attr"`
	Body          ConnectionManagerBodyResponse `xml:"Body"`
}
type ConnectionManagerBodyResponse struct {
	XMLName                  xml.Name                                           `xml:"Body"`
	GetProtocolInfo          *ConnectionManagerGetProtocolInfoResponse          `xml:"GetProtocolInfoResponse"`
	GetCurrentConnectionIDs  *ConnectionManagerGetCurrentConnectionIDsResponse  `xml:"GetCurrentConnectionIDsResponse"`
	GetCurrentConnectionInfo *ConnectionManagerGetCurrentConnectionInfoResponse `xml:"GetCurrentConnectionInfoResponse"`
}

func (s *ConnectionManagerService) _ConnectionManagerExec(soapAction string, httpClient *http.Client, envelope *ConnectionManagerEnvelope) (*ConnectionManagerEnvelopeResponse, error) {
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
	var envelopeResponse ConnectionManagerEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type ConnectionManagerGetProtocolInfoArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ConnectionManagerGetProtocolInfoResponse struct {
	Source string `xml:"Source"`
	Sink   string `xml:"Sink"`
}

func (s *ConnectionManagerService) GetProtocolInfo(httpClient *http.Client, args *ConnectionManagerGetProtocolInfoArgs) (*ConnectionManagerGetProtocolInfoResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ConnectionManager:1"
	r, err := s._ConnectionManagerExec("urn:schemas-upnp-org:service:ConnectionManager:1#GetProtocolInfo", httpClient,
		&ConnectionManagerEnvelope{
			Body:          ConnectionManagerBody{GetProtocolInfo: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetProtocolInfo == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetProtocolInfo, nil
}

type ConnectionManagerGetCurrentConnectionIDsArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ConnectionManagerGetCurrentConnectionIDsResponse struct {
	ConnectionIDs string `xml:"ConnectionIDs"`
}

func (s *ConnectionManagerService) GetCurrentConnectionIDs(httpClient *http.Client, args *ConnectionManagerGetCurrentConnectionIDsArgs) (*ConnectionManagerGetCurrentConnectionIDsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ConnectionManager:1"
	r, err := s._ConnectionManagerExec("urn:schemas-upnp-org:service:ConnectionManager:1#GetCurrentConnectionIDs", httpClient,
		&ConnectionManagerEnvelope{
			Body:          ConnectionManagerBody{GetCurrentConnectionIDs: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCurrentConnectionIDs == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetCurrentConnectionIDs, nil
}

type ConnectionManagerGetCurrentConnectionInfoArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ConnectionID int32  `xml:"ConnectionID"`
}
type ConnectionManagerGetCurrentConnectionInfoResponse struct {
	RcsID                 int32  `xml:"RcsID"`
	AVTransportID         int32  `xml:"AVTransportID"`
	ProtocolInfo          string `xml:"ProtocolInfo"`
	PeerConnectionManager string `xml:"PeerConnectionManager"`
	PeerConnectionID      int32  `xml:"PeerConnectionID"`
	Direction             string `xml:"Direction"`
	Status                string `xml:"Status"`
}

func (s *ConnectionManagerService) GetCurrentConnectionInfo(httpClient *http.Client, args *ConnectionManagerGetCurrentConnectionInfoArgs) (*ConnectionManagerGetCurrentConnectionInfoResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ConnectionManager:1"
	r, err := s._ConnectionManagerExec("urn:schemas-upnp-org:service:ConnectionManager:1#GetCurrentConnectionInfo", httpClient,
		&ConnectionManagerEnvelope{
			Body:          ConnectionManagerBody{GetCurrentConnectionInfo: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCurrentConnectionInfo == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetCurrentConnectionInfo, nil
}

type ConnectionManagerUpnpEvent struct {
	XMLName      xml.Name                    `xml:"propertyset"`
	XMLNameSpace string                      `xml:"xmlns:e,attr"`
	Properties   []ConnectionManagerProperty `xml:"property"`
}
type ConnectionManagerProperty struct {
	XMLName              xml.Name `xml:"property"`
	SourceProtocolInfo   *string  `xml:"SourceProtocolInfo"`
	SinkProtocolInfo     *string  `xml:"SinkProtocolInfo"`
	CurrentConnectionIDs *string  `xml:"CurrentConnectionIDs"`
}

func ConnectionManagerDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt ConnectionManagerUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.SourceProtocolInfo != nil:
			dispatchConnectionManagerSourceProtocolInfo(*prop.SourceProtocolInfo) // string
		case prop.SinkProtocolInfo != nil:
			dispatchConnectionManagerSinkProtocolInfo(*prop.SinkProtocolInfo) // string
		case prop.CurrentConnectionIDs != nil:
			dispatchConnectionManagerCurrentConnectionIDs(*prop.CurrentConnectionIDs) // string
		}
	}
}
