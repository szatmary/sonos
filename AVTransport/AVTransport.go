package avtransport

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:AVTransport:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/MediaRenderer/AVTransport/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/MediaRenderer/AVTransport/Event`)
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
	XMLName                            xml.Name                                `xml:"s:Body"`
	SetAVTransportURI                  *SetAVTransportURIArgs                  `xml:"u:SetAVTransportURI,omitempty"`
	SetNextAVTransportURI              *SetNextAVTransportURIArgs              `xml:"u:SetNextAVTransportURI,omitempty"`
	AddURIToQueue                      *AddURIToQueueArgs                      `xml:"u:AddURIToQueue,omitempty"`
	AddMultipleURIsToQueue             *AddMultipleURIsToQueueArgs             `xml:"u:AddMultipleURIsToQueue,omitempty"`
	ReorderTracksInQueue               *ReorderTracksInQueueArgs               `xml:"u:ReorderTracksInQueue,omitempty"`
	RemoveTrackFromQueue               *RemoveTrackFromQueueArgs               `xml:"u:RemoveTrackFromQueue,omitempty"`
	RemoveTrackRangeFromQueue          *RemoveTrackRangeFromQueueArgs          `xml:"u:RemoveTrackRangeFromQueue,omitempty"`
	RemoveAllTracksFromQueue           *RemoveAllTracksFromQueueArgs           `xml:"u:RemoveAllTracksFromQueue,omitempty"`
	SaveQueue                          *SaveQueueArgs                          `xml:"u:SaveQueue,omitempty"`
	BackupQueue                        *BackupQueueArgs                        `xml:"u:BackupQueue,omitempty"`
	CreateSavedQueue                   *CreateSavedQueueArgs                   `xml:"u:CreateSavedQueue,omitempty"`
	AddURIToSavedQueue                 *AddURIToSavedQueueArgs                 `xml:"u:AddURIToSavedQueue,omitempty"`
	ReorderTracksInSavedQueue          *ReorderTracksInSavedQueueArgs          `xml:"u:ReorderTracksInSavedQueue,omitempty"`
	GetMediaInfo                       *GetMediaInfoArgs                       `xml:"u:GetMediaInfo,omitempty"`
	GetTransportInfo                   *GetTransportInfoArgs                   `xml:"u:GetTransportInfo,omitempty"`
	GetPositionInfo                    *GetPositionInfoArgs                    `xml:"u:GetPositionInfo,omitempty"`
	GetDeviceCapabilities              *GetDeviceCapabilitiesArgs              `xml:"u:GetDeviceCapabilities,omitempty"`
	GetTransportSettings               *GetTransportSettingsArgs               `xml:"u:GetTransportSettings,omitempty"`
	GetCrossfadeMode                   *GetCrossfadeModeArgs                   `xml:"u:GetCrossfadeMode,omitempty"`
	Stop                               *StopArgs                               `xml:"u:Stop,omitempty"`
	Play                               *PlayArgs                               `xml:"u:Play,omitempty"`
	Pause                              *PauseArgs                              `xml:"u:Pause,omitempty"`
	Seek                               *SeekArgs                               `xml:"u:Seek,omitempty"`
	Next                               *NextArgs                               `xml:"u:Next,omitempty"`
	Previous                           *PreviousArgs                           `xml:"u:Previous,omitempty"`
	SetPlayMode                        *SetPlayModeArgs                        `xml:"u:SetPlayMode,omitempty"`
	SetCrossfadeMode                   *SetCrossfadeModeArgs                   `xml:"u:SetCrossfadeMode,omitempty"`
	NotifyDeletedURI                   *NotifyDeletedURIArgs                   `xml:"u:NotifyDeletedURI,omitempty"`
	GetCurrentTransportActions         *GetCurrentTransportActionsArgs         `xml:"u:GetCurrentTransportActions,omitempty"`
	BecomeCoordinatorOfStandaloneGroup *BecomeCoordinatorOfStandaloneGroupArgs `xml:"u:BecomeCoordinatorOfStandaloneGroup,omitempty"`
	DelegateGroupCoordinationTo        *DelegateGroupCoordinationToArgs        `xml:"u:DelegateGroupCoordinationTo,omitempty"`
	BecomeGroupCoordinator             *BecomeGroupCoordinatorArgs             `xml:"u:BecomeGroupCoordinator,omitempty"`
	BecomeGroupCoordinatorAndSource    *BecomeGroupCoordinatorAndSourceArgs    `xml:"u:BecomeGroupCoordinatorAndSource,omitempty"`
	ChangeCoordinator                  *ChangeCoordinatorArgs                  `xml:"u:ChangeCoordinator,omitempty"`
	ChangeTransportSettings            *ChangeTransportSettingsArgs            `xml:"u:ChangeTransportSettings,omitempty"`
	ConfigureSleepTimer                *ConfigureSleepTimerArgs                `xml:"u:ConfigureSleepTimer,omitempty"`
	GetRemainingSleepTimerDuration     *GetRemainingSleepTimerDurationArgs     `xml:"u:GetRemainingSleepTimerDuration,omitempty"`
	RunAlarm                           *RunAlarmArgs                           `xml:"u:RunAlarm,omitempty"`
	StartAutoplay                      *StartAutoplayArgs                      `xml:"u:StartAutoplay,omitempty"`
	GetRunningAlarmProperties          *GetRunningAlarmPropertiesArgs          `xml:"u:GetRunningAlarmProperties,omitempty"`
	SnoozeAlarm                        *SnoozeAlarmArgs                        `xml:"u:SnoozeAlarm,omitempty"`
	EndDirectControlSession            *EndDirectControlSessionArgs            `xml:"u:EndDirectControlSession,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                            xml.Name                                    `xml:"Body"`
	SetAVTransportURI                  *SetAVTransportURIResponse                  `xml:"SetAVTransportURIResponse,omitempty"`
	SetNextAVTransportURI              *SetNextAVTransportURIResponse              `xml:"SetNextAVTransportURIResponse,omitempty"`
	AddURIToQueue                      *AddURIToQueueResponse                      `xml:"AddURIToQueueResponse,omitempty"`
	AddMultipleURIsToQueue             *AddMultipleURIsToQueueResponse             `xml:"AddMultipleURIsToQueueResponse,omitempty"`
	ReorderTracksInQueue               *ReorderTracksInQueueResponse               `xml:"ReorderTracksInQueueResponse,omitempty"`
	RemoveTrackFromQueue               *RemoveTrackFromQueueResponse               `xml:"RemoveTrackFromQueueResponse,omitempty"`
	RemoveTrackRangeFromQueue          *RemoveTrackRangeFromQueueResponse          `xml:"RemoveTrackRangeFromQueueResponse,omitempty"`
	RemoveAllTracksFromQueue           *RemoveAllTracksFromQueueResponse           `xml:"RemoveAllTracksFromQueueResponse,omitempty"`
	SaveQueue                          *SaveQueueResponse                          `xml:"SaveQueueResponse,omitempty"`
	BackupQueue                        *BackupQueueResponse                        `xml:"BackupQueueResponse,omitempty"`
	CreateSavedQueue                   *CreateSavedQueueResponse                   `xml:"CreateSavedQueueResponse,omitempty"`
	AddURIToSavedQueue                 *AddURIToSavedQueueResponse                 `xml:"AddURIToSavedQueueResponse,omitempty"`
	ReorderTracksInSavedQueue          *ReorderTracksInSavedQueueResponse          `xml:"ReorderTracksInSavedQueueResponse,omitempty"`
	GetMediaInfo                       *GetMediaInfoResponse                       `xml:"GetMediaInfoResponse,omitempty"`
	GetTransportInfo                   *GetTransportInfoResponse                   `xml:"GetTransportInfoResponse,omitempty"`
	GetPositionInfo                    *GetPositionInfoResponse                    `xml:"GetPositionInfoResponse,omitempty"`
	GetDeviceCapabilities              *GetDeviceCapabilitiesResponse              `xml:"GetDeviceCapabilitiesResponse,omitempty"`
	GetTransportSettings               *GetTransportSettingsResponse               `xml:"GetTransportSettingsResponse,omitempty"`
	GetCrossfadeMode                   *GetCrossfadeModeResponse                   `xml:"GetCrossfadeModeResponse,omitempty"`
	Stop                               *StopResponse                               `xml:"StopResponse,omitempty"`
	Play                               *PlayResponse                               `xml:"PlayResponse,omitempty"`
	Pause                              *PauseResponse                              `xml:"PauseResponse,omitempty"`
	Seek                               *SeekResponse                               `xml:"SeekResponse,omitempty"`
	Next                               *NextResponse                               `xml:"NextResponse,omitempty"`
	Previous                           *PreviousResponse                           `xml:"PreviousResponse,omitempty"`
	SetPlayMode                        *SetPlayModeResponse                        `xml:"SetPlayModeResponse,omitempty"`
	SetCrossfadeMode                   *SetCrossfadeModeResponse                   `xml:"SetCrossfadeModeResponse,omitempty"`
	NotifyDeletedURI                   *NotifyDeletedURIResponse                   `xml:"NotifyDeletedURIResponse,omitempty"`
	GetCurrentTransportActions         *GetCurrentTransportActionsResponse         `xml:"GetCurrentTransportActionsResponse,omitempty"`
	BecomeCoordinatorOfStandaloneGroup *BecomeCoordinatorOfStandaloneGroupResponse `xml:"BecomeCoordinatorOfStandaloneGroupResponse,omitempty"`
	DelegateGroupCoordinationTo        *DelegateGroupCoordinationToResponse        `xml:"DelegateGroupCoordinationToResponse,omitempty"`
	BecomeGroupCoordinator             *BecomeGroupCoordinatorResponse             `xml:"BecomeGroupCoordinatorResponse,omitempty"`
	BecomeGroupCoordinatorAndSource    *BecomeGroupCoordinatorAndSourceResponse    `xml:"BecomeGroupCoordinatorAndSourceResponse,omitempty"`
	ChangeCoordinator                  *ChangeCoordinatorResponse                  `xml:"ChangeCoordinatorResponse,omitempty"`
	ChangeTransportSettings            *ChangeTransportSettingsResponse            `xml:"ChangeTransportSettingsResponse,omitempty"`
	ConfigureSleepTimer                *ConfigureSleepTimerResponse                `xml:"ConfigureSleepTimerResponse,omitempty"`
	GetRemainingSleepTimerDuration     *GetRemainingSleepTimerDurationResponse     `xml:"GetRemainingSleepTimerDurationResponse,omitempty"`
	RunAlarm                           *RunAlarmResponse                           `xml:"RunAlarmResponse,omitempty"`
	StartAutoplay                      *StartAutoplayResponse                      `xml:"StartAutoplayResponse,omitempty"`
	GetRunningAlarmProperties          *GetRunningAlarmPropertiesResponse          `xml:"GetRunningAlarmPropertiesResponse,omitempty"`
	SnoozeAlarm                        *SnoozeAlarmResponse                        `xml:"SnoozeAlarmResponse,omitempty"`
	EndDirectControlSession            *EndDirectControlSessionResponse            `xml:"EndDirectControlSessionResponse,omitempty"`
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

