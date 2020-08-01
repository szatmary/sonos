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

type MusicServicesService struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewMusicServicesService(deviceUrl *url.URL) *MusicServicesService {
	c, _ := url.Parse("/MusicServices/Control")
	e, _ := url.Parse("/MusicServices/Event")
	return &MusicServicesService{
		ControlEndpoint: deviceUrl.ResolveReference(c),
		EventEndpoint:   deviceUrl.ResolveReference(e),
	}
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
func (s *MusicServicesService) MusicServicesSubscribe(callback url.URL) error {
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
type MusicServicesServiceListVersion string
type MusicServicesUpnpEvent struct {
	XMLName      xml.Name                `xml:"propertyset"`
	XMLNameSpace string                  `xml:"xmlns:e,attr"`
	Properties   []MusicServicesProperty `xml:"property"`
}
type MusicServicesProperty struct {
	XMLName            xml.Name                         `xml:"property"`
	ServiceListVersion *MusicServicesServiceListVersion `xml:"ServiceListVersion"`
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
			zp.EventCallback(*prop.ServiceListVersion)
		}
	}
}
