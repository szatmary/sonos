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

type AVTransportService struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewAVTransportService(deviceUrl *url.URL) *AVTransportService {
	c, _ := url.Parse("/MediaRenderer/AVTransport/Control")
	e, _ := url.Parse("/MediaRenderer/AVTransport/Event")
	return &AVTransportService{
		ControlEndpoint: deviceUrl.ResolveReference(c),
		EventEndpoint:   deviceUrl.ResolveReference(e),
	}
}

type AVTransportEnvelope struct {
	XMLName       xml.Name        `xml:"s:Envelope"`
	XMLNameSpace  string          `xml:"xmlns:s,attr"`
	EncodingStyle string          `xml:"s:encodingStyle,attr"`
	Body          AVTransportBody `xml:"s:Body"`
}
type AVTransportBody struct {
	XMLName                            xml.Name                                           `xml:"s:Body"`
	SetAVTransportURI                  *AVTransportSetAVTransportURIArgs                  `xml:"u:SetAVTransportURI,omitempty"`
	SetNextAVTransportURI              *AVTransportSetNextAVTransportURIArgs              `xml:"u:SetNextAVTransportURI,omitempty"`
	AddURIToQueue                      *AVTransportAddURIToQueueArgs                      `xml:"u:AddURIToQueue,omitempty"`
	AddMultipleURIsToQueue             *AVTransportAddMultipleURIsToQueueArgs             `xml:"u:AddMultipleURIsToQueue,omitempty"`
	ReorderTracksInQueue               *AVTransportReorderTracksInQueueArgs               `xml:"u:ReorderTracksInQueue,omitempty"`
	RemoveTrackFromQueue               *AVTransportRemoveTrackFromQueueArgs               `xml:"u:RemoveTrackFromQueue,omitempty"`
	RemoveTrackRangeFromQueue          *AVTransportRemoveTrackRangeFromQueueArgs          `xml:"u:RemoveTrackRangeFromQueue,omitempty"`
	RemoveAllTracksFromQueue           *AVTransportRemoveAllTracksFromQueueArgs           `xml:"u:RemoveAllTracksFromQueue,omitempty"`
	SaveQueue                          *AVTransportSaveQueueArgs                          `xml:"u:SaveQueue,omitempty"`
	BackupQueue                        *AVTransportBackupQueueArgs                        `xml:"u:BackupQueue,omitempty"`
	CreateSavedQueue                   *AVTransportCreateSavedQueueArgs                   `xml:"u:CreateSavedQueue,omitempty"`
	AddURIToSavedQueue                 *AVTransportAddURIToSavedQueueArgs                 `xml:"u:AddURIToSavedQueue,omitempty"`
	ReorderTracksInSavedQueue          *AVTransportReorderTracksInSavedQueueArgs          `xml:"u:ReorderTracksInSavedQueue,omitempty"`
	GetMediaInfo                       *AVTransportGetMediaInfoArgs                       `xml:"u:GetMediaInfo,omitempty"`
	GetTransportInfo                   *AVTransportGetTransportInfoArgs                   `xml:"u:GetTransportInfo,omitempty"`
	GetPositionInfo                    *AVTransportGetPositionInfoArgs                    `xml:"u:GetPositionInfo,omitempty"`
	GetDeviceCapabilities              *AVTransportGetDeviceCapabilitiesArgs              `xml:"u:GetDeviceCapabilities,omitempty"`
	GetTransportSettings               *AVTransportGetTransportSettingsArgs               `xml:"u:GetTransportSettings,omitempty"`
	GetCrossfadeMode                   *AVTransportGetCrossfadeModeArgs                   `xml:"u:GetCrossfadeMode,omitempty"`
	Stop                               *AVTransportStopArgs                               `xml:"u:Stop,omitempty"`
	Play                               *AVTransportPlayArgs                               `xml:"u:Play,omitempty"`
	Pause                              *AVTransportPauseArgs                              `xml:"u:Pause,omitempty"`
	Seek                               *AVTransportSeekArgs                               `xml:"u:Seek,omitempty"`
	Next                               *AVTransportNextArgs                               `xml:"u:Next,omitempty"`
	Previous                           *AVTransportPreviousArgs                           `xml:"u:Previous,omitempty"`
	SetPlayMode                        *AVTransportSetPlayModeArgs                        `xml:"u:SetPlayMode,omitempty"`
	SetCrossfadeMode                   *AVTransportSetCrossfadeModeArgs                   `xml:"u:SetCrossfadeMode,omitempty"`
	NotifyDeletedURI                   *AVTransportNotifyDeletedURIArgs                   `xml:"u:NotifyDeletedURI,omitempty"`
	GetCurrentTransportActions         *AVTransportGetCurrentTransportActionsArgs         `xml:"u:GetCurrentTransportActions,omitempty"`
	BecomeCoordinatorOfStandaloneGroup *AVTransportBecomeCoordinatorOfStandaloneGroupArgs `xml:"u:BecomeCoordinatorOfStandaloneGroup,omitempty"`
	DelegateGroupCoordinationTo        *AVTransportDelegateGroupCoordinationToArgs        `xml:"u:DelegateGroupCoordinationTo,omitempty"`
	BecomeGroupCoordinator             *AVTransportBecomeGroupCoordinatorArgs             `xml:"u:BecomeGroupCoordinator,omitempty"`
	BecomeGroupCoordinatorAndSource    *AVTransportBecomeGroupCoordinatorAndSourceArgs    `xml:"u:BecomeGroupCoordinatorAndSource,omitempty"`
	ChangeCoordinator                  *AVTransportChangeCoordinatorArgs                  `xml:"u:ChangeCoordinator,omitempty"`
	ChangeTransportSettings            *AVTransportChangeTransportSettingsArgs            `xml:"u:ChangeTransportSettings,omitempty"`
	ConfigureSleepTimer                *AVTransportConfigureSleepTimerArgs                `xml:"u:ConfigureSleepTimer,omitempty"`
	GetRemainingSleepTimerDuration     *AVTransportGetRemainingSleepTimerDurationArgs     `xml:"u:GetRemainingSleepTimerDuration,omitempty"`
	RunAlarm                           *AVTransportRunAlarmArgs                           `xml:"u:RunAlarm,omitempty"`
	StartAutoplay                      *AVTransportStartAutoplayArgs                      `xml:"u:StartAutoplay,omitempty"`
	GetRunningAlarmProperties          *AVTransportGetRunningAlarmPropertiesArgs          `xml:"u:GetRunningAlarmProperties,omitempty"`
	SnoozeAlarm                        *AVTransportSnoozeAlarmArgs                        `xml:"u:SnoozeAlarm,omitempty"`
	EndDirectControlSession            *AVTransportEndDirectControlSessionArgs            `xml:"u:EndDirectControlSession,omitempty"`
}
type AVTransportEnvelopeResponse struct {
	XMLName       xml.Name                `xml:"Envelope"`
	XMLNameSpace  string                  `xml:"xmlns:s,attr"`
	EncodingStyle string                  `xml:"encodingStyle,attr"`
	Body          AVTransportBodyResponse `xml:"Body"`
}
type AVTransportBodyResponse struct {
	XMLName                            xml.Name                                               `xml:"Body"`
	SetAVTransportURI                  *AVTransportSetAVTransportURIResponse                  `xml:"SetAVTransportURIResponse"`
	SetNextAVTransportURI              *AVTransportSetNextAVTransportURIResponse              `xml:"SetNextAVTransportURIResponse"`
	AddURIToQueue                      *AVTransportAddURIToQueueResponse                      `xml:"AddURIToQueueResponse"`
	AddMultipleURIsToQueue             *AVTransportAddMultipleURIsToQueueResponse             `xml:"AddMultipleURIsToQueueResponse"`
	ReorderTracksInQueue               *AVTransportReorderTracksInQueueResponse               `xml:"ReorderTracksInQueueResponse"`
	RemoveTrackFromQueue               *AVTransportRemoveTrackFromQueueResponse               `xml:"RemoveTrackFromQueueResponse"`
	RemoveTrackRangeFromQueue          *AVTransportRemoveTrackRangeFromQueueResponse          `xml:"RemoveTrackRangeFromQueueResponse"`
	RemoveAllTracksFromQueue           *AVTransportRemoveAllTracksFromQueueResponse           `xml:"RemoveAllTracksFromQueueResponse"`
	SaveQueue                          *AVTransportSaveQueueResponse                          `xml:"SaveQueueResponse"`
	BackupQueue                        *AVTransportBackupQueueResponse                        `xml:"BackupQueueResponse"`
	CreateSavedQueue                   *AVTransportCreateSavedQueueResponse                   `xml:"CreateSavedQueueResponse"`
	AddURIToSavedQueue                 *AVTransportAddURIToSavedQueueResponse                 `xml:"AddURIToSavedQueueResponse"`
	ReorderTracksInSavedQueue          *AVTransportReorderTracksInSavedQueueResponse          `xml:"ReorderTracksInSavedQueueResponse"`
	GetMediaInfo                       *AVTransportGetMediaInfoResponse                       `xml:"GetMediaInfoResponse"`
	GetTransportInfo                   *AVTransportGetTransportInfoResponse                   `xml:"GetTransportInfoResponse"`
	GetPositionInfo                    *AVTransportGetPositionInfoResponse                    `xml:"GetPositionInfoResponse"`
	GetDeviceCapabilities              *AVTransportGetDeviceCapabilitiesResponse              `xml:"GetDeviceCapabilitiesResponse"`
	GetTransportSettings               *AVTransportGetTransportSettingsResponse               `xml:"GetTransportSettingsResponse"`
	GetCrossfadeMode                   *AVTransportGetCrossfadeModeResponse                   `xml:"GetCrossfadeModeResponse"`
	Stop                               *AVTransportStopResponse                               `xml:"StopResponse"`
	Play                               *AVTransportPlayResponse                               `xml:"PlayResponse"`
	Pause                              *AVTransportPauseResponse                              `xml:"PauseResponse"`
	Seek                               *AVTransportSeekResponse                               `xml:"SeekResponse"`
	Next                               *AVTransportNextResponse                               `xml:"NextResponse"`
	Previous                           *AVTransportPreviousResponse                           `xml:"PreviousResponse"`
	SetPlayMode                        *AVTransportSetPlayModeResponse                        `xml:"SetPlayModeResponse"`
	SetCrossfadeMode                   *AVTransportSetCrossfadeModeResponse                   `xml:"SetCrossfadeModeResponse"`
	NotifyDeletedURI                   *AVTransportNotifyDeletedURIResponse                   `xml:"NotifyDeletedURIResponse"`
	GetCurrentTransportActions         *AVTransportGetCurrentTransportActionsResponse         `xml:"GetCurrentTransportActionsResponse"`
	BecomeCoordinatorOfStandaloneGroup *AVTransportBecomeCoordinatorOfStandaloneGroupResponse `xml:"BecomeCoordinatorOfStandaloneGroupResponse"`
	DelegateGroupCoordinationTo        *AVTransportDelegateGroupCoordinationToResponse        `xml:"DelegateGroupCoordinationToResponse"`
	BecomeGroupCoordinator             *AVTransportBecomeGroupCoordinatorResponse             `xml:"BecomeGroupCoordinatorResponse"`
	BecomeGroupCoordinatorAndSource    *AVTransportBecomeGroupCoordinatorAndSourceResponse    `xml:"BecomeGroupCoordinatorAndSourceResponse"`
	ChangeCoordinator                  *AVTransportChangeCoordinatorResponse                  `xml:"ChangeCoordinatorResponse"`
	ChangeTransportSettings            *AVTransportChangeTransportSettingsResponse            `xml:"ChangeTransportSettingsResponse"`
	ConfigureSleepTimer                *AVTransportConfigureSleepTimerResponse                `xml:"ConfigureSleepTimerResponse"`
	GetRemainingSleepTimerDuration     *AVTransportGetRemainingSleepTimerDurationResponse     `xml:"GetRemainingSleepTimerDurationResponse"`
	RunAlarm                           *AVTransportRunAlarmResponse                           `xml:"RunAlarmResponse"`
	StartAutoplay                      *AVTransportStartAutoplayResponse                      `xml:"StartAutoplayResponse"`
	GetRunningAlarmProperties          *AVTransportGetRunningAlarmPropertiesResponse          `xml:"GetRunningAlarmPropertiesResponse"`
	SnoozeAlarm                        *AVTransportSnoozeAlarmResponse                        `xml:"SnoozeAlarmResponse"`
	EndDirectControlSession            *AVTransportEndDirectControlSessionResponse            `xml:"EndDirectControlSessionResponse"`
}