type SetAVTransportURIArgs struct {
	Xmlns              string `xml:"xmlns:u,attr"`
	InstanceID         uint32 `xml:"InstanceID"`
	CurrentURI         string `xml:"CurrentURI"`
	CurrentURIMetaData string `xml:"CurrentURIMetaData"`
}
type SetAVTransportURIResponse struct {
}

func (s *Service) SetAVTransportURI(httpClient *http.Client, args *SetAVTransportURIArgs) (*SetAVTransportURIResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetAVTransportURI`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetAVTransportURI: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAVTransportURI == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.SetAVTransportURI()`)
	}

	return r.Body.SetAVTransportURI, nil
}

type SetNextAVTransportURIArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
	InstanceID      uint32 `xml:"InstanceID"`
	NextURI         string `xml:"NextURI"`
	NextURIMetaData string `xml:"NextURIMetaData"`
}
type SetNextAVTransportURIResponse struct {
}

func (s *Service) SetNextAVTransportURI(httpClient *http.Client, args *SetNextAVTransportURIArgs) (*SetNextAVTransportURIResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetNextAVTransportURI`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetNextAVTransportURI: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetNextAVTransportURI == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.SetNextAVTransportURI()`)
	}

	return r.Body.SetNextAVTransportURI, nil
}

