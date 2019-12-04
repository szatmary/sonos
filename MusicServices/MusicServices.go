package musicservices

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:MusicServices:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/MusicServices/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/MusicServices/Event`)
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
	XMLName                 xml.Name                     `xml:"s:Body"`
	GetSessionId            *GetSessionIdArgs            `xml:"u:GetSessionId,omitempty"`
	ListAvailableServices   *ListAvailableServicesArgs   `xml:"u:ListAvailableServices,omitempty"`
	UpdateAvailableServices *UpdateAvailableServicesArgs `xml:"u:UpdateAvailableServices,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                 xml.Name                         `xml:"Body"`
	GetSessionId            *GetSessionIdResponse            `xml:"GetSessionIdResponse,omitempty"`
	ListAvailableServices   *ListAvailableServicesResponse   `xml:"ListAvailableServicesResponse,omitempty"`
	UpdateAvailableServices *UpdateAvailableServicesResponse `xml:"UpdateAvailableServicesResponse,omitempty"`
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

type GetSessionIdArgs struct {
	Xmlns     string `xml:"xmlns:u,attr"`
	ServiceId uint32 `xml:"ServiceId"`
	Username  string `xml:"Username"`
}
type GetSessionIdResponse struct {
	SessionId string `xml:"SessionId"`
}

func (s *Service) GetSessionId(httpClient *http.Client, args *GetSessionIdArgs) (*GetSessionIdResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetSessionId`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetSessionId: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSessionId == nil {
		return nil, errors.New(`unexpected respose from service calling musicservices.GetSessionId()`)
	}

	return r.Body.GetSessionId, nil
}

type ListAvailableServicesArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type ListAvailableServicesResponse struct {
	AvailableServiceDescriptorList string `xml:"AvailableServiceDescriptorList"`
	AvailableServiceTypeList       string `xml:"AvailableServiceTypeList"`
	AvailableServiceListVersion    string `xml:"AvailableServiceListVersion"`
}

func (s *Service) ListAvailableServices(httpClient *http.Client, args *ListAvailableServicesArgs) (*ListAvailableServicesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ListAvailableServices`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ListAvailableServices: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ListAvailableServices == nil {
		return nil, errors.New(`unexpected respose from service calling musicservices.ListAvailableServices()`)
	}

	return r.Body.ListAvailableServices, nil
}

type UpdateAvailableServicesArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type UpdateAvailableServicesResponse struct {
}

func (s *Service) UpdateAvailableServices(httpClient *http.Client, args *UpdateAvailableServicesArgs) (*UpdateAvailableServicesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`UpdateAvailableServices`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{UpdateAvailableServices: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.UpdateAvailableServices == nil {
		return nil, errors.New(`unexpected respose from service calling musicservices.UpdateAvailableServices()`)
	}

	return r.Body.UpdateAvailableServices, nil
}
