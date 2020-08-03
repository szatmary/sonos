package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// State Variables

type QPlayService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
	// State

}

func NewQPlayService(deviceUrl *url.URL) *QPlayService {
	c, _ := url.Parse("/QPlay/Control")
	e, _ := url.Parse("/QPlay/Event")
	return &QPlayService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *QPlayService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *QPlayService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type QPlayEnvelope struct {
	XMLName       xml.Name  `xml:"s:Envelope"`
	XMLNameSpace  string    `xml:"xmlns:s,attr"`
	EncodingStyle string    `xml:"s:encodingStyle,attr"`
	Body          QPlayBody `xml:"s:Body"`
}
type QPlayBody struct {
	XMLName   xml.Name            `xml:"s:Body"`
	QPlayAuth *QPlayQPlayAuthArgs `xml:"u:QPlayAuth,omitempty"`
}
type QPlayEnvelopeResponse struct {
	XMLName       xml.Name          `xml:"Envelope"`
	XMLNameSpace  string            `xml:"xmlns:s,attr"`
	EncodingStyle string            `xml:"encodingStyle,attr"`
	Body          QPlayBodyResponse `xml:"Body"`
}
type QPlayBodyResponse struct {
	XMLName   xml.Name                `xml:"Body"`
	QPlayAuth *QPlayQPlayAuthResponse `xml:"QPlayAuthResponse"`
}

func (s *QPlayService) _QPlayExec(soapAction string, httpClient *http.Client, envelope *QPlayEnvelope) (*QPlayEnvelopeResponse, error) {
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
	var envelopeResponse QPlayEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type QPlayQPlayAuthArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Seed         string `xml:"Seed"`
}
type QPlayQPlayAuthResponse struct {
	Code string `xml:"Code"`
	MID  string `xml:"MID"`
	DID  string `xml:"DID"`
}

func (s *QPlayService) QPlayAuth(httpClient *http.Client, args *QPlayQPlayAuthArgs) (*QPlayQPlayAuthResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:QPlay:1"
	r, err := s._QPlayExec("urn:schemas-upnp-org:service:QPlay:1#QPlayAuth", httpClient,
		&QPlayEnvelope{
			Body:          QPlayBody{QPlayAuth: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.QPlayAuth == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.QPlayAuth, nil
}
func (zp *QPlayService) ParseEvent([]byte) []interface{} {
	return []interface{}{}
}
