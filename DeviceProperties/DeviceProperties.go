package deviceproperties

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:DeviceProperties:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/DeviceProperties/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/DeviceProperties/Event`)
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
	XMLName                xml.Name                    `xml:"s:Body"`
	SetLEDState            *SetLEDStateArgs            `xml:"u:SetLEDState,omitempty"`
	GetLEDState            *GetLEDStateArgs            `xml:"u:GetLEDState,omitempty"`
	AddBondedZones         *AddBondedZonesArgs         `xml:"u:AddBondedZones,omitempty"`
	RemoveBondedZones      *RemoveBondedZonesArgs      `xml:"u:RemoveBondedZones,omitempty"`
	CreateStereoPair       *CreateStereoPairArgs       `xml:"u:CreateStereoPair,omitempty"`
	SeparateStereoPair     *SeparateStereoPairArgs     `xml:"u:SeparateStereoPair,omitempty"`
	SetZoneAttributes      *SetZoneAttributesArgs      `xml:"u:SetZoneAttributes,omitempty"`
	GetZoneAttributes      *GetZoneAttributesArgs      `xml:"u:GetZoneAttributes,omitempty"`
	GetHouseholdID         *GetHouseholdIDArgs         `xml:"u:GetHouseholdID,omitempty"`
	GetZoneInfo            *GetZoneInfoArgs            `xml:"u:GetZoneInfo,omitempty"`
	SetAutoplayLinkedZones *SetAutoplayLinkedZonesArgs `xml:"u:SetAutoplayLinkedZones,omitempty"`
	GetAutoplayLinkedZones *GetAutoplayLinkedZonesArgs `xml:"u:GetAutoplayLinkedZones,omitempty"`
	SetAutoplayRoomUUID    *SetAutoplayRoomUUIDArgs    `xml:"u:SetAutoplayRoomUUID,omitempty"`
	GetAutoplayRoomUUID    *GetAutoplayRoomUUIDArgs    `xml:"u:GetAutoplayRoomUUID,omitempty"`
	SetAutoplayVolume      *SetAutoplayVolumeArgs      `xml:"u:SetAutoplayVolume,omitempty"`
	GetAutoplayVolume      *GetAutoplayVolumeArgs      `xml:"u:GetAutoplayVolume,omitempty"`
	SetUseAutoplayVolume   *SetUseAutoplayVolumeArgs   `xml:"u:SetUseAutoplayVolume,omitempty"`
	GetUseAutoplayVolume   *GetUseAutoplayVolumeArgs   `xml:"u:GetUseAutoplayVolume,omitempty"`
	AddHTSatellite         *AddHTSatelliteArgs         `xml:"u:AddHTSatellite,omitempty"`
	RemoveHTSatellite      *RemoveHTSatelliteArgs      `xml:"u:RemoveHTSatellite,omitempty"`
	EnterConfigMode        *EnterConfigModeArgs        `xml:"u:EnterConfigMode,omitempty"`
	ExitConfigMode         *ExitConfigModeArgs         `xml:"u:ExitConfigMode,omitempty"`
	GetButtonState         *GetButtonStateArgs         `xml:"u:GetButtonState,omitempty"`
	SetButtonLockState     *SetButtonLockStateArgs     `xml:"u:SetButtonLockState,omitempty"`
	GetButtonLockState     *GetButtonLockStateArgs     `xml:"u:GetButtonLockState,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                xml.Name                        `xml:"Body"`
	SetLEDState            *SetLEDStateResponse            `xml:"SetLEDStateResponse,omitempty"`
	GetLEDState            *GetLEDStateResponse            `xml:"GetLEDStateResponse,omitempty"`
	AddBondedZones         *AddBondedZonesResponse         `xml:"AddBondedZonesResponse,omitempty"`
	RemoveBondedZones      *RemoveBondedZonesResponse      `xml:"RemoveBondedZonesResponse,omitempty"`
	CreateStereoPair       *CreateStereoPairResponse       `xml:"CreateStereoPairResponse,omitempty"`
	SeparateStereoPair     *SeparateStereoPairResponse     `xml:"SeparateStereoPairResponse,omitempty"`
	SetZoneAttributes      *SetZoneAttributesResponse      `xml:"SetZoneAttributesResponse,omitempty"`
	GetZoneAttributes      *GetZoneAttributesResponse      `xml:"GetZoneAttributesResponse,omitempty"`
	GetHouseholdID         *GetHouseholdIDResponse         `xml:"GetHouseholdIDResponse,omitempty"`
	GetZoneInfo            *GetZoneInfoResponse            `xml:"GetZoneInfoResponse,omitempty"`
	SetAutoplayLinkedZones *SetAutoplayLinkedZonesResponse `xml:"SetAutoplayLinkedZonesResponse,omitempty"`
	GetAutoplayLinkedZones *GetAutoplayLinkedZonesResponse `xml:"GetAutoplayLinkedZonesResponse,omitempty"`
	SetAutoplayRoomUUID    *SetAutoplayRoomUUIDResponse    `xml:"SetAutoplayRoomUUIDResponse,omitempty"`
	GetAutoplayRoomUUID    *GetAutoplayRoomUUIDResponse    `xml:"GetAutoplayRoomUUIDResponse,omitempty"`
	SetAutoplayVolume      *SetAutoplayVolumeResponse      `xml:"SetAutoplayVolumeResponse,omitempty"`
	GetAutoplayVolume      *GetAutoplayVolumeResponse      `xml:"GetAutoplayVolumeResponse,omitempty"`
	SetUseAutoplayVolume   *SetUseAutoplayVolumeResponse   `xml:"SetUseAutoplayVolumeResponse,omitempty"`
	GetUseAutoplayVolume   *GetUseAutoplayVolumeResponse   `xml:"GetUseAutoplayVolumeResponse,omitempty"`
	AddHTSatellite         *AddHTSatelliteResponse         `xml:"AddHTSatelliteResponse,omitempty"`
	RemoveHTSatellite      *RemoveHTSatelliteResponse      `xml:"RemoveHTSatelliteResponse,omitempty"`
	EnterConfigMode        *EnterConfigModeResponse        `xml:"EnterConfigModeResponse,omitempty"`
	ExitConfigMode         *ExitConfigModeResponse         `xml:"ExitConfigModeResponse,omitempty"`
	GetButtonState         *GetButtonStateResponse         `xml:"GetButtonStateResponse,omitempty"`
	SetButtonLockState     *SetButtonLockStateResponse     `xml:"SetButtonLockStateResponse,omitempty"`
	GetButtonLockState     *GetButtonLockStateResponse     `xml:"GetButtonLockStateResponse,omitempty"`
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

type SetLEDStateArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
	// Allowed Value: On
	// Allowed Value: Off
	DesiredLEDState string `xml:"DesiredLEDState"`
}
type SetLEDStateResponse struct {
}

func (s *Service) SetLEDState(httpClient *http.Client, args *SetLEDStateArgs) (*SetLEDStateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetLEDState`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetLEDState: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetLEDState == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SetLEDState()`)
	}

	return r.Body.SetLEDState, nil
}

type GetLEDStateArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetLEDStateResponse struct {
	CurrentLEDState string `xml:"CurrentLEDState"`
}

func (s *Service) GetLEDState(httpClient *http.Client, args *GetLEDStateArgs) (*GetLEDStateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetLEDState`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetLEDState: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetLEDState == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetLEDState()`)
	}

	return r.Body.GetLEDState, nil
}

type AddBondedZonesArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
}
type AddBondedZonesResponse struct {
}

func (s *Service) AddBondedZones(httpClient *http.Client, args *AddBondedZonesArgs) (*AddBondedZonesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddBondedZones`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddBondedZones: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddBondedZones == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.AddBondedZones()`)
	}

	return r.Body.AddBondedZones, nil
}

type RemoveBondedZonesArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
	KeepGrouped   bool   `xml:"KeepGrouped"`
}
type RemoveBondedZonesResponse struct {
}