func (s *AVTransportService) _AVTransportExec(soapAction string, httpClient *http.Client, envelope *AVTransportEnvelope) (*AVTransportEnvelopeResponse, error) {
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
	var envelopeResponse AVTransportEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type AVTransportSetAVTransportURIArgs struct {
	XMLNameSpace       string `xml:"xmlns:u,attr"`
	InstanceID         uint32 `xml:"InstanceID"`
	CurrentURI         string `xml:"CurrentURI"`
	CurrentURIMetaData string `xml:"CurrentURIMetaData"`
}
type AVTransportSetAVTransportURIResponse struct {
}

func (s *AVTransportService) SetAVTransportURI(httpClient *http.Client, args *AVTransportSetAVTransportURIArgs) (*AVTransportSetAVTransportURIResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#SetAVTransportURI", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{SetAVTransportURI: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAVTransportURI == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetAVTransportURI, nil
}

type AVTransportSetNextAVTransportURIArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	InstanceID      uint32 `xml:"InstanceID"`
	NextURI         string `xml:"NextURI"`
	NextURIMetaData string `xml:"NextURIMetaData"`
}
type AVTransportSetNextAVTransportURIResponse struct {
}

func (s *AVTransportService) SetNextAVTransportURI(httpClient *http.Client, args *AVTransportSetNextAVTransportURIArgs) (*AVTransportSetNextAVTransportURIResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#SetNextAVTransportURI", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{SetNextAVTransportURI: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetNextAVTransportURI == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetNextAVTransportURI, nil
}

