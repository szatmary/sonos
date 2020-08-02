package sonos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SystemPropertiesService struct {
	controlEndpoint *url.URL
	eventEndpoint   *url.URL
}

func NewSystemPropertiesService(deviceUrl *url.URL) *SystemPropertiesService {
	c, _ := url.Parse("/SystemProperties/Control")
	e, _ := url.Parse("/SystemProperties/Event")
	return &SystemPropertiesService{
		controlEndpoint: deviceUrl.ResolveReference(c),
		eventEndpoint:   deviceUrl.ResolveReference(e),
	}
}
func (s *SystemPropertiesService) ControlEndpoint() *url.URL {
	return s.controlEndpoint
}
func (s *SystemPropertiesService) EventEndpoint() *url.URL {
	return s.eventEndpoint
}

type SystemPropertiesEnvelope struct {
	XMLName       xml.Name             `xml:"s:Envelope"`
	XMLNameSpace  string               `xml:"xmlns:s,attr"`
	EncodingStyle string               `xml:"s:encodingStyle,attr"`
	Body          SystemPropertiesBody `xml:"s:Body"`
}
type SystemPropertiesBody struct {
	XMLName                            xml.Name                                                `xml:"s:Body"`
	SetString                          *SystemPropertiesSetStringArgs                          `xml:"u:SetString,omitempty"`
	GetString                          *SystemPropertiesGetStringArgs                          `xml:"u:GetString,omitempty"`
	Remove                             *SystemPropertiesRemoveArgs                             `xml:"u:Remove,omitempty"`
	GetWebCode                         *SystemPropertiesGetWebCodeArgs                         `xml:"u:GetWebCode,omitempty"`
	ProvisionCredentialedTrialAccountX *SystemPropertiesProvisionCredentialedTrialAccountXArgs `xml:"u:ProvisionCredentialedTrialAccountX,omitempty"`
	AddAccountX                        *SystemPropertiesAddAccountXArgs                        `xml:"u:AddAccountX,omitempty"`
	AddOAuthAccountX                   *SystemPropertiesAddOAuthAccountXArgs                   `xml:"u:AddOAuthAccountX,omitempty"`
	RemoveAccount                      *SystemPropertiesRemoveAccountArgs                      `xml:"u:RemoveAccount,omitempty"`
	EditAccountPasswordX               *SystemPropertiesEditAccountPasswordXArgs               `xml:"u:EditAccountPasswordX,omitempty"`
	SetAccountNicknameX                *SystemPropertiesSetAccountNicknameXArgs                `xml:"u:SetAccountNicknameX,omitempty"`
	RefreshAccountCredentialsX         *SystemPropertiesRefreshAccountCredentialsXArgs         `xml:"u:RefreshAccountCredentialsX,omitempty"`
	EditAccountMd                      *SystemPropertiesEditAccountMdArgs                      `xml:"u:EditAccountMd,omitempty"`
	DoPostUpdateTasks                  *SystemPropertiesDoPostUpdateTasksArgs                  `xml:"u:DoPostUpdateTasks,omitempty"`
	ResetThirdPartyCredentials         *SystemPropertiesResetThirdPartyCredentialsArgs         `xml:"u:ResetThirdPartyCredentials,omitempty"`
	EnableRDM                          *SystemPropertiesEnableRDMArgs                          `xml:"u:EnableRDM,omitempty"`
	GetRDM                             *SystemPropertiesGetRDMArgs                             `xml:"u:GetRDM,omitempty"`
	ReplaceAccountX                    *SystemPropertiesReplaceAccountXArgs                    `xml:"u:ReplaceAccountX,omitempty"`
}
type SystemPropertiesEnvelopeResponse struct {
	XMLName       xml.Name                     `xml:"Envelope"`
	XMLNameSpace  string                       `xml:"xmlns:s,attr"`
	EncodingStyle string                       `xml:"encodingStyle,attr"`
	Body          SystemPropertiesBodyResponse `xml:"Body"`
}
type SystemPropertiesBodyResponse struct {
	XMLName                            xml.Name                                                    `xml:"Body"`
	SetString                          *SystemPropertiesSetStringResponse                          `xml:"SetStringResponse"`
	GetString                          *SystemPropertiesGetStringResponse                          `xml:"GetStringResponse"`
	Remove                             *SystemPropertiesRemoveResponse                             `xml:"RemoveResponse"`
	GetWebCode                         *SystemPropertiesGetWebCodeResponse                         `xml:"GetWebCodeResponse"`
	ProvisionCredentialedTrialAccountX *SystemPropertiesProvisionCredentialedTrialAccountXResponse `xml:"ProvisionCredentialedTrialAccountXResponse"`
	AddAccountX                        *SystemPropertiesAddAccountXResponse                        `xml:"AddAccountXResponse"`
	AddOAuthAccountX                   *SystemPropertiesAddOAuthAccountXResponse                   `xml:"AddOAuthAccountXResponse"`
	RemoveAccount                      *SystemPropertiesRemoveAccountResponse                      `xml:"RemoveAccountResponse"`
	EditAccountPasswordX               *SystemPropertiesEditAccountPasswordXResponse               `xml:"EditAccountPasswordXResponse"`
	SetAccountNicknameX                *SystemPropertiesSetAccountNicknameXResponse                `xml:"SetAccountNicknameXResponse"`
	RefreshAccountCredentialsX         *SystemPropertiesRefreshAccountCredentialsXResponse         `xml:"RefreshAccountCredentialsXResponse"`
	EditAccountMd                      *SystemPropertiesEditAccountMdResponse                      `xml:"EditAccountMdResponse"`
	DoPostUpdateTasks                  *SystemPropertiesDoPostUpdateTasksResponse                  `xml:"DoPostUpdateTasksResponse"`
	ResetThirdPartyCredentials         *SystemPropertiesResetThirdPartyCredentialsResponse         `xml:"ResetThirdPartyCredentialsResponse"`
	EnableRDM                          *SystemPropertiesEnableRDMResponse                          `xml:"EnableRDMResponse"`
	GetRDM                             *SystemPropertiesGetRDMResponse                             `xml:"GetRDMResponse"`
	ReplaceAccountX                    *SystemPropertiesReplaceAccountXResponse                    `xml:"ReplaceAccountXResponse"`
}

