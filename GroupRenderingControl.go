package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GroupRenderingControlService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewGroupRenderingControlService(deviceUrl *url.URL) *GroupRenderingControlService {
	c, _ := url.Parse("/MediaRenderer/GroupRenderingControl/Control")
	e, _ := url.Parse("/MediaRenderer/GroupRenderingControl/Event")
	return &GroupRenderingControlService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *GroupRenderingControlService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *GroupRenderingControlService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type GroupRenderingControlEnvelope struct {
	XMLName       xml.Name                  `xml:"s:Envelope"`
	XMLNameSpace  string                    `xml:"xmlns:s,attr"`
	EncodingStyle string                    `xml:"s:encodingStyle,attr"`
	Body          GroupRenderingControlBody `xml:"s:Body"`
}
type GroupRenderingControlBody struct {
	XMLName                xml.Name                                         `xml:"s:Body"`
	GetGroupMute           *GroupRenderingControlGetGroupMuteArgs           `xml:"u:GetGroupMute,omitempty"`
	SetGroupMute           *GroupRenderingControlSetGroupMuteArgs           `xml:"u:SetGroupMute,omitempty"`
	GetGroupVolume         *GroupRenderingControlGetGroupVolumeArgs         `xml:"u:GetGroupVolume,omitempty"`
	SetGroupVolume         *GroupRenderingControlSetGroupVolumeArgs         `xml:"u:SetGroupVolume,omitempty"`
	SetRelativeGroupVolume *GroupRenderingControlSetRelativeGroupVolumeArgs `xml:"u:SetRelativeGroupVolume,omitempty"`
	SnapshotGroupVolume    *GroupRenderingControlSnapshotGroupVolumeArgs    `xml:"u:SnapshotGroupVolume,omitempty"`
}
type GroupRenderingControlEnvelopeResponse struct {
	XMLName       xml.Name                          `xml:"Envelope"`
	XMLNameSpace  string                            `xml:"xmlns:s,attr"`
	EncodingStyle string                            `xml:"encodingStyle,attr"`
	Body          GroupRenderingControlBodyResponse `xml:"Body"`
}
type GroupRenderingControlBodyResponse struct {
	XMLName                xml.Name                                             `xml:"Body"`
	GetGroupMute           *GroupRenderingControlGetGroupMuteResponse           `xml:"GetGroupMuteResponse"`
	SetGroupMute           *GroupRenderingControlSetGroupMuteResponse           `xml:"SetGroupMuteResponse"`
	GetGroupVolume         *GroupRenderingControlGetGroupVolumeResponse         `xml:"GetGroupVolumeResponse"`
	SetGroupVolume         *GroupRenderingControlSetGroupVolumeResponse         `xml:"SetGroupVolumeResponse"`
	SetRelativeGroupVolume *GroupRenderingControlSetRelativeGroupVolumeResponse `xml:"SetRelativeGroupVolumeResponse"`
	SnapshotGroupVolume    *GroupRenderingControlSnapshotGroupVolumeResponse    `xml:"SnapshotGroupVolumeResponse"`
}

func (s *GroupRenderingControlService) _GroupRenderingControlExec(soapAction string, httpClient *http.Client, envelope *GroupRenderingControlEnvelope) (*GroupRenderingControlEnvelopeResponse, error) {
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
	var envelopeResponse GroupRenderingControlEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type GroupRenderingControlGetGroupMuteArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type GroupRenderingControlGetGroupMuteResponse struct {
	CurrentMute bool `xml:"CurrentMute"`
}

func (s *GroupRenderingControlService) GetGroupMute(httpClient *http.Client, args *GroupRenderingControlGetGroupMuteArgs) (*GroupRenderingControlGetGroupMuteResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupRenderingControl:1"
	r, err := s._GroupRenderingControlExec("urn:schemas-upnp-org:service:GroupRenderingControl:1#GetGroupMute", httpClient,
		&GroupRenderingControlEnvelope{
			Body:          GroupRenderingControlBody{GetGroupMute: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetGroupMute == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetGroupMute, nil
}

type GroupRenderingControlSetGroupMuteArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	DesiredMute  bool   `xml:"DesiredMute"`
}
type GroupRenderingControlSetGroupMuteResponse struct {
}

func (s *GroupRenderingControlService) SetGroupMute(httpClient *http.Client, args *GroupRenderingControlSetGroupMuteArgs) (*GroupRenderingControlSetGroupMuteResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupRenderingControl:1"
	r, err := s._GroupRenderingControlExec("urn:schemas-upnp-org:service:GroupRenderingControl:1#SetGroupMute", httpClient,
		&GroupRenderingControlEnvelope{
			Body:          GroupRenderingControlBody{SetGroupMute: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetGroupMute == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetGroupMute, nil
}

type GroupRenderingControlGetGroupVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type GroupRenderingControlGetGroupVolumeResponse struct {
	CurrentVolume uint16 `xml:"CurrentVolume"`
}

func (s *GroupRenderingControlService) GetGroupVolume(httpClient *http.Client, args *GroupRenderingControlGetGroupVolumeArgs) (*GroupRenderingControlGetGroupVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupRenderingControl:1"
	r, err := s._GroupRenderingControlExec("urn:schemas-upnp-org:service:GroupRenderingControl:1#GetGroupVolume", httpClient,
		&GroupRenderingControlEnvelope{
			Body:          GroupRenderingControlBody{GetGroupVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetGroupVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetGroupVolume, nil
}

type GroupRenderingControlSetGroupVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Range: 0 -> 100 step: 1
	DesiredVolume uint16 `xml:"DesiredVolume"`
}
type GroupRenderingControlSetGroupVolumeResponse struct {
}

func (s *GroupRenderingControlService) SetGroupVolume(httpClient *http.Client, args *GroupRenderingControlSetGroupVolumeArgs) (*GroupRenderingControlSetGroupVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupRenderingControl:1"
	r, err := s._GroupRenderingControlExec("urn:schemas-upnp-org:service:GroupRenderingControl:1#SetGroupVolume", httpClient,
		&GroupRenderingControlEnvelope{
			Body:          GroupRenderingControlBody{SetGroupVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetGroupVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetGroupVolume, nil
}

type GroupRenderingControlSetRelativeGroupVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	Adjustment   int32  `xml:"Adjustment"`
}
type GroupRenderingControlSetRelativeGroupVolumeResponse struct {
	NewVolume uint16 `xml:"NewVolume"`
}

func (s *GroupRenderingControlService) SetRelativeGroupVolume(httpClient *http.Client, args *GroupRenderingControlSetRelativeGroupVolumeArgs) (*GroupRenderingControlSetRelativeGroupVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupRenderingControl:1"
	r, err := s._GroupRenderingControlExec("urn:schemas-upnp-org:service:GroupRenderingControl:1#SetRelativeGroupVolume", httpClient,
		&GroupRenderingControlEnvelope{
			Body:          GroupRenderingControlBody{SetRelativeGroupVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetRelativeGroupVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetRelativeGroupVolume, nil
}

type GroupRenderingControlSnapshotGroupVolumeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type GroupRenderingControlSnapshotGroupVolumeResponse struct {
}

func (s *GroupRenderingControlService) SnapshotGroupVolume(httpClient *http.Client, args *GroupRenderingControlSnapshotGroupVolumeArgs) (*GroupRenderingControlSnapshotGroupVolumeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupRenderingControl:1"
	r, err := s._GroupRenderingControlExec("urn:schemas-upnp-org:service:GroupRenderingControl:1#SnapshotGroupVolume", httpClient,
		&GroupRenderingControlEnvelope{
			Body:          GroupRenderingControlBody{SnapshotGroupVolume: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SnapshotGroupVolume == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SnapshotGroupVolume, nil
}

// Events
type GroupRenderingControlGroupMute bool
type GroupRenderingControlGroupVolume uint16
type GroupRenderingControlGroupVolumeChangeable bool
type GroupRenderingControlUpnpEvent struct {
	XMLName      xml.Name                        `xml:"propertyset"`
	XMLNameSpace string                          `xml:"xmlns:e,attr"`
	Properties   []GroupRenderingControlProperty `xml:"property"`
}
type GroupRenderingControlProperty struct {
	XMLName               xml.Name                                    `xml:"property"`
	GroupMute             *GroupRenderingControlGroupMute             `xml:"GroupMute"`
	GroupVolume           *GroupRenderingControlGroupVolume           `xml:"GroupVolume"`
	GroupVolumeChangeable *GroupRenderingControlGroupVolumeChangeable `xml:"GroupVolumeChangeable"`
}

func GroupRenderingControlDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt GroupRenderingControlUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.GroupMute != nil:
			zp.EventCallback(*prop.GroupMute)
		case prop.GroupVolume != nil:
			zp.EventCallback(*prop.GroupVolume)
		case prop.GroupVolumeChangeable != nil:
			zp.EventCallback(*prop.GroupVolumeChangeable)
		}
	}
}
