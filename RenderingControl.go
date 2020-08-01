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

type RenderingControlService struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewRenderingControlService(deviceUrl *url.URL) *RenderingControlService {
	c, _ := url.Parse("/MediaRenderer/RenderingControl/Control")
	e, _ := url.Parse("/MediaRenderer/RenderingControl/Event")
	return &RenderingControlService{
		ControlEndpoint: deviceUrl.ResolveReference(c),
		EventEndpoint:   deviceUrl.ResolveReference(e),
	}
}

type RenderingControlEnvelope struct {
	XMLName       xml.Name             `xml:"s:Envelope"`
	XMLNameSpace  string               `xml:"xmlns:s,attr"`
	EncodingStyle string               `xml:"s:encodingStyle,attr"`
	Body          RenderingControlBody `xml:"s:Body"`
}
type RenderingControlBody struct {
	XMLName                  xml.Name                                      `xml:"s:Body"`
	GetMute                  *RenderingControlGetMuteArgs                  `xml:"u:GetMute,omitempty"`
	SetMute                  *RenderingControlSetMuteArgs                  `xml:"u:SetMute,omitempty"`
	ResetBasicEQ             *RenderingControlResetBasicEQArgs             `xml:"u:ResetBasicEQ,omitempty"`
	ResetExtEQ               *RenderingControlResetExtEQArgs               `xml:"u:ResetExtEQ,omitempty"`
	GetVolume                *RenderingControlGetVolumeArgs                `xml:"u:GetVolume,omitempty"`
	SetVolume                *RenderingControlSetVolumeArgs                `xml:"u:SetVolume,omitempty"`
	SetRelativeVolume        *RenderingControlSetRelativeVolumeArgs        `xml:"u:SetRelativeVolume,omitempty"`
	GetVolumeDB              *RenderingControlGetVolumeDBArgs              `xml:"u:GetVolumeDB,omitempty"`
	SetVolumeDB              *RenderingControlSetVolumeDBArgs              `xml:"u:SetVolumeDB,omitempty"`
	GetVolumeDBRange         *RenderingControlGetVolumeDBRangeArgs         `xml:"u:GetVolumeDBRange,omitempty"`
	GetBass                  *RenderingControlGetBassArgs                  `xml:"u:GetBass,omitempty"`
	SetBass                  *RenderingControlSetBassArgs                  `xml:"u:SetBass,omitempty"`
	GetTreble                *RenderingControlGetTrebleArgs                `xml:"u:GetTreble,omitempty"`
	SetTreble                *RenderingControlSetTrebleArgs                `xml:"u:SetTreble,omitempty"`
	GetEQ                    *RenderingControlGetEQArgs                    `xml:"u:GetEQ,omitempty"`
	SetEQ                    *RenderingControlSetEQArgs                    `xml:"u:SetEQ,omitempty"`
	GetLoudness              *RenderingControlGetLoudnessArgs              `xml:"u:GetLoudness,omitempty"`
	SetLoudness              *RenderingControlSetLoudnessArgs              `xml:"u:SetLoudness,omitempty"`
	GetSupportsOutputFixed   *RenderingControlGetSupportsOutputFixedArgs   `xml:"u:GetSupportsOutputFixed,omitempty"`
	GetOutputFixed           *RenderingControlGetOutputFixedArgs           `xml:"u:GetOutputFixed,omitempty"`
	SetOutputFixed           *RenderingControlSetOutputFixedArgs           `xml:"u:SetOutputFixed,omitempty"`
	GetHeadphoneConnected    *RenderingControlGetHeadphoneConnectedArgs    `xml:"u:GetHeadphoneConnected,omitempty"`
	RampToVolume             *RenderingControlRampToVolumeArgs             `xml:"u:RampToVolume,omitempty"`
	RestoreVolumePriorToRamp *RenderingControlRestoreVolumePriorToRampArgs `xml:"u:RestoreVolumePriorToRamp,omitempty"`
	SetChannelMap            *RenderingControlSetChannelMapArgs            `xml:"u:SetChannelMap,omitempty"`
	SetRoomCalibrationX      *RenderingControlSetRoomCalibrationXArgs      `xml:"u:SetRoomCalibrationX,omitempty"`
	GetRoomCalibrationStatus *RenderingControlGetRoomCalibrationStatusArgs `xml:"u:GetRoomCalibrationStatus,omitempty"`
	SetRoomCalibrationStatus *RenderingControlSetRoomCalibrationStatusArgs `xml:"u:SetRoomCalibrationStatus,omitempty"`
}
type RenderingControlEnvelopeResponse struct {
	XMLName       xml.Name                     `xml:"Envelope"`
	XMLNameSpace  string                       `xml:"xmlns:s,attr"`
	EncodingStyle string                       `xml:"encodingStyle,attr"`
	Body          RenderingControlBodyResponse `xml:"Body"`
}
type RenderingControlBodyResponse struct {
	XMLName                  xml.Name                                          `xml:"Body"`
	GetMute                  *RenderingControlGetMuteResponse                  `xml:"GetMuteResponse"`
	SetMute                  *RenderingControlSetMuteResponse                  `xml:"SetMuteResponse"`
	ResetBasicEQ             *RenderingControlResetBasicEQResponse             `xml:"ResetBasicEQResponse"`
	ResetExtEQ               *RenderingControlResetExtEQResponse               `xml:"ResetExtEQResponse"`
	GetVolume                *RenderingControlGetVolumeResponse                `xml:"GetVolumeResponse"`
	SetVolume                *RenderingControlSetVolumeResponse                `xml:"SetVolumeResponse"`
	SetRelativeVolume        *RenderingControlSetRelativeVolumeResponse        `xml:"SetRelativeVolumeResponse"`
	GetVolumeDB              *RenderingControlGetVolumeDBResponse              `xml:"GetVolumeDBResponse"`
	SetVolumeDB              *RenderingControlSetVolumeDBResponse              `xml:"SetVolumeDBResponse"`
	GetVolumeDBRange         *RenderingControlGetVolumeDBRangeResponse         `xml:"GetVolumeDBRangeResponse"`
	GetBass                  *RenderingControlGetBassResponse                  `xml:"GetBassResponse"`
	SetBass                  *RenderingControlSetBassResponse                  `xml:"SetBassResponse"`
	GetTreble                *RenderingControlGetTrebleResponse                `xml:"GetTrebleResponse"`
	SetTreble                *RenderingControlSetTrebleResponse                `xml:"SetTrebleResponse"`
	GetEQ                    *RenderingControlGetEQResponse                    `xml:"GetEQResponse"`
	SetEQ                    *RenderingControlSetEQResponse                    `xml:"SetEQResponse"`
	GetLoudness              *RenderingControlGetLoudnessResponse              `xml:"GetLoudnessResponse"`
	SetLoudness              *RenderingControlSetLoudnessResponse              `xml:"SetLoudnessResponse"`
	GetSupportsOutputFixed   *RenderingControlGetSupportsOutputFixedResponse   `xml:"GetSupportsOutputFixedResponse"`
	GetOutputFixed           *RenderingControlGetOutputFixedResponse           `xml:"GetOutputFixedResponse"`
	SetOutputFixed           *RenderingControlSetOutputFixedResponse           `xml:"SetOutputFixedResponse"`
	GetHeadphoneConnected    *RenderingControlGetHeadphoneConnectedResponse    `xml:"GetHeadphoneConnectedResponse"`
	RampToVolume             *RenderingControlRampToVolumeResponse             `xml:"RampToVolumeResponse"`
	RestoreVolumePriorToRamp *RenderingControlRestoreVolumePriorToRampResponse `xml:"RestoreVolumePriorToRampResponse"`
	SetChannelMap            *RenderingControlSetChannelMapResponse            `xml:"SetChannelMapResponse"`
	SetRoomCalibrationX      *RenderingControlSetRoomCalibrationXResponse      `xml:"SetRoomCalibrationXResponse"`
	GetRoomCalibrationStatus *RenderingControlGetRoomCalibrationStatusResponse `xml:"GetRoomCalibrationStatusResponse"`
	SetRoomCalibrationStatus *RenderingControlSetRoomCalibrationStatusResponse `xml:"SetRoomCalibrationStatusResponse"`
}