func (s *SystemPropertiesService) _SystemPropertiesExec(soapAction string, httpClient *http.Client, envelope *SystemPropertiesEnvelope) (*SystemPropertiesEnvelopeResponse, error) {
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
	var envelopeResponse SystemPropertiesEnvelopeResponse
	err = xml.Unmarshal(responseBody, &envelopeResponse)
	if err != nil {
		return nil, err
	}
	return &envelopeResponse, nil
}

type SystemPropertiesSetStringArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	VariableName string `xml:"VariableName"`
	StringValue  string `xml:"StringValue"`
}
type SystemPropertiesSetStringResponse struct {
}

func (s *SystemPropertiesService) SetString(httpClient *http.Client, args *SystemPropertiesSetStringArgs) (*SystemPropertiesSetStringResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#SetString", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{SetString: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetString == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetString, nil
}

type SystemPropertiesGetStringArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	VariableName string `xml:"VariableName"`
}
type SystemPropertiesGetStringResponse struct {
	StringValue string `xml:"StringValue"`
}

func (s *SystemPropertiesService) GetString(httpClient *http.Client, args *SystemPropertiesGetStringArgs) (*SystemPropertiesGetStringResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#GetString", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{GetString: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetString == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetString, nil
}

type SystemPropertiesRemoveArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	VariableName string `xml:"VariableName"`
}
type SystemPropertiesRemoveResponse struct {
}

func (s *SystemPropertiesService) Remove(httpClient *http.Client, args *SystemPropertiesRemoveArgs) (*SystemPropertiesRemoveResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#Remove", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{Remove: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Remove == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.Remove, nil
}

type SystemPropertiesGetWebCodeArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	AccountType  uint32 `xml:"AccountType"`
}
type SystemPropertiesGetWebCodeResponse struct {
	WebCode string `xml:"WebCode"`
}

func (s *SystemPropertiesService) GetWebCode(httpClient *http.Client, args *SystemPropertiesGetWebCodeArgs) (*SystemPropertiesGetWebCodeResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#GetWebCode", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{GetWebCode: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetWebCode == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetWebCode, nil
}

type SystemPropertiesProvisionCredentialedTrialAccountXArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	AccountType     uint32 `xml:"AccountType"`
	AccountID       string `xml:"AccountID"`
	AccountPassword string `xml:"AccountPassword"`
}
type SystemPropertiesProvisionCredentialedTrialAccountXResponse struct {
	IsExpired  bool   `xml:"IsExpired"`
	AccountUDN string `xml:"AccountUDN"`
}

func (s *SystemPropertiesService) ProvisionCredentialedTrialAccountX(httpClient *http.Client, args *SystemPropertiesProvisionCredentialedTrialAccountXArgs) (*SystemPropertiesProvisionCredentialedTrialAccountXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#ProvisionCredentialedTrialAccountX", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{ProvisionCredentialedTrialAccountX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ProvisionCredentialedTrialAccountX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ProvisionCredentialedTrialAccountX, nil
}

type SystemPropertiesAddAccountXArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	AccountType     uint32 `xml:"AccountType"`
	AccountID       string `xml:"AccountID"`
	AccountPassword string `xml:"AccountPassword"`
}
type SystemPropertiesAddAccountXResponse struct {
	AccountUDN string `xml:"AccountUDN"`
}

