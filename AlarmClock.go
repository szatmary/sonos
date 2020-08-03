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
type AlarmClock_TimeZone string
type AlarmClock_TimeServer string
type AlarmClock_TimeGeneration uint32
type AlarmClock_AlarmListVersion string
type AlarmClock_DailyIndexRefreshTime string
type AlarmClock_TimeFormat string
type AlarmClock_DateFormat string

type AlarmClockService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
	// State
	TimeZone              *AlarmClock_TimeZone
	TimeServer            *AlarmClock_TimeServer
	TimeGeneration        *AlarmClock_TimeGeneration
	AlarmListVersion      *AlarmClock_AlarmListVersion
	DailyIndexRefreshTime *AlarmClock_DailyIndexRefreshTime
	TimeFormat            *AlarmClock_TimeFormat
	DateFormat            *AlarmClock_DateFormat
}

func NewAlarmClockService(deviceUrl *url.URL) *AlarmClockService {
	c, _ := url.Parse("/AlarmClock/Control")
	e, _ := url.Parse("/AlarmClock/Event")
	return &AlarmClockService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *AlarmClockService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *AlarmClockService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type AlarmClockEnvelope struct {
	XMLName       xml.Name       `xml:"s:Envelope"`
	XMLNameSpace  string         `xml:"xmlns:s,attr"`
	EncodingStyle string         `xml:"s:encodingStyle,attr"`
	Body          AlarmClockBody `xml:"s:Body"`
}
type AlarmClockBody struct {
	XMLName                  xml.Name                                `xml:"s:Body"`
	SetFormat                *AlarmClockSetFormatArgs                `xml:"u:SetFormat,omitempty"`
	GetFormat                *AlarmClockGetFormatArgs                `xml:"u:GetFormat,omitempty"`
	SetTimeZone              *AlarmClockSetTimeZoneArgs              `xml:"u:SetTimeZone,omitempty"`
	GetTimeZone              *AlarmClockGetTimeZoneArgs              `xml:"u:GetTimeZone,omitempty"`
	GetTimeZoneAndRule       *AlarmClockGetTimeZoneAndRuleArgs       `xml:"u:GetTimeZoneAndRule,omitempty"`
	GetTimeZoneRule          *AlarmClockGetTimeZoneRuleArgs          `xml:"u:GetTimeZoneRule,omitempty"`
	SetTimeServer            *AlarmClockSetTimeServerArgs            `xml:"u:SetTimeServer,omitempty"`
	GetTimeServer            *AlarmClockGetTimeServerArgs            `xml:"u:GetTimeServer,omitempty"`
	SetTimeNow               *AlarmClockSetTimeNowArgs               `xml:"u:SetTimeNow,omitempty"`
	GetHouseholdTimeAtStamp  *AlarmClockGetHouseholdTimeAtStampArgs  `xml:"u:GetHouseholdTimeAtStamp,omitempty"`
	GetTimeNow               *AlarmClockGetTimeNowArgs               `xml:"u:GetTimeNow,omitempty"`
	CreateAlarm              *AlarmClockCreateAlarmArgs              `xml:"u:CreateAlarm,omitempty"`
	UpdateAlarm              *AlarmClockUpdateAlarmArgs              `xml:"u:UpdateAlarm,omitempty"`
	DestroyAlarm             *AlarmClockDestroyAlarmArgs             `xml:"u:DestroyAlarm,omitempty"`
	ListAlarms               *AlarmClockListAlarmsArgs               `xml:"u:ListAlarms,omitempty"`
	SetDailyIndexRefreshTime *AlarmClockSetDailyIndexRefreshTimeArgs `xml:"u:SetDailyIndexRefreshTime,omitempty"`
	GetDailyIndexRefreshTime *AlarmClockGetDailyIndexRefreshTimeArgs `xml:"u:GetDailyIndexRefreshTime,omitempty"`
}
type AlarmClockEnvelopeResponse struct {
	XMLName       xml.Name               `xml:"Envelope"`
	XMLNameSpace  string                 `xml:"xmlns:s,attr"`
	EncodingStyle string                 `xml:"encodingStyle,attr"`
	Body          AlarmClockBodyResponse `xml:"Body"`
}
type AlarmClockBodyResponse struct {
	XMLName                  xml.Name                                    `xml:"Body"`
	SetFormat                *AlarmClockSetFormatResponse                `xml:"SetFormatResponse"`
	GetFormat                *AlarmClockGetFormatResponse                `xml:"GetFormatResponse"`
	SetTimeZone              *AlarmClockSetTimeZoneResponse              `xml:"SetTimeZoneResponse"`
	GetTimeZone              *AlarmClockGetTimeZoneResponse              `xml:"GetTimeZoneResponse"`
	GetTimeZoneAndRule       *AlarmClockGetTimeZoneAndRuleResponse       `xml:"GetTimeZoneAndRuleResponse"`
	GetTimeZoneRule          *AlarmClockGetTimeZoneRuleResponse          `xml:"GetTimeZoneRuleResponse"`
	SetTimeServer            *AlarmClockSetTimeServerResponse            `xml:"SetTimeServerResponse"`
	GetTimeServer            *AlarmClockGetTimeServerResponse            `xml:"GetTimeServerResponse"`
	SetTimeNow               *AlarmClockSetTimeNowResponse               `xml:"SetTimeNowResponse"`
	GetHouseholdTimeAtStamp  *AlarmClockGetHouseholdTimeAtStampResponse  `xml:"GetHouseholdTimeAtStampResponse"`
	GetTimeNow               *AlarmClockGetTimeNowResponse               `xml:"GetTimeNowResponse"`
	CreateAlarm              *AlarmClockCreateAlarmResponse              `xml:"CreateAlarmResponse"`
	UpdateAlarm              *AlarmClockUpdateAlarmResponse              `xml:"UpdateAlarmResponse"`
	DestroyAlarm             *AlarmClockDestroyAlarmResponse             `xml:"DestroyAlarmResponse"`
	ListAlarms               *AlarmClockListAlarmsResponse               `xml:"ListAlarmsResponse"`
	SetDailyIndexRefreshTime *AlarmClockSetDailyIndexRefreshTimeResponse `xml:"SetDailyIndexRefreshTimeResponse"`
	GetDailyIndexRefreshTime *AlarmClockGetDailyIndexRefreshTimeResponse `xml:"GetDailyIndexRefreshTimeResponse"`
}

func (s *AlarmClockService) _AlarmClockExec(soapAction string, httpClient *http.Client, envelope *AlarmClockEnvelope) (*AlarmClockEnvelopeResponse, error) {
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
	var envelopeResponse AlarmClockEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type AlarmClockSetFormatArgs struct {
	XMLNameSpace      string `xml:"xmlns:u,attr"`
	DesiredTimeFormat string `xml:"DesiredTimeFormat"`
	DesiredDateFormat string `xml:"DesiredDateFormat"`
}
type AlarmClockSetFormatResponse struct {
}

func (s *AlarmClockService) SetFormat(httpClient *http.Client, args *AlarmClockSetFormatArgs) (*AlarmClockSetFormatResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#SetFormat", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{SetFormat: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetFormat == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetFormat, nil
}

type AlarmClockGetFormatArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type AlarmClockGetFormatResponse struct {
	CurrentTimeFormat string `xml:"CurrentTimeFormat"`
	CurrentDateFormat string `xml:"CurrentDateFormat"`
}

func (s *AlarmClockService) GetFormat(httpClient *http.Client, args *AlarmClockGetFormatArgs) (*AlarmClockGetFormatResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetFormat", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetFormat: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetFormat == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetFormat, nil
}

type AlarmClockSetTimeZoneArgs struct {
	XMLNameSpace  string `xml:"xmlns:u,attr"`
	Index         int32  `xml:"Index"`
	AutoAdjustDst bool   `xml:"AutoAdjustDst"`
}
type AlarmClockSetTimeZoneResponse struct {
}

func (s *AlarmClockService) SetTimeZone(httpClient *http.Client, args *AlarmClockSetTimeZoneArgs) (*AlarmClockSetTimeZoneResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#SetTimeZone", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{SetTimeZone: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetTimeZone == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetTimeZone, nil
}

type AlarmClockGetTimeZoneArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type AlarmClockGetTimeZoneResponse struct {
	Index         int32 `xml:"Index"`
	AutoAdjustDst bool  `xml:"AutoAdjustDst"`
}

func (s *AlarmClockService) GetTimeZone(httpClient *http.Client, args *AlarmClockGetTimeZoneArgs) (*AlarmClockGetTimeZoneResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetTimeZone", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetTimeZone: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeZone == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTimeZone, nil
}

type AlarmClockGetTimeZoneAndRuleArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type AlarmClockGetTimeZoneAndRuleResponse struct {
	Index           int32  `xml:"Index"`
	AutoAdjustDst   bool   `xml:"AutoAdjustDst"`
	CurrentTimeZone string `xml:"CurrentTimeZone"`
}

func (s *AlarmClockService) GetTimeZoneAndRule(httpClient *http.Client, args *AlarmClockGetTimeZoneAndRuleArgs) (*AlarmClockGetTimeZoneAndRuleResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetTimeZoneAndRule", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetTimeZoneAndRule: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeZoneAndRule == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTimeZoneAndRule, nil
}

type AlarmClockGetTimeZoneRuleArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	Index        int32  `xml:"Index"`
}
type AlarmClockGetTimeZoneRuleResponse struct {
	TimeZone string `xml:"TimeZone"`
}

func (s *AlarmClockService) GetTimeZoneRule(httpClient *http.Client, args *AlarmClockGetTimeZoneRuleArgs) (*AlarmClockGetTimeZoneRuleResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetTimeZoneRule", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetTimeZoneRule: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeZoneRule == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTimeZoneRule, nil
}

type AlarmClockSetTimeServerArgs struct {
	XMLNameSpace      string `xml:"xmlns:u,attr"`
	DesiredTimeServer string `xml:"DesiredTimeServer"`
}
type AlarmClockSetTimeServerResponse struct {
}

func (s *AlarmClockService) SetTimeServer(httpClient *http.Client, args *AlarmClockSetTimeServerArgs) (*AlarmClockSetTimeServerResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#SetTimeServer", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{SetTimeServer: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetTimeServer == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetTimeServer, nil
}

type AlarmClockGetTimeServerArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type AlarmClockGetTimeServerResponse struct {
	CurrentTimeServer string `xml:"CurrentTimeServer"`
}

func (s *AlarmClockService) GetTimeServer(httpClient *http.Client, args *AlarmClockGetTimeServerArgs) (*AlarmClockGetTimeServerResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetTimeServer", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetTimeServer: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeServer == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTimeServer, nil
}

type AlarmClockSetTimeNowArgs struct {
	XMLNameSpace           string `xml:"xmlns:u,attr"`
	DesiredTime            string `xml:"DesiredTime"`
	TimeZoneForDesiredTime string `xml:"TimeZoneForDesiredTime"`
}
type AlarmClockSetTimeNowResponse struct {
}

func (s *AlarmClockService) SetTimeNow(httpClient *http.Client, args *AlarmClockSetTimeNowArgs) (*AlarmClockSetTimeNowResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#SetTimeNow", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{SetTimeNow: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetTimeNow == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetTimeNow, nil
}

type AlarmClockGetHouseholdTimeAtStampArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	TimeStamp    string `xml:"TimeStamp"`
}
type AlarmClockGetHouseholdTimeAtStampResponse struct {
	HouseholdUTCTime string `xml:"HouseholdUTCTime"`
}