func (s *RenderingControlService) _RenderingControlExec(soapAction string, httpClient *http.Client, envelope *RenderingControlEnvelope) (*RenderingControlEnvelopeResponse, error) {
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
	var envelopeResponse RenderingControlEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type RenderingControlGetMuteArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	// Allowed Value: SpeakerOnly
	Channel string `xml:"Channel"`
}
type RenderingControlGetMuteResponse struct {
	CurrentMute bool `xml:"CurrentMute"`
}

func (s *RenderingControlService) GetMute(httpClient *http.Client, args *RenderingControlGetMuteArgs) (*RenderingControlGetMuteResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetMute", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetMute: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetMute == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetMute, nil
}

type RenderingControlSetMuteArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	// Allowed Value: SpeakerOnly
	Channel     string `xml:"Channel"`
	DesiredMute bool   `xml:"DesiredMute"`
}
type RenderingControlSetMuteResponse struct {
}

func (s *RenderingControlService) SetMute(httpClient *http.Client, args *RenderingControlSetMuteArgs) (*RenderingControlSetMuteResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetMute", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetMute: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetMute == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetMute, nil
}

type RenderingControlResetBasicEQArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type RenderingControlResetBasicEQResponse struct {
	Bass        int16  `xml:"Bass"`
	Treble      int16  `xml:"Treble"`
	Loudness    bool   `xml:"Loudness"`
	LeftVolume  uint16 `xml:"LeftVolume"`
	RightVolume uint16 `xml:"RightVolume"`
}

