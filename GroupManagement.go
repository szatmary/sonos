package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// State Variables
type GroupManagement_GroupCoordinatorIsLocal bool
type GroupManagement_LocalGroupUUID string
type GroupManagement_VirtualLineInGroupID string
type GroupManagement_ResetVolumeAfter bool
type GroupManagement_VolumeAVTransportURI string

type GroupManagementService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
	// State
	GroupCoordinatorIsLocal *GroupManagement_GroupCoordinatorIsLocal
	LocalGroupUUID          *GroupManagement_LocalGroupUUID
	VirtualLineInGroupID    *GroupManagement_VirtualLineInGroupID
	ResetVolumeAfter        *GroupManagement_ResetVolumeAfter
	VolumeAVTransportURI    *GroupManagement_VolumeAVTransportURI
}

func NewGroupManagementService(deviceUrl *url.URL) *GroupManagementService {
	c, _ := url.Parse("/GroupManagement/Control")
	e, _ := url.Parse("/GroupManagement/Event")
	return &GroupManagementService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *GroupManagementService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *GroupManagementService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type GroupManagementEnvelope struct {
	XMLName       xml.Name            `xml:"s:Envelope"`
	XMLNameSpace  string              `xml:"xmlns:s,attr"`
	EncodingStyle string              `xml:"s:encodingStyle,attr"`
	Body          GroupManagementBody `xml:"s:Body"`
}
type GroupManagementBody struct {
	XMLName                    xml.Name                                       `xml:"s:Body"`
	AddMember                  *GroupManagementAddMemberArgs                  `xml:"u:AddMember,omitempty"`
	RemoveMember               *GroupManagementRemoveMemberArgs               `xml:"u:RemoveMember,omitempty"`
	ReportTrackBufferingResult *GroupManagementReportTrackBufferingResultArgs `xml:"u:ReportTrackBufferingResult,omitempty"`
	SetSourceAreaIds           *GroupManagementSetSourceAreaIdsArgs           `xml:"u:SetSourceAreaIds,omitempty"`
}
type GroupManagementEnvelopeResponse struct {
	XMLName       xml.Name                    `xml:"Envelope"`
	XMLNameSpace  string                      `xml:"xmlns:s,attr"`
	EncodingStyle string                      `xml:"encodingStyle,attr"`
	Body          GroupManagementBodyResponse `xml:"Body"`
}
type GroupManagementBodyResponse struct {
	XMLName                    xml.Name                                           `xml:"Body"`
	AddMember                  *GroupManagementAddMemberResponse                  `xml:"AddMemberResponse"`
	RemoveMember               *GroupManagementRemoveMemberResponse               `xml:"RemoveMemberResponse"`
	ReportTrackBufferingResult *GroupManagementReportTrackBufferingResultResponse `xml:"ReportTrackBufferingResultResponse"`
	SetSourceAreaIds           *GroupManagementSetSourceAreaIdsResponse           `xml:"SetSourceAreaIdsResponse"`
}

func (s *GroupManagementService) _GroupManagementExec(soapAction string, httpClient *http.Client, envelope *GroupManagementEnvelope) (*GroupManagementEnvelopeResponse, error) {
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
	var envelopeResponse GroupManagementEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type GroupManagementAddMemberArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	MemberID     string `xml:"MemberID"`
	BootSeq      uint32 `xml:"BootSeq"`
}
type GroupManagementAddMemberResponse struct {
	CurrentTransportSettings string `xml:"CurrentTransportSettings"`
	CurrentURI               string `xml:"CurrentURI"`
	GroupUUIDJoined          string `xml:"GroupUUIDJoined"`
	ResetVolumeAfter         bool   `xml:"ResetVolumeAfter"`
	VolumeAVTransportURI     string `xml:"VolumeAVTransportURI"`
}

func (s *GroupManagementService) AddMember(httpClient *http.Client, args *GroupManagementAddMemberArgs) (*GroupManagementAddMemberResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupManagement:1"
	r, err := s._GroupManagementExec("urn:schemas-upnp-org:service:GroupManagement:1#AddMember", httpClient,
		&GroupManagementEnvelope{
			Body:          GroupManagementBody{AddMember: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddMember == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddMember, nil
}

type GroupManagementRemoveMemberArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	MemberID     string `xml:"MemberID"`
}
type GroupManagementRemoveMemberResponse struct {
}

func (s *GroupManagementService) RemoveMember(httpClient *http.Client, args *GroupManagementRemoveMemberArgs) (*GroupManagementRemoveMemberResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupManagement:1"
	r, err := s._GroupManagementExec("urn:schemas-upnp-org:service:GroupManagement:1#RemoveMember", httpClient,
		&GroupManagementEnvelope{
			Body:          GroupManagementBody{RemoveMember: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveMember == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveMember, nil
}

type GroupManagementReportTrackBufferingResultArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	MemberID     string `xml:"MemberID"`
	ResultCode   int32  `xml:"ResultCode"`
}
type GroupManagementReportTrackBufferingResultResponse struct {
}

func (s *GroupManagementService) ReportTrackBufferingResult(httpClient *http.Client, args *GroupManagementReportTrackBufferingResultArgs) (*GroupManagementReportTrackBufferingResultResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupManagement:1"
	r, err := s._GroupManagementExec("urn:schemas-upnp-org:service:GroupManagement:1#ReportTrackBufferingResult", httpClient,
		&GroupManagementEnvelope{
			Body:          GroupManagementBody{ReportTrackBufferingResult: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReportTrackBufferingResult == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReportTrackBufferingResult, nil
}

type GroupManagementSetSourceAreaIdsArgs struct {
	XMLNameSpace         string `xml:"xmlns:u,attr"`
	DesiredSourceAreaIds string `xml:"DesiredSourceAreaIds"`
}
type GroupManagementSetSourceAreaIdsResponse struct {
}

func (s *GroupManagementService) SetSourceAreaIds(httpClient *http.Client, args *GroupManagementSetSourceAreaIdsArgs) (*GroupManagementSetSourceAreaIdsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:GroupManagement:1"
	r, err := s._GroupManagementExec("urn:schemas-upnp-org:service:GroupManagement:1#SetSourceAreaIds", httpClient,
		&GroupManagementEnvelope{
			Body:          GroupManagementBody{SetSourceAreaIds: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetSourceAreaIds == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetSourceAreaIds, nil
}

type GroupManagementUpnpEvent struct {
	XMLName      xml.Name                  `xml:"propertyset"`
	XMLNameSpace string                    `xml:"xmlns:e,attr"`
	Properties   []GroupManagementProperty `xml:"property"`
}
type GroupManagementProperty struct {
	XMLName                 xml.Name                                 `xml:"property"`
	GroupCoordinatorIsLocal *GroupManagement_GroupCoordinatorIsLocal `xml:"GroupCoordinatorIsLocal"`
	LocalGroupUUID          *GroupManagement_LocalGroupUUID          `xml:"LocalGroupUUID"`
	VirtualLineInGroupID    *GroupManagement_VirtualLineInGroupID    `xml:"VirtualLineInGroupID"`
	ResetVolumeAfter        *GroupManagement_ResetVolumeAfter        `xml:"ResetVolumeAfter"`
	VolumeAVTransportURI    *GroupManagement_VolumeAVTransportURI    `xml:"VolumeAVTransportURI"`
}

func (zp *GroupManagementService) ParseEvent(body []byte) []interface{} {
	var evt GroupManagementUpnpEvent
	var events []interface{}
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return events
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.GroupCoordinatorIsLocal != nil:
			zp.GroupCoordinatorIsLocal = prop.GroupCoordinatorIsLocal
			events = append(events, *prop.GroupCoordinatorIsLocal)
		case prop.LocalGroupUUID != nil:
			zp.LocalGroupUUID = prop.LocalGroupUUID
			events = append(events, *prop.LocalGroupUUID)
		case prop.VirtualLineInGroupID != nil:
			zp.VirtualLineInGroupID = prop.VirtualLineInGroupID
			events = append(events, *prop.VirtualLineInGroupID)
		case prop.ResetVolumeAfter != nil:
			zp.ResetVolumeAfter = prop.ResetVolumeAfter
			events = append(events, *prop.ResetVolumeAfter)
		case prop.VolumeAVTransportURI != nil:
			zp.VolumeAVTransportURI = prop.VolumeAVTransportURI
			events = append(events, *prop.VolumeAVTransportURI)
		}
	}
	return events
}
