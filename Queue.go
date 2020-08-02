package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type QueueService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewQueueService(deviceUrl *url.URL) *QueueService {
	c, _ := url.Parse("/MediaRenderer/Queue/Control")
	e, _ := url.Parse("/MediaRenderer/Queue/Event")
	return &QueueService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *QueueService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *QueueService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type QueueEnvelope struct {
	XMLName       xml.Name  `xml:"s:Envelope"`
	XMLNameSpace  string    `xml:"xmlns:s,attr"`
	EncodingStyle string    `xml:"s:encodingStyle,attr"`
	Body          QueueBody `xml:"s:Body"`
}
type QueueBody struct {
	XMLName             xml.Name                      `xml:"s:Body"`
	AddURI              *QueueAddURIArgs              `xml:"u:AddURI,omitempty"`
	AddMultipleURIs     *QueueAddMultipleURIsArgs     `xml:"u:AddMultipleURIs,omitempty"`
	AttachQueue         *QueueAttachQueueArgs         `xml:"u:AttachQueue,omitempty"`
	Backup              *QueueBackupArgs              `xml:"u:Backup,omitempty"`
	Browse              *QueueBrowseArgs              `xml:"u:Browse,omitempty"`
	CreateQueue         *QueueCreateQueueArgs         `xml:"u:CreateQueue,omitempty"`
	RemoveAllTracks     *QueueRemoveAllTracksArgs     `xml:"u:RemoveAllTracks,omitempty"`
	RemoveTrackRange    *QueueRemoveTrackRangeArgs    `xml:"u:RemoveTrackRange,omitempty"`
	ReorderTracks       *QueueReorderTracksArgs       `xml:"u:ReorderTracks,omitempty"`
	ReplaceAllTracks    *QueueReplaceAllTracksArgs    `xml:"u:ReplaceAllTracks,omitempty"`
	SaveAsSonosPlaylist *QueueSaveAsSonosPlaylistArgs `xml:"u:SaveAsSonosPlaylist,omitempty"`
}
type QueueEnvelopeResponse struct {
	XMLName       xml.Name          `xml:"Envelope"`
	XMLNameSpace  string            `xml:"xmlns:s,attr"`
	EncodingStyle string            `xml:"encodingStyle,attr"`
	Body          QueueBodyResponse `xml:"Body"`
}
type QueueBodyResponse struct {
	XMLName             xml.Name                          `xml:"Body"`
	AddURI              *QueueAddURIResponse              `xml:"AddURIResponse"`
	AddMultipleURIs     *QueueAddMultipleURIsResponse     `xml:"AddMultipleURIsResponse"`
	AttachQueue         *QueueAttachQueueResponse         `xml:"AttachQueueResponse"`
	Backup              *QueueBackupResponse              `xml:"BackupResponse"`
	Browse              *QueueBrowseResponse              `xml:"BrowseResponse"`
	CreateQueue         *QueueCreateQueueResponse         `xml:"CreateQueueResponse"`
	RemoveAllTracks     *QueueRemoveAllTracksResponse     `xml:"RemoveAllTracksResponse"`
	RemoveTrackRange    *QueueRemoveTrackRangeResponse    `xml:"RemoveTrackRangeResponse"`
	ReorderTracks       *QueueReorderTracksResponse       `xml:"ReorderTracksResponse"`
	ReplaceAllTracks    *QueueReplaceAllTracksResponse    `xml:"ReplaceAllTracksResponse"`
	SaveAsSonosPlaylist *QueueSaveAsSonosPlaylistResponse `xml:"SaveAsSonosPlaylistResponse"`
}

func (s *QueueService) _QueueExec(soapAction string, httpClient *http.Client, envelope *QueueEnvelope) (*QueueEnvelopeResponse, error) {
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
	var envelopeResponse QueueEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type QueueAddURIArgs struct {
	XMLNameSpace                    string `xml:"xmlns:u,attr"`
	QueueID                         uint32 `xml:"QueueID"`
	UpdateID                        uint32 `xml:"UpdateID"`
	EnqueuedURI                     string `xml:"EnqueuedURI"`
	EnqueuedURIMetaData             string `xml:"EnqueuedURIMetaData"`
	DesiredFirstTrackNumberEnqueued uint32 `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                   bool   `xml:"EnqueueAsNext"`
}
type QueueAddURIResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
	NewUpdateID              uint32 `xml:"NewUpdateID"`
}

func (s *QueueService) AddURI(httpClient *http.Client, args *QueueAddURIArgs) (*QueueAddURIResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#AddURI", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{AddURI: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddURI == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddURI, nil
}

type QueueAddMultipleURIsArgs struct {
	XMLNameSpace                    string `xml:"xmlns:u,attr"`
	QueueID                         uint32 `xml:"QueueID"`
	UpdateID                        uint32 `xml:"UpdateID"`
	ContainerURI                    string `xml:"ContainerURI"`
	ContainerMetaData               string `xml:"ContainerMetaData"`
	DesiredFirstTrackNumberEnqueued uint32 `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                   bool   `xml:"EnqueueAsNext"`
	NumberOfURIs                    uint32 `xml:"NumberOfURIs"`
	EnqueuedURIsAndMetaData         string `xml:"EnqueuedURIsAndMetaData"`
}
type QueueAddMultipleURIsResponse struct {
	FirstTrackNumberEnqueued uint32 `xml:"FirstTrackNumberEnqueued"`
	NumTracksAdded           uint32 `xml:"NumTracksAdded"`
	NewQueueLength           uint32 `xml:"NewQueueLength"`
	NewUpdateID              uint32 `xml:"NewUpdateID"`
}

func (s *QueueService) AddMultipleURIs(httpClient *http.Client, args *QueueAddMultipleURIsArgs) (*QueueAddMultipleURIsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#AddMultipleURIs", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{AddMultipleURIs: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddMultipleURIs == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddMultipleURIs, nil
}

type QueueAttachQueueArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	QueueOwnerID string `xml:"QueueOwnerID"`
}
type QueueAttachQueueResponse struct {
	QueueID           uint32 `xml:"QueueID"`
	QueueOwnerContext string `xml:"QueueOwnerContext"`
}

func (s *QueueService) AttachQueue(httpClient *http.Client, args *QueueAttachQueueArgs) (*QueueAttachQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#AttachQueue", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{AttachQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AttachQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AttachQueue, nil
}

type QueueBackupArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type QueueBackupResponse struct {
}

func (s *QueueService) Backup(httpClient *http.Client, args *QueueBackupArgs) (*QueueBackupResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#Backup", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{Backup: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Backup == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Backup, nil
}

type QueueBrowseArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
	QueueID        uint32 `xml:"QueueID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	RequestedCount uint32 `xml:"RequestedCount"`
}
type QueueBrowseResponse struct {
	Result         string `xml:"Result"`
	NumberReturned uint32 `xml:"NumberReturned"`
	TotalMatches   uint32 `xml:"TotalMatches"`
	UpdateID       uint32 `xml:"UpdateID"`
}

func (s *QueueService) Browse(httpClient *http.Client, args *QueueBrowseArgs) (*QueueBrowseResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#Browse", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{Browse: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Browse == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Browse, nil
}

type QueueCreateQueueArgs struct {
	XMLNameSpace      string `xml:"xmlns:u,attr"`
	QueueOwnerID      string `xml:"QueueOwnerID"`
	QueueOwnerContext string `xml:"QueueOwnerContext"`
	QueuePolicy       string `xml:"QueuePolicy"`
}
type QueueCreateQueueResponse struct {
	QueueID uint32 `xml:"QueueID"`
}

func (s *QueueService) CreateQueue(httpClient *http.Client, args *QueueCreateQueueArgs) (*QueueCreateQueueResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#CreateQueue", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{CreateQueue: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateQueue == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.CreateQueue, nil
}

type QueueRemoveAllTracksArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	QueueID      uint32 `xml:"QueueID"`
	UpdateID     uint32 `xml:"UpdateID"`
}
type QueueRemoveAllTracksResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *QueueService) RemoveAllTracks(httpClient *http.Client, args *QueueRemoveAllTracksArgs) (*QueueRemoveAllTracksResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#RemoveAllTracks", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{RemoveAllTracks: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveAllTracks == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveAllTracks, nil
}

type QueueRemoveTrackRangeArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
	QueueID        uint32 `xml:"QueueID"`
	UpdateID       uint32 `xml:"UpdateID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
}
type QueueRemoveTrackRangeResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *QueueService) RemoveTrackRange(httpClient *http.Client, args *QueueRemoveTrackRangeArgs) (*QueueRemoveTrackRangeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#RemoveTrackRange", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{RemoveTrackRange: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveTrackRange == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveTrackRange, nil
}

type QueueReorderTracksArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
	QueueID        uint32 `xml:"QueueID"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	NumberOfTracks uint32 `xml:"NumberOfTracks"`
	InsertBefore   uint32 `xml:"InsertBefore"`
	UpdateID       uint32 `xml:"UpdateID"`
}
type QueueReorderTracksResponse struct {
	NewUpdateID uint32 `xml:"NewUpdateID"`
}

func (s *QueueService) ReorderTracks(httpClient *http.Client, args *QueueReorderTracksArgs) (*QueueReorderTracksResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#ReorderTracks", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{ReorderTracks: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReorderTracks == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReorderTracks, nil
}

type QueueReplaceAllTracksArgs struct {
	XMLNameSpace            string `xml:"xmlns:u,attr"`
	QueueID                 uint32 `xml:"QueueID"`
	UpdateID                uint32 `xml:"UpdateID"`
	ContainerURI            string `xml:"ContainerURI"`
	ContainerMetaData       string `xml:"ContainerMetaData"`
	CurrentTrackIndex       uint32 `xml:"CurrentTrackIndex"`
	NewCurrentTrackIndices  string `xml:"NewCurrentTrackIndices"`
	NumberOfURIs            uint32 `xml:"NumberOfURIs"`
	EnqueuedURIsAndMetaData string `xml:"EnqueuedURIsAndMetaData"`
}
type QueueReplaceAllTracksResponse struct {
	NewQueueLength uint32 `xml:"NewQueueLength"`
	NewUpdateID    uint32 `xml:"NewUpdateID"`
}

func (s *QueueService) ReplaceAllTracks(httpClient *http.Client, args *QueueReplaceAllTracksArgs) (*QueueReplaceAllTracksResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#ReplaceAllTracks", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{ReplaceAllTracks: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReplaceAllTracks == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReplaceAllTracks, nil
}

type QueueSaveAsSonosPlaylistArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	QueueID      uint32 `xml:"QueueID"`
	Title        string `xml:"Title"`
	ObjectID     string `xml:"ObjectID"`
}
type QueueSaveAsSonosPlaylistResponse struct {
	AssignedObjectID string `xml:"AssignedObjectID"`
}

func (s *QueueService) SaveAsSonosPlaylist(httpClient *http.Client, args *QueueSaveAsSonosPlaylistArgs) (*QueueSaveAsSonosPlaylistResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:Queue:1"
	r, err := s._QueueExec("urn:schemas-upnp-org:service:Queue:1#SaveAsSonosPlaylist", httpClient,
		&QueueEnvelope{
			Body:          QueueBody{SaveAsSonosPlaylist: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SaveAsSonosPlaylist == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SaveAsSonosPlaylist, nil
}

// Events
type QueueLastChange string
type QueueUpnpEvent struct {
	XMLName      xml.Name        `xml:"propertyset"`
	XMLNameSpace string          `xml:"xmlns:e,attr"`
	Properties   []QueueProperty `xml:"property"`
}
type QueueProperty struct {
	XMLName    xml.Name         `xml:"property"`
	LastChange *QueueLastChange `xml:"LastChange"`
}

func QueueDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt QueueUpnpEvent
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