type AVTransportAddURIToQueueArgs struct {
	XMLNameSpace                    string `xml:"xmlns:u,attr"`
	InstanceID                      uint32 `xml:"InstanceID"`
	EnqueuedURI                     string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData             string `xml:"EnqueuedURIMetaData"`
	DesiredFirstTrackNumberEnqueued uint32 `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                   bool   `xml:"EnqueueAsNext"`
}
type AVTransportAddURIToQueueResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
}

func (s *AVTransportService) AddURIToQueue(httpClient *http.Client, args *AVTransportAddURIToQueueArgs) (*AVTransportAddURIToQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#AddURIToQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{AddURIToQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddURIToQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddURIToQueue, nil
}

type AVTransportAddMultipleURIsToQueueArgs struct {
	XMLNameSpace                    string `xml:"xmlns:u,attr"`
	InstanceID                      uint32 `xml:"InstanceID"`
	UpdateID                        uint32 `xml:"UpdateID"`
	NumberOfURIs                    uint32 `xml:"NumberOfURIs"`
	EnqueuedURIs                    string `xml:"EnqueuedURIs"`
	EnqueuedURIsMetaData            string `xml:"EnqueuedURIsMetaData"`
	ContainerURI                    string `xml:"ContainerURI"`
	ContainerMetaData               string `xml:"ContainerMetaData"`
	DesiredFirstTrackNumberEnqueued uint32 `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                   bool   `xml:"EnqueueAsNext"`
}
type AVTransportAddMultipleURIsToQueueResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
	NewUpdateID              uint32 `xml:"NewUpdateID"`
}