type AddURIToQueueArgs struct {
	Xmlns                           string `xml:"xmlns:u,attr"`
	InstanceID                      uint32 `xml:"InstanceID"`
	EnqueuedURI                     string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData             string `xml:"EnqueuedURIMetaData"`
	DesiredFirstTrackNumberEnqueued uint32 `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                   bool   `xml:"EnqueueAsNext"`
}
type AddURIToQueueResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
}

func (s *Service) AddURIToQueue(httpClient *http.Client, args *AddURIToQueueArgs) (*AddURIToQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddURIToQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddURIToQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddURIToQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.AddURIToQueue()`)
	}

	return r.Body.AddURIToQueue, nil
}

type AddMultipleURIsToQueueArgs struct {
	Xmlns                           string `xml:"xmlns:u,attr"`
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
type AddMultipleURIsToQueueResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
	NewUpdateID              uint32 `xml:"NewUpdateID"`
}

func (s *Service) AddMultipleURIsToQueue(httpClient *http.Client, args *AddMultipleURIsToQueueArgs) (*AddMultipleURIsToQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddMultipleURIsToQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddMultipleURIsToQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddMultipleURIsToQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.AddMultipleURIsToQueue()`)
	}

	return r.Body.AddMultipleURIsToQueue, nil
}

type ReorderTracksInQueueArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	InstanceID     uint32 `xml:"InstanceID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
	InsertBefore   uint32 `xml:"InsertBefore"`
	UpdateID       uint32 `xml:"UpdateID"`
}
type ReorderTracksInQueueResponse struct {
}

