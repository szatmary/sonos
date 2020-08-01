package sonos

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

type ConnectionManagerService struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewConnectionManagerService(deviceUrl *url.URL) *ConnectionManagerService {
	c, _ := url.Parse("/MediaServer/ConnectionManager/Control")
	e, _ := url.Parse("/MediaServer/ConnectionManager/Event")
	return &ConnectionManagerService{
		ControlEndpoint: deviceUrl.ResolveReference(c),
		EventEndpoint:   deviceUrl.ResolveReference(e),
	}
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
	req, err := http.NewRequest("POST", s.ControlEndpoint.String(), bytes.NewBuffer(postBody))
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
func (s *ConnectionManagerService) ConnectionManagerSubscribe(callback url.URL) error {
	var req string
	req += fmt.Sprintf("SUBSCRIBE %s HTTP/1.0\r\n", s.EventEndpoint.String())
	req += fmt.Sprintf("HOST: %s\r\n", s.EventEndpoint.Host)
	req += fmt.Sprintf("USER-AGENT: Unknown UPnP/1.0 Gonos/1.0\r\n")
	req += fmt.Sprintf("CALLBACK: <%s>\r\n", callback.String())
	req += fmt.Sprintf("NT: upnp:event\r\n")
	req += fmt.Sprintf("TIMEOUT: Second-300\r\n")
	conn, err := net.Dial("tcp", s.EventEndpoint.Host)
	if err != nil {
		return err
	}
	fmt.Fprintf(conn, req+"\r\n")
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if 200 != res.StatusCode {
		fmt.Printf("%v\n", res)
		return errors.New(string(body))
	}
	return nil
}

// Events
type ConnectionManagerSourceProtocolInfo string
type ConnectionManagerSinkProtocolInfo string
type ConnectionManagerCurrentConnectionIDs string
type ConnectionManagerUpnpEvent struct {
	XMLName      xml.Name                    `xml:"propertyset"`
	XMLNameSpace string                      `xml:"xmlns:e,attr"`
	Properties   []ConnectionManagerProperty `xml:"property"`
}
type ConnectionManagerProperty struct {
	XMLName              xml.Name                               `xml:"property"`
	SourceProtocolInfo   *ConnectionManagerSourceProtocolInfo   `xml:"SourceProtocolInfo"`
	SinkProtocolInfo     *ConnectionManagerSinkProtocolInfo     `xml:"SinkProtocolInfo"`
	CurrentConnectionIDs *ConnectionManagerCurrentConnectionIDs `xml:"CurrentConnectionIDs"`
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
			zp.EventCallback(*prop.SourceProtocolInfo)
		case prop.SinkProtocolInfo != nil:
			zp.EventCallback(*prop.SinkProtocolInfo)
		case prop.CurrentConnectionIDs != nil:
			zp.EventCallback(*prop.CurrentConnectionIDs)
		}
	}
}
