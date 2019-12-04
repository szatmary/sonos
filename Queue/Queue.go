package queue

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:Queue:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/MediaRenderer/Queue/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/MediaRenderer/Queue/Event`)
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
	XMLName             xml.Name                 `xml:"s:Body"`
	AddURI              *AddURIArgs              `xml:"u:AddURI,omitempty"`
	AddMultipleURIs     *AddMultipleURIsArgs     `xml:"u:AddMultipleURIs,omitempty"`
	AttachQueue         *AttachQueueArgs         `xml:"u:AttachQueue,omitempty"`
	Backup              *BackupArgs              `xml:"u:Backup,omitempty"`
	Browse              *BrowseArgs              `xml:"u:Browse,omitempty"`
	CreateQueue         *CreateQueueArgs         `xml:"u:CreateQueue,omitempty"`
	RemoveAllTracks     *RemoveAllTracksArgs     `xml:"u:RemoveAllTracks,omitempty"`
	RemoveTrackRange    *RemoveTrackRangeArgs    `xml:"u:RemoveTrackRange,omitempty"`
	ReorderTracks       *ReorderTracksArgs       `xml:"u:ReorderTracks,omitempty"`
	ReplaceAllTracks    *ReplaceAllTracksArgs    `xml:"u:ReplaceAllTracks,omitempty"`
	SaveAsSonosPlaylist *SaveAsSonosPlaylistArgs `xml:"u:SaveAsSonosPlaylist,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName             xml.Name                     `xml:"Body"`
	AddURI              *AddURIResponse              `xml:"AddURIResponse,omitempty"`
	AddMultipleURIs     *AddMultipleURIsResponse     `xml:"AddMultipleURIsResponse,omitempty"`
	AttachQueue         *AttachQueueResponse         `xml:"AttachQueueResponse,omitempty"`
	Backup              *BackupResponse              `xml:"BackupResponse,omitempty"`
	Browse              *BrowseResponse              `xml:"BrowseResponse,omitempty"`
	CreateQueue         *CreateQueueResponse         `xml:"CreateQueueResponse,omitempty"`
	RemoveAllTracks     *RemoveAllTracksResponse     `xml:"RemoveAllTracksResponse,omitempty"`
	RemoveTrackRange    *RemoveTrackRangeResponse    `xml:"RemoveTrackRangeResponse,omitempty"`
	ReorderTracks       *ReorderTracksResponse       `xml:"ReorderTracksResponse,omitempty"`
	ReplaceAllTracks    *ReplaceAllTracksResponse    `xml:"ReplaceAllTracksResponse,omitempty"`
	SaveAsSonosPlaylist *SaveAsSonosPlaylistResponse `xml:"SaveAsSonosPlaylistResponse,omitempty"`
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

type AddURIArgs struct {
	Xmlns                           string `xml:"xmlns:u,attr"`
	QueueID                         uint32 `xml:"QueueID"`
	UpdateID                        uint32 `xml:"UpdateID"`
	EnqueuedURI                     string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData             string `xml:"EnqueuedURIMetaData"`
	DesiredFirstTrackNumberEnqueued uint32 `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                   bool   `xml:"EnqueueAsNext"`
}
type AddURIResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
	NewUpdateID              uint32 `xml:"NewUpdateID"`
}

func (s *Service) AddURI(httpClient *http.Client, args *AddURIArgs) (*AddURIResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddURI`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddURI: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddURI == nil {
		return nil, errors.New(`unexpected respose from service calling queue.AddURI()`)
	}

	return r.Body.AddURI, nil
}

type AddMultipleURIsArgs struct {
	Xmlns                           string `xml:"xmlns:u,attr"`
	QueueID                         uint32 `xml:"QueueID"`
	UpdateID                        uint32 `xml:"UpdateID"`
	ContainerURI                    string `xml:"ContainerURI"`
	ContainerMetaData               string `xml:"ContainerMetaData"`
	DesiredFirstTrackNumberEnqueued uint32 `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                   bool   `xml:"EnqueueAsNext"`
	NumberOfURIs                    uint32 `xml:"NumberOfURIs"`
	EnqueuedURIsAndMetaData         string `xml:"EnqueuedURIsAndMetaData"`
}
type AddMultipleURIsResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
	NewUpdateID              uint32 `xml:"NewUpdateID"`
}

func (s *Service) AddMultipleURIs(httpClient *http.Client, args *AddMultipleURIsArgs) (*AddMultipleURIsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddMultipleURIs`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddMultipleURIs: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddMultipleURIs == nil {
		return nil, errors.New(`unexpected respose from service calling queue.AddMultipleURIs()`)
	}

	return r.Body.AddMultipleURIs, nil
}

type AttachQueueArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	QueueOwnerID string `xml:"QueueOwnerID"`
}
type AttachQueueResponse struct {
	QueueID           uint32 `xml:"QueueID"`
	QueueOwnerContext string `xml:"QueueOwnerContext"`
}

func (s *Service) AttachQueue(httpClient *http.Client, args *AttachQueueArgs) (*AttachQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AttachQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AttachQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AttachQueue == nil {
		return nil, errors.New(`unexpected respose from service calling queue.AttachQueue()`)
	}

	return r.Body.AttachQueue, nil
}

type BackupArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type BackupResponse struct {
}