func (s *AlarmClockService) GetHouseholdTimeAtStamp(httpClient *http.Client, args *AlarmClockGetHouseholdTimeAtStampArgs) (*AlarmClockGetHouseholdTimeAtStampResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetHouseholdTimeAtStamp", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetHouseholdTimeAtStamp: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetHouseholdTimeAtStamp == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetHouseholdTimeAtStamp, nil
}

type AlarmClockGetTimeNowArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type AlarmClockGetTimeNowResponse struct {
	CurrentUTCTime        string `xml:"CurrentUTCTime"`
	CurrentLocalTime      string `xml:"CurrentLocalTime"`
	CurrentTimeZone       string `xml:"CurrentTimeZone"`
	CurrentTimeGeneration uint32 `xml:"CurrentTimeGeneration"`
}

func (s *AlarmClockService) GetTimeNow(httpClient *http.Client, args *AlarmClockGetTimeNowArgs) (*AlarmClockGetTimeNowResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetTimeNow", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetTimeNow: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetTimeNow == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetTimeNow, nil
}

type AlarmClockCreateAlarmArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
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
type AlarmClockCreateAlarmResponse struct {
	AssignedID uint32 `xml:"AssignedID"`
}

func (s *AlarmClockService) CreateAlarm(httpClient *http.Client, args *AlarmClockCreateAlarmArgs) (*AlarmClockCreateAlarmResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#CreateAlarm", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{CreateAlarm: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.CreateAlarm == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.CreateAlarm, nil
}