func (s *RenderingControlService) ResetBasicEQ(httpClient *http.Client, args *RenderingControlResetBasicEQArgs) (*RenderingControlResetBasicEQResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#ResetBasicEQ", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{ResetBasicEQ: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ResetBasicEQ == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ResetBasicEQ, nil
}

type RenderingControlResetExtEQArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	EQType       string `xml:"EQType"`
}
type RenderingControlResetExtEQResponse struct {
}

func (s *RenderingControlService) ResetExtEQ(httpClient *http.Client, args *RenderingControlResetExtEQArgs) (*RenderingControlResetExtEQResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#ResetExtEQ", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{ResetExtEQ: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ResetExtEQ == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ResetExtEQ, nil
}

type RenderingControlGetVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel string `xml:"Channel"`
}
type RenderingControlGetVolumeResponse struct {
	CurrentVolume uint16 `xml:"CurrentVolume"`
}

func (s *RenderingControlService) GetVolume(httpClient *http.Client, args *RenderingControlGetVolumeArgs) (*RenderingControlGetVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetVolume", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetVolume, nil
}

type RenderingControlSetVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel string `xml:"Channel"`
	// Allowed Range: 0 -> 100 step: 1
	DesiredVolume uint16 `xml:"DesiredVolume"`
}
type RenderingControlSetVolumeResponse struct {
}

func (s *RenderingControlService) SetVolume(httpClient *http.Client, args *RenderingControlSetVolumeArgs) (*RenderingControlSetVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetVolume", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetVolume, nil
}

type RenderingControlSetRelativeVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel    string `xml:"Channel"`
	Adjustment int32  `xml:"Adjustment"`
}
type RenderingControlSetRelativeVolumeResponse struct {
	NewVolume uint16 `xml:"NewVolume"`
}

func (s *RenderingControlService) SetRelativeVolume(httpClient *http.Client, args *RenderingControlSetRelativeVolumeArgs) (*RenderingControlSetRelativeVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetRelativeVolume", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetRelativeVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetRelativeVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetRelativeVolume, nil
}

type RenderingControlGetVolumeDBArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel string `xml:"Channel"`
}
type RenderingControlGetVolumeDBResponse struct {
	CurrentVolume int16 `xml:"CurrentVolume"`
}

func (s *RenderingControlService) GetVolumeDB(httpClient *http.Client, args *RenderingControlGetVolumeDBArgs) (*RenderingControlGetVolumeDBResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetVolumeDB", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetVolumeDB: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetVolumeDB == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetVolumeDB, nil
}

type RenderingControlSetVolumeDBArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel       string `xml:"Channel"`
	DesiredVolume int16  `xml:"DesiredVolume"`
}
type RenderingControlSetVolumeDBResponse struct {
}

func (s *RenderingControlService) SetVolumeDB(httpClient *http.Client, args *RenderingControlSetVolumeDBArgs) (*RenderingControlSetVolumeDBResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetVolumeDB", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetVolumeDB: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetVolumeDB == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetVolumeDB, nil
}

type RenderingControlGetVolumeDBRangeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel string `xml:"Channel"`
}
type RenderingControlGetVolumeDBRangeResponse struct {
	MinValue int16 `xml:"MinValue"`
	MaxValue int16 `xml:"MaxValue"`
}

func (s *RenderingControlService) GetVolumeDBRange(httpClient *http.Client, args *RenderingControlGetVolumeDBRangeArgs) (*RenderingControlGetVolumeDBRangeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetVolumeDBRange", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetVolumeDBRange: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetVolumeDBRange == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetVolumeDBRange, nil
}

type RenderingControlGetBassArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type RenderingControlGetBassResponse struct {
	CurrentBass int16 `xml:"CurrentBass"`
}

func (s *RenderingControlService) GetBass(httpClient *http.Client, args *RenderingControlGetBassArgs) (*RenderingControlGetBassResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetBass", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetBass: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetBass == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetBass, nil
}

type RenderingControlSetBassArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Range: -10 -> 10 step: 1
	DesiredBass int16 `xml:"DesiredBass"`
}
type RenderingControlSetBassResponse struct {
}

func (s *RenderingControlService) SetBass(httpClient *http.Client, args *RenderingControlSetBassArgs) (*RenderingControlSetBassResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetBass", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetBass: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetBass == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetBass, nil
}

type RenderingControlGetTrebleArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type RenderingControlGetTrebleResponse struct {
	CurrentTreble int16 `xml:"CurrentTreble"`
}

func (s *RenderingControlService) GetTreble(httpClient *http.Client, args *RenderingControlGetTrebleArgs) (*RenderingControlGetTrebleResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetTreble", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetTreble: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTreble == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTreble, nil
}

type RenderingControlSetTrebleArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Range: -10 -> 10 step: 1
	DesiredTreble int16 `xml:"DesiredTreble"`
}
type RenderingControlSetTrebleResponse struct {
}

func (s *RenderingControlService) SetTreble(httpClient *http.Client, args *RenderingControlSetTrebleArgs) (*RenderingControlSetTrebleResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetTreble", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetTreble: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetTreble == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetTreble, nil
}

type RenderingControlGetEQArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	EQType       string `xml:"EQType"`
}
type RenderingControlGetEQResponse struct {
	CurrentValue int16 `xml:"CurrentValue"`
}

func (s *RenderingControlService) GetEQ(httpClient *http.Client, args *RenderingControlGetEQArgs) (*RenderingControlGetEQResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetEQ", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetEQ: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetEQ == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetEQ, nil
}

type RenderingControlSetEQArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	EQType       string `xml:"EQType"`
	DesiredValue int16  `xml:"DesiredValue"`
}
type RenderingControlSetEQResponse struct {
}

func (s *RenderingControlService) SetEQ(httpClient *http.Client, args *RenderingControlSetEQArgs) (*RenderingControlSetEQResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetEQ", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetEQ: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetEQ == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetEQ, nil
}

type RenderingControlGetLoudnessArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel string `xml:"Channel"`
}
type RenderingControlGetLoudnessResponse struct {
	CurrentLoudness bool `xml:"CurrentLoudness"`
}

func (s *RenderingControlService) GetLoudness(httpClient *http.Client, args *RenderingControlGetLoudnessArgs) (*RenderingControlGetLoudnessResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetLoudness", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetLoudness: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetLoudness == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetLoudness, nil
}

type RenderingControlSetLoudnessArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel         string `xml:"Channel"`
	DesiredLoudness bool   `xml:"DesiredLoudness"`
}
type RenderingControlSetLoudnessResponse struct {
}

func (s *RenderingControlService) SetLoudness(httpClient *http.Client, args *RenderingControlSetLoudnessArgs) (*RenderingControlSetLoudnessResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetLoudness", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetLoudness: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetLoudness == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetLoudness, nil
}

type RenderingControlGetSupportsOutputFixedArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type RenderingControlGetSupportsOutputFixedResponse struct {
	CurrentSupportsFixed bool `xml:"CurrentSupportsFixed"`
}

func (s *RenderingControlService) GetSupportsOutputFixed(httpClient *http.Client, args *RenderingControlGetSupportsOutputFixedArgs) (*RenderingControlGetSupportsOutputFixedResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetSupportsOutputFixed", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetSupportsOutputFixed: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSupportsOutputFixed == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetSupportsOutputFixed, nil
}

type RenderingControlGetOutputFixedArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type RenderingControlGetOutputFixedResponse struct {
	CurrentFixed bool `xml:"CurrentFixed"`
}

func (s *RenderingControlService) GetOutputFixed(httpClient *http.Client, args *RenderingControlGetOutputFixedArgs) (*RenderingControlGetOutputFixedResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetOutputFixed", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetOutputFixed: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetOutputFixed == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetOutputFixed, nil
}

type RenderingControlSetOutputFixedArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	DesiredFixed bool   `xml:"DesiredFixed"`
}
type RenderingControlSetOutputFixedResponse struct {
}

func (s *RenderingControlService) SetOutputFixed(httpClient *http.Client, args *RenderingControlSetOutputFixedArgs) (*RenderingControlSetOutputFixedResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetOutputFixed", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetOutputFixed: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetOutputFixed == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetOutputFixed, nil
}

type RenderingControlGetHeadphoneConnectedArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type RenderingControlGetHeadphoneConnectedResponse struct {
	CurrentHeadphoneConnected bool `xml:"CurrentHeadphoneConnected"`
}

func (s *RenderingControlService) GetHeadphoneConnected(httpClient *http.Client, args *RenderingControlGetHeadphoneConnectedArgs) (*RenderingControlGetHeadphoneConnectedResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetHeadphoneConnected", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetHeadphoneConnected: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetHeadphoneConnected == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetHeadphoneConnected, nil
}

type RenderingControlRampToVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel string `xml:"Channel"`
	// Allowed Value: SLEEP_TIMER_RAMP_TYPE
	// Allowed Value: ALARM_RAMP_TYPE
	// Allowed Value: AUTOPLAY_RAMP_TYPE
	RampType string `xml:"RampType"`
	// Allowed Range: 0 -> 100 step: 1
	DesiredVolume    uint16 `xml:"DesiredVolume"`
	ResetVolumeAfter bool   `xml:"ResetVolumeAfter"`
	ProgramURI       string `xml:"ProgramURI"`
}
type RenderingControlRampToVolumeResponse struct {
	RampTime uint32 `xml:"RampTime"`
}

func (s *RenderingControlService) RampToVolume(httpClient *http.Client, args *RenderingControlRampToVolumeArgs) (*RenderingControlRampToVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#RampToVolume", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{RampToVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RampToVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RampToVolume, nil
}

type RenderingControlRestoreVolumePriorToRampArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: Master
	// Allowed Value: LF
	// Allowed Value: RF
	Channel string `xml:"Channel"`
}
type RenderingControlRestoreVolumePriorToRampResponse struct {
}

func (s *RenderingControlService) RestoreVolumePriorToRamp(httpClient *http.Client, args *RenderingControlRestoreVolumePriorToRampArgs) (*RenderingControlRestoreVolumePriorToRampResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#RestoreVolumePriorToRamp", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{RestoreVolumePriorToRamp: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RestoreVolumePriorToRamp == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RestoreVolumePriorToRamp, nil
}

type RenderingControlSetChannelMapArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	ChannelMap   string `xml:"ChannelMap"`
}
type RenderingControlSetChannelMapResponse struct {
}

func (s *RenderingControlService) SetChannelMap(httpClient *http.Client, args *RenderingControlSetChannelMapArgs) (*RenderingControlSetChannelMapResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetChannelMap", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetChannelMap: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetChannelMap == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetChannelMap, nil
}

type RenderingControlSetRoomCalibrationXArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	InstanceID      uint32 `xml:"InstanceID"`
	CalibrationID   string `xml:"CalibrationID"`
	Coefficients    string `xml:"Coefficients"`
	CalibrationMode string `xml:"CalibrationMode"`
}
type RenderingControlSetRoomCalibrationXResponse struct {
}

func (s *RenderingControlService) SetRoomCalibrationX(httpClient *http.Client, args *RenderingControlSetRoomCalibrationXArgs) (*RenderingControlSetRoomCalibrationXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetRoomCalibrationX", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetRoomCalibrationX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetRoomCalibrationX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetRoomCalibrationX, nil
}

type RenderingControlGetRoomCalibrationStatusArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type RenderingControlGetRoomCalibrationStatusResponse struct {
	RoomCalibrationEnabled   bool `xml:"RoomCalibrationEnabled"`
	RoomCalibrationAvailable bool `xml:"RoomCalibrationAvailable"`
}

func (s *RenderingControlService) GetRoomCalibrationStatus(httpClient *http.Client, args *RenderingControlGetRoomCalibrationStatusArgs) (*RenderingControlGetRoomCalibrationStatusResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#GetRoomCalibrationStatus", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{GetRoomCalibrationStatus: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetRoomCalibrationStatus == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetRoomCalibrationStatus, nil
}

type RenderingControlSetRoomCalibrationStatusArgs struct {
	XMLNameSpace           string `xml:"xmlns:u,attr"`
	InstanceID             uint32 `xml:"InstanceID"`
	RoomCalibrationEnabled bool   `xml:"RoomCalibrationEnabled"`
}
type RenderingControlSetRoomCalibrationStatusResponse struct {
}

func (s *RenderingControlService) SetRoomCalibrationStatus(httpClient *http.Client, args *RenderingControlSetRoomCalibrationStatusArgs) (*RenderingControlSetRoomCalibrationStatusResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:RenderingControl:1"
	r, err := s._RenderingControlExec("urn:schemas-upnp-org:service:RenderingControl:1#SetRoomCalibrationStatus", httpClient,
		&RenderingControlEnvelope{
			Body:          RenderingControlBody{SetRoomCalibrationStatus: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetRoomCalibrationStatus == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetRoomCalibrationStatus, nil
}
func (s *RenderingControlService) RenderingControlSubscribe(callback url.URL) error {
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
type RenderingControlLastChange string
type RenderingControlUpnpEvent struct {
	XMLName      xml.Name                   `xml:"propertyset"`
	XMLNameSpace string                     `xml:"xmlns:e,attr"`
	Properties   []RenderingControlProperty `xml:"property"`
}
type RenderingControlProperty struct {
	XMLName    xml.Name                    `xml:"property"`
	LastChange *RenderingControlLastChange `xml:"LastChange"`
}

func RenderingControlDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt RenderingControlUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.LastChange != nil:
			zp.EventCallback(*prop.LastChange)
		}
	}
}
