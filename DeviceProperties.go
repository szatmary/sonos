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

type DevicePropertiesService struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewDevicePropertiesService(deviceUrl *url.URL) *DevicePropertiesService {
	c, _ := url.Parse("/DeviceProperties/Control")
	e, _ := url.Parse("/DeviceProperties/Event")
	return &DevicePropertiesService{
		ControlEndpoint: deviceUrl.ResolveReference(c),
		EventEndpoint:   deviceUrl.ResolveReference(e),
	}
}

type DevicePropertiesEnvelope struct {
	XMLName       xml.Name             `xml:"s:Envelope"`
	XMLNameSpace  string               `xml:"xmlns:s,attr"`
	EncodingStyle string               `xml:"s:encodingStyle,attr"`
	Body          DevicePropertiesBody `xml:"s:Body"`
}
type DevicePropertiesBody struct {
	XMLName                xml.Name                                    `xml:"s:Body"`
	SetLEDState            *DevicePropertiesSetLEDStateArgs            `xml:"u:SetLEDState,omitempty"`
	GetLEDState            *DevicePropertiesGetLEDStateArgs            `xml:"u:GetLEDState,omitempty"`
	AddBondedZones         *DevicePropertiesAddBondedZonesArgs         `xml:"u:AddBondedZones,omitempty"`
	RemoveBondedZones      *DevicePropertiesRemoveBondedZonesArgs      `xml:"u:RemoveBondedZones,omitempty"`
	CreateStereoPair       *DevicePropertiesCreateStereoPairArgs       `xml:"u:CreateStereoPair,omitempty"`
	SeparateStereoPair     *DevicePropertiesSeparateStereoPairArgs     `xml:"u:SeparateStereoPair,omitempty"`
	SetZoneAttributes      *DevicePropertiesSetZoneAttributesArgs      `xml:"u:SetZoneAttributes,omitempty"`
	GetZoneAttributes      *DevicePropertiesGetZoneAttributesArgs      `xml:"u:GetZoneAttributes,omitempty"`
	GetHouseholdID         *DevicePropertiesGetHouseholdIDArgs         `xml:"u:GetHouseholdID,omitempty"`
	GetZoneInfo            *DevicePropertiesGetZoneInfoArgs            `xml:"u:GetZoneInfo,omitempty"`
	SetAutoplayLinkedZones *DevicePropertiesSetAutoplayLinkedZonesArgs `xml:"u:SetAutoplayLinkedZones,omitempty"`
	GetAutoplayLinkedZones *DevicePropertiesGetAutoplayLinkedZonesArgs `xml:"u:GetAutoplayLinkedZones,omitempty"`
	SetAutoplayRoomUUID    *DevicePropertiesSetAutoplayRoomUUIDArgs    `xml:"u:SetAutoplayRoomUUID,omitempty"`
	GetAutoplayRoomUUID    *DevicePropertiesGetAutoplayRoomUUIDArgs    `xml:"u:GetAutoplayRoomUUID,omitempty"`
	SetAutoplayVolume      *DevicePropertiesSetAutoplayVolumeArgs      `xml:"u:SetAutoplayVolume,omitempty"`
	GetAutoplayVolume      *DevicePropertiesGetAutoplayVolumeArgs      `xml:"u:GetAutoplayVolume,omitempty"`
	SetUseAutoplayVolume   *DevicePropertiesSetUseAutoplayVolumeArgs   `xml:"u:SetUseAutoplayVolume,omitempty"`
	GetUseAutoplayVolume   *DevicePropertiesGetUseAutoplayVolumeArgs   `xml:"u:GetUseAutoplayVolume,omitempty"`
	AddHTSatellite         *DevicePropertiesAddHTSatelliteArgs         `xml:"u:AddHTSatellite,omitempty"`
	RemoveHTSatellite      *DevicePropertiesRemoveHTSatelliteArgs      `xml:"u:RemoveHTSatellite,omitempty"`
	EnterConfigMode        *DevicePropertiesEnterConfigModeArgs        `xml:"u:EnterConfigMode,omitempty"`
	ExitConfigMode         *DevicePropertiesExitConfigModeArgs         `xml:"u:ExitConfigMode,omitempty"`
	GetButtonState         *DevicePropertiesGetButtonStateArgs         `xml:"u:GetButtonState,omitempty"`
	SetButtonLockState     *DevicePropertiesSetButtonLockStateArgs     `xml:"u:SetButtonLockState,omitempty"`
	GetButtonLockState     *DevicePropertiesGetButtonLockStateArgs     `xml:"u:GetButtonLockState,omitempty"`
}
type DevicePropertiesEnvelopeResponse struct {
	XMLName       xml.Name                     `xml:"Envelope"`
	XMLNameSpace  string                       `xml:"xmlns:s,attr"`
	EncodingStyle string                       `xml:"encodingStyle,attr"`
	Body          DevicePropertiesBodyResponse `xml:"Body"`
}
type DevicePropertiesBodyResponse struct {
	XMLName                xml.Name                                        `xml:"Body"`
	SetLEDState            *DevicePropertiesSetLEDStateResponse            `xml:"SetLEDStateResponse"`
	GetLEDState            *DevicePropertiesGetLEDStateResponse            `xml:"GetLEDStateResponse"`
	AddBondedZones         *DevicePropertiesAddBondedZonesResponse         `xml:"AddBondedZonesResponse"`
	RemoveBondedZones      *DevicePropertiesRemoveBondedZonesResponse      `xml:"RemoveBondedZonesResponse"`
	CreateStereoPair       *DevicePropertiesCreateStereoPairResponse       `xml:"CreateStereoPairResponse"`
	SeparateStereoPair     *DevicePropertiesSeparateStereoPairResponse     `xml:"SeparateStereoPairResponse"`
	SetZoneAttributes      *DevicePropertiesSetZoneAttributesResponse      `xml:"SetZoneAttributesResponse"`
	GetZoneAttributes      *DevicePropertiesGetZoneAttributesResponse      `xml:"GetZoneAttributesResponse"`
	GetHouseholdID         *DevicePropertiesGetHouseholdIDResponse         `xml:"GetHouseholdIDResponse"`
	GetZoneInfo            *DevicePropertiesGetZoneInfoResponse            `xml:"GetZoneInfoResponse"`
	SetAutoplayLinkedZones *DevicePropertiesSetAutoplayLinkedZonesResponse `xml:"SetAutoplayLinkedZonesResponse"`
	GetAutoplayLinkedZones *DevicePropertiesGetAutoplayLinkedZonesResponse `xml:"GetAutoplayLinkedZonesResponse"`
	SetAutoplayRoomUUID    *DevicePropertiesSetAutoplayRoomUUIDResponse    `xml:"SetAutoplayRoomUUIDResponse"`
	GetAutoplayRoomUUID    *DevicePropertiesGetAutoplayRoomUUIDResponse    `xml:"GetAutoplayRoomUUIDResponse"`
	SetAutoplayVolume      *DevicePropertiesSetAutoplayVolumeResponse      `xml:"SetAutoplayVolumeResponse"`
	GetAutoplayVolume      *DevicePropertiesGetAutoplayVolumeResponse      `xml:"GetAutoplayVolumeResponse"`
	SetUseAutoplayVolume   *DevicePropertiesSetUseAutoplayVolumeResponse   `xml:"SetUseAutoplayVolumeResponse"`
	GetUseAutoplayVolume   *DevicePropertiesGetUseAutoplayVolumeResponse   `xml:"GetUseAutoplayVolumeResponse"`
	AddHTSatellite         *DevicePropertiesAddHTSatelliteResponse         `xml:"AddHTSatelliteResponse"`
	RemoveHTSatellite      *DevicePropertiesRemoveHTSatelliteResponse      `xml:"RemoveHTSatelliteResponse"`
	EnterConfigMode        *DevicePropertiesEnterConfigModeResponse        `xml:"EnterConfigModeResponse"`
	ExitConfigMode         *DevicePropertiesExitConfigModeResponse         `xml:"ExitConfigModeResponse"`
	GetButtonState         *DevicePropertiesGetButtonStateResponse         `xml:"GetButtonStateResponse"`
	SetButtonLockState     *DevicePropertiesSetButtonLockStateResponse     `xml:"SetButtonLockStateResponse"`
	GetButtonLockState     *DevicePropertiesGetButtonLockStateResponse     `xml:"GetButtonLockStateResponse"`
}