func (s *SystemPropertiesService) AddAccountX(httpClient *http.Client, args *SystemPropertiesAddAccountXArgs) (*SystemPropertiesAddAccountXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#AddAccountX", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{AddAccountX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddAccountX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddAccountX, nil
}

type SystemPropertiesAddOAuthAccountXArgs struct {
	XMLNameSpace      string `xml:"xmlns:u,attr"`
	AccountType       uint32 `xml:"AccountType"`
	AccountToken      string `xml:"AccountToken"`
	AccountKey        string `xml:"AccountKey"`
	OAuthDeviceID     string `xml:"OAuthDeviceID"`
	AuthorizationCode string `xml:"AuthorizationCode"`
	RedirectURI       string `xml:"RedirectURI"`
	UserIdHashCode    string `xml:"UserIdHashCode"`
	AccountTier       uint32 `xml:"AccountTier"`
}
type SystemPropertiesAddOAuthAccountXResponse struct {
	AccountUDN      string `xml:"AccountUDN"`
	AccountNickname string `xml:"AccountNickname"`
}

func (s *SystemPropertiesService) AddOAuthAccountX(httpClient *http.Client, args *SystemPropertiesAddOAuthAccountXArgs) (*SystemPropertiesAddOAuthAccountXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#AddOAuthAccountX", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{AddOAuthAccountX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddOAuthAccountX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.AddOAuthAccountX, nil
}

type SystemPropertiesRemoveAccountArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	AccountType  uint32 `xml:"AccountType"`
	AccountID    string `xml:"AccountID"`
}
type SystemPropertiesRemoveAccountResponse struct {
}

func (s *SystemPropertiesService) RemoveAccount(httpClient *http.Client, args *SystemPropertiesRemoveAccountArgs) (*SystemPropertiesRemoveAccountResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#RemoveAccount", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{RemoveAccount: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveAccount == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RemoveAccount, nil
}

type SystemPropertiesEditAccountPasswordXArgs struct {
	XMLNameSpace       string `xml:"xmlns:u,attr"`
	AccountType        uint32 `xml:"AccountType"`
	AccountID          string `xml:"AccountID"`
	NewAccountPassword string `xml:"NewAccountPassword"`
}
type SystemPropertiesEditAccountPasswordXResponse struct {
}

func (s *SystemPropertiesService) EditAccountPasswordX(httpClient *http.Client, args *SystemPropertiesEditAccountPasswordXArgs) (*SystemPropertiesEditAccountPasswordXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#EditAccountPasswordX", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{EditAccountPasswordX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EditAccountPasswordX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.EditAccountPasswordX, nil
}

type SystemPropertiesSetAccountNicknameXArgs struct {
	XMLNameSpace    string `xml:"xmlns:u,attr"`
	AccountUDN      string `xml:"AccountUDN"`
	AccountNickname string `xml:"AccountNickname"`
}
type SystemPropertiesSetAccountNicknameXResponse struct {
}

func (s *SystemPropertiesService) SetAccountNicknameX(httpClient *http.Client, args *SystemPropertiesSetAccountNicknameXArgs) (*SystemPropertiesSetAccountNicknameXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#SetAccountNicknameX", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{SetAccountNicknameX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAccountNicknameX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.SetAccountNicknameX, nil
}

type SystemPropertiesRefreshAccountCredentialsXArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	AccountType  uint32 `xml:"AccountType"`
	AccountUID   uint32 `xml:"AccountUID"`
	AccountToken string `xml:"AccountToken"`
	AccountKey   string `xml:"AccountKey"`
}
type SystemPropertiesRefreshAccountCredentialsXResponse struct {
}

func (s *SystemPropertiesService) RefreshAccountCredentialsX(httpClient *http.Client, args *SystemPropertiesRefreshAccountCredentialsXArgs) (*SystemPropertiesRefreshAccountCredentialsXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#RefreshAccountCredentialsX", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{RefreshAccountCredentialsX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RefreshAccountCredentialsX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.RefreshAccountCredentialsX, nil
}

type SystemPropertiesEditAccountMdArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	AccountType  uint32 `xml:"AccountType"`
	AccountID    string `xml:"AccountID"`
	NewAccountMd string `xml:"NewAccountMd"`
}
type SystemPropertiesEditAccountMdResponse struct {
}

func (s *SystemPropertiesService) EditAccountMd(httpClient *http.Client, args *SystemPropertiesEditAccountMdArgs) (*SystemPropertiesEditAccountMdResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#EditAccountMd", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{EditAccountMd: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EditAccountMd == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.EditAccountMd, nil
}

type SystemPropertiesDoPostUpdateTasksArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type SystemPropertiesDoPostUpdateTasksResponse struct {
}