func (s *Service) Backup(httpClient *http.Client, args *BackupArgs) (*BackupResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Backup`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Backup: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Backup == nil {
		return nil, errors.New(`unexpected respose from service calling queue.Backup()`)
	}

	return r.Body.Backup, nil
}

type BrowseArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	QueueID        uint32 `xml:"QueueID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	RequestedCount uint32 `xml:"RequestedCount"`
}
type BrowseResponse struct {
	Result         string `xml:"Result"`
	NumberReturned uint32 `xml:"NumberReturned"`
	TotalMatches   uint32 `xml:"TotalMatches"`
	UpdateID       uint32 `xml:"UpdateID"`
}

func (s *Service) Browse(httpClient *http.Client, args *BrowseArgs) (*BrowseResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Browse`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Browse: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Browse == nil {
		return nil, errors.New(`unexpected respose from service calling queue.Browse()`)
	}

	return r.Body.Browse, nil
}

type CreateQueueArgs struct {
	Xmlns             string `xml:"xmlns:u,attr"`
	QueueOwnerID      string `xml:"QueueOwnerID"`
	QueueOwnerContext string `xml:"QueueOwnerContext"`
	QueuePolicy       string `xml:"QueuePolicy"`
}
type CreateQueueResponse struct {
	QueueID uint32 `xml:"QueueID"`
}

func (s *Service) CreateQueue(httpClient *http.Client, args *CreateQueueArgs) (*CreateQueueResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`CreateQueue`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{CreateQueue: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateQueue == nil {
		return nil, errors.New(`unexpected respose from service calling queue.CreateQueue()`)
	}

	return r.Body.CreateQueue, nil
}

type RemoveAllTracksArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	QueueID  uint32 `xml:"QueueID"`
	UpdateID uint32 `xml:"UpdateID"`
}
type RemoveAllTracksResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *Service) RemoveAllTracks(httpClient *http.Client, args *RemoveAllTracksArgs) (*RemoveAllTracksResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveAllTracks`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveAllTracks: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveAllTracks == nil {
		return nil, errors.New(`unexpected respose from service calling queue.RemoveAllTracks()`)
	}

	return r.Body.RemoveAllTracks, nil
}

type RemoveTrackRangeArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	QueueID        uint32 `xml:"QueueID"`
	UpdateID       uint32 `xml:"UpdateID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
}
type RemoveTrackRangeResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *Service) RemoveTrackRange(httpClient *http.Client, args *RemoveTrackRangeArgs) (*RemoveTrackRangeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveTrackRange`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveTrackRange: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveTrackRange == nil {
		return nil, errors.New(`unexpected respose from service calling queue.RemoveTrackRange()`)
	}

	return r.Body.RemoveTrackRange, nil
}

type ReorderTracksArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	QueueID        uint32 `xml:"QueueID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
	InsertBefore   uint32 `xml:"InsertBefore"`
	UpdateID       uint32 `xml:"UpdateID"`
}
type ReorderTracksResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *Service) ReorderTracks(httpClient *http.Client, args *ReorderTracksArgs) (*ReorderTracksResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReorderTracks`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReorderTracks: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReorderTracks == nil {
		return nil, errors.New(`unexpected respose from service calling queue.ReorderTracks()`)
	}

	return r.Body.ReorderTracks, nil
}

type ReplaceAllTracksArgs struct {
	Xmlns                   string `xml:"xmlns:u,attr"`
	QueueID                 uint32 `xml:"QueueID"`
	UpdateID                uint32 `xml:"UpdateID"`
	ContainerURI            string `xml:"ContainerURI"`
	ContainerMetaData       string `xml:"ContainerMetaData"`
	CurrentTrackIndex       uint32 `xml:"CurrentTrackIndex"`
	NewCurrentTrackIndices  string `xml:"NewCurrentTrackIndices"`
	NumberOfURIs            uint32 `xml:"NumberOfURIs"`
	EnqueuedURIsAndMetaData string `xml:"EnqueuedURIsAndMetaData"`
}
type ReplaceAllTracksResponse struct {
	NewQueueLength uint32 `xml:"NewQueueLength"`
	NewUpdateID    uint32 `xml:"NewUpdateID"`
}

func (s *Service) ReplaceAllTracks(httpClient *http.Client, args *ReplaceAllTracksArgs) (*ReplaceAllTracksResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReplaceAllTracks`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReplaceAllTracks: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReplaceAllTracks == nil {
		return nil, errors.New(`unexpected respose from service calling queue.ReplaceAllTracks()`)
	}

	return r.Body.ReplaceAllTracks, nil
}

type SaveAsSonosPlaylistArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	QueueID  uint32 `xml:"QueueID"`
	Title    string `xml:"Title"`
	ObjectID string `xml:"ObjectID"`
}
type SaveAsSonosPlaylistResponse struct {
	AssignedObjectID string `xml:"AssignedObjectID"`
}

func (s *Service) SaveAsSonosPlaylist(httpClient *http.Client, args *SaveAsSonosPlaylistArgs) (*SaveAsSonosPlaylistResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SaveAsSonosPlaylist`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SaveAsSonosPlaylist: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SaveAsSonosPlaylist == nil {
		return nil, errors.New(`unexpected respose from service calling queue.SaveAsSonosPlaylist()`)
	}

	return r.Body.SaveAsSonosPlaylist, nil
}