func (s *DevicePropertiesService) _DevicePropertiesExec(soapAction string, httpClient *http.Client, envelope *DevicePropertiesEnvelope) (*DevicePropertiesEnvelopeResponse, error) {
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
	var envelopeResponse DevicePropertiesEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type DevicePropertiesSetLEDStateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	// Allowed Value: On
	// Allowed Value: Off
	DesiredLEDState string `xml:"DesiredLEDState"`
}
type DevicePropertiesSetLEDStateResponse struct {
}

func (s *DevicePropertiesService) SetLEDState(httpClient *http.Client, args *DevicePropertiesSetLEDStateArgs) (*DevicePropertiesSetLEDStateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SetLEDState", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SetLEDState: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetLEDState == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetLEDState, nil
}

type DevicePropertiesGetLEDStateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type DevicePropertiesGetLEDStateResponse struct {
	CurrentLEDState string `xml:"CurrentLEDState"`
}

func (s *DevicePropertiesService) GetLEDState(httpClient *http.Client, args *DevicePropertiesGetLEDStateArgs) (*DevicePropertiesGetLEDStateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetLEDState", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetLEDState: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetLEDState == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetLEDState, nil
}

type DevicePropertiesAddBondedZonesArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
}
type DevicePropertiesAddBondedZonesResponse struct {
}