func (s *SystemPropertiesService) DoPostUpdateTasks(httpClient *http.Client, args *SystemPropertiesDoPostUpdateTasksArgs) (*SystemPropertiesDoPostUpdateTasksResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#DoPostUpdateTasks", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{DoPostUpdateTasks: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DoPostUpdateTasks == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.DoPostUpdateTasks, nil
}

type SystemPropertiesResetThirdPartyCredentialsArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type SystemPropertiesResetThirdPartyCredentialsResponse struct {
}

func (s *SystemPropertiesService) ResetThirdPartyCredentials(httpClient *http.Client, args *SystemPropertiesResetThirdPartyCredentialsArgs) (*SystemPropertiesResetThirdPartyCredentialsResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#ResetThirdPartyCredentials", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{ResetThirdPartyCredentials: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ResetThirdPartyCredentials == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ResetThirdPartyCredentials, nil
}

type SystemPropertiesEnableRDMArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
	RDMValue     bool   `xml:"RDMValue"`
}
type SystemPropertiesEnableRDMResponse struct {
}

func (s *SystemPropertiesService) EnableRDM(httpClient *http.Client, args *SystemPropertiesEnableRDMArgs) (*SystemPropertiesEnableRDMResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#EnableRDM", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{EnableRDM: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EnableRDM == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.EnableRDM, nil
}

type SystemPropertiesGetRDMArgs struct {
	XMLNameSpace string `xml:"xmlns:u,attr"`
}
type SystemPropertiesGetRDMResponse struct {
	RDMValue bool `xml:"RDMValue"`
}

func (s *SystemPropertiesService) GetRDM(httpClient *http.Client, args *SystemPropertiesGetRDMArgs) (*SystemPropertiesGetRDMResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#GetRDM", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{GetRDM: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetRDM == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.GetRDM, nil
}

type SystemPropertiesReplaceAccountXArgs struct {
	XMLNameSpace       string `xml:"xmlns:u,attr"`
	AccountUDN         string `xml:"AccountUDN"`
	NewAccountID       string `xml:"NewAccountID"`
	NewAccountPassword string `xml:"NewAccountPassword"`
	AccountToken       string `xml:"AccountToken"`
	AccountKey         string `xml:"AccountKey"`
	OAuthDeviceID      string `xml:"OAuthDeviceID"`
}
type SystemPropertiesReplaceAccountXResponse struct {
	NewAccountUDN string `xml:"NewAccountUDN"`
}

func (s *SystemPropertiesService) ReplaceAccountX(httpClient *http.Client, args *SystemPropertiesReplaceAccountXArgs) (*SystemPropertiesReplaceAccountXResponse, error) {
	args.XMLNameSpace = "urn:schemas-upnp-org:service:SystemProperties:1"
	r, err := s._SystemPropertiesExec("urn:schemas-upnp-org:service:SystemProperties:1#ReplaceAccountX", httpClient,
		&SystemPropertiesEnvelope{
			Body:          SystemPropertiesBody{ReplaceAccountX: args},
			EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/", XMLNameSpace: "http://schemas.xmlsoap.org/soap/envelope/",
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReplaceAccountX == nil {
		return nil, errors.New("unexpected respose from service")
	}
	return r.Body.ReplaceAccountX, nil
}

type SystemPropertiesUpnpEvent struct {
	XMLName      xml.Name                   `xml:"propertyset"`
	XMLNameSpace string                     `xml:"xmlns:e,attr"`
	Properties   []SystemPropertiesProperty `xml:"property"`
}
type SystemPropertiesProperty struct {
	XMLName        xml.Name `xml:"property"`
	CustomerID     *string  `xml:"CustomerID"`
	UpdateID       *uint32  `xml:"UpdateID"`
	UpdateIDX      *uint32  `xml:"UpdateIDX"`
	VoiceUpdateID  *uint32  `xml:"VoiceUpdateID"`
	ThirdPartyHash *string  `xml:"ThirdPartyHash"`
}

func SystemPropertiesDispatchEvent(zp *ZonePlayer, body []byte) {
	var evt SystemPropertiesUpnpEvent
	err := xml.Unmarshal(body, &evt)
	if err != nil {
		return
	}
	for _, prop := range evt.Properties {
		switch {
		case prop.CustomerID != nil:
			dispatchSystemPropertiesCustomerID(zp, *prop.CustomerID) // string
		case prop.UpdateID != nil:
			dispatchSystemPropertiesUpdateID(zp, *prop.UpdateID) // uint32
		case prop.UpdateIDX != nil:
			dispatchSystemPropertiesUpdateIDX(zp, *prop.UpdateIDX) // uint32
		case prop.VoiceUpdateID != nil:
			dispatchSystemPropertiesVoiceUpdateID(zp, *prop.VoiceUpdateID) // uint32
		case prop.ThirdPartyHash != nil:
			dispatchSystemPropertiesThirdPartyHash(zp, *prop.ThirdPartyHash) // string
		}
	}
}