func (s *Service) RemoveBondedZones(httpClient *http.Client, args *RemoveBondedZonesArgs) (*RemoveBondedZonesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveBondedZones`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveBondedZones: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveBondedZones == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.RemoveBondedZones()`)
	}

	return r.Body.RemoveBondedZones, nil
}

type CreateStereoPairArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
}
type CreateStereoPairResponse struct {
}

func (s *Service) CreateStereoPair(httpClient *http.Client, args *CreateStereoPairArgs) (*CreateStereoPairResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`CreateStereoPair`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{CreateStereoPair: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateStereoPair == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.CreateStereoPair()`)
	}

	return r.Body.CreateStereoPair, nil
}

type SeparateStereoPairArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
}
type SeparateStereoPairResponse struct {
}

func (s *Service) SeparateStereoPair(httpClient *http.Client, args *SeparateStereoPairArgs) (*SeparateStereoPairResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SeparateStereoPair`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SeparateStereoPair: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SeparateStereoPair == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SeparateStereoPair()`)
	}

	return r.Body.SeparateStereoPair, nil
}

type SetZoneAttributesArgs struct {
	Xmlns                string `xml:"xmlns:u,attr"`
	DesiredZoneName      string `xml:"DesiredZoneName"`
	DesiredIcon          string `xml:"DesiredIcon"`
	DesiredConfiguration string `xml:"DesiredConfiguration"`
}
type SetZoneAttributesResponse struct {
}

func (s *Service) SetZoneAttributes(httpClient *http.Client, args *SetZoneAttributesArgs) (*SetZoneAttributesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetZoneAttributes`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetZoneAttributes: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetZoneAttributes == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SetZoneAttributes()`)
	}

	return r.Body.SetZoneAttributes, nil
}

type GetZoneAttributesArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetZoneAttributesResponse struct {
	CurrentZoneName      string `xml:"CurrentZoneName"`
	CurrentIcon          string `xml:"CurrentIcon"`
	CurrentConfiguration string `xml:"CurrentConfiguration"`
}

func (s *Service) GetZoneAttributes(httpClient *http.Client, args *GetZoneAttributesArgs) (*GetZoneAttributesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetZoneAttributes`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetZoneAttributes: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneAttributes == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetZoneAttributes()`)
	}

	return r.Body.GetZoneAttributes, nil
}

type GetHouseholdIDArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetHouseholdIDResponse struct {
	CurrentHouseholdID string `xml:"CurrentHouseholdID"`
}

func (s *Service) GetHouseholdID(httpClient *http.Client, args *GetHouseholdIDArgs) (*GetHouseholdIDResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetHouseholdID`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetHouseholdID: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetHouseholdID == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetHouseholdID()`)
	}

	return r.Body.GetHouseholdID, nil
}

type GetZoneInfoArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetZoneInfoResponse struct {
	SerialNumber           string `xml:"SerialNumber"`
	SoftwareVersion        string `xml:"SoftwareVersion"`
	DisplaySoftwareVersion string `xml:"DisplaySoftwareVersion"`
	HardwareVersion        string `xml:"HardwareVersion"`
	IPAddress              string `xml:"IPAddress"`
	MACAddress             string `xml:"MACAddress"`
	CopyrightInfo          string `xml:"CopyrightInfo"`
	ExtraInfo              string `xml:"ExtraInfo"`
	HTAudioIn              uint32 `xml:"HTAudioIn"`
	Flags                  uint32 `xml:"Flags"`
}

func (s *Service) GetZoneInfo(httpClient *http.Client, args *GetZoneInfoArgs) (*GetZoneInfoResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetZoneInfo`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetZoneInfo: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneInfo == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetZoneInfo()`)
	}

	return r.Body.GetZoneInfo, nil
}

type SetAutoplayLinkedZonesArgs struct {
	Xmlns              string `xml:"xmlns:u,attr"`
	IncludeLinkedZones bool   `xml:"IncludeLinkedZones"`
	Source             string `xml:"Source"`
}
type SetAutoplayLinkedZonesResponse struct {
}

func (s *Service) SetAutoplayLinkedZones(httpClient *http.Client, args *SetAutoplayLinkedZonesArgs) (*SetAutoplayLinkedZonesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetAutoplayLinkedZones`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetAutoplayLinkedZones: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAutoplayLinkedZones == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SetAutoplayLinkedZones()`)
	}

	return r.Body.SetAutoplayLinkedZones, nil
}

type GetAutoplayLinkedZonesArgs struct {
	Xmlns  string `xml:"xmlns:u,attr"`
	Source string `xml:"Source"`
}
type GetAutoplayLinkedZonesResponse struct {
	IncludeLinkedZones bool `xml:"IncludeLinkedZones"`
}

func (s *Service) GetAutoplayLinkedZones(httpClient *http.Client, args *GetAutoplayLinkedZonesArgs) (*GetAutoplayLinkedZonesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetAutoplayLinkedZones`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetAutoplayLinkedZones: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAutoplayLinkedZones == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetAutoplayLinkedZones()`)
	}

	return r.Body.GetAutoplayLinkedZones, nil
}

type SetAutoplayRoomUUIDArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	RoomUUID string `xml:"RoomUUID"`
	Source   string `xml:"Source"`
}
type SetAutoplayRoomUUIDResponse struct {
}

func (s *Service) SetAutoplayRoomUUID(httpClient *http.Client, args *SetAutoplayRoomUUIDArgs) (*SetAutoplayRoomUUIDResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetAutoplayRoomUUID`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetAutoplayRoomUUID: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAutoplayRoomUUID == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SetAutoplayRoomUUID()`)
	}

	return r.Body.SetAutoplayRoomUUID, nil
}

type GetAutoplayRoomUUIDArgs struct {
	Xmlns  string `xml:"xmlns:u,attr"`
	Source string `xml:"Source"`
}
type GetAutoplayRoomUUIDResponse struct {
	RoomUUID string `xml:"RoomUUID"`
}

func (s *Service) GetAutoplayRoomUUID(httpClient *http.Client, args *GetAutoplayRoomUUIDArgs) (*GetAutoplayRoomUUIDResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetAutoplayRoomUUID`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetAutoplayRoomUUID: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAutoplayRoomUUID == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetAutoplayRoomUUID()`)
	}

	return r.Body.GetAutoplayRoomUUID, nil
}

type SetAutoplayVolumeArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
	// Allowed Range: 0 -> 100 step: 1
	Volume uint16 `xml:"Volume"`
	Source string `xml:"Source"`
}
type SetAutoplayVolumeResponse struct {
}

func (s *Service) SetAutoplayVolume(httpClient *http.Client, args *SetAutoplayVolumeArgs) (*SetAutoplayVolumeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetAutoplayVolume`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetAutoplayVolume: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAutoplayVolume == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SetAutoplayVolume()`)
	}

	return r.Body.SetAutoplayVolume, nil
}

type GetAutoplayVolumeArgs struct {
	Xmlns  string `xml:"xmlns:u,attr"`
	Source string `xml:"Source"`
}
type GetAutoplayVolumeResponse struct {
	CurrentVolume uint16 `xml:"CurrentVolume"`
}

func (s *Service) GetAutoplayVolume(httpClient *http.Client, args *GetAutoplayVolumeArgs) (*GetAutoplayVolumeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetAutoplayVolume`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetAutoplayVolume: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAutoplayVolume == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetAutoplayVolume()`)
	}

	return r.Body.GetAutoplayVolume, nil
}

type SetUseAutoplayVolumeArgs struct {
	Xmlns     string `xml:"xmlns:u,attr"`
	UseVolume bool   `xml:"UseVolume"`
	Source    string `xml:"Source"`
}
type SetUseAutoplayVolumeResponse struct {
}

func (s *Service) SetUseAutoplayVolume(httpClient *http.Client, args *SetUseAutoplayVolumeArgs) (*SetUseAutoplayVolumeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetUseAutoplayVolume`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetUseAutoplayVolume: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetUseAutoplayVolume == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SetUseAutoplayVolume()`)
	}

	return r.Body.SetUseAutoplayVolume, nil
}

type GetUseAutoplayVolumeArgs struct {
	Xmlns  string `xml:"xmlns:u,attr"`
	Source string `xml:"Source"`
}
type GetUseAutoplayVolumeResponse struct {
	UseVolume bool `xml:"UseVolume"`
}

func (s *Service) GetUseAutoplayVolume(httpClient *http.Client, args *GetUseAutoplayVolumeArgs) (*GetUseAutoplayVolumeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetUseAutoplayVolume`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetUseAutoplayVolume: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetUseAutoplayVolume == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetUseAutoplayVolume()`)
	}

	return r.Body.GetUseAutoplayVolume, nil
}

type AddHTSatelliteArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
	HTSatChanMapSet string `xml:"HTSatChanMapSet"`
}
type AddHTSatelliteResponse struct {
}

func (s *Service) AddHTSatellite(httpClient *http.Client, args *AddHTSatelliteArgs) (*AddHTSatelliteResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddHTSatellite`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddHTSatellite: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddHTSatellite == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.AddHTSatellite()`)
	}

	return r.Body.AddHTSatellite, nil
}

type RemoveHTSatelliteArgs struct {
	Xmlns       string `xml:"xmlns:u,attr"`
	SatRoomUUID string `xml:"SatRoomUUID"`
}
type RemoveHTSatelliteResponse struct {
}

func (s *Service) RemoveHTSatellite(httpClient *http.Client, args *RemoveHTSatelliteArgs) (*RemoveHTSatelliteResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveHTSatellite`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveHTSatellite: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveHTSatellite == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.RemoveHTSatellite()`)
	}

	return r.Body.RemoveHTSatellite, nil
}

type EnterConfigModeArgs struct {
	Xmlns   string `xml:"xmlns:u,attr"`
	Mode    string `xml:"Mode"`
	Options string `xml:"Options"`
}
type EnterConfigModeResponse struct {
	State string `xml:"State"`
}

func (s *Service) EnterConfigMode(httpClient *http.Client, args *EnterConfigModeArgs) (*EnterConfigModeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`EnterConfigMode`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{EnterConfigMode: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EnterConfigMode == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.EnterConfigMode()`)
	}

	return r.Body.EnterConfigMode, nil
}

