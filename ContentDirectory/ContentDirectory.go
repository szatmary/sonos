package contentdirectory

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:ContentDirectory:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/MediaServer/ContentDirectory/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/MediaServer/ContentDirectory/Event`)
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
	XMLName                     xml.Name                         `xml:"s:Body"`
	GetSearchCapabilities       *GetSearchCapabilitiesArgs       `xml:"u:GetSearchCapabilities,omitempty"`
	GetSortCapabilities         *GetSortCapabilitiesArgs         `xml:"u:GetSortCapabilities,omitempty"`
	GetSystemUpdateID           *GetSystemUpdateIDArgs           `xml:"u:GetSystemUpdateID,omitempty"`
	GetAlbumArtistDisplayOption *GetAlbumArtistDisplayOptionArgs `xml:"u:GetAlbumArtistDisplayOption,omitempty"`
	GetLastIndexChange          *GetLastIndexChangeArgs          `xml:"u:GetLastIndexChange,omitempty"`
	Browse                      *BrowseArgs                      `xml:"u:Browse,omitempty"`
	FindPrefix                  *FindPrefixArgs                  `xml:"u:FindPrefix,omitempty"`
	GetAllPrefixLocations       *GetAllPrefixLocationsArgs       `xml:"u:GetAllPrefixLocations,omitempty"`
	CreateObject                *CreateObjectArgs                `xml:"u:CreateObject,omitempty"`
	UpdateObject                *UpdateObjectArgs                `xml:"u:UpdateObject,omitempty"`
	DestroyObject               *DestroyObjectArgs               `xml:"u:DestroyObject,omitempty"`
	RefreshShareIndex           *RefreshShareIndexArgs           `xml:"u:RefreshShareIndex,omitempty"`
	RequestResort               *RequestResortArgs               `xml:"u:RequestResort,omitempty"`
	GetShareIndexInProgress     *GetShareIndexInProgressArgs     `xml:"u:GetShareIndexInProgress,omitempty"`
	GetBrowseable               *GetBrowseableArgs               `xml:"u:GetBrowseable,omitempty"`
	SetBrowseable               *SetBrowseableArgs               `xml:"u:SetBrowseable,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                     xml.Name                             `xml:"Body"`
	GetSearchCapabilities       *GetSearchCapabilitiesResponse       `xml:"GetSearchCapabilitiesResponse,omitempty"`
	GetSortCapabilities         *GetSortCapabilitiesResponse         `xml:"GetSortCapabilitiesResponse,omitempty"`
	GetSystemUpdateID           *GetSystemUpdateIDResponse           `xml:"GetSystemUpdateIDResponse,omitempty"`
	GetAlbumArtistDisplayOption *GetAlbumArtistDisplayOptionResponse `xml:"GetAlbumArtistDisplayOptionResponse,omitempty"`
	GetLastIndexChange          *GetLastIndexChangeResponse          `xml:"GetLastIndexChangeResponse,omitempty"`
	Browse                      *BrowseResponse                      `xml:"BrowseResponse,omitempty"`
	FindPrefix                  *FindPrefixResponse                  `xml:"FindPrefixResponse,omitempty"`
	GetAllPrefixLocations       *GetAllPrefixLocationsResponse       `xml:"GetAllPrefixLocationsResponse,omitempty"`
	CreateObject                *CreateObjectResponse                `xml:"CreateObjectResponse,omitempty"`
	UpdateObject                *UpdateObjectResponse                `xml:"UpdateObjectResponse,omitempty"`
	DestroyObject               *DestroyObjectResponse               `xml:"DestroyObjectResponse,omitempty"`
	RefreshShareIndex           *RefreshShareIndexResponse           `xml:"RefreshShareIndexResponse,omitempty"`
	RequestResort               *RequestResortResponse               `xml:"RequestResortResponse,omitempty"`
	GetShareIndexInProgress     *GetShareIndexInProgressResponse     `xml:"GetShareIndexInProgressResponse,omitempty"`
	GetBrowseable               *GetBrowseableResponse               `xml:"GetBrowseableResponse,omitempty"`
	SetBrowseable               *SetBrowseableResponse               `xml:"SetBrowseableResponse,omitempty"`
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

type GetSearchCapabilitiesArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetSearchCapabilitiesResponse struct {
	SearchCaps string `xml:"SearchCaps"`
}

func (s *Service) GetSearchCapabilities(httpClient *http.Client, args *GetSearchCapabilitiesArgs) (*GetSearchCapabilitiesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetSearchCapabilities`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetSearchCapabilities: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSearchCapabilities == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetSearchCapabilities()`)
	}

	return r.Body.GetSearchCapabilities, nil
}

type GetSortCapabilitiesArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetSortCapabilitiesResponse struct {
	SortCaps string `xml:"SortCaps"`
}

func (s *Service) GetSortCapabilities(httpClient *http.Client, args *GetSortCapabilitiesArgs) (*GetSortCapabilitiesResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetSortCapabilities`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetSortCapabilities: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSortCapabilities == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetSortCapabilities()`)
	}

	return r.Body.GetSortCapabilities, nil
}

type GetSystemUpdateIDArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetSystemUpdateIDResponse struct {
	Id uint32 `xml:"Id"`
}

func (s *Service) GetSystemUpdateID(httpClient *http.Client, args *GetSystemUpdateIDArgs) (*GetSystemUpdateIDResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetSystemUpdateID`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetSystemUpdateID: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetSystemUpdateID == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetSystemUpdateID()`)
	}

	return r.Body.GetSystemUpdateID, nil
}

type GetAlbumArtistDisplayOptionArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetAlbumArtistDisplayOptionResponse struct {
	AlbumArtistDisplayOption string `xml:"AlbumArtistDisplayOption"`
}

func (s *Service) GetAlbumArtistDisplayOption(httpClient *http.Client, args *GetAlbumArtistDisplayOptionArgs) (*GetAlbumArtistDisplayOptionResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetAlbumArtistDisplayOption`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetAlbumArtistDisplayOption: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAlbumArtistDisplayOption == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetAlbumArtistDisplayOption()`)
	}

	return r.Body.GetAlbumArtistDisplayOption, nil
}

type GetLastIndexChangeArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetLastIndexChangeResponse struct {
	LastIndexChange string `xml:"LastIndexChange"`
}

func (s *Service) GetLastIndexChange(httpClient *http.Client, args *GetLastIndexChangeArgs) (*GetLastIndexChangeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetLastIndexChange`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetLastIndexChange: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetLastIndexChange == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetLastIndexChange()`)
	}

	return r.Body.GetLastIndexChange, nil
}

type BrowseArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	ObjectID string `xml:"ObjectID"`
	// Allowed Value: BrowseMetadata
	// Allowed Value: BrowseDirectChildren
	BrowseFlag     string `xml:"BrowseFlag"`
	Filter         string `xml:"Filter"`
	StartingIndex  uint32 `xml:"StartingIndex"`
	RequestedCount uint32 `xml:"RequestedCount"`
	SortCriteria   string `xml:"SortCriteria"`
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
		return nil, errors.New(`unexpected respose from service calling contentdirectory.Browse()`)
	}

	return r.Body.Browse, nil
}

type FindPrefixArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	ObjectID string `xml:"ObjectID"`
	Prefix   string `xml:"Prefix"`
}
type FindPrefixResponse struct {
	StartingIndex uint32 `xml:"StartingIndex"`
	UpdateID      uint32 `xml:"UpdateID"`
}

func (s *Service) FindPrefix(httpClient *http.Client, args *FindPrefixArgs) (*FindPrefixResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`FindPrefix`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{FindPrefix: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.FindPrefix == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.FindPrefix()`)
	}

	return r.Body.FindPrefix, nil
}

type GetAllPrefixLocationsArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	ObjectID string `xml:"ObjectID"`
}
type GetAllPrefixLocationsResponse struct {
	TotalPrefixes     uint32 `xml:"TotalPrefixes"`
	PrefixAndIndexCSV string `xml:"PrefixAndIndexCSV"`
	UpdateID          uint32 `xml:"UpdateID"`
}

func (s *Service) GetAllPrefixLocations(httpClient *http.Client, args *GetAllPrefixLocationsArgs) (*GetAllPrefixLocationsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetAllPrefixLocations`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetAllPrefixLocations: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetAllPrefixLocations == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetAllPrefixLocations()`)
	}

	return r.Body.GetAllPrefixLocations, nil
}

type CreateObjectArgs struct {
	Xmlns       string `xml:"xmlns:u,attr"`
	ContainerID string `xml:"ContainerID"`
	Elements    string `xml:"Elements"`
}
type CreateObjectResponse struct {
	ObjectID string `xml:"ObjectID"`
	Result   string `xml:"Result"`
}

func (s *Service) CreateObject(httpClient *http.Client, args *CreateObjectArgs) (*CreateObjectResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`CreateObject`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{CreateObject: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateObject == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.CreateObject()`)
	}

	return r.Body.CreateObject, nil
}

type UpdateObjectArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
	ObjectID        string `xml:"ObjectID"`
	CurrentTagValue string `xml:"CurrentTagValue"`
	NewTagValue     string `xml:"NewTagValue"`
}
type UpdateObjectResponse struct {
}

func (s *Service) UpdateObject(httpClient *http.Client, args *UpdateObjectArgs) (*UpdateObjectResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`UpdateObject`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{UpdateObject: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.UpdateObject == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.UpdateObject()`)
	}

	return r.Body.UpdateObject, nil
}

type DestroyObjectArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	ObjectID string `xml:"ObjectID"`
}
type DestroyObjectResponse struct {
}

func (s *Service) DestroyObject(httpClient *http.Client, args *DestroyObjectArgs) (*DestroyObjectResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`DestroyObject`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{DestroyObject: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DestroyObject == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.DestroyObject()`)
	}

	return r.Body.DestroyObject, nil
}

type RefreshShareIndexArgs struct {
	Xmlns                    string `xml:"xmlns:u,attr"`
	AlbumArtistDisplayOption string `xml:"AlbumArtistDisplayOption"`
}
type RefreshShareIndexResponse struct {
}

func (s *Service) RefreshShareIndex(httpClient *http.Client, args *RefreshShareIndexArgs) (*RefreshShareIndexResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RefreshShareIndex`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RefreshShareIndex: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RefreshShareIndex == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.RefreshShareIndex()`)
	}

	return r.Body.RefreshShareIndex, nil
}

type RequestResortArgs struct {
	Xmlns     string `xml:"xmlns:u,attr"`
	SortOrder string `xml:"SortOrder"`
}
type RequestResortResponse struct {
}

func (s *Service) RequestResort(httpClient *http.Client, args *RequestResortArgs) (*RequestResortResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RequestResort`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RequestResort: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RequestResort == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.RequestResort()`)
	}

	return r.Body.RequestResort, nil
}

type GetShareIndexInProgressArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetShareIndexInProgressResponse struct {
	IsIndexing bool `xml:"IsIndexing"`
}

func (s *Service) GetShareIndexInProgress(httpClient *http.Client, args *GetShareIndexInProgressArgs) (*GetShareIndexInProgressResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetShareIndexInProgress`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetShareIndexInProgress: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetShareIndexInProgress == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetShareIndexInProgress()`)
	}

	return r.Body.GetShareIndexInProgress, nil
}

type GetBrowseableArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetBrowseableResponse struct {
	IsBrowseable bool `xml:"IsBrowseable"`
}

func (s *Service) GetBrowseable(httpClient *http.Client, args *GetBrowseableArgs) (*GetBrowseableResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetBrowseable`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetBrowseable: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetBrowseable == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.GetBrowseable()`)
	}

	return r.Body.GetBrowseable, nil
}

type SetBrowseableArgs struct {
	Xmlns      string `xml:"xmlns:u,attr"`
	Browseable bool   `xml:"Browseable"`
}
type SetBrowseableResponse struct {
}

func (s *Service) SetBrowseable(httpClient *http.Client, args *SetBrowseableArgs) (*SetBrowseableResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetBrowseable`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetBrowseable: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetBrowseable == nil {
		return nil, errors.New(`unexpected respose from service calling contentdirectory.SetBrowseable()`)
	}

	return r.Body.SetBrowseable, nil
}