type AlarmClockUpdateAlarmArgs struct {
	XMLNameSpace   string `xml:"xmlns:u,attr"`
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
type AlarmClockUpdateAlarmResponse struct {
}

func (s *AlarmClockService) UpdateAlarm(httpClient *http.Client, args *AlarmClockUpdateAlarmArgs) (*AlarmClockUpdateAlarmResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#UpdateAlarm", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{UpdateAlarm: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.UpdateAlarm == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.UpdateAlarm, nil
}

type AlarmClockDestroyAlarmArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	ID           uint32 `xml:"ID"`
}
type AlarmClockDestroyAlarmResponse struct {
}

func (s *AlarmClockService) DestroyAlarm(httpClient *http.Client, args *AlarmClockDestroyAlarmArgs) (*AlarmClockDestroyAlarmResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#DestroyAlarm", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{DestroyAlarm: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DestroyAlarm == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.DestroyAlarm, nil
}

type AlarmClockListAlarmsArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type AlarmClockListAlarmsResponse struct {
	CurrentAlarmList        string `xml:"CurrentAlarmList"`
	CurrentAlarmListVersion string `xml:"CurrentAlarmListVersion"`
}

func (s *AlarmClockService) ListAlarms(httpClient *http.Client, args *AlarmClockListAlarmsArgs) (*AlarmClockListAlarmsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#ListAlarms", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{ListAlarms: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ListAlarms == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ListAlarms, nil
}

