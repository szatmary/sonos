package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type VirtualLineInService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewVirtualLineInService(deviceUrl *url.URL) *VirtualLineInService {
	c, _ := url.Parse("/MediaRenderer/VirtualLineIn/Control")
	e, _ := url.Parse("/MediaRenderer/VirtualLineIn/Event")
	return &VirtualLineInService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *VirtualLineInService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *VirtualLineInService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type VirtualLineInEnvelope struct {
	XMLName       xml.Name          `xml:"s:Envelope"`
	XMLNameSpace  string            `xml:"xmlns:s,attr"`
	EncodingStyle string            `xml:"s:encodingStyle,attr"`
	Body          VirtualLineInBody `xml:"s:Body"`
}
type VirtualLineInBody struct {
	XMLName           xml.Name                            `xml:"s:Body"`
	StartTransmission *VirtualLineInStartTransmissionArgs `xml:"u:StartTransmission,omitempty"`
	StopTransmission  *VirtualLineInStopTransmissionArgs  `xml:"u:StopTransmission,omitempty"`
	Play              *VirtualLineInPlayArgs              `xml:"u:Play,omitempty"`
	Pause             *VirtualLineInPauseArgs             `xml:"u:Pause,omitempty"`
	Next              *VirtualLineInNextArgs              `xml:"u:Next,omitempty"`
	Previous          *VirtualLineInPreviousArgs          `xml:"u:Previous,omitempty"`
	Stop              *VirtualLineInStopArgs              `xml:"u:Stop,omitempty"`
	SetVolume         *VirtualLineInSetVolumeArgs         `xml:"u:SetVolume,omitempty"`
}
type VirtualLineInEnvelopeResponse struct {
	XMLName       xml.Name                  `xml:"Envelope"`
	XMLNameSpace  string                    `xml:"xmlns:s,attr"`
	EncodingStyle string                    `xml:"encodingStyle,attr"`
	Body          VirtualLineInBodyResponse `xml:"Body"`
}
type VirtualLineInBodyResponse struct {
	XMLName           xml.Name                                `xml:"Body"`
	StartTransmission *VirtualLineInStartTransmissionResponse `xml:"StartTransmissionResponse"`
	StopTransmission  *VirtualLineInStopTransmissionResponse  `xml:"StopTransmissionResponse"`
	Play              *VirtualLineInPlayResponse              `xml:"PlayResponse"`
	Pause             *VirtualLineInPauseResponse             `xml:"PauseResponse"`
	Next              *VirtualLineInNextResponse              `xml:"NextResponse"`
	Previous          *VirtualLineInPreviousResponse          `xml:"PreviousResponse"`
	Stop              *VirtualLineInStopResponse              `xml:"StopResponse"`
	SetVolume         *VirtualLineInSetVolumeResponse         `xml:"SetVolumeResponse"`
}

func (s *VirtualLineInService) _VirtualLineInExec(soapAction string, httpClient *http.Client, envelope *VirtualLineInEnvelope) (*VirtualLineInEnvelopeResponse, error) {
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
	var envelopeResponse VirtualLineInEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type VirtualLineInStartTransmissionArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	CoordinatorID string `xml:"CoordinatorID"`
}
type VirtualLineInStartTransmissionResponse struct {
	CurrentTransportSettings string `xml:"CurrentTransportSettings"`
}

func (s *VirtualLineInService) StartTransmission(httpClient *http.Client, args *VirtualLineInStartTransmissionArgs) (*VirtualLineInStartTransmissionResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#StartTransmission", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{StartTransmission: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.StartTransmission == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.StartTransmission, nil
}

type VirtualLineInStopTransmissionArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	CoordinatorID string `xml:"CoordinatorID"`
}
type VirtualLineInStopTransmissionResponse struct {
}

func (s *VirtualLineInService) StopTransmission(httpClient *http.Client, args *VirtualLineInStopTransmissionArgs) (*VirtualLineInStopTransmissionResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#StopTransmission", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{StopTransmission: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.StopTransmission == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.StopTransmission, nil
}

type VirtualLineInPlayArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	Speed        string `xml:"Speed"`
}
type VirtualLineInPlayResponse struct {
}

func (s *VirtualLineInService) Play(httpClient *http.Client, args *VirtualLineInPlayArgs) (*VirtualLineInPlayResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#Play", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{Play: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Play == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Play, nil
}

type VirtualLineInPauseArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type VirtualLineInPauseResponse struct {
}

func (s *VirtualLineInService) Pause(httpClient *http.Client, args *VirtualLineInPauseArgs) (*VirtualLineInPauseResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#Pause", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{Pause: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Pause == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Pause, nil
}

type VirtualLineInNextArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type VirtualLineInNextResponse struct {
}

func (s *VirtualLineInService) Next(httpClient *http.Client, args *VirtualLineInNextArgs) (*VirtualLineInNextResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#Next", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{Next: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Next == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Next, nil
}

type VirtualLineInPreviousArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type VirtualLineInPreviousResponse struct {
}

func (s *VirtualLineInService) Previous(httpClient *http.Client, args *VirtualLineInPreviousArgs) (*VirtualLineInPreviousResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#Previous", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{Previous: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Previous == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Previous, nil
}

type VirtualLineInStopArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type VirtualLineInStopResponse struct {
}

func (s *VirtualLineInService) Stop(httpClient *http.Client, args *VirtualLineInStopArgs) (*VirtualLineInStopResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#Stop", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{Stop: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Stop == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Stop, nil
}

type VirtualLineInSetVolumeArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	DesiredVolume uint16 `xml:"DesiredVolume"`
}
type VirtualLineInSetVolumeResponse struct {
}

func (s *VirtualLineInService) SetVolume(httpClient *http.Client, args *VirtualLineInSetVolumeArgs) (*VirtualLineInSetVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:VirtualLineIn:1"
	r, err := s._VirtualLineInExec("urn:schemas-upnp-org:service:VirtualLineIn:1#SetVolume", httpClient,
		&VirtualLineInEnvelope{
			Body:          VirtualLineInBody{SetVolume: args},
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

type VirtualLineInUpnpEvent struct {
	XMLName      xml.Name                `xml:"propertyset"`
	XMLNameSpace string                  `xml:"xmlns:e,attr"`
	Properties   []VirtualLineInProperty `xml:"property"`
}
type VirtualLineInProperty struct {
	XMLName    xml.Name `xml:"property"`
	LastChange *string  `xml:"LastChange"`
}

func VirtualLineInDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt VirtualLineInUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.LastChange != nil:
			dispatchVirtualLineInLastChange(*prop.LastChange) // string
		}
	}
}