func (s *AVTransportService) AddMultipleURIsToQueue(httpClient *http.Client, args *AVTransportAddMultipleURIsToQueueArgs) (*AVTransportAddMultipleURIsToQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#AddMultipleURIsToQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{AddMultipleURIsToQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddMultipleURIsToQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddMultipleURIsToQueue, nil
}

type AVTransportReorderTracksInQueueArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
	InstanceID     uint32 `xml:"InstanceID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
	InsertBefore   uint32 `xml:"InsertBefore"`
	UpdateID       uint32 `xml:"UpdateID"`
}
type AVTransportReorderTracksInQueueResponse struct {
}

func (s *AVTransportService) ReorderTracksInQueue(httpClient *http.Client, args *AVTransportReorderTracksInQueueArgs) (*AVTransportReorderTracksInQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#ReorderTracksInQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{ReorderTracksInQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReorderTracksInQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReorderTracksInQueue, nil
}

type AVTransportRemoveTrackFromQueueArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	ObjectID     string `xml:"ObjectID"`
	UpdateID     uint32 `xml:"UpdateID"`
}
type AVTransportRemoveTrackFromQueueResponse struct {
}

func (s *AVTransportService) RemoveTrackFromQueue(httpClient *http.Client, args *AVTransportRemoveTrackFromQueueArgs) (*AVTransportRemoveTrackFromQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#RemoveTrackFromQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{RemoveTrackFromQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveTrackFromQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveTrackFromQueue, nil
}

type AVTransportRemoveTrackRangeFromQueueArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
	InstanceID     uint32 `xml:"InstanceID"`
	UpdateID       uint32 `xml:"UpdateID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
}
type AVTransportRemoveTrackRangeFromQueueResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *AVTransportService) RemoveTrackRangeFromQueue(httpClient *http.Client, args *AVTransportRemoveTrackRangeFromQueueArgs) (*AVTransportRemoveTrackRangeFromQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#RemoveTrackRangeFromQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{RemoveTrackRangeFromQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveTrackRangeFromQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveTrackRangeFromQueue, nil
}

type AVTransportRemoveAllTracksFromQueueArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportRemoveAllTracksFromQueueResponse struct {
}

func (s *AVTransportService) RemoveAllTracksFromQueue(httpClient *http.Client, args *AVTransportRemoveAllTracksFromQueueArgs) (*AVTransportRemoveAllTracksFromQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#RemoveAllTracksFromQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{RemoveAllTracksFromQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveAllTracksFromQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveAllTracksFromQueue, nil
}

type AVTransportSaveQueueArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	Title        string `xml:"Title"`
	ObjectID     string `xml:"ObjectID"`
}
type AVTransportSaveQueueResponse struct {
	AssignedObjectID string `xml:"AssignedObjectID"`
}

func (s *AVTransportService) SaveQueue(httpClient *http.Client, args *AVTransportSaveQueueArgs) (*AVTransportSaveQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#SaveQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{SaveQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SaveQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SaveQueue, nil
}

type AVTransportBackupQueueArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportBackupQueueResponse struct {
}

func (s *AVTransportService) BackupQueue(httpClient *http.Client, args *AVTransportBackupQueueArgs) (*AVTransportBackupQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#BackupQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{BackupQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BackupQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.BackupQueue, nil
}

type AVTransportCreateSavedQueueArgs struct {
	XMLNameSpace        string `xml:"xmlns:u,attr"`
	InstanceID          uint32 `xml:"InstanceID"`
	Title               string `xml:"Title"`
	EnqueuedURI         string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData string `xml:"EnqueuedURIMetaData"`
}
type AVTransportCreateSavedQueueResponse struct {
	NumTracksAdded   uint32 `xml:"NumTracksAdded"`
	NewQueueLength   uint32 `xml:"NewQueueLength"`
	AssignedObjectID string `xml:"AssignedObjectID"`
	NewUpdateID      uint32 `xml:"NewUpdateID"`
}

func (s *AVTransportService) CreateSavedQueue(httpClient *http.Client, args *AVTransportCreateSavedQueueArgs) (*AVTransportCreateSavedQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#CreateSavedQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{CreateSavedQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateSavedQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.CreateSavedQueue, nil
}