func (s *Service) ReorderTracksInQueue(httpClient *http.Client, args *ReorderTracksInQueueArgs) (*ReorderTracksInQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReorderTracksInQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReorderTracksInQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReorderTracksInQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.ReorderTracksInQueue()`)
	}

	return r.Body.ReorderTracksInQueue, nil
}

type RemoveTrackFromQueueArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	ObjectID   string `xml:"ObjectID"`
	UpdateID   uint32 `xml:"UpdateID"`
}
type RemoveTrackFromQueueResponse struct {
}

func (s *Service) RemoveTrackFromQueue(httpClient *http.Client, args *RemoveTrackFromQueueArgs) (*RemoveTrackFromQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveTrackFromQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveTrackFromQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveTrackFromQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.RemoveTrackFromQueue()`)
	}

	return r.Body.RemoveTrackFromQueue, nil
}

type RemoveTrackRangeFromQueueArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	InstanceID     uint32 `xml:"InstanceID"`
	UpdateID       uint32 `xml:"UpdateID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
}
type RemoveTrackRangeFromQueueResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *Service) RemoveTrackRangeFromQueue(httpClient *http.Client, args *RemoveTrackRangeFromQueueArgs) (*RemoveTrackRangeFromQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveTrackRangeFromQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveTrackRangeFromQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveTrackRangeFromQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.RemoveTrackRangeFromQueue()`)
	}

	return r.Body.RemoveTrackRangeFromQueue, nil
}

type RemoveAllTracksFromQueueArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type RemoveAllTracksFromQueueResponse struct {
}

func (s *Service) RemoveAllTracksFromQueue(httpClient *http.Client, args *RemoveAllTracksFromQueueArgs) (*RemoveAllTracksFromQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveAllTracksFromQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveAllTracksFromQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveAllTracksFromQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.RemoveAllTracksFromQueue()`)
	}

	return r.Body.RemoveAllTracksFromQueue, nil
}

type SaveQueueArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	Title      string `xml:"Title"`
	ObjectID   string `xml:"ObjectID"`
}
type SaveQueueResponse struct {
	AssignedObjectID string `xml:"AssignedObjectID"`
}

func (s *Service) SaveQueue(httpClient *http.Client, args *SaveQueueArgs) (*SaveQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SaveQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SaveQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SaveQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.SaveQueue()`)
	}

	return r.Body.SaveQueue, nil
}

type BackupQueueArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type BackupQueueResponse struct {
}

func (s *Service) BackupQueue(httpClient *http.Client, args *BackupQueueArgs) (*BackupQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`BackupQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{BackupQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BackupQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.BackupQueue()`)
	}

	return r.Body.BackupQueue, nil
}

type CreateSavedQueueArgs struct {
	Xmlns               string `xml:"xmlns:u,attr"`
	InstanceID          uint32 `xml:"InstanceID"`
	Title               string `xml:"Title"`
	EnqueuedURI         string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData string `xml:"EnqueuedURIMetaData"`
}
type CreateSavedQueueResponse struct {
	NumTracksAdded   uint32 `xml:"NumTracksAdded"`
	NewQueueLength   uint32 `xml:"NewQueueLength"`
	AssignedObjectID string `xml:"AssignedObjectID"`
	NewUpdateID      uint32 `xml:"NewUpdateID"`
}

func (s *Service) CreateSavedQueue(httpClient *http.Client, args *CreateSavedQueueArgs) (*CreateSavedQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`CreateSavedQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{CreateSavedQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateSavedQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.CreateSavedQueue()`)
	}

	return r.Body.CreateSavedQueue, nil
}