type AlarmClockSetDailyIndexRefreshTimeArgs struct {
	XMLNameSpace                 string `xml:"xmlns:u,attr"`
	DesiredDailyIndexRefreshTime string `xml:"DesiredDailyIndexRefreshTime"`
}
type AlarmClockSetDailyIndexRefreshTimeResponse struct {
}

func (s *AlarmClockService) SetDailyIndexRefreshTime(httpClient *http.Client, args *AlarmClockSetDailyIndexRefreshTimeArgs) (*AlarmClockSetDailyIndexRefreshTimeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#SetDailyIndexRefreshTime", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{SetDailyIndexRefreshTime: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetDailyIndexRefreshTime == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetDailyIndexRefreshTime, nil
}

type AlarmClockGetDailyIndexRefreshTimeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type AlarmClockGetDailyIndexRefreshTimeResponse struct {
	CurrentDailyIndexRefreshTime string `xml:"CurrentDailyIndexRefreshTime"`
}

func (s *AlarmClockService) GetDailyIndexRefreshTime(httpClient *http.Client, args *AlarmClockGetDailyIndexRefreshTimeArgs) (*AlarmClockGetDailyIndexRefreshTimeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:AlarmClock:1"
	r, err := s._AlarmClockExec("urn:schemas-upnp-org:service:AlarmClock:1#GetDailyIndexRefreshTime", httpClient,
		&AlarmClockEnvelope{
			Body:          AlarmClockBody{GetDailyIndexRefreshTime: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetDailyIndexRefreshTime == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetDailyIndexRefreshTime, nil
}

type AlarmClockUpnpEvent struct {
	XMLName      xml.Name             `xml:"propertyset"`
	XMLNameSpace string               `xml:"xmlns:e,attr"`
	Properties   []AlarmClockProperty `xml:"property"`
}
type AlarmClockProperty struct {
	XMLName               xml.Name                          `xml:"property"`
	TimeZone              *AlarmClock_TimeZone              `xml:"TimeZone"`
	TimeServer            *AlarmClock_TimeServer            `xml:"TimeServer"`
	TimeGeneration        *AlarmClock_TimeGeneration        `xml:"TimeGeneration"`
	AlarmListVersion      *AlarmClock_AlarmListVersion      `xml:"AlarmListVersion"`
	DailyIndexRefreshTime *AlarmClock_DailyIndexRefreshTime `xml:"DailyIndexRefreshTime"`
	TimeFormat            *AlarmClock_TimeFormat            `xml:"TimeFormat"`
	DateFormat            *AlarmClock_DateFormat            `xml:"DateFormat"`
}

func (zp *AlarmClockService) ParseEvent(body []byte) []interface{} {
	var evt AlarmClockUpnpEvent
	var events []interface{}
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return events
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.TimeZone != nil:
			zp.TimeZone = prop.TimeZone
			events = append(events, *prop.TimeZone)
		case prop.TimeServer != nil:
			zp.TimeServer = prop.TimeServer
			events = append(events, *prop.TimeServer)
		case prop.TimeGeneration != nil:
			zp.TimeGeneration = prop.TimeGeneration
			events = append(events, *prop.TimeGeneration)
		case prop.AlarmListVersion != nil:
			zp.AlarmListVersion = prop.AlarmListVersion
			events = append(events, *prop.AlarmListVersion)
		case prop.DailyIndexRefreshTime != nil:
			zp.DailyIndexRefreshTime = prop.DailyIndexRefreshTime
			events = append(events, *prop.DailyIndexRefreshTime)
		case prop.TimeFormat != nil:
			zp.TimeFormat = prop.TimeFormat
			events = append(events, *prop.TimeFormat)
		case prop.DateFormat != nil:
			zp.DateFormat = prop.DateFormat
			events = append(events, *prop.DateFormat)
		}
	}
	return events
}