func (s *DevicePropertiesService) AddBondedZones(httpClient *http.Client, args *DevicePropertiesAddBondedZonesArgs) (*DevicePropertiesAddBondedZonesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#AddBondedZones", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{AddBondedZones: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddBondedZones == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddBondedZones, nil
}

type DevicePropertiesRemoveBondedZonesArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
	KeepGrouped   bool   `xml:"KeepGrouped"`
}
type DevicePropertiesRemoveBondedZonesResponse struct {
}

func (s *DevicePropertiesService) RemoveBondedZones(httpClient *http.Client, args *DevicePropertiesRemoveBondedZonesArgs) (*DevicePropertiesRemoveBondedZonesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#RemoveBondedZones", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{RemoveBondedZones: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveBondedZones == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveBondedZones, nil
}

type DevicePropertiesCreateStereoPairArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
}
type DevicePropertiesCreateStereoPairResponse struct {
}

func (s *DevicePropertiesService) CreateStereoPair(httpClient *http.Client, args *DevicePropertiesCreateStereoPairArgs) (*DevicePropertiesCreateStereoPairResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#CreateStereoPair", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{CreateStereoPair: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateStereoPair == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.CreateStereoPair, nil
}

type DevicePropertiesSeparateStereoPairArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	ChannelMapSet string `xml:"ChannelMapSet"`
}
type DevicePropertiesSeparateStereoPairResponse struct {
}

func (s *DevicePropertiesService) SeparateStereoPair(httpClient *http.Client, args *DevicePropertiesSeparateStereoPairArgs) (*DevicePropertiesSeparateStereoPairResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SeparateStereoPair", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SeparateStereoPair: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SeparateStereoPair == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SeparateStereoPair, nil
}

type DevicePropertiesSetZoneAttributesArgs struct {
	XMLNameSpace         string `xml:"xmlns:u,attr"`
	DesiredZoneName      string `xml:"DesiredZoneName"`
	DesiredIcon          string `xml:"DesiredIcon"`
	DesiredConfiguration string `xml:"DesiredConfiguration"`
}
type DevicePropertiesSetZoneAttributesResponse struct {
}

func (s *DevicePropertiesService) SetZoneAttributes(httpClient *http.Client, args *DevicePropertiesSetZoneAttributesArgs) (*DevicePropertiesSetZoneAttributesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SetZoneAttributes", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SetZoneAttributes: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetZoneAttributes == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetZoneAttributes, nil
}

type DevicePropertiesGetZoneAttributesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type DevicePropertiesGetZoneAttributesResponse struct {
	CurrentZoneName      string `xml:"CurrentZoneName"`
	CurrentIcon          string `xml:"CurrentIcon"`
	CurrentConfiguration string `xml:"CurrentConfiguration"`
}

func (s *DevicePropertiesService) GetZoneAttributes(httpClient *http.Client, args *DevicePropertiesGetZoneAttributesArgs) (*DevicePropertiesGetZoneAttributesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetZoneAttributes", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetZoneAttributes: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneAttributes == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetZoneAttributes, nil
}

type DevicePropertiesGetHouseholdIDArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type DevicePropertiesGetHouseholdIDResponse struct {
	CurrentHouseholdID string `xml:"CurrentHouseholdID"`
}

func (s *DevicePropertiesService) GetHouseholdID(httpClient *http.Client, args *DevicePropertiesGetHouseholdIDArgs) (*DevicePropertiesGetHouseholdIDResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetHouseholdID", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetHouseholdID: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetHouseholdID == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetHouseholdID, nil
}

type DevicePropertiesGetZoneInfoArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type DevicePropertiesGetZoneInfoResponse struct {
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

func (s *DevicePropertiesService) GetZoneInfo(httpClient *http.Client, args *DevicePropertiesGetZoneInfoArgs) (*DevicePropertiesGetZoneInfoResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetZoneInfo", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetZoneInfo: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetZoneInfo == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetZoneInfo, nil
}

type DevicePropertiesSetAutoplayLinkedZonesArgs struct {
	XMLNameSpace       string `xml:"xmlns:u,attr"`
	IncludeLinkedZones bool   `xml:"IncludeLinkedZones"`
	Source             string `xml:"Source"`
}
type DevicePropertiesSetAutoplayLinkedZonesResponse struct {
}

func (s *DevicePropertiesService) SetAutoplayLinkedZones(httpClient *http.Client, args *DevicePropertiesSetAutoplayLinkedZonesArgs) (*DevicePropertiesSetAutoplayLinkedZonesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SetAutoplayLinkedZones", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SetAutoplayLinkedZones: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAutoplayLinkedZones == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetAutoplayLinkedZones, nil
}

type DevicePropertiesGetAutoplayLinkedZonesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Source       string `xml:"Source"`
}
type DevicePropertiesGetAutoplayLinkedZonesResponse struct {
	IncludeLinkedZones bool `xml:"IncludeLinkedZones"`
}

func (s *DevicePropertiesService) GetAutoplayLinkedZones(httpClient *http.Client, args *DevicePropertiesGetAutoplayLinkedZonesArgs) (*DevicePropertiesGetAutoplayLinkedZonesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetAutoplayLinkedZones", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetAutoplayLinkedZones: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAutoplayLinkedZones == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetAutoplayLinkedZones, nil
}

type DevicePropertiesSetAutoplayRoomUUIDArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	RoomUUID     string `xml:"RoomUUID"`
	Source       string `xml:"Source"`
}
type DevicePropertiesSetAutoplayRoomUUIDResponse struct {
}

func (s *DevicePropertiesService) SetAutoplayRoomUUID(httpClient *http.Client, args *DevicePropertiesSetAutoplayRoomUUIDArgs) (*DevicePropertiesSetAutoplayRoomUUIDResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SetAutoplayRoomUUID", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SetAutoplayRoomUUID: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAutoplayRoomUUID == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetAutoplayRoomUUID, nil
}

type DevicePropertiesGetAutoplayRoomUUIDArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Source       string `xml:"Source"`
}
type DevicePropertiesGetAutoplayRoomUUIDResponse struct {
	RoomUUID string `xml:"RoomUUID"`
}

func (s *DevicePropertiesService) GetAutoplayRoomUUID(httpClient *http.Client, args *DevicePropertiesGetAutoplayRoomUUIDArgs) (*DevicePropertiesGetAutoplayRoomUUIDResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetAutoplayRoomUUID", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetAutoplayRoomUUID: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAutoplayRoomUUID == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetAutoplayRoomUUID, nil
}

type DevicePropertiesSetAutoplayVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	// Allowed Range: 0 -> 100 step: 1
	Volume uint16 `xml:"Volume"`
	Source string `xml:"Source"`
}
type DevicePropertiesSetAutoplayVolumeResponse struct {
}

func (s *DevicePropertiesService) SetAutoplayVolume(httpClient *http.Client, args *DevicePropertiesSetAutoplayVolumeArgs) (*DevicePropertiesSetAutoplayVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SetAutoplayVolume", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SetAutoplayVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAutoplayVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetAutoplayVolume, nil
}

type DevicePropertiesGetAutoplayVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Source       string `xml:"Source"`
}
type DevicePropertiesGetAutoplayVolumeResponse struct {
	CurrentVolume uint16 `xml:"CurrentVolume"`
}

func (s *DevicePropertiesService) GetAutoplayVolume(httpClient *http.Client, args *DevicePropertiesGetAutoplayVolumeArgs) (*DevicePropertiesGetAutoplayVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetAutoplayVolume", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetAutoplayVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAutoplayVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetAutoplayVolume, nil
}

type DevicePropertiesSetUseAutoplayVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	UseVolume    bool   `xml:"UseVolume"`
	Source       string `xml:"Source"`
}
type DevicePropertiesSetUseAutoplayVolumeResponse struct {
}