type AVTransportAddURIToSavedQueueArgs struct {
	XMLNameSpace        string `xml:"xmlns:u,attr"`
	InstanceID          uint32 `xml:"InstanceID"`
	ObjectID            string `xml:"ObjectID"`
	UpdateID            uint32 `xml:"UpdateID"`
	EnqueuedURI         string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData string `xml:"EnqueuedURIMetaData"`
	AddAtIndex          uint32 `xml:"AddAtIndex"`
}
type AVTransportAddURIToSavedQueueResponse struct {
	NumTracksAdded uint32 `xml:"NumTracksAdded"`
	NewQueueLength uint32 `xml:"NewQueueLength"`
	NewUpdateID    uint32 `xml:"NewUpdateID"`
}

func (s *AVTransportService) AddURIToSavedQueue(httpClient *http.Client, args *AVTransportAddURIToSavedQueueArgs) (*AVTransportAddURIToSavedQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#AddURIToSavedQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{AddURIToSavedQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddURIToSavedQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddURIToSavedQueue, nil
}

type AVTransportReorderTracksInSavedQueueArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	InstanceID      uint32 `xml:"InstanceID"`
	ObjectID        string `xml:"ObjectID"`
	UpdateID        uint32 `xml:"UpdateID"`
	TrackList       string `xml:"TrackList"`
	NewPositionList string `xml:"NewPositionList"`
}
type AVTransportReorderTracksInSavedQueueResponse struct {
	QueueLengthChange int32  `xml:"QueueLengthChange"`
	NewQueueLength    uint32 `xml:"NewQueueLength"`
	NewUpdateID       uint32 `xml:"NewUpdateID"`
}

func (s *AVTransportService) ReorderTracksInSavedQueue(httpClient *http.Client, args *AVTransportReorderTracksInSavedQueueArgs) (*AVTransportReorderTracksInSavedQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#ReorderTracksInSavedQueue", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{ReorderTracksInSavedQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReorderTracksInSavedQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReorderTracksInSavedQueue, nil
}

type AVTransportGetMediaInfoArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetMediaInfoResponse struct {
	NrTracks           uint32 `xml:"NrTracks"`
	MediaDuration      string `xml:"MediaDuration"`
	CurrentURI         string `xml:"CurrentURI"`
	CurrentURIMetaData string `xml:"CurrentURIMetaData"`
	NextURI            string `xml:"NextURI"`
	NextURIMetaData    string `xml:"NextURIMetaData"`
	PlayMedium         string `xml:"PlayMedium"`
	RecordMedium       string `xml:"RecordMedium"`
	WriteStatus        string `xml:"WriteStatus"`
}

func (s *AVTransportService) GetMediaInfo(httpClient *http.Client, args *AVTransportGetMediaInfoArgs) (*AVTransportGetMediaInfoResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetMediaInfo", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetMediaInfo: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetMediaInfo == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetMediaInfo, nil
}

type AVTransportGetTransportInfoArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetTransportInfoResponse struct {
	CurrentTransportState  string `xml:"CurrentTransportState"`
	CurrentTransportStatus string `xml:"CurrentTransportStatus"`
	CurrentSpeed           string `xml:"CurrentSpeed"`
}

func (s *AVTransportService) GetTransportInfo(httpClient *http.Client, args *AVTransportGetTransportInfoArgs) (*AVTransportGetTransportInfoResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetTransportInfo", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetTransportInfo: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTransportInfo == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTransportInfo, nil
}

type AVTransportGetPositionInfoArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetPositionInfoResponse struct {
	Track         uint32 `xml:"Track"`
	TrackDuration string `xml:"TrackDuration"`
	TrackMetaData string `xml:"TrackMetaData"`
	TrackURI      string `xml:"TrackURI"`
	RelTime       string `xml:"RelTime"`
	AbsTime       string `xml:"AbsTime"`
	RelCount      int32  `xml:"RelCount"`
	AbsCount      int32  `xml:"AbsCount"`
}

func (s *AVTransportService) GetPositionInfo(httpClient *http.Client, args *AVTransportGetPositionInfoArgs) (*AVTransportGetPositionInfoResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetPositionInfo", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetPositionInfo: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetPositionInfo == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetPositionInfo, nil
}

type AVTransportGetDeviceCapabilitiesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetDeviceCapabilitiesResponse struct {
	PlayMedia       string `xml:"PlayMedia"`
	RecMedia        string `xml:"RecMedia"`
	RecQualityModes string `xml:"RecQualityModes"`
}

func (s *AVTransportService) GetDeviceCapabilities(httpClient *http.Client, args *AVTransportGetDeviceCapabilitiesArgs) (*AVTransportGetDeviceCapabilitiesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetDeviceCapabilities", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetDeviceCapabilities: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetDeviceCapabilities == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetDeviceCapabilities, nil
}

type AVTransportGetTransportSettingsArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetTransportSettingsResponse struct {
	PlayMode       string `xml:"PlayMode"`
	RecQualityMode string `xml:"RecQualityMode"`
}

func (s *AVTransportService) GetTransportSettings(httpClient *http.Client, args *AVTransportGetTransportSettingsArgs) (*AVTransportGetTransportSettingsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetTransportSettings", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetTransportSettings: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTransportSettings == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTransportSettings, nil
}

type AVTransportGetCrossfadeModeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetCrossfadeModeResponse struct {
	CrossfadeMode bool `xml:"CrossfadeMode"`
}