type AddURIToSavedQueueArgs struct {
	Xmlns               string `xml:"xmlns:u,attr"`
	InstanceID          uint32 `xml:"InstanceID"`
	ObjectID            string `xml:"ObjectID"`
	UpdateID            uint32 `xml:"UpdateID"`
	EnqueuedURI         string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData string `xml:"EnqueuedURIMetaData"`
	AddAtIndex          uint32 `xml:"AddAtIndex"`
}
type AddURIToSavedQueueResponse struct {
	NumTracksAdded uint32 `xml:"NumTracksAdded"`
	NewQueueLength uint32 `xml:"NewQueueLength"`
	NewUpdateID    uint32 `xml:"NewUpdateID"`
}

func (s *Service) AddURIToSavedQueue(httpClient *http.Client, args *AddURIToSavedQueueArgs) (*AddURIToSavedQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddURIToSavedQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddURIToSavedQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddURIToSavedQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.AddURIToSavedQueue()`)
	}

	return r.Body.AddURIToSavedQueue, nil
}

type ReorderTracksInSavedQueueArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
	InstanceID      uint32 `xml:"InstanceID"`
	ObjectID        string `xml:"ObjectID"`
	UpdateID        uint32 `xml:"UpdateID"`
	TrackList       string `xml:"TrackList"`
	NewPositionList string `xml:"NewPositionList"`
}
type ReorderTracksInSavedQueueResponse struct {
	QueueLengthChange int32  `xml:"QueueLengthChange"`
	NewQueueLength    uint32 `xml:"NewQueueLength"`
	NewUpdateID       uint32 `xml:"NewUpdateID"`
}

func (s *Service) ReorderTracksInSavedQueue(httpClient *http.Client, args *ReorderTracksInSavedQueueArgs) (*ReorderTracksInSavedQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReorderTracksInSavedQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReorderTracksInSavedQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReorderTracksInSavedQueue == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.ReorderTracksInSavedQueue()`)
	}

	return r.Body.ReorderTracksInSavedQueue, nil
}

type GetMediaInfoArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetMediaInfoResponse struct {
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

func (s *Service) GetMediaInfo(httpClient *http.Client, args *GetMediaInfoArgs) (*GetMediaInfoResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetMediaInfo`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetMediaInfo: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetMediaInfo == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetMediaInfo()`)
	}

	return r.Body.GetMediaInfo, nil
}

type GetTransportInfoArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetTransportInfoResponse struct {
	CurrentTransportState  string `xml:"CurrentTransportState"`
	CurrentTransportStatus string `xml:"CurrentTransportStatus"`
	CurrentSpeed           string `xml:"CurrentSpeed"`
}

func (s *Service) GetTransportInfo(httpClient *http.Client, args *GetTransportInfoArgs) (*GetTransportInfoResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetTransportInfo`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetTransportInfo: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTransportInfo == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetTransportInfo()`)
	}

	return r.Body.GetTransportInfo, nil
}

type GetPositionInfoArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetPositionInfoResponse struct {
	Track         uint32 `xml:"Track"`
	TrackDuration string `xml:"TrackDuration"`
	TrackMetaData string `xml:"TrackMetaData"`
	TrackURI      string `xml:"TrackURI"`
	RelTime       string `xml:"RelTime"`
	AbsTime       string `xml:"AbsTime"`
	RelCount      int32  `xml:"RelCount"`
	AbsCount      int32  `xml:"AbsCount"`
}

func (s *Service) GetPositionInfo(httpClient *http.Client, args *GetPositionInfoArgs) (*GetPositionInfoResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetPositionInfo`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetPositionInfo: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetPositionInfo == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetPositionInfo()`)
	}

	return r.Body.GetPositionInfo, nil
}

type GetDeviceCapabilitiesArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetDeviceCapabilitiesResponse struct {
	PlayMedia       string `xml:"PlayMedia"`
	RecMedia        string `xml:"RecMedia"`
	RecQualityModes string `xml:"RecQualityModes"`
}

func (s *Service) GetDeviceCapabilities(httpClient *http.Client, args *GetDeviceCapabilitiesArgs) (*GetDeviceCapabilitiesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetDeviceCapabilities`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetDeviceCapabilities: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetDeviceCapabilities == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetDeviceCapabilities()`)
	}

	return r.Body.GetDeviceCapabilities, nil
}

type GetTransportSettingsArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetTransportSettingsResponse struct {
	PlayMode       string `xml:"PlayMode"`
	RecQualityMode string `xml:"RecQualityMode"`
}

func (s *Service) GetTransportSettings(httpClient *http.Client, args *GetTransportSettingsArgs) (*GetTransportSettingsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetTransportSettings`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetTransportSettings: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTransportSettings == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetTransportSettings()`)
	}

	return r.Body.GetTransportSettings, nil
}

type GetCrossfadeModeArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetCrossfadeModeResponse struct {
	CrossfadeMode bool `xml:"CrossfadeMode"`
}

func (s *Service) GetCrossfadeMode(httpClient *http.Client, args *GetCrossfadeModeArgs) (*GetCrossfadeModeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetCrossfadeMode`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetCrossfadeMode: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCrossfadeMode == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetCrossfadeMode()`)
	}

	return r.Body.GetCrossfadeMode, nil
}

type StopArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type StopResponse struct {
}

func (s *Service) Stop(httpClient *http.Client, args *StopArgs) (*StopResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Stop`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Stop: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Stop == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.Stop()`)
	}

	return r.Body.Stop, nil
}

type PlayArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	// Allowed Value: 1
	Speed string `xml:"Speed"`
}
type PlayResponse struct {
}

func (s *Service) Play(httpClient *http.Client, args *PlayArgs) (*PlayResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Play`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Play: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Play == nil {
		return nil, errors.New(`unexpected response from service calling avtransport.Play()`)
	}

	return r.Body.Play, nil
}

type PauseArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type PauseResponse struct {
}

func (s *Service) Pause(httpClient *http.Client, args *PauseArgs) (*PauseResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Pause`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Pause: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Pause == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.Pause()`)
	}

	return r.Body.Pause, nil
}

type SeekArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	// Allowed Value: TRACK_NR
	// Allowed Value: REL_TIME
	// Allowed Value: TIME_DELTA
	Unit   string `xml:"Unit"`
	Target string `xml:"Target"`
}
type SeekResponse struct {
}

func (s *Service) Seek(httpClient *http.Client, args *SeekArgs) (*SeekResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Seek`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Seek: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Seek == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.Seek()`)
	}

	return r.Body.Seek, nil
}

type NextArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type NextResponse struct {
}

func (s *Service) Next(httpClient *http.Client, args *NextArgs) (*NextResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Next`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Next: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Next == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.Next()`)
	}

	return r.Body.Next, nil
}

type PreviousArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type PreviousResponse struct {
}

func (s *Service) Previous(httpClient *http.Client, args *PreviousArgs) (*PreviousResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Previous`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Previous: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Previous == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.Previous()`)
	}

	return r.Body.Previous, nil
}

type SetPlayModeArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	// Allowed Value: NORMAL
	// Allowed Value: REPEAT_ALL
	// Allowed Value: REPEAT_ONE
	// Allowed Value: SHUFFLE_NOREPEAT
	// Allowed Value: SHUFFLE
	// Allowed Value: SHUFFLE_REPEAT_ONE
	NewPlayMode string `xml:"NewPlayMode"`
}
type SetPlayModeResponse struct {
}

func (s *Service) SetPlayMode(httpClient *http.Client, args *SetPlayModeArgs) (*SetPlayModeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetPlayMode`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetPlayMode: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetPlayMode == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.SetPlayMode()`)
	}

	return r.Body.SetPlayMode, nil
}

type SetCrossfadeModeArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	InstanceID    uint32 `xml:"InstanceID"`
	CrossfadeMode bool   `xml:"CrossfadeMode"`
}
type SetCrossfadeModeResponse struct {
}

func (s *Service) SetCrossfadeMode(httpClient *http.Client, args *SetCrossfadeModeArgs) (*SetCrossfadeModeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetCrossfadeMode`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetCrossfadeMode: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetCrossfadeMode == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.SetCrossfadeMode()`)
	}

	return r.Body.SetCrossfadeMode, nil
}

type NotifyDeletedURIArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	DeletedURI string `xml:"DeletedURI"`
}
type NotifyDeletedURIResponse struct {
}

func (s *Service) NotifyDeletedURI(httpClient *http.Client, args *NotifyDeletedURIArgs) (*NotifyDeletedURIResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`NotifyDeletedURI`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{NotifyDeletedURI: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.NotifyDeletedURI == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.NotifyDeletedURI()`)
	}

	return r.Body.NotifyDeletedURI, nil
}

type GetCurrentTransportActionsArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetCurrentTransportActionsResponse struct {
	Actions string `xml:"Actions"`
}

func (s *Service) GetCurrentTransportActions(httpClient *http.Client, args *GetCurrentTransportActionsArgs) (*GetCurrentTransportActionsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetCurrentTransportActions`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetCurrentTransportActions: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetCurrentTransportActions == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetCurrentTransportActions()`)
	}

	return r.Body.GetCurrentTransportActions, nil
}

type BecomeCoordinatorOfStandaloneGroupArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type BecomeCoordinatorOfStandaloneGroupResponse struct {
	DelegatedGroupCoordinatorID string `xml:"DelegatedGroupCoordinatorID"`
	NewGroupID                  string `xml:"NewGroupID"`
}

func (s *Service) BecomeCoordinatorOfStandaloneGroup(httpClient *http.Client, args *BecomeCoordinatorOfStandaloneGroupArgs) (*BecomeCoordinatorOfStandaloneGroupResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`BecomeCoordinatorOfStandaloneGroup`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{BecomeCoordinatorOfStandaloneGroup: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BecomeCoordinatorOfStandaloneGroup == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.BecomeCoordinatorOfStandaloneGroup()`)
	}

	return r.Body.BecomeCoordinatorOfStandaloneGroup, nil
}

type DelegateGroupCoordinationToArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	InstanceID     uint32 `xml:"InstanceID"`
	NewCoordinator string `xml:"NewCoordinator"`
	RejoinGroup    bool   `xml:"RejoinGroup"`
}
type DelegateGroupCoordinationToResponse struct {
}

func (s *Service) DelegateGroupCoordinationTo(httpClient *http.Client, args *DelegateGroupCoordinationToArgs) (*DelegateGroupCoordinationToResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`DelegateGroupCoordinationTo`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{DelegateGroupCoordinationTo: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DelegateGroupCoordinationTo == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.DelegateGroupCoordinationTo()`)
	}

	return r.Body.DelegateGroupCoordinationTo, nil
}

type BecomeGroupCoordinatorArgs struct {
	Xmlns                 string `xml:"xmlns:u,attr"`
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
type BecomeGroupCoordinatorResponse struct {
}

func (s *Service) BecomeGroupCoordinator(httpClient *http.Client, args *BecomeGroupCoordinatorArgs) (*BecomeGroupCoordinatorResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`BecomeGroupCoordinator`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{BecomeGroupCoordinator: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BecomeGroupCoordinator == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.BecomeGroupCoordinator()`)
	}

	return r.Body.BecomeGroupCoordinator, nil
}

type BecomeGroupCoordinatorAndSourceArgs struct {
	Xmlns                 string `xml:"xmlns:u,attr"`
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
type BecomeGroupCoordinatorAndSourceResponse struct {
}

func (s *Service) BecomeGroupCoordinatorAndSource(httpClient *http.Client, args *BecomeGroupCoordinatorAndSourceArgs) (*BecomeGroupCoordinatorAndSourceResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`BecomeGroupCoordinatorAndSource`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{BecomeGroupCoordinatorAndSource: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.BecomeGroupCoordinatorAndSource == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.BecomeGroupCoordinatorAndSource()`)
	}

	return r.Body.BecomeGroupCoordinatorAndSource, nil
}

type ChangeCoordinatorArgs struct {
	Xmlns                 string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	CurrentCoordinator    string `xml:"CurrentCoordinator"`
	NewCoordinator        string `xml:"NewCoordinator"`
	NewTransportSettings  string `xml:"NewTransportSettings"`
	CurrentAVTransportURI string `xml:"CurrentAVTransportURI"`
}
type ChangeCoordinatorResponse struct {
}

func (s *Service) ChangeCoordinator(httpClient *http.Client, args *ChangeCoordinatorArgs) (*ChangeCoordinatorResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ChangeCoordinator`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ChangeCoordinator: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ChangeCoordinator == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.ChangeCoordinator()`)
	}

	return r.Body.ChangeCoordinator, nil
}

type ChangeTransportSettingsArgs struct {
	Xmlns                 string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	NewTransportSettings  string `xml:"NewTransportSettings"`
	CurrentAVTransportURI string `xml:"CurrentAVTransportURI"`
}
type ChangeTransportSettingsResponse struct {
}