func (s *DevicePropertiesService) SetUseAutoplayVolume(httpClient *http.Client, args *DevicePropertiesSetUseAutoplayVolumeArgs) (*DevicePropertiesSetUseAutoplayVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SetUseAutoplayVolume", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SetUseAutoplayVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetUseAutoplayVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetUseAutoplayVolume, nil
}

type DevicePropertiesGetUseAutoplayVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Source       string `xml:"Source"`
}
type DevicePropertiesGetUseAutoplayVolumeResponse struct {
	UseVolume bool `xml:"UseVolume"`
}

func (s *DevicePropertiesService) GetUseAutoplayVolume(httpClient *http.Client, args *DevicePropertiesGetUseAutoplayVolumeArgs) (*DevicePropertiesGetUseAutoplayVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetUseAutoplayVolume", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetUseAutoplayVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetUseAutoplayVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetUseAutoplayVolume, nil
}

type DevicePropertiesAddHTSatelliteArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	HTSatChanMapSet string `xml:"HTSatChanMapSet"`
}
type DevicePropertiesAddHTSatelliteResponse struct {
}

func (s *DevicePropertiesService) AddHTSatellite(httpClient *http.Client, args *DevicePropertiesAddHTSatelliteArgs) (*DevicePropertiesAddHTSatelliteResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#AddHTSatellite", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{AddHTSatellite: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddHTSatellite == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddHTSatellite, nil
}

type DevicePropertiesRemoveHTSatelliteArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	SatRoomUUID  string `xml:"SatRoomUUID"`
}
type DevicePropertiesRemoveHTSatelliteResponse struct {
}

func (s *DevicePropertiesService) RemoveHTSatellite(httpClient *http.Client, args *DevicePropertiesRemoveHTSatelliteArgs) (*DevicePropertiesRemoveHTSatelliteResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#RemoveHTSatellite", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{RemoveHTSatellite: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveHTSatellite == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveHTSatellite, nil
}

type DevicePropertiesEnterConfigModeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Mode         string `xml:"Mode"`
	Options      string `xml:"Options"`
}
type DevicePropertiesEnterConfigModeResponse struct {
	State string `xml:"State"`
}

func (s *DevicePropertiesService) EnterConfigMode(httpClient *http.Client, args *DevicePropertiesEnterConfigModeArgs) (*DevicePropertiesEnterConfigModeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#EnterConfigMode", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{EnterConfigMode: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EnterConfigMode == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.EnterConfigMode, nil
}

type DevicePropertiesExitConfigModeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Options      string `xml:"Options"`
}
type DevicePropertiesExitConfigModeResponse struct {
}

func (s *DevicePropertiesService) ExitConfigMode(httpClient *http.Client, args *DevicePropertiesExitConfigModeArgs) (*DevicePropertiesExitConfigModeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#ExitConfigMode", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{ExitConfigMode: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ExitConfigMode == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ExitConfigMode, nil
}

type DevicePropertiesGetButtonStateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type DevicePropertiesGetButtonStateResponse struct {
	State string `xml:"State"`
}

func (s *DevicePropertiesService) GetButtonState(httpClient *http.Client, args *DevicePropertiesGetButtonStateArgs) (*DevicePropertiesGetButtonStateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetButtonState", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetButtonState: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetButtonState == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetButtonState, nil
}

type DevicePropertiesSetButtonLockStateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	// Allowed Value: On
	// Allowed Value: Off
	DesiredButtonLockState string `xml:"DesiredButtonLockState"`
}
type DevicePropertiesSetButtonLockStateResponse struct {
}

func (s *DevicePropertiesService) SetButtonLockState(httpClient *http.Client, args *DevicePropertiesSetButtonLockStateArgs) (*DevicePropertiesSetButtonLockStateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#SetButtonLockState", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{SetButtonLockState: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetButtonLockState == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetButtonLockState, nil
}

type DevicePropertiesGetButtonLockStateArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type DevicePropertiesGetButtonLockStateResponse struct {
	CurrentButtonLockState string `xml:"CurrentButtonLockState"`
}

