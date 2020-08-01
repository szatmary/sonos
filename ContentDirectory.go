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

type ContentDirectoryService struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewContentDirectoryService(deviceUrl *url.URL) *ContentDirectoryService {
	c, _ := url.Parse("/MediaServer/ContentDirectory/Control")
	e, _ := url.Parse("/MediaServer/ContentDirectory/Event")
	return &ContentDirectoryService{
		ControlEndpoint: deviceUrl.ResolveReference(c),
		EventEndpoint:   deviceUrl.ResolveReference(e),
	}
}

type ContentDirectoryEnvelope struct {
	XMLName       xml.Name             `xml:"s:Envelope"`
	XMLNameSpace  string               `xml:"xmlns:s,attr"`
	EncodingStyle string               `xml:"s:encodingStyle,attr"`
	Body          ContentDirectoryBody `xml:"s:Body"`
}
type ContentDirectoryBody struct {
	XMLName                     xml.Name                                         `xml:"s:Body"`
	GetSearchCapabilities       *ContentDirectoryGetSearchCapabilitiesArgs       `xml:"u:GetSearchCapabilities,omitempty"`
	GetSortCapabilities         *ContentDirectoryGetSortCapabilitiesArgs         `xml:"u:GetSortCapabilities,omitempty"`
	GetSystemUpdateID           *ContentDirectoryGetSystemUpdateIDArgs           `xml:"u:GetSystemUpdateID,omitempty"`
	GetAlbumArtistDisplayOption *ContentDirectoryGetAlbumArtistDisplayOptionArgs `xml:"u:GetAlbumArtistDisplayOption,omitempty"`
	GetLastIndexChange          *ContentDirectoryGetLastIndexChangeArgs          `xml:"u:GetLastIndexChange,omitempty"`
	Browse                      *ContentDirectoryBrowseArgs                      `xml:"u:Browse,omitempty"`
	FindPrefix                  *ContentDirectoryFindPrefixArgs                  `xml:"u:FindPrefix,omitempty"`
	GetAllPrefixLocations       *ContentDirectoryGetAllPrefixLocationsArgs       `xml:"u:GetAllPrefixLocations,omitempty"`
	CreateObject                *ContentDirectoryCreateObjectArgs                `xml:"u:CreateObject,omitempty"`
	UpdateObject                *ContentDirectoryUpdateObjectArgs                `xml:"u:UpdateObject,omitempty"`
	DestroyObject               *ContentDirectoryDestroyObjectArgs               `xml:"u:DestroyObject,omitempty"`
	RefreshShareIndex           *ContentDirectoryRefreshShareIndexArgs           `xml:"u:RefreshShareIndex,omitempty"`
	RequestResort               *ContentDirectoryRequestResortArgs               `xml:"u:RequestResort,omitempty"`
	GetShareIndexInProgress     *ContentDirectoryGetShareIndexInProgressArgs     `xml:"u:GetShareIndexInProgress,omitempty"`
	GetBrowseable               *ContentDirectoryGetBrowseableArgs               `xml:"u:GetBrowseable,omitempty"`
	SetBrowseable               *ContentDirectorySetBrowseableArgs               `xml:"u:SetBrowseable,omitempty"`
}
type ContentDirectoryEnvelopeResponse struct {
	XMLName       xml.Name                     `xml:"Envelope"`
	XMLNameSpace  string                       `xml:"xmlns:s,attr"`
	EncodingStyle string                       `xml:"encodingStyle,attr"`
	Body          ContentDirectoryBodyResponse `xml:"Body"`
}
type ContentDirectoryBodyResponse struct {
	XMLName                     xml.Name                                             `xml:"Body"`
	GetSearchCapabilities       *ContentDirectoryGetSearchCapabilitiesResponse       `xml:"GetSearchCapabilitiesResponse"`
	GetSortCapabilities         *ContentDirectoryGetSortCapabilitiesResponse         `xml:"GetSortCapabilitiesResponse"`
	GetSystemUpdateID           *ContentDirectoryGetSystemUpdateIDResponse           `xml:"GetSystemUpdateIDResponse"`
	GetAlbumArtistDisplayOption *ContentDirectoryGetAlbumArtistDisplayOptionResponse `xml:"GetAlbumArtistDisplayOptionResponse"`
	GetLastIndexChange          *ContentDirectoryGetLastIndexChangeResponse          `xml:"GetLastIndexChangeResponse"`
	Browse                      *ContentDirectoryBrowseResponse                      `xml:"BrowseResponse"`
	FindPrefix                  *ContentDirectoryFindPrefixResponse                  `xml:"FindPrefixResponse"`
	GetAllPrefixLocations       *ContentDirectoryGetAllPrefixLocationsResponse       `xml:"GetAllPrefixLocationsResponse"`
	CreateObject                *ContentDirectoryCreateObjectResponse                `xml:"CreateObjectResponse"`
	UpdateObject                *ContentDirectoryUpdateObjectResponse                `xml:"UpdateObjectResponse"`
	DestroyObject               *ContentDirectoryDestroyObjectResponse               `xml:"DestroyObjectResponse"`
	RefreshShareIndex           *ContentDirectoryRefreshShareIndexResponse           `xml:"RefreshShareIndexResponse"`
	RequestResort               *ContentDirectoryRequestResortResponse               `xml:"RequestResortResponse"`
	GetShareIndexInProgress     *ContentDirectoryGetShareIndexInProgressResponse     `xml:"GetShareIndexInProgressResponse"`
	GetBrowseable               *ContentDirectoryGetBrowseableResponse               `xml:"GetBrowseableResponse"`
	SetBrowseable               *ContentDirectorySetBrowseableResponse               `xml:"SetBrowseableResponse"`
}

