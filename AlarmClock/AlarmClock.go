package alarmclock

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:AlarmClock:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/AlarmClock/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/AlarmClock/Event`)
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
	XMLName                  xml.Name                      `xml:"s:Body"`
	SetFormat                *SetFormatArgs                `xml:"u:SetFormat,omitempty"`
	GetFormat                *GetFormatArgs                `xml:"u:GetFormat,omitempty"`
	SetTimeZone              *SetTimeZoneArgs              `xml:"u:SetTimeZone,omitempty"`
	GetTimeZone              *GetTimeZoneArgs              `xml:"u:GetTimeZone,omitempty"`
	GetTimeZoneAndRule       *GetTimeZoneAndRuleArgs       `xml:"u:GetTimeZoneAndRule,omitempty"`
	GetTimeZoneRule          *GetTimeZoneRuleArgs          `xml:"u:GetTimeZoneRule,omitempty"`
	SetTimeServer            *SetTimeServerArgs            `xml:"u:SetTimeServer,omitempty"`
	GetTimeServer            *GetTimeServerArgs            `xml:"u:GetTimeServer,omitempty"`
	SetTimeNow               *SetTimeNowArgs               `xml:"u:SetTimeNow,omitempty"`
	GetHouseholdTimeAtStamp  *GetHouseholdTimeAtStampArgs  `xml:"u:GetHouseholdTimeAtStamp,omitempty"`
	GetTimeNow               *GetTimeNowArgs               `xml:"u:GetTimeNow,omitempty"`
	CreateAlarm              *CreateAlarmArgs              `xml:"u:CreateAlarm,omitempty"`
	UpdateAlarm              *UpdateAlarmArgs              `xml:"u:UpdateAlarm,omitempty"`
	DestroyAlarm             *DestroyAlarmArgs             `xml:"u:DestroyAlarm,omitempty"`
	ListAlarms               *ListAlarmsArgs               `xml:"u:ListAlarms,omitempty"`
	SetDailyIndexRefreshTime *SetDailyIndexRefreshTimeArgs `xml:"u:SetDailyIndexRefreshTime,omitempty"`
	GetDailyIndexRefreshTime *GetDailyIndexRefreshTimeArgs `xml:"u:GetDailyIndexRefreshTime,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                  xml.Name                          `xml:"Body"`
	SetFormat                *SetFormatResponse                `xml:"SetFormatResponse,omitempty"`
	GetFormat                *GetFormatResponse                `xml:"GetFormatResponse,omitempty"`
	SetTimeZone              *SetTimeZoneResponse              `xml:"SetTimeZoneResponse,omitempty"`
	GetTimeZone              *GetTimeZoneResponse              `xml:"GetTimeZoneResponse,omitempty"`
	GetTimeZoneAndRule       *GetTimeZoneAndRuleResponse       `xml:"GetTimeZoneAndRuleResponse,omitempty"`
	GetTimeZoneRule          *GetTimeZoneRuleResponse          `xml:"GetTimeZoneRuleResponse,omitempty"`
	SetTimeServer            *SetTimeServerResponse            `xml:"SetTimeServerResponse,omitempty"`
	GetTimeServer            *GetTimeServerResponse            `xml:"GetTimeServerResponse,omitempty"`
	SetTimeNow               *SetTimeNowResponse               `xml:"SetTimeNowResponse,omitempty"`
	GetHouseholdTimeAtStamp  *GetHouseholdTimeAtStampResponse  `xml:"GetHouseholdTimeAtStampResponse,omitempty"`
	GetTimeNow               *GetTimeNowResponse               `xml:"GetTimeNowResponse,omitempty"`
	CreateAlarm              *CreateAlarmResponse              `xml:"CreateAlarmResponse,omitempty"`
	UpdateAlarm              *UpdateAlarmResponse              `xml:"UpdateAlarmResponse,omitempty"`
	DestroyAlarm             *DestroyAlarmResponse             `xml:"DestroyAlarmResponse,omitempty"`
	ListAlarms               *ListAlarmsResponse               `xml:"ListAlarmsResponse,omitempty"`
	SetDailyIndexRefreshTime *SetDailyIndexRefreshTimeResponse `xml:"SetDailyIndexRefreshTimeResponse,omitempty"`
	GetDailyIndexRefreshTime *GetDailyIndexRefreshTimeResponse `xml:"GetDailyIndexRefreshTimeResponse,omitempty"`
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

type SetFormatArgs struct {
	Xmlns             string `xml:"xmlns:u,attr"`
	DesiredTimeFormat string `xml:"DesiredTimeFormat"`
	DesiredDateFormat string `xml:"DesiredDateFormat"`
}
type SetFormatResponse struct {
}

func (s *Service) SetFormat(httpClient *http.Client, args *SetFormatArgs) (*SetFormatResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetFormat`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetFormat: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetFormat == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.SetFormat()`)
	}

	return r.Body.SetFormat, nil
}

type GetFormatArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetFormatResponse struct {
	CurrentTimeFormat string `xml:"CurrentTimeFormat"`
	CurrentDateFormat string `xml:"CurrentDateFormat"`
}

