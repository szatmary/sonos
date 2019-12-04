package groupmanagement

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:GroupManagement:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/GroupManagement/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/GroupManagement/Event`)
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
	XMLName                    xml.Name                        `xml:"s:Body"`
	AddMember                  *AddMemberArgs                  `xml:"u:AddMember,omitempty"`
	RemoveMember               *RemoveMemberArgs               `xml:"u:RemoveMember,omitempty"`
	ReportTrackBufferingResult *ReportTrackBufferingResultArgs `xml:"u:ReportTrackBufferingResult,omitempty"`
	SetSourceAreaIds           *SetSourceAreaIdsArgs           `xml:"u:SetSourceAreaIds,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                    xml.Name                            `xml:"Body"`
	AddMember                  *AddMemberResponse                  `xml:"AddMemberResponse,omitempty"`
	RemoveMember               *RemoveMemberResponse               `xml:"RemoveMemberResponse,omitempty"`
	ReportTrackBufferingResult *ReportTrackBufferingResultResponse `xml:"ReportTrackBufferingResultResponse,omitempty"`
	SetSourceAreaIds           *SetSourceAreaIdsResponse           `xml:"SetSourceAreaIdsResponse,omitempty"`
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

type AddMemberArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	MemberID string `xml:"MemberID"`
	BootSeq  uint32 `xml:"BootSeq"`
}
type AddMemberResponse struct {
	CurrentTransportSettings string `xml:"CurrentTransportSettings"`
	CurrentURI               string `xml:"CurrentURI"`
	GroupUUIDJoined          string `xml:"GroupUUIDJoined"`
	ResetVolumeAfter         bool   `xml:"ResetVolumeAfter"`
	VolumeAVTransportURI     string `xml:"VolumeAVTransportURI"`
}

func (s *Service) AddMember(httpClient *http.Client, args *AddMemberArgs) (*AddMemberResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddMember`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddMember: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddMember == nil {
		return nil, errors.New(`unexpected respose from service calling groupmanagement.AddMember()`)
	}

	return r.Body.AddMember, nil
}

type RemoveMemberArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	MemberID string `xml:"MemberID"`
}
type RemoveMemberResponse struct {
}

func (s *Service) RemoveMember(httpClient *http.Client, args *RemoveMemberArgs) (*RemoveMemberResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveMember`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveMember: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveMember == nil {
		return nil, errors.New(`unexpected respose from service calling groupmanagement.RemoveMember()`)
	}

	return r.Body.RemoveMember, nil
}

type ReportTrackBufferingResultArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	MemberID   string `xml:"MemberID"`
	ResultCode int32  `xml:"ResultCode"`
}
type ReportTrackBufferingResultResponse struct {
}

func (s *Service) ReportTrackBufferingResult(httpClient *http.Client, args *ReportTrackBufferingResultArgs) (*ReportTrackBufferingResultResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReportTrackBufferingResult`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReportTrackBufferingResult: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReportTrackBufferingResult == nil {
		return nil, errors.New(`unexpected respose from service calling groupmanagement.ReportTrackBufferingResult()`)
	}

	return r.Body.ReportTrackBufferingResult, nil
}

type SetSourceAreaIdsArgs struct {
	Xmlns                string `xml:"xmlns:u,attr"`
	DesiredSourceAreaIds string `xml:"DesiredSourceAreaIds"`
}
type SetSourceAreaIdsResponse struct {
}

func (s *Service) SetSourceAreaIds(httpClient *http.Client, args *SetSourceAreaIdsArgs) (*SetSourceAreaIdsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetSourceAreaIds`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetSourceAreaIds: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetSourceAreaIds == nil {
		return nil, errors.New(`unexpected respose from service calling groupmanagement.SetSourceAreaIds()`)
	}

	return r.Body.SetSourceAreaIds, nil
}
