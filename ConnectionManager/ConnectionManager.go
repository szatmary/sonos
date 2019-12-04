package connectionmanager

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:ConnectionManager:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/MediaServer/ConnectionManager/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/MediaServer/ConnectionManager/Event`)
	if nil != err {
		panic(err)
	}
	return &Service{
		ControlEndpoint: deviceUrl.ResolveReference(c),
		EventEndpoint:   deviceUrl.ResolveReference(e),
	}
}

type Envelope struct {
	XMLName       xml.Name `xml:"s:Envelope"`
	Xmlns         string   `xml:"xmlns:s,attr"`
	EncodingStyle string   `xml:"s:encodingStyle,attr"`
	Body          Body     `xml:"s:Body"`
}
type Body struct {
	XMLName                  xml.Name                      `xml:"s:Body"`
	GetProtocolInfo          *GetProtocolInfoArgs          `xml:"u:GetProtocolInfo,omitempty"`
	GetCurrentConnectionIDs  *GetCurrentConnectionIDsArgs  `xml:"u:GetCurrentConnectionIDs,omitempty"`
	GetCurrentConnectionInfo *GetCurrentConnectionInfoArgs `xml:"u:GetCurrentConnectionInfo,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                  xml.Name                          `xml:"Body"`
	GetProtocolInfo          *GetProtocolInfoResponse          `xml:"GetProtocolInfoResponse,omitempty"`
	GetCurrentConnectionIDs  *GetCurrentConnectionIDsResponse  `xml:"GetCurrentConnectionIDsResponse,omitempty"`
	GetCurrentConnectionInfo *GetCurrentConnectionInfoResponse `xml:"GetCurrentConnectionInfoResponse,omitempty"`
}

func (s *Service) exec(actionName string, httpClient *http.Client, envelope *Envelope) (*EnvelopeResponse, error) {
	marshaled, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	postBody := []byte(`<?xml version="1.0"?>`)
	postBody = append(postBody, marshaled...)
	req, err := http.NewRequest(`POST`, s.ControlEndpoint.String(), bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set(`Content-Type`, `text/xml; charset="utf-8"`)
	req.Header.Set(`SOAPAction`, _ServiceURN+`#`+actionName)
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var envelopeResponse EnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type GetProtocolInfoArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetProtocolInfoResponse struct {
	Source string `xml:"Source"`
	Sink   string `xml:"Sink"`
}

func (s *Service) GetProtocolInfo(httpClient *http.Client, args *GetProtocolInfoArgs) (*GetProtocolInfoResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetProtocolInfo`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetProtocolInfo: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetProtocolInfo == nil {
		return nil, errors.New(`unexpected respose from service calling connectionmanager.GetProtocolInfo()`)
	}

	return r.Body.GetProtocolInfo, nil
}

type GetCurrentConnectionIDsArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetCurrentConnectionIDsResponse struct {
	ConnectionIDs string `xml:"ConnectionIDs"`
}

func (s *Service) GetCurrentConnectionIDs(httpClient *http.Client, args *GetCurrentConnectionIDsArgs) (*GetCurrentConnectionIDsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetCurrentConnectionIDs`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetCurrentConnectionIDs: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCurrentConnectionIDs == nil {
		return nil, errors.New(`unexpected respose from service calling connectionmanager.GetCurrentConnectionIDs()`)
	}

	return r.Body.GetCurrentConnectionIDs, nil
}

type GetCurrentConnectionInfoArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	ConnectionID int32  `xml:"ConnectionID"`
}
type GetCurrentConnectionInfoResponse struct {
	RcsID                 int32  `xml:"RcsID"`
	AVTransportID         int32  `xml:"AVTransportID"`
	ProtocolInfo          string `xml:"ProtocolInfo"`
	PeerConnectionManager string `xml:"PeerConnectionManager"`
	PeerConnectionID      int32  `xml:"PeerConnectionID"`
	Direction             string `xml:"Direction"`
	Status                string `xml:"Status"`
}

func (s *Service) GetCurrentConnectionInfo(httpClient *http.Client, args *GetCurrentConnectionInfoArgs) (*GetCurrentConnectionInfoResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetCurrentConnectionInfo`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetCurrentConnectionInfo: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCurrentConnectionInfo == nil {
		return nil, errors.New(`unexpected respose from service calling connectionmanager.GetCurrentConnectionInfo()`)
	}

	return r.Body.GetCurrentConnectionInfo, nil
}