func (s *Service) ChangeTransportSettings(httpClient *http.Client, args *ChangeTransportSettingsArgs) (*ChangeTransportSettingsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ChangeTransportSettings`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ChangeTransportSettings: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ChangeTransportSettings == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.ChangeTransportSettings()`)
	}

	return r.Body.ChangeTransportSettings, nil
}

type ConfigureSleepTimerArgs struct {
	Xmlns                 string `xml:"xmlns:u,attr"`
	InstanceID            uint32 `xml:"InstanceID"`
	NewSleepTimerDuration string `xml:"NewSleepTimerDuration"`
}
type ConfigureSleepTimerResponse struct {
}

func (s *Service) ConfigureSleepTimer(httpClient *http.Client, args *ConfigureSleepTimerArgs) (*ConfigureSleepTimerResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ConfigureSleepTimer`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ConfigureSleepTimer: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ConfigureSleepTimer == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.ConfigureSleepTimer()`)
	}

	return r.Body.ConfigureSleepTimer, nil
}

type GetRemainingSleepTimerDurationArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetRemainingSleepTimerDurationResponse struct {
	RemainingSleepTimerDuration string `xml:"RemainingSleepTimerDuration"`
	CurrentSleepTimerGeneration uint32 `xml:"CurrentSleepTimerGeneration"`
}

func (s *Service) GetRemainingSleepTimerDuration(httpClient *http.Client, args *GetRemainingSleepTimerDurationArgs) (*GetRemainingSleepTimerDurationResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetRemainingSleepTimerDuration`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetRemainingSleepTimerDuration: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetRemainingSleepTimerDuration == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetRemainingSleepTimerDuration()`)
	}

	return r.Body.GetRemainingSleepTimerDuration, nil
}

type RunAlarmArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
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
type RunAlarmResponse struct {
}

func (s *Service) RunAlarm(httpClient *http.Client, args *RunAlarmArgs) (*RunAlarmResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RunAlarm`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RunAlarm: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RunAlarm == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.RunAlarm()`)
	}

	return r.Body.RunAlarm, nil
}

type StartAutoplayArgs struct {
	Xmlns              string `xml:"xmlns:u,attr"`
	InstanceID         uint32 `xml:"InstanceID"`
	ProgramURI         string `xml:"ProgramURI"`
	ProgramMetaData    string `xml:"ProgramMetaData"`
	Volume             uint16 `xml:"Volume"`
	IncludeLinkedZones bool   `xml:"IncludeLinkedZones"`
	ResetVolumeAfter   bool   `xml:"ResetVolumeAfter"`
}
type StartAutoplayResponse struct {
}

func (s *Service) StartAutoplay(httpClient *http.Client, args *StartAutoplayArgs) (*StartAutoplayResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`StartAutoplay`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{StartAutoplay: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.StartAutoplay == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.StartAutoplay()`)
	}

	return r.Body.StartAutoplay, nil
}

type GetRunningAlarmPropertiesArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type GetRunningAlarmPropertiesResponse struct {
	AlarmID         uint32 `xml:"AlarmID"`
	GroupID         string `xml:"GroupID"`
	LoggedStartTime string `xml:"LoggedStartTime"`
}

func (s *Service) GetRunningAlarmProperties(httpClient *http.Client, args *GetRunningAlarmPropertiesArgs) (*GetRunningAlarmPropertiesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetRunningAlarmProperties`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetRunningAlarmProperties: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetRunningAlarmProperties == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.GetRunningAlarmProperties()`)
	}

	return r.Body.GetRunningAlarmProperties, nil
}

type SnoozeAlarmArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
	Duration   string `xml:"Duration"`
}
type SnoozeAlarmResponse struct {
}

func (s *Service) SnoozeAlarm(httpClient *http.Client, args *SnoozeAlarmArgs) (*SnoozeAlarmResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SnoozeAlarm`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SnoozeAlarm: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SnoozeAlarm == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.SnoozeAlarm()`)
	}

	return r.Body.SnoozeAlarm, nil
}

type EndDirectControlSessionArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	InstanceID uint32 `xml:"InstanceID"`
}
type EndDirectControlSessionResponse struct {
}

func (s *Service) EndDirectControlSession(httpClient *http.Client, args *EndDirectControlSessionArgs) (*EndDirectControlSessionResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`EndDirectControlSession`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{EndDirectControlSession: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EndDirectControlSession == nil {
		return nil, errors.New(`unexpected respose from service calling avtransport.EndDirectControlSession()`)
	}

	return r.Body.EndDirectControlSession, nil
}