func (s *Service) GetFormat(httpClient *http.Client, args *GetFormatArgs) (*GetFormatResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetFormat`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetFormat: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetFormat == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetFormat()`)
	}

	return r.Body.GetFormat, nil
}

type SetTimeZoneArgs struct {
	Xmlns         string `xml:"xmlns:u,attr"`
	Index         int32  `xml:"Index"`
	AutoAdjustDst bool   `xml:"AutoAdjustDst"`
}
type SetTimeZoneResponse struct {
}

func (s *Service) SetTimeZone(httpClient *http.Client, args *SetTimeZoneArgs) (*SetTimeZoneResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetTimeZone`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetTimeZone: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetTimeZone == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.SetTimeZone()`)
	}

	return r.Body.SetTimeZone, nil
}

type GetTimeZoneArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetTimeZoneResponse struct {
	Index         int32 `xml:"Index"`
	AutoAdjustDst bool  `xml:"AutoAdjustDst"`
}

func (s *Service) GetTimeZone(httpClient *http.Client, args *GetTimeZoneArgs) (*GetTimeZoneResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetTimeZone`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetTimeZone: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeZone == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetTimeZone()`)
	}

	return r.Body.GetTimeZone, nil
}

type GetTimeZoneAndRuleArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetTimeZoneAndRuleResponse struct {
	Index           int32  `xml:"Index"`
	AutoAdjustDst   bool   `xml:"AutoAdjustDst"`
	CurrentTimeZone string `xml:"CurrentTimeZone"`
}

func (s *Service) GetTimeZoneAndRule(httpClient *http.Client, args *GetTimeZoneAndRuleArgs) (*GetTimeZoneAndRuleResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetTimeZoneAndRule`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetTimeZoneAndRule: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeZoneAndRule == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetTimeZoneAndRule()`)
	}

	return r.Body.GetTimeZoneAndRule, nil
}

type GetTimeZoneRuleArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
	Index int32  `xml:"Index"`
}
type GetTimeZoneRuleResponse struct {
	TimeZone string `xml:"TimeZone"`
}

func (s *Service) GetTimeZoneRule(httpClient *http.Client, args *GetTimeZoneRuleArgs) (*GetTimeZoneRuleResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetTimeZoneRule`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetTimeZoneRule: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeZoneRule == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetTimeZoneRule()`)
	}

	return r.Body.GetTimeZoneRule, nil
}

type SetTimeServerArgs struct {
	Xmlns             string `xml:"xmlns:u,attr"`
	DesiredTimeServer string `xml:"DesiredTimeServer"`
}
type SetTimeServerResponse struct {
}

func (s *Service) SetTimeServer(httpClient *http.Client, args *SetTimeServerArgs) (*SetTimeServerResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetTimeServer`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetTimeServer: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetTimeServer == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.SetTimeServer()`)
	}

	return r.Body.SetTimeServer, nil
}

type GetTimeServerArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetTimeServerResponse struct {
	CurrentTimeServer string `xml:"CurrentTimeServer"`
}

func (s *Service) GetTimeServer(httpClient *http.Client, args *GetTimeServerArgs) (*GetTimeServerResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetTimeServer`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetTimeServer: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeServer == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetTimeServer()`)
	}

	return r.Body.GetTimeServer, nil
}

type SetTimeNowArgs struct {
	Xmlns                  string `xml:"xmlns:u,attr"`
	DesiredTime            string `xml:"DesiredTime"`
	TimeZoneForDesiredTime string `xml:"TimeZoneForDesiredTime"`
}
type SetTimeNowResponse struct {
}

func (s *Service) SetTimeNow(httpClient *http.Client, args *SetTimeNowArgs) (*SetTimeNowResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetTimeNow`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetTimeNow: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetTimeNow == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.SetTimeNow()`)
	}

	return r.Body.SetTimeNow, nil
}

type GetHouseholdTimeAtStampArgs struct {
	Xmlns     string `xml:"xmlns:u,attr"`
	TimeStamp string `xml:"TimeStamp"`
}
type GetHouseholdTimeAtStampResponse struct {
	HouseholdUTCTime string `xml:"HouseholdUTCTime"`
}

func (s *Service) GetHouseholdTimeAtStamp(httpClient *http.Client, args *GetHouseholdTimeAtStampArgs) (*GetHouseholdTimeAtStampResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetHouseholdTimeAtStamp`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetHouseholdTimeAtStamp: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetHouseholdTimeAtStamp == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetHouseholdTimeAtStamp()`)
	}

	return r.Body.GetHouseholdTimeAtStamp, nil
}

type GetTimeNowArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetTimeNowResponse struct {
	CurrentUTCTime        string `xml:"CurrentUTCTime"`
	CurrentLocalTime      string `xml:"CurrentLocalTime"`
	CurrentTimeZone       string `xml:"CurrentTimeZone"`
	CurrentTimeGeneration uint32 `xml:"CurrentTimeGeneration"`
}

func (s *Service) GetTimeNow(httpClient *http.Client, args *GetTimeNowArgs) (*GetTimeNowResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetTimeNow`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetTimeNow: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeNow == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetTimeNow()`)
	}

	return r.Body.GetTimeNow, nil
}

type CreateAlarmArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	StartLocalTime string `xml:"StartLocalTime"`
	Duration       string `xml:"Duration"`
	// Allowed Value: ONCE
	// Allowed Value: WEEKDAYS
	// Allowed Value: WEEKENDS
	// Allowed Value: DAILY
	Recurrence      string `xml:"Recurrence"`
	Enabled         bool   `xml:"Enabled"`
	RoomUUID        string `xml:"RoomUUID"`
	ProgramURI      string `xml:"ProgramURI"`
	ProgramMetaData string `xml:"ProgramMetaData"`
	// Allowed Value: NORMAL
	// Allowed Value: REPEAT_ALL
	// Allowed Value: SHUFFLE_NOREPEAT
	// Allowed Value: SHUFFLE
	PlayMode           string `xml:"PlayMode"`
	Volume             uint16 `xml:"Volume"`
	IncludeLinkedZones bool   `xml:"IncludeLinkedZones"`
}
type CreateAlarmResponse struct {
	AssignedID uint32 `xml:"AssignedID"`
}

func (s *Service) CreateAlarm(httpClient *http.Client, args *CreateAlarmArgs) (*CreateAlarmResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`CreateAlarm`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{CreateAlarm: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateAlarm == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.CreateAlarm()`)
	}

	return r.Body.CreateAlarm, nil
}

type UpdateAlarmArgs struct {
	Xmlns          string `xml:"xmlns:u,attr"`
	ID             uint32 `xml:"ID"`
	StartLocalTime string `xml:"StartLocalTime"`
	Duration       string `xml:"Duration"`
	// Allowed Value: ONCE
	// Allowed Value: WEEKDAYS
	// Allowed Value: WEEKENDS
	// Allowed Value: DAILY
	Recurrence      string `xml:"Recurrence"`
	Enabled         bool   `xml:"Enabled"`
	RoomUUID        string `xml:"RoomUUID"`
	ProgramURI      string `xml:"ProgramURI"`
	ProgramMetaData string `xml:"ProgramMetaData"`
	// Allowed Value: NORMAL
	// Allowed Value: REPEAT_ALL
	// Allowed Value: SHUFFLE_NOREPEAT
	// Allowed Value: SHUFFLE
	PlayMode           string `xml:"PlayMode"`
	Volume             uint16 `xml:"Volume"`
	IncludeLinkedZones bool   `xml:"IncludeLinkedZones"`
}
type UpdateAlarmResponse struct {
}

func (s *Service) UpdateAlarm(httpClient *http.Client, args *UpdateAlarmArgs) (*UpdateAlarmResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`UpdateAlarm`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{UpdateAlarm: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.UpdateAlarm == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.UpdateAlarm()`)
	}

	return r.Body.UpdateAlarm, nil
}

type DestroyAlarmArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
	ID    uint32 `xml:"ID"`
}
type DestroyAlarmResponse struct {
}

func (s *Service) DestroyAlarm(httpClient *http.Client, args *DestroyAlarmArgs) (*DestroyAlarmResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`DestroyAlarm`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{DestroyAlarm: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DestroyAlarm == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.DestroyAlarm()`)
	}

	return r.Body.DestroyAlarm, nil
}

type ListAlarmsArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type ListAlarmsResponse struct {
	CurrentAlarmList        string `xml:"CurrentAlarmList"`
	CurrentAlarmListVersion string `xml:"CurrentAlarmListVersion"`
}

func (s *Service) ListAlarms(httpClient *http.Client, args *ListAlarmsArgs) (*ListAlarmsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ListAlarms`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ListAlarms: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ListAlarms == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.ListAlarms()`)
	}

	return r.Body.ListAlarms, nil
}

type SetDailyIndexRefreshTimeArgs struct {
	Xmlns                        string `xml:"xmlns:u,attr"`
	DesiredDailyIndexRefreshTime string `xml:"DesiredDailyIndexRefreshTime"`
}
type SetDailyIndexRefreshTimeResponse struct {
}

func (s *Service) SetDailyIndexRefreshTime(httpClient *http.Client, args *SetDailyIndexRefreshTimeArgs) (*SetDailyIndexRefreshTimeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetDailyIndexRefreshTime`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetDailyIndexRefreshTime: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetDailyIndexRefreshTime == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.SetDailyIndexRefreshTime()`)
	}

	return r.Body.SetDailyIndexRefreshTime, nil
}

type GetDailyIndexRefreshTimeArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetDailyIndexRefreshTimeResponse struct {
	CurrentDailyIndexRefreshTime string `xml:"CurrentDailyIndexRefreshTime"`
}

func (s *Service) GetDailyIndexRefreshTime(httpClient *http.Client, args *GetDailyIndexRefreshTimeArgs) (*GetDailyIndexRefreshTimeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetDailyIndexRefreshTime`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetDailyIndexRefreshTime: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetDailyIndexRefreshTime == nil {
		return nil, errors.New(`unexpected respose from service calling alarmclock.GetDailyIndexRefreshTime()`)
	}

	return r.Body.GetDailyIndexRefreshTime, nil
}