func (s *AVTransportService) GetCrossfadeMode(httpClient *http.Client, args *AVTransportGetCrossfadeModeArgs) (*AVTransportGetCrossfadeModeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetCrossfadeMode", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetCrossfadeMode: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCrossfadeMode == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetCrossfadeMode, nil
}

type AVTransportStopArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportStopResponse struct {
}

func (s *AVTransportService) Stop(httpClient *http.Client, args *AVTransportStopArgs) (*AVTransportStopResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#Stop", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{Stop: args},
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

type AVTransportPlayArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: 1
	Speed string `xml:"Speed"`
}
type AVTransportPlayResponse struct {
}

func (s *AVTransportService) Play(httpClient *http.Client, args *AVTransportPlayArgs) (*AVTransportPlayResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#Play", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{Play: args},
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

type AVTransportPauseArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportPauseResponse struct {
}

func (s *AVTransportService) Pause(httpClient *http.Client, args *AVTransportPauseArgs) (*AVTransportPauseResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#Pause", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{Pause: args},
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

type AVTransportSeekArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: TRACK_NR
	// Allowed Value: REL_TIME
	// Allowed Value: TIME_DELTA
	Unit   string `xml:"Unit"`
	Target string `xml:"Target"`
}
type AVTransportSeekResponse struct {
}

func (s *AVTransportService) Seek(httpClient *http.Client, args *AVTransportSeekArgs) (*AVTransportSeekResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#Seek", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{Seek: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Seek == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Seek, nil
}

type AVTransportNextArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportNextResponse struct {
}

func (s *AVTransportService) Next(httpClient *http.Client, args *AVTransportNextArgs) (*AVTransportNextResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#Next", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{Next: args},
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

type AVTransportPreviousArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportPreviousResponse struct {
}

func (s *AVTransportService) Previous(httpClient *http.Client, args *AVTransportPreviousArgs) (*AVTransportPreviousResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#Previous", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{Previous: args},
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

type AVTransportSetPlayModeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	// Allowed Value: NORMAL
	// Allowed Value: REPEAT_ALL
	// Allowed Value: REPEAT_ONE
	// Allowed Value: SHUFFLE_NOREPEAT
	// Allowed Value: SHUFFLE
	// Allowed Value: SHUFFLE_REPEAT_ONE
	NewPlayMode string `xml:"NewPlayMode"`
}
type AVTransportSetPlayModeResponse struct {
}

func (s *AVTransportService) SetPlayMode(httpClient *http.Client, args *AVTransportSetPlayModeArgs) (*AVTransportSetPlayModeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#SetPlayMode", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{SetPlayMode: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetPlayMode == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetPlayMode, nil
}

type AVTransportSetCrossfadeModeArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	CrossfadeMode bool   `xml:"CrossfadeMode"`
}
type AVTransportSetCrossfadeModeResponse struct {
}

func (s *AVTransportService) SetCrossfadeMode(httpClient *http.Client, args *AVTransportSetCrossfadeModeArgs) (*AVTransportSetCrossfadeModeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#SetCrossfadeMode", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{SetCrossfadeMode: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetCrossfadeMode == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetCrossfadeMode, nil
}

type AVTransportNotifyDeletedURIArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	DeletedURI   string `xml:"DeletedURI"`
}
type AVTransportNotifyDeletedURIResponse struct {
}

func (s *AVTransportService) NotifyDeletedURI(httpClient *http.Client, args *AVTransportNotifyDeletedURIArgs) (*AVTransportNotifyDeletedURIResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#NotifyDeletedURI", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{NotifyDeletedURI: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.NotifyDeletedURI == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.NotifyDeletedURI, nil
}

type AVTransportGetCurrentTransportActionsArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetCurrentTransportActionsResponse struct {
	Actions string `xml:"Actions"`
}

func (s *AVTransportService) GetCurrentTransportActions(httpClient *http.Client, args *AVTransportGetCurrentTransportActionsArgs) (*AVTransportGetCurrentTransportActionsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetCurrentTransportActions", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetCurrentTransportActions: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCurrentTransportActions == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetCurrentTransportActions, nil
}

type AVTransportBecomeCoordinatorOfStandaloneGroupArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportBecomeCoordinatorOfStandaloneGroupResponse struct {
	DelegatedGroupCoordinatorID string `xml:"DelegatedGroupCoordinatorID"`
	NewGroupID                  string `xml:"NewGroupID"`
}

func (s *AVTransportService) BecomeCoordinatorOfStandaloneGroup(httpClient *http.Client, args *AVTransportBecomeCoordinatorOfStandaloneGroupArgs) (*AVTransportBecomeCoordinatorOfStandaloneGroupResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#BecomeCoordinatorOfStandaloneGroup", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{BecomeCoordinatorOfStandaloneGroup: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BecomeCoordinatorOfStandaloneGroup == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.BecomeCoordinatorOfStandaloneGroup, nil
}

type AVTransportDelegateGroupCoordinationToArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
	InstanceID     uint32 `xml:"InstanceID"`
	NewCoordinator string `xml:"NewCoordinator"`
	RejoinGroup    bool   `xml:"RejoinGroup"`
}
type AVTransportDelegateGroupCoordinationToResponse struct {
}

func (s *AVTransportService) DelegateGroupCoordinationTo(httpClient *http.Client, args *AVTransportDelegateGroupCoordinationToArgs) (*AVTransportDelegateGroupCoordinationToResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#DelegateGroupCoordinationTo", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{DelegateGroupCoordinationTo: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DelegateGroupCoordinationTo == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.DelegateGroupCoordinationTo, nil
}

type AVTransportBecomeGroupCoordinatorArgs struct {
	XMLNameSpace          string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	CurrentCoordinator    string `xml:"CurrentCoordinator"`
	CurrentGroupID        string `xml:"CurrentGroupID"`
	OtherMembers          string `xml:"OtherMembers"`
	TransportSettings     string `xml:"TransportSettings"`
	CurrentURI            string `xml:"CurrentURI"`
	CurrentURIMetaData    string `xml:"CurrentURIMetaData"`
	SleepTimerState       string `xml:"SleepTimerState"`
	AlarmState            string `xml:"AlarmState"`
	StreamRestartState    string `xml:"StreamRestartState"`
	CurrentQueueTrackList string `xml:"CurrentQueueTrackList"`
	CurrentVLIState       string `xml:"CurrentVLIState"`
}
type AVTransportBecomeGroupCoordinatorResponse struct {
}

func (s *AVTransportService) BecomeGroupCoordinator(httpClient *http.Client, args *AVTransportBecomeGroupCoordinatorArgs) (*AVTransportBecomeGroupCoordinatorResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#BecomeGroupCoordinator", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{BecomeGroupCoordinator: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BecomeGroupCoordinator == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.BecomeGroupCoordinator, nil
}

type AVTransportBecomeGroupCoordinatorAndSourceArgs struct {
	XMLNameSpace          string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	CurrentCoordinator    string `xml:"CurrentCoordinator"`
	CurrentGroupID        string `xml:"CurrentGroupID"`
	OtherMembers          string `xml:"OtherMembers"`
	CurrentURI            string `xml:"CurrentURI"`
	CurrentURIMetaData    string `xml:"CurrentURIMetaData"`
	SleepTimerState       string `xml:"SleepTimerState"`
	AlarmState            string `xml:"AlarmState"`
	StreamRestartState    string `xml:"StreamRestartState"`
	CurrentAVTTrackList   string `xml:"CurrentAVTTrackList"`
	CurrentQueueTrackList string `xml:"CurrentQueueTrackList"`
	CurrentSourceState    string `xml:"CurrentSourceState"`
	ResumePlayback        bool   `xml:"ResumePlayback"`
}
type AVTransportBecomeGroupCoordinatorAndSourceResponse struct {
}

func (s *AVTransportService) BecomeGroupCoordinatorAndSource(httpClient *http.Client, args *AVTransportBecomeGroupCoordinatorAndSourceArgs) (*AVTransportBecomeGroupCoordinatorAndSourceResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#BecomeGroupCoordinatorAndSource", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{BecomeGroupCoordinatorAndSource: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BecomeGroupCoordinatorAndSource == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.BecomeGroupCoordinatorAndSource, nil
}

type AVTransportChangeCoordinatorArgs struct {
	XMLNameSpace          string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	CurrentCoordinator    string `xml:"CurrentCoordinator"`
	NewCoordinator        string `xml:"NewCoordinator"`
	NewTransportSettings  string `xml:"NewTransportSettings"`
	CurrentAVTransportURI string `xml:"CurrentAVTransportURI"`
}
type AVTransportChangeCoordinatorResponse struct {
}

func (s *AVTransportService) ChangeCoordinator(httpClient *http.Client, args *AVTransportChangeCoordinatorArgs) (*AVTransportChangeCoordinatorResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#ChangeCoordinator", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{ChangeCoordinator: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ChangeCoordinator == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ChangeCoordinator, nil
}

type AVTransportChangeTransportSettingsArgs struct {
	XMLNameSpace          string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	NewTransportSettings  string `xml:"NewTransportSettings"`
	CurrentAVTransportURI string `xml:"CurrentAVTransportURI"`
}
type AVTransportChangeTransportSettingsResponse struct {
}

func (s *AVTransportService) ChangeTransportSettings(httpClient *http.Client, args *AVTransportChangeTransportSettingsArgs) (*AVTransportChangeTransportSettingsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#ChangeTransportSettings", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{ChangeTransportSettings: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ChangeTransportSettings == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ChangeTransportSettings, nil
}

type AVTransportConfigureSleepTimerArgs struct {
	XMLNameSpace          string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	NewSleepTimerDuration string `xml:"NewSleepTimerDuration"`
}
type AVTransportConfigureSleepTimerResponse struct {
}

func (s *AVTransportService) ConfigureSleepTimer(httpClient *http.Client, args *AVTransportConfigureSleepTimerArgs) (*AVTransportConfigureSleepTimerResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#ConfigureSleepTimer", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{ConfigureSleepTimer: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ConfigureSleepTimer == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ConfigureSleepTimer, nil
}