type ExitConfigModeArgs struct {
	Xmlns   string `xml:"xmlns:u,attr"`
	Options string `xml:"Options"`
}
type ExitConfigModeResponse struct {
}

func (s *Service) ExitConfigMode(httpClient *http.Client, args *ExitConfigModeArgs) (*ExitConfigModeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ExitConfigMode`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ExitConfigMode: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ExitConfigMode == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.ExitConfigMode()`)
	}

	return r.Body.ExitConfigMode, nil
}

type GetButtonStateArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetButtonStateResponse struct {
	State string `xml:"State"`
}

func (s *Service) GetButtonState(httpClient *http.Client, args *GetButtonStateArgs) (*GetButtonStateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetButtonState`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetButtonState: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetButtonState == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetButtonState()`)
	}

	return r.Body.GetButtonState, nil
}

type SetButtonLockStateArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
	// Allowed Value: On
	// Allowed Value: Off
	DesiredButtonLockState string `xml:"DesiredButtonLockState"`
}
type SetButtonLockStateResponse struct {
}

func (s *Service) SetButtonLockState(httpClient *http.Client, args *SetButtonLockStateArgs) (*SetButtonLockStateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetButtonLockState`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetButtonLockState: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetButtonLockState == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.SetButtonLockState()`)
	}

	return r.Body.SetButtonLockState, nil
}

type GetButtonLockStateArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetButtonLockStateResponse struct {
	CurrentButtonLockState string `xml:"CurrentButtonLockState"`
}

func (s *Service) GetButtonLockState(httpClient *http.Client, args *GetButtonLockStateArgs) (*GetButtonLockStateResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetButtonLockState`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetButtonLockState: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetButtonLockState == nil {
		return nil, errors.New(`unexpected respose from service calling deviceproperties.GetButtonLockState()`)
	}

	return r.Body.GetButtonLockState, nil
}
