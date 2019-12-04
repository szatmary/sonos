package qplay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:QPlay:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/QPlay/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/QPlay/Event`)
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
	XMLName   xml.Name       `xml:"s:Body"`
	QPlayAuth *QPlayAuthArgs `xml:"u:QPlayAuth,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName   xml.Name           `xml:"Body"`
	QPlayAuth *QPlayAuthResponse `xml:"QPlayAuthResponse,omitempty"`
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

type QPlayAuthArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
	Seed  string `xml:"Seed"`
}
type QPlayAuthResponse struct {
	Code string `xml:"Code"`
	MID  string `xml:"MID"`
	DID  string `xml:"DID"`
}

func (s *Service) QPlayAuth(httpClient *http.Client, args *QPlayAuthArgs) (*QPlayAuthResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`QPlayAuth`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{QPlayAuth: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.QPlayAuth == nil {
		return nil, errors.New(`unexpected respose from service calling qplay.QPlayAuth()`)
	}

	return r.Body.QPlayAuth, nil
}
