package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DevicePropertiesService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewDevicePropertiesService(deviceUrl *url.URL) *DevicePropertiesService {
	c, _ := url.Parse("/DeviceProperties/Control")
	e, _ := url.Parse("/DeviceProperties/Event")
	return &DevicePropertiesService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *DevicePropertiesService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *DevicePropertiesService) EventEndpoint() *url.URL {
	return s.eventEndpoint
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

type DevicePropertiesUpnpEvent struct {
	XMLName      xml.Name                   `xml:"propertyset"`
	XMLNameSpace string                     `xml:"xmlns:e,attr"`
	Properties   []DevicePropertiesProperty `xml:"property"`
}
type DevicePropertiesProperty struct {
	XMLName                  xml.Name `xml:"property"`
	SettingsReplicationState *string  `xml:"SettingsReplicationState"`
	ZoneName                 *string  `xml:"ZoneName"`
	Icon                     *string  `xml:"Icon"`
	Configuration            *string  `xml:"Configuration"`
	Invisible                *bool    `xml:"Invisible"`
	IsZoneBridge             *bool    `xml:"IsZoneBridge"`
	AirPlayEnabled           *bool    `xml:"AirPlayEnabled"`
	SupportsAudioIn          *bool    `xml:"SupportsAudioIn"`
	SupportsAudioClip        *bool    `xml:"SupportsAudioClip"`
	IsIdle                   *bool    `xml:"IsIdle"`
	MoreInfo                 *string  `xml:"MoreInfo"`
	ChannelMapSet            *string  `xml:"ChannelMapSet"`
	HTSatChanMapSet          *string  `xml:"HTSatChanMapSet"`
	HTFreq                   *uint32  `xml:"HTFreq"`
	HTBondedZoneCommitState  *uint32  `xml:"HTBondedZoneCommitState"`
	Orientation              *int32   `xml:"Orientation"`
	LastChangedPlayState     *string  `xml:"LastChangedPlayState"`
	RoomCalibrationState     *int32   `xml:"RoomCalibrationState"`
	AvailableRoomCalibration *string  `xml:"AvailableRoomCalibration"`
	TVConfigurationError     *bool    `xml:"TVConfigurationError"`
	HdmiCecAvailable         *bool    `xml:"HdmiCecAvailable"`
	WirelessMode             *uint32  `xml:"WirelessMode"`
	WirelessLeafOnly         *bool    `xml:"WirelessLeafOnly"`
	HasConfiguredSSID        *bool    `xml:"HasConfiguredSSID"`
	ChannelFreq              *uint32  `xml:"ChannelFreq"`
	BehindWifiExtender       *uint32  `xml:"BehindWifiExtender"`
	WifiEnabled              *bool    `xml:"WifiEnabled"`
	ConfigMode               *string  `xml:"ConfigMode"`
	SecureRegState           *uint32  `xml:"SecureRegState"`
	VoiceConfigState         *uint32  `xml:"VoiceConfigState"`
	MicEnabled               *uint32  `xml:"MicEnabled"`
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
			dispatchDevicePropertiesSettingsReplicationState(*prop.SettingsReplicationState) // string
		case prop.ZoneName != nil:
			dispatchDevicePropertiesZoneName(*prop.ZoneName) // string
		case prop.Icon != nil:
			dispatchDevicePropertiesIcon(*prop.Icon) // string
		case prop.Configuration != nil:
			dispatchDevicePropertiesConfiguration(*prop.Configuration) // string
		case prop.Invisible != nil:
			dispatchDevicePropertiesInvisible(*prop.Invisible) // bool
		case prop.IsZoneBridge != nil:
			dispatchDevicePropertiesIsZoneBridge(*prop.IsZoneBridge) // bool
		case prop.AirPlayEnabled != nil:
			dispatchDevicePropertiesAirPlayEnabled(*prop.AirPlayEnabled) // bool
		case prop.SupportsAudioIn != nil:
			dispatchDevicePropertiesSupportsAudioIn(*prop.SupportsAudioIn) // bool
		case prop.SupportsAudioClip != nil:
			dispatchDevicePropertiesSupportsAudioClip(*prop.SupportsAudioClip) // bool
		case prop.IsIdle != nil:
			dispatchDevicePropertiesIsIdle(*prop.IsIdle) // bool
		case prop.MoreInfo != nil:
			dispatchDevicePropertiesMoreInfo(*prop.MoreInfo) // string
		case prop.ChannelMapSet != nil:
			dispatchDevicePropertiesChannelMapSet(*prop.ChannelMapSet) // string
		case prop.HTSatChanMapSet != nil:
			dispatchDevicePropertiesHTSatChanMapSet(*prop.HTSatChanMapSet) // string
		case prop.HTFreq != nil:
			dispatchDevicePropertiesHTFreq(*prop.HTFreq) // uint32
		case prop.HTBondedZoneCommitState != nil:
			dispatchDevicePropertiesHTBondedZoneCommitState(*prop.HTBondedZoneCommitState) // uint32
		case prop.Orientation != nil:
			dispatchDevicePropertiesOrientation(*prop.Orientation) // int32
		case prop.LastChangedPlayState != nil:
			dispatchDevicePropertiesLastChangedPlayState(*prop.LastChangedPlayState) // string
		case prop.RoomCalibrationState != nil:
			dispatchDevicePropertiesRoomCalibrationState(*prop.RoomCalibrationState) // int32
		case prop.AvailableRoomCalibration != nil:
			dispatchDevicePropertiesAvailableRoomCalibration(*prop.AvailableRoomCalibration) // string
		case prop.TVConfigurationError != nil:
			dispatchDevicePropertiesTVConfigurationError(*prop.TVConfigurationError) // bool
		case prop.HdmiCecAvailable != nil:
			dispatchDevicePropertiesHdmiCecAvailable(*prop.HdmiCecAvailable) // bool
		case prop.WirelessMode != nil:
			dispatchDevicePropertiesWirelessMode(*prop.WirelessMode) // uint32
		case prop.WirelessLeafOnly != nil:
			dispatchDevicePropertiesWirelessLeafOnly(*prop.WirelessLeafOnly) // bool
		case prop.HasConfiguredSSID != nil:
			dispatchDevicePropertiesHasConfiguredSSID(*prop.HasConfiguredSSID) // bool
		case prop.ChannelFreq != nil:
			dispatchDevicePropertiesChannelFreq(*prop.ChannelFreq) // uint32
		case prop.BehindWifiExtender != nil:
			dispatchDevicePropertiesBehindWifiExtender(*prop.BehindWifiExtender) // uint32
		case prop.WifiEnabled != nil:
			dispatchDevicePropertiesWifiEnabled(*prop.WifiEnabled) // bool
		case prop.ConfigMode != nil:
			dispatchDevicePropertiesConfigMode(*prop.ConfigMode) // string
		case prop.SecureRegState != nil:
			dispatchDevicePropertiesSecureRegState(*prop.SecureRegState) // uint32
		case prop.VoiceConfigState != nil:
			dispatchDevicePropertiesVoiceConfigState(*prop.VoiceConfigState) // uint32
		case prop.MicEnabled != nil:
			dispatchDevicePropertiesMicEnabled(*prop.MicEnabled) // uint32
		}
	}
}