func (s *DevicePropertiesService) GetButtonLockState(httpClient *http.Client, args *DevicePropertiesGetButtonLockStateArgs) (*DevicePropertiesGetButtonLockStateResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:DeviceProperties:1"
	r, err := s._DevicePropertiesExec("urn:schemas-upnp-org:service:DeviceProperties:1#GetButtonLockState", httpClient,
		&DevicePropertiesEnvelope{
			Body:          DevicePropertiesBody{GetButtonLockState: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetButtonLockState == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetButtonLockState, nil
}
func (s *DevicePropertiesService) DevicePropertiesSubscribe(callback url.URL) error {
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
type DevicePropertiesSettingsReplicationState string
type DevicePropertiesZoneName string
type DevicePropertiesIcon string
type DevicePropertiesConfiguration string
type DevicePropertiesInvisible bool
type DevicePropertiesIsZoneBridge bool
type DevicePropertiesAirPlayEnabled bool
type DevicePropertiesSupportsAudioIn bool
type DevicePropertiesSupportsAudioClip bool
type DevicePropertiesIsIdle bool
type DevicePropertiesMoreInfo string
type DevicePropertiesChannelMapSet string
type DevicePropertiesHTSatChanMapSet string
type DevicePropertiesHTFreq uint32
type DevicePropertiesHTBondedZoneCommitState uint32
type DevicePropertiesOrientation int32
type DevicePropertiesLastChangedPlayState string
type DevicePropertiesRoomCalibrationState int32
type DevicePropertiesAvailableRoomCalibration string
type DevicePropertiesTVConfigurationError bool
type DevicePropertiesHdmiCecAvailable bool
type DevicePropertiesWirelessMode uint32
type DevicePropertiesWirelessLeafOnly bool
type DevicePropertiesHasConfiguredSSID bool
type DevicePropertiesChannelFreq uint32
type DevicePropertiesBehindWifiExtender uint32
type DevicePropertiesWifiEnabled bool
type DevicePropertiesConfigMode string
type DevicePropertiesSecureRegState uint32
type DevicePropertiesVoiceConfigState uint32
type DevicePropertiesMicEnabled uint32
type DevicePropertiesUpnpEvent struct {
	XMLName      xml.Name                   `xml:"propertyset"`
	XMLNameSpace string                     `xml:"xmlns:e,attr"`
	Properties   []DevicePropertiesProperty `xml:"property"`
}
type DevicePropertiesProperty struct {
	XMLName                  xml.Name                                  `xml:"property"`
	SettingsReplicationState *DevicePropertiesSettingsReplicationState `xml:"SettingsReplicationState"`
	ZoneName                 *DevicePropertiesZoneName                 `xml:"ZoneName"`
	Icon                     *DevicePropertiesIcon                     `xml:"Icon"`
	Configuration            *DevicePropertiesConfiguration            `xml:"Configuration"`
	Invisible                *DevicePropertiesInvisible                `xml:"Invisible"`
	IsZoneBridge             *DevicePropertiesIsZoneBridge             `xml:"IsZoneBridge"`
	AirPlayEnabled           *DevicePropertiesAirPlayEnabled           `xml:"AirPlayEnabled"`
	SupportsAudioIn          *DevicePropertiesSupportsAudioIn          `xml:"SupportsAudioIn"`
	SupportsAudioClip        *DevicePropertiesSupportsAudioClip        `xml:"SupportsAudioClip"`
	IsIdle                   *DevicePropertiesIsIdle                   `xml:"IsIdle"`
	MoreInfo                 *DevicePropertiesMoreInfo                 `xml:"MoreInfo"`
	ChannelMapSet            *DevicePropertiesChannelMapSet            `xml:"ChannelMapSet"`
	HTSatChanMapSet          *DevicePropertiesHTSatChanMapSet          `xml:"HTSatChanMapSet"`
	HTFreq                   *DevicePropertiesHTFreq                   `xml:"HTFreq"`
	HTBondedZoneCommitState  *DevicePropertiesHTBondedZoneCommitState  `xml:"HTBondedZoneCommitState"`
	Orientation              *DevicePropertiesOrientation              `xml:"Orientation"`
	LastChangedPlayState     *DevicePropertiesLastChangedPlayState     `xml:"LastChangedPlayState"`
	RoomCalibrationState     *DevicePropertiesRoomCalibrationState     `xml:"RoomCalibrationState"`
	AvailableRoomCalibration *DevicePropertiesAvailableRoomCalibration `xml:"AvailableRoomCalibration"`
	TVConfigurationError     *DevicePropertiesTVConfigurationError     `xml:"TVConfigurationError"`
	HdmiCecAvailable         *DevicePropertiesHdmiCecAvailable         `xml:"HdmiCecAvailable"`
	WirelessMode             *DevicePropertiesWirelessMode             `xml:"WirelessMode"`
	WirelessLeafOnly         *DevicePropertiesWirelessLeafOnly         `xml:"WirelessLeafOnly"`
	HasConfiguredSSID        *DevicePropertiesHasConfiguredSSID        `xml:"HasConfiguredSSID"`
	ChannelFreq              *DevicePropertiesChannelFreq              `xml:"ChannelFreq"`
	BehindWifiExtender       *DevicePropertiesBehindWifiExtender       `xml:"BehindWifiExtender"`
	WifiEnabled              *DevicePropertiesWifiEnabled              `xml:"WifiEnabled"`
	ConfigMode               *DevicePropertiesConfigMode               `xml:"ConfigMode"`
	SecureRegState           *DevicePropertiesSecureRegState           `xml:"SecureRegState"`
	VoiceConfigState         *DevicePropertiesVoiceConfigState         `xml:"VoiceConfigState"`
	MicEnabled               *DevicePropertiesMicEnabled               `xml:"MicEnabled"`
}

func DevicePropertiesDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt DevicePropertiesUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.SettingsReplicationState != nil:
			zp.EventCallback(*prop.SettingsReplicationState)
		case prop.ZoneName != nil:
			zp.EventCallback(*prop.ZoneName)
		case prop.Icon != nil:
			zp.EventCallback(*prop.Icon)
		case prop.Configuration != nil:
			zp.EventCallback(*prop.Configuration)
		case prop.Invisible != nil:
			zp.EventCallback(*prop.Invisible)
		case prop.IsZoneBridge != nil:
			zp.EventCallback(*prop.IsZoneBridge)
		case prop.AirPlayEnabled != nil:
			zp.EventCallback(*prop.AirPlayEnabled)
		case prop.SupportsAudioIn != nil:
			zp.EventCallback(*prop.SupportsAudioIn)
		case prop.SupportsAudioClip != nil:
			zp.EventCallback(*prop.SupportsAudioClip)
		case prop.IsIdle != nil:
			zp.EventCallback(*prop.IsIdle)
		case prop.MoreInfo != nil:
			zp.EventCallback(*prop.MoreInfo)
		case prop.ChannelMapSet != nil:
			zp.EventCallback(*prop.ChannelMapSet)
		case prop.HTSatChanMapSet != nil:
			zp.EventCallback(*prop.HTSatChanMapSet)
		case prop.HTFreq != nil:
			zp.EventCallback(*prop.HTFreq)
		case prop.HTBondedZoneCommitState != nil:
			zp.EventCallback(*prop.HTBondedZoneCommitState)
		case prop.Orientation != nil:
			zp.EventCallback(*prop.Orientation)
		case prop.LastChangedPlayState != nil:
			zp.EventCallback(*prop.LastChangedPlayState)
		case prop.RoomCalibrationState != nil:
			zp.EventCallback(*prop.RoomCalibrationState)
		case prop.AvailableRoomCalibration != nil:
			zp.EventCallback(*prop.AvailableRoomCalibration)
		case prop.TVConfigurationError != nil:
			zp.EventCallback(*prop.TVConfigurationError)
		case prop.HdmiCecAvailable != nil:
			zp.EventCallback(*prop.HdmiCecAvailable)
		case prop.WirelessMode != nil:
			zp.EventCallback(*prop.WirelessMode)
		case prop.WirelessLeafOnly != nil:
			zp.EventCallback(*prop.WirelessLeafOnly)
		case prop.HasConfiguredSSID != nil:
			zp.EventCallback(*prop.HasConfiguredSSID)
		case prop.ChannelFreq != nil:
			zp.EventCallback(*prop.ChannelFreq)
		case prop.BehindWifiExtender != nil:
			zp.EventCallback(*prop.BehindWifiExtender)
		case prop.WifiEnabled != nil:
			zp.EventCallback(*prop.WifiEnabled)
		case prop.ConfigMode != nil:
			zp.EventCallback(*prop.ConfigMode)
		case prop.SecureRegState != nil:
			zp.EventCallback(*prop.SecureRegState)
		case prop.VoiceConfigState != nil:
			zp.EventCallback(*prop.VoiceConfigState)
		case prop.MicEnabled != nil:
			zp.EventCallback(*prop.MicEnabled)
		}
	}
}