func (s *ContentDirectoryService) _ContentDirectoryExec(soapAction string, httpClient *http.Client, envelope *ContentDirectoryEnvelope) (*ContentDirectoryEnvelopeResponse, error) {
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
	var envelopeResponse ContentDirectoryEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type ContentDirectoryGetSearchCapabilitiesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ContentDirectoryGetSearchCapabilitiesResponse struct {
	SearchCaps string `xml:"SearchCaps"`
}

func (s *ContentDirectoryService) GetSearchCapabilities(httpClient *http.Client, args *ContentDirectoryGetSearchCapabilitiesArgs) (*ContentDirectoryGetSearchCapabilitiesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetSearchCapabilities", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetSearchCapabilities: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSearchCapabilities == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetSearchCapabilities, nil
}

type ContentDirectoryGetSortCapabilitiesArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ContentDirectoryGetSortCapabilitiesResponse struct {
	SortCaps string `xml:"SortCaps"`
}

func (s *ContentDirectoryService) GetSortCapabilities(httpClient *http.Client, args *ContentDirectoryGetSortCapabilitiesArgs) (*ContentDirectoryGetSortCapabilitiesResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetSortCapabilities", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetSortCapabilities: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSortCapabilities == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetSortCapabilities, nil
}

type ContentDirectoryGetSystemUpdateIDArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ContentDirectoryGetSystemUpdateIDResponse struct {
	Id uint32 `xml:"Id"`
}

func (s *ContentDirectoryService) GetSystemUpdateID(httpClient *http.Client, args *ContentDirectoryGetSystemUpdateIDArgs) (*ContentDirectoryGetSystemUpdateIDResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetSystemUpdateID", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetSystemUpdateID: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSystemUpdateID == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetSystemUpdateID, nil
}

type ContentDirectoryGetAlbumArtistDisplayOptionArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ContentDirectoryGetAlbumArtistDisplayOptionResponse struct {
	AlbumArtistDisplayOption string `xml:"AlbumArtistDisplayOption"`
}

func (s *ContentDirectoryService) GetAlbumArtistDisplayOption(httpClient *http.Client, args *ContentDirectoryGetAlbumArtistDisplayOptionArgs) (*ContentDirectoryGetAlbumArtistDisplayOptionResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetAlbumArtistDisplayOption", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetAlbumArtistDisplayOption: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAlbumArtistDisplayOption == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetAlbumArtistDisplayOption, nil
}

type ContentDirectoryGetLastIndexChangeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ContentDirectoryGetLastIndexChangeResponse struct {
	LastIndexChange string `xml:"LastIndexChange"`
}

func (s *ContentDirectoryService) GetLastIndexChange(httpClient *http.Client, args *ContentDirectoryGetLastIndexChangeArgs) (*ContentDirectoryGetLastIndexChangeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetLastIndexChange", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetLastIndexChange: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetLastIndexChange == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetLastIndexChange, nil
}

type ContentDirectoryBrowseArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ObjectID     string `xml:"ObjectID"`
	// Allowed Value: BrowseMetadata
	// Allowed Value: BrowseDirectChildren
	BrowseFlag     string `xml:"BrowseFlag"`
	Filter         string `xml:"Filter"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	RequestedCount uint32 `xml:"RequestedCount"`
	SortCriteria   string `xml:"SortCriteria"`
}
type ContentDirectoryBrowseResponse struct {
	Result         string `xml:"Result"`
	NumberReturned uint32 `xml:"NumberReturned"`
	TotalMatches   uint32 `xml:"TotalMatches"`
	UpdateID       uint32 `xml:"UpdateID"`
}

func (s *ContentDirectoryService) Browse(httpClient *http.Client, args *ContentDirectoryBrowseArgs) (*ContentDirectoryBrowseResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#Browse", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{Browse: args},
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

type ContentDirectoryFindPrefixArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ObjectID     string `xml:"ObjectID"`
	Prefix       string `xml:"Prefix"`
}
type ContentDirectoryFindPrefixResponse struct {
	StartingIndex uint32 `xml:"StartingIndex"`
	UpdateID      uint32 `xml:"UpdateID"`
}

func (s *ContentDirectoryService) FindPrefix(httpClient *http.Client, args *ContentDirectoryFindPrefixArgs) (*ContentDirectoryFindPrefixResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#FindPrefix", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{FindPrefix: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.FindPrefix == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.FindPrefix, nil
}

type ContentDirectoryGetAllPrefixLocationsArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ObjectID     string `xml:"ObjectID"`
}
type ContentDirectoryGetAllPrefixLocationsResponse struct {
	TotalPrefixes     uint32 `xml:"TotalPrefixes"`
	PrefixAndIndexCSV string `xml:"PrefixAndIndexCSV"`
	UpdateID          uint32 `xml:"UpdateID"`
}

func (s *ContentDirectoryService) GetAllPrefixLocations(httpClient *http.Client, args *ContentDirectoryGetAllPrefixLocationsArgs) (*ContentDirectoryGetAllPrefixLocationsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetAllPrefixLocations", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetAllPrefixLocations: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAllPrefixLocations == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetAllPrefixLocations, nil
}

type ContentDirectoryCreateObjectArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ContainerID  string `xml:"ContainerID"`
	Elements     string `xml:"Elements"`
}
type ContentDirectoryCreateObjectResponse struct {
	ObjectID string `xml:"ObjectID"`
	Result   string `xml:"Result"`
}

func (s *ContentDirectoryService) CreateObject(httpClient *http.Client, args *ContentDirectoryCreateObjectArgs) (*ContentDirectoryCreateObjectResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#CreateObject", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{CreateObject: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateObject == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.CreateObject, nil
}

type ContentDirectoryUpdateObjectArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	ObjectID        string `xml:"ObjectID"`
	CurrentTagValue string `xml:"CurrentTagValue"`
	NewTagValue     string `xml:"NewTagValue"`
}
type ContentDirectoryUpdateObjectResponse struct {
}

func (s *ContentDirectoryService) UpdateObject(httpClient *http.Client, args *ContentDirectoryUpdateObjectArgs) (*ContentDirectoryUpdateObjectResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#UpdateObject", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{UpdateObject: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.UpdateObject == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.UpdateObject, nil
}

type ContentDirectoryDestroyObjectArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ObjectID     string `xml:"ObjectID"`
}
type ContentDirectoryDestroyObjectResponse struct {
}

func (s *ContentDirectoryService) DestroyObject(httpClient *http.Client, args *ContentDirectoryDestroyObjectArgs) (*ContentDirectoryDestroyObjectResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#DestroyObject", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{DestroyObject: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DestroyObject == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.DestroyObject, nil
}

type ContentDirectoryRefreshShareIndexArgs struct {
	XMLNameSpace             string `xml:"xmlns:u,attr"`
	AlbumArtistDisplayOption string `xml:"AlbumArtistDisplayOption"`
}
type ContentDirectoryRefreshShareIndexResponse struct {
}

func (s *ContentDirectoryService) RefreshShareIndex(httpClient *http.Client, args *ContentDirectoryRefreshShareIndexArgs) (*ContentDirectoryRefreshShareIndexResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#RefreshShareIndex", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{RefreshShareIndex: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RefreshShareIndex == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RefreshShareIndex, nil
}

type ContentDirectoryRequestResortArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	SortOrder    string `xml:"SortOrder"`
}
type ContentDirectoryRequestResortResponse struct {
}

func (s *ContentDirectoryService) RequestResort(httpClient *http.Client, args *ContentDirectoryRequestResortArgs) (*ContentDirectoryRequestResortResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#RequestResort", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{RequestResort: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RequestResort == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RequestResort, nil
}

type ContentDirectoryGetShareIndexInProgressArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ContentDirectoryGetShareIndexInProgressResponse struct {
	IsIndexing bool `xml:"IsIndexing"`
}

func (s *ContentDirectoryService) GetShareIndexInProgress(httpClient *http.Client, args *ContentDirectoryGetShareIndexInProgressArgs) (*ContentDirectoryGetShareIndexInProgressResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetShareIndexInProgress", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetShareIndexInProgress: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetShareIndexInProgress == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetShareIndexInProgress, nil
}

type ContentDirectoryGetBrowseableArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type ContentDirectoryGetBrowseableResponse struct {
	IsBrowseable bool `xml:"IsBrowseable"`
}

func (s *ContentDirectoryService) GetBrowseable(httpClient *http.Client, args *ContentDirectoryGetBrowseableArgs) (*ContentDirectoryGetBrowseableResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#GetBrowseable", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{GetBrowseable: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetBrowseable == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetBrowseable, nil
}

type ContentDirectorySetBrowseableArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Browseable   bool   `xml:"Browseable"`
}
type ContentDirectorySetBrowseableResponse struct {
}

func (s *ContentDirectoryService) SetBrowseable(httpClient *http.Client, args *ContentDirectorySetBrowseableArgs) (*ContentDirectorySetBrowseableResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:ContentDirectory:1"
	r, err := s._ContentDirectoryExec("urn:schemas-upnp-org:service:ContentDirectory:1#SetBrowseable", httpClient,
		&ContentDirectoryEnvelope{
			Body:          ContentDirectoryBody{SetBrowseable: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetBrowseable == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetBrowseable, nil
}
func (s *ContentDirectoryService) ContentDirectorySubscribe(callback url.URL) error {
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
type ContentDirectorySystemUpdateID uint32
type ContentDirectoryContainerUpdateIDs string
type ContentDirectoryShareIndexInProgress bool
type ContentDirectoryShareIndexLastError string
type ContentDirectoryUserRadioUpdateID string
type ContentDirectorySavedQueuesUpdateID string
type ContentDirectoryShareListUpdateID string
type ContentDirectoryRecentlyPlayedUpdateID string
type ContentDirectoryBrowseable bool
type ContentDirectoryRadioFavoritesUpdateID uint32
type ContentDirectoryRadioLocationUpdateID uint32
type ContentDirectoryFavoritesUpdateID string
type ContentDirectoryFavoritePresetsUpdateID string
type ContentDirectoryUpnpEvent struct {
	XMLName      xml.Name                   `xml:"propertyset"`
	XMLNameSpace string                     `xml:"xmlns:e,attr"`
	Properties   []ContentDirectoryProperty `xml:"property"`
}
type ContentDirectoryProperty struct {
	XMLName                 xml.Name                                 `xml:"property"`
	SystemUpdateID          *ContentDirectorySystemUpdateID          `xml:"SystemUpdateID"`
	ContainerUpdateIDs      *ContentDirectoryContainerUpdateIDs      `xml:"ContainerUpdateIDs"`
	ShareIndexInProgress    *ContentDirectoryShareIndexInProgress    `xml:"ShareIndexInProgress"`
	ShareIndexLastError     *ContentDirectoryShareIndexLastError     `xml:"ShareIndexLastError"`
	UserRadioUpdateID       *ContentDirectoryUserRadioUpdateID       `xml:"UserRadioUpdateID"`
	SavedQueuesUpdateID     *ContentDirectorySavedQueuesUpdateID     `xml:"SavedQueuesUpdateID"`
	ShareListUpdateID       *ContentDirectoryShareListUpdateID       `xml:"ShareListUpdateID"`
	RecentlyPlayedUpdateID  *ContentDirectoryRecentlyPlayedUpdateID  `xml:"RecentlyPlayedUpdateID"`
	Browseable              *ContentDirectoryBrowseable              `xml:"Browseable"`
	RadioFavoritesUpdateID  *ContentDirectoryRadioFavoritesUpdateID  `xml:"RadioFavoritesUpdateID"`
	RadioLocationUpdateID   *ContentDirectoryRadioLocationUpdateID   `xml:"RadioLocationUpdateID"`
	FavoritesUpdateID       *ContentDirectoryFavoritesUpdateID       `xml:"FavoritesUpdateID"`
	FavoritePresetsUpdateID *ContentDirectoryFavoritePresetsUpdateID `xml:"FavoritePresetsUpdateID"`
}

func ContentDirectoryDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt ContentDirectoryUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.SystemUpdateID != nil:
			zp.EventCallback(*prop.SystemUpdateID)
		case prop.ContainerUpdateIDs != nil:
			zp.EventCallback(*prop.ContainerUpdateIDs)
		case prop.ShareIndexInProgress != nil:
			zp.EventCallback(*prop.ShareIndexInProgress)
		case prop.ShareIndexLastError != nil:
			zp.EventCallback(*prop.ShareIndexLastError)
		case prop.UserRadioUpdateID != nil:
			zp.EventCallback(*prop.UserRadioUpdateID)
		case prop.SavedQueuesUpdateID != nil:
			zp.EventCallback(*prop.SavedQueuesUpdateID)
		case prop.ShareListUpdateID != nil:
			zp.EventCallback(*prop.ShareListUpdateID)
		case prop.RecentlyPlayedUpdateID != nil:
			zp.EventCallback(*prop.RecentlyPlayedUpdateID)
		case prop.Browseable != nil:
			zp.EventCallback(*prop.Browseable)
		case prop.RadioFavoritesUpdateID != nil:
			zp.EventCallback(*prop.RadioFavoritesUpdateID)
		case prop.RadioLocationUpdateID != nil:
			zp.EventCallback(*prop.RadioLocationUpdateID)
		case prop.FavoritesUpdateID != nil:
			zp.EventCallback(*prop.FavoritesUpdateID)
		case prop.FavoritePresetsUpdateID != nil:
			zp.EventCallback(*prop.FavoritePresetsUpdateID)
		}
	}
}