type AVTransportGetRemainingSleepTimerDurationArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetRemainingSleepTimerDurationResponse struct {
	RemainingSleepTimerDuration string `xml:"RemainingSleepTimerDuration"`
	CurrentSleepTimerGeneration uint32 `xml:"CurrentSleepTimerGeneration"`
}

func (s *AVTransportService) GetRemainingSleepTimerDuration(httpClient *http.Client, args *AVTransportGetRemainingSleepTimerDurationArgs) (*AVTransportGetRemainingSleepTimerDurationResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetRemainingSleepTimerDuration", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetRemainingSleepTimerDuration: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetRemainingSleepTimerDuration == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetRemainingSleepTimerDuration, nil
}

type AVTransportRunAlarmArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	InstanceID      uint32 `xml:"InstanceID"`
	AlarmID         uint32 `xml:"AlarmID"`
	LoggedStartTime string `xml:"LoggedStartTime"`
	Duration        string `xml:"Duration"`
	ProgramURI      string `xml:"ProgramURI"`
	ProgramMetaData string `xml:"ProgramMetaData"`
	// Allowed Value: NORMAL
	// Allowed Value: REPEAT_ALL
	// Allowed Value: REPEAT_ONE
	// Allowed Value: SHUFFLE_NOREPEAT
	// Allowed Value: SHUFFLE
	// Allowed Value: SHUFFLE_REPEAT_ONE
	PlayMode           string `xml:"PlayMode"`
	Volume             uint16 `xml:"Volume"`
	IncludeLinkedZones bool   `xml:"IncludeLinkedZones"`
}
type AVTransportRunAlarmResponse struct {
}

func (s *AVTransportService) RunAlarm(httpClient *http.Client, args *AVTransportRunAlarmArgs) (*AVTransportRunAlarmResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#RunAlarm", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{RunAlarm: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RunAlarm == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RunAlarm, nil
}

type AVTransportStartAutoplayArgs struct {
	XMLNameSpace       string `xml:"xmlns:u,attr"`
	InstanceID         uint32 `xml:"InstanceID"`
	ProgramURI         string `xml:"ProgramURI"`
	ProgramMetaData    string `xml:"ProgramMetaData"`
	Volume             uint16 `xml:"Volume"`
	IncludeLinkedZones bool   `xml:"IncludeLinkedZones"`
	ResetVolumeAfter   bool   `xml:"ResetVolumeAfter"`
}
type AVTransportStartAutoplayResponse struct {
}

func (s *AVTransportService) StartAutoplay(httpClient *http.Client, args *AVTransportStartAutoplayArgs) (*AVTransportStartAutoplayResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#StartAutoplay", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{StartAutoplay: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.StartAutoplay == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.StartAutoplay, nil
}

type AVTransportGetRunningAlarmPropertiesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportGetRunningAlarmPropertiesResponse struct {
	AlarmID         uint32 `xml:"AlarmID"`
	GroupID         string `xml:"GroupID"`
	LoggedStartTime string `xml:"LoggedStartTime"`
}

func (s *AVTransportService) GetRunningAlarmProperties(httpClient *http.Client, args *AVTransportGetRunningAlarmPropertiesArgs) (*AVTransportGetRunningAlarmPropertiesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#GetRunningAlarmProperties", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{GetRunningAlarmProperties: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetRunningAlarmProperties == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetRunningAlarmProperties, nil
}

type AVTransportSnoozeAlarmArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
	Duration     string `xml:"Duration"`
}
type AVTransportSnoozeAlarmResponse struct {
}

func (s *AVTransportService) SnoozeAlarm(httpClient *http.Client, args *AVTransportSnoozeAlarmArgs) (*AVTransportSnoozeAlarmResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#SnoozeAlarm", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{SnoozeAlarm: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SnoozeAlarm == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SnoozeAlarm, nil
}

type AVTransportEndDirectControlSessionArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	InstanceID   uint32 `xml:"InstanceID"`
}
type AVTransportEndDirectControlSessionResponse struct {
}

func (s *AVTransportService) EndDirectControlSession(httpClient *http.Client, args *AVTransportEndDirectControlSessionArgs) (*AVTransportEndDirectControlSessionResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AVTransport:1"
	r, err := s._AVTransportExec("urn:schemas-upnp-org:service:AVTransport:1#EndDirectControlSession", httpClient,
		&AVTransportEnvelope{
			Body:          AVTransportBody{EndDirectControlSession: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EndDirectControlSession == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.EndDirectControlSession, nil
}
func (s *AVTransportService) AVTransportSubscribe(callback url.URL) error {
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
type AVTransportLastChange string
type AVTransportUpnpEvent struct {
	XMLName      xml.Name              `xml:"propertyset"`
	XMLNameSpace string                `xml:"xmlns:e,attr"`
	Properties   []AVTransportProperty `xml:"property"`
}
type AVTransportProperty struct {
	XMLName    xml.Name               `xml:"property"`
	LastChange *AVTransportLastChange `xml:"LastChange"`
}

func AVTransportDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt AVTransportUpnpEvent
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
