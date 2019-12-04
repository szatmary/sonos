package zonegrouptopology

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:ZoneGroupTopology:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/ZoneGroupTopology/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/ZoneGroupTopology/Event`)
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
	XMLName                   xml.Name                       `xml:"s:Body"`
	CheckForUpdate            *CheckForUpdateArgs            `xml:"u:CheckForUpdate,omitempty"`
	BeginSoftwareUpdate       *BeginSoftwareUpdateArgs       `xml:"u:BeginSoftwareUpdate,omitempty"`
	ReportUnresponsiveDevice  *ReportUnresponsiveDeviceArgs  `xml:"u:ReportUnresponsiveDevice,omitempty"`
	ReportAlarmStartedRunning *ReportAlarmStartedRunningArgs `xml:"u:ReportAlarmStartedRunning,omitempty"`
	SubmitDiagnostics         *SubmitDiagnosticsArgs         `xml:"u:SubmitDiagnostics,omitempty"`
	RegisterMobileDevice      *RegisterMobileDeviceArgs      `xml:"u:RegisterMobileDevice,omitempty"`
	GetZoneGroupAttributes    *GetZoneGroupAttributesArgs    `xml:"u:GetZoneGroupAttributes,omitempty"`
	GetZoneGroupState         *GetZoneGroupStateArgs         `xml:"u:GetZoneGroupState,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                   xml.Name                           `xml:"Body"`
	CheckForUpdate            *CheckForUpdateResponse            `xml:"CheckForUpdateResponse,omitempty"`
	BeginSoftwareUpdate       *BeginSoftwareUpdateResponse       `xml:"BeginSoftwareUpdateResponse,omitempty"`
	ReportUnresponsiveDevice  *ReportUnresponsiveDeviceResponse  `xml:"ReportUnresponsiveDeviceResponse,omitempty"`
	ReportAlarmStartedRunning *ReportAlarmStartedRunningResponse `xml:"ReportAlarmStartedRunningResponse,omitempty"`
	SubmitDiagnostics         *SubmitDiagnosticsResponse         `xml:"SubmitDiagnosticsResponse,omitempty"`
	RegisterMobileDevice      *RegisterMobileDeviceResponse      `xml:"RegisterMobileDeviceResponse,omitempty"`
	GetZoneGroupAttributes    *GetZoneGroupAttributesResponse    `xml:"GetZoneGroupAttributesResponse,omitempty"`
	GetZoneGroupState         *GetZoneGroupStateResponse         `xml:"GetZoneGroupStateResponse,omitempty"`
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

type CheckForUpdateArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
	// Allowed Value: All
	// Allowed Value: Software
	UpdateType string `xml:"UpdateType"`
	CachedOnly bool   `xml:"CachedOnly"`
	Version    string `xml:"Version"`
}
type CheckForUpdateResponse struct {
	UpdateItem string `xml:"UpdateItem"`
}

func (s *Service) CheckForUpdate(httpClient *http.Client, args *CheckForUpdateArgs) (*CheckForUpdateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`CheckForUpdate`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{CheckForUpdate: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CheckForUpdate == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.CheckForUpdate()`)
	}

	return r.Body.CheckForUpdate, nil
}

type BeginSoftwareUpdateArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	UpdateURL    string `xml:"UpdateURL"`
	Flags        uint32 `xml:"Flags"`
	ExtraOptions string `xml:"ExtraOptions"`
}
type BeginSoftwareUpdateResponse struct {
}

func (s *Service) BeginSoftwareUpdate(httpClient *http.Client, args *BeginSoftwareUpdateArgs) (*BeginSoftwareUpdateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`BeginSoftwareUpdate`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{BeginSoftwareUpdate: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BeginSoftwareUpdate == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.BeginSoftwareUpdate()`)
	}

	return r.Body.BeginSoftwareUpdate, nil
}

type ReportUnresponsiveDeviceArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	DeviceUUID string `xml:"DeviceUUID"`
	// Allowed Value: Remove
	// Allowed Value: TopologyMonitorProbe
	// Allowed Value: VerifyThenRemoveSystemwide
	DesiredAction string `xml:"DesiredAction"`
}
type ReportUnresponsiveDeviceResponse struct {
}

func (s *Service) ReportUnresponsiveDevice(httpClient *http.Client, args *ReportUnresponsiveDeviceArgs) (*ReportUnresponsiveDeviceResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReportUnresponsiveDevice`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReportUnresponsiveDevice: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReportUnresponsiveDevice == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.ReportUnresponsiveDevice()`)
	}

	return r.Body.ReportUnresponsiveDevice, nil
}

type ReportAlarmStartedRunningArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type ReportAlarmStartedRunningResponse struct {
}

func (s *Service) ReportAlarmStartedRunning(httpClient *http.Client, args *ReportAlarmStartedRunningArgs) (*ReportAlarmStartedRunningResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReportAlarmStartedRunning`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReportAlarmStartedRunning: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReportAlarmStartedRunning == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.ReportAlarmStartedRunning()`)
	}

	return r.Body.ReportAlarmStartedRunning, nil
}

type SubmitDiagnosticsArgs struct {
	Xmlns              string `xml:"xmlns:u,attr"`
	IncludeControllers bool   `xml:"IncludeControllers"`
	Type               string `xml:"Type"`
}
type SubmitDiagnosticsResponse struct {
	DiagnosticID uint32 `xml:"DiagnosticID"`
}

func (s *Service) SubmitDiagnostics(httpClient *http.Client, args *SubmitDiagnosticsArgs) (*SubmitDiagnosticsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SubmitDiagnostics`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SubmitDiagnostics: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SubmitDiagnostics == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.SubmitDiagnostics()`)
	}

	return r.Body.SubmitDiagnostics, nil
}

type RegisterMobileDeviceArgs struct {
	Xmlns            string `xml:"xmlns:u,attr"`
	MobileDeviceName string `xml:"MobileDeviceName"`
	MobileDeviceUDN  string `xml:"MobileDeviceUDN"`
	MobileIPAndPort  string `xml:"MobileIPAndPort"`
}
type RegisterMobileDeviceResponse struct {
}

func (s *Service) RegisterMobileDevice(httpClient *http.Client, args *RegisterMobileDeviceArgs) (*RegisterMobileDeviceResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RegisterMobileDevice`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RegisterMobileDevice: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RegisterMobileDevice == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.RegisterMobileDevice()`)
	}

	return r.Body.RegisterMobileDevice, nil
}

type GetZoneGroupAttributesArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetZoneGroupAttributesResponse struct {
	CurrentZoneGroupName          string `xml:"CurrentZoneGroupName"`
	CurrentZoneGroupID            string `xml:"CurrentZoneGroupID"`
	CurrentZonePlayerUUIDsInGroup string `xml:"CurrentZonePlayerUUIDsInGroup"`
	CurrentMuseHouseholdId        string `xml:"CurrentMuseHouseholdId"`
}

func (s *Service) GetZoneGroupAttributes(httpClient *http.Client, args *GetZoneGroupAttributesArgs) (*GetZoneGroupAttributesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetZoneGroupAttributes`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetZoneGroupAttributes: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneGroupAttributes == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.GetZoneGroupAttributes()`)
	}

	return r.Body.GetZoneGroupAttributes, nil
}

type GetZoneGroupStateArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetZoneGroupStateResponse struct {
	ZoneGroupState string `xml:"ZoneGroupState"`
}

func (s *Service) GetZoneGroupState(httpClient *http.Client, args *GetZoneGroupStateArgs) (*GetZoneGroupStateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetZoneGroupState`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetZoneGroupState: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneGroupState == nil {
		return nil, errors.New(`unexpected respose from service calling zonegrouptopology.GetZoneGroupState()`)
	}

	return r.Body.GetZoneGroupState, nil
}
