package systemproperties

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	_ServiceURN     = "urn:schemas-upnp-org:service:SystemProperties:1"
	_EncodingSchema = "http://schemas.xmlsoap.org/soap/encoding/"
	_EnvelopeSchema = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Service struct {
	ControlEndpoint *url.URL
	EventEndpoint   *url.URL
}

func NewService(deviceUrl *url.URL) *Service {
	c, err := url.Parse(`/SystemProperties/Control`)
	if nil != err {
		panic(err)
	}
	e, err := url.Parse(`/SystemProperties/Event`)
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
	SetString                          *SetStringArgs                          `xml:"u:SetString,omitempty"`
	GetString                          *GetStringArgs                          `xml:"u:GetString,omitempty"`
	Remove                             *RemoveArgs                             `xml:"u:Remove,omitempty"`
	GetWebCode                         *GetWebCodeArgs                         `xml:"u:GetWebCode,omitempty"`
	ProvisionCredentialedTrialAccountX *ProvisionCredentialedTrialAccountXArgs `xml:"u:ProvisionCredentialedTrialAccountX,omitempty"`
	AddAccountX                        *AddAccountXArgs                        `xml:"u:AddAccountX,omitempty"`
	AddOAuthAccountX                   *AddOAuthAccountXArgs                   `xml:"u:AddOAuthAccountX,omitempty"`
	RemoveAccount                      *RemoveAccountArgs                      `xml:"u:RemoveAccount,omitempty"`
	EditAccountPasswordX               *EditAccountPasswordXArgs               `xml:"u:EditAccountPasswordX,omitempty"`
	SetAccountNicknameX                *SetAccountNicknameXArgs                `xml:"u:SetAccountNicknameX,omitempty"`
	RefreshAccountCredentialsX         *RefreshAccountCredentialsXArgs         `xml:"u:RefreshAccountCredentialsX,omitempty"`
	EditAccountMd                      *EditAccountMdArgs                      `xml:"u:EditAccountMd,omitempty"`
	DoPostUpdateTasks                  *DoPostUpdateTasksArgs                  `xml:"u:DoPostUpdateTasks,omitempty"`
	ResetThirdPartyCredentials         *ResetThirdPartyCredentialsArgs         `xml:"u:ResetThirdPartyCredentials,omitempty"`
	EnableRDM                          *EnableRDMArgs                          `xml:"u:EnableRDM,omitempty"`
	GetRDM                             *GetRDMArgs                             `xml:"u:GetRDM,omitempty"`
	ReplaceAccountX                    *ReplaceAccountXArgs                    `xml:"u:ReplaceAccountX,omitempty"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"Envelope"`
	Xmlns         string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          BodyResponse `xml:"Body"`
}
type BodyResponse struct {
	XMLName                            xml.Name                                    `xml:"Body"`
	SetString                          *SetStringResponse                          `xml:"SetStringResponse,omitempty"`
	GetString                          *GetStringResponse                          `xml:"GetStringResponse,omitempty"`
	Remove                             *RemoveResponse                             `xml:"RemoveResponse,omitempty"`
	GetWebCode                         *GetWebCodeResponse                         `xml:"GetWebCodeResponse,omitempty"`
	ProvisionCredentialedTrialAccountX *ProvisionCredentialedTrialAccountXResponse `xml:"ProvisionCredentialedTrialAccountXResponse,omitempty"`
	AddAccountX                        *AddAccountXResponse                        `xml:"AddAccountXResponse,omitempty"`
	AddOAuthAccountX                   *AddOAuthAccountXResponse                   `xml:"AddOAuthAccountXResponse,omitempty"`
	RemoveAccount                      *RemoveAccountResponse                      `xml:"RemoveAccountResponse,omitempty"`
	EditAccountPasswordX               *EditAccountPasswordXResponse               `xml:"EditAccountPasswordXResponse,omitempty"`
	SetAccountNicknameX                *SetAccountNicknameXResponse                `xml:"SetAccountNicknameXResponse,omitempty"`
	RefreshAccountCredentialsX         *RefreshAccountCredentialsXResponse         `xml:"RefreshAccountCredentialsXResponse,omitempty"`
	EditAccountMd                      *EditAccountMdResponse                      `xml:"EditAccountMdResponse,omitempty"`
	DoPostUpdateTasks                  *DoPostUpdateTasksResponse                  `xml:"DoPostUpdateTasksResponse,omitempty"`
	ResetThirdPartyCredentials         *ResetThirdPartyCredentialsResponse         `xml:"ResetThirdPartyCredentialsResponse,omitempty"`
	EnableRDM                          *EnableRDMResponse                          `xml:"EnableRDMResponse,omitempty"`
	GetRDM                             *GetRDMResponse                             `xml:"GetRDMResponse,omitempty"`
	ReplaceAccountX                    *ReplaceAccountXResponse                    `xml:"ReplaceAccountXResponse,omitempty"`
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

type SetStringArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	VariableName string `xml:"VariableName"`
	StringValue  string `xml:"StringValue"`
}
type SetStringResponse struct {
}

func (s *Service) SetString(httpClient *http.Client, args *SetStringArgs) (*SetStringResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetString`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetString: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetString == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.SetString()`)
	}

	return r.Body.SetString, nil
}

type GetStringArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	VariableName string `xml:"VariableName"`
}
type GetStringResponse struct {
	StringValue string `xml:"StringValue"`
}

func (s *Service) GetString(httpClient *http.Client, args *GetStringArgs) (*GetStringResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetString`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetString: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetString == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.GetString()`)
	}

	return r.Body.GetString, nil
}

type RemoveArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	VariableName string `xml:"VariableName"`
}
type RemoveResponse struct {
}

func (s *Service) Remove(httpClient *http.Client, args *RemoveArgs) (*RemoveResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`Remove`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{Remove: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.Remove == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.Remove()`)
	}

	return r.Body.Remove, nil
}

type GetWebCodeArgs struct {
	Xmlns       string `xml:"xmlns:u,attr"`
	AccountType uint32 `xml:"AccountType"`
}
type GetWebCodeResponse struct {
	WebCode string `xml:"WebCode"`
}

func (s *Service) GetWebCode(httpClient *http.Client, args *GetWebCodeArgs) (*GetWebCodeResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetWebCode`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetWebCode: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetWebCode == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.GetWebCode()`)
	}

	return r.Body.GetWebCode, nil
}

type ProvisionCredentialedTrialAccountXArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
	AccountType     uint32 `xml:"AccountType"`
	AccountID       string `xml:"AccountID"`
	AccountPassword string `xml:"AccountPassword"`
}
type ProvisionCredentialedTrialAccountXResponse struct {
	IsExpired  bool   `xml:"IsExpired"`
	AccountUDN string `xml:"AccountUDN"`
}

func (s *Service) ProvisionCredentialedTrialAccountX(httpClient *http.Client, args *ProvisionCredentialedTrialAccountXArgs) (*ProvisionCredentialedTrialAccountXResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ProvisionCredentialedTrialAccountX`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ProvisionCredentialedTrialAccountX: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ProvisionCredentialedTrialAccountX == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.ProvisionCredentialedTrialAccountX()`)
	}

	return r.Body.ProvisionCredentialedTrialAccountX, nil
}

type AddAccountXArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
	AccountType     uint32 `xml:"AccountType"`
	AccountID       string `xml:"AccountID"`
	AccountPassword string `xml:"AccountPassword"`
}
type AddAccountXResponse struct {
	AccountUDN string `xml:"AccountUDN"`
}

func (s *Service) AddAccountX(httpClient *http.Client, args *AddAccountXArgs) (*AddAccountXResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddAccountX`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddAccountX: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddAccountX == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.AddAccountX()`)
	}

	return r.Body.AddAccountX, nil
}

type AddOAuthAccountXArgs struct {
	Xmlns             string `xml:"xmlns:u,attr"`
	AccountType       uint32 `xml:"AccountType"`
	AccountToken      string `xml:"AccountToken"`
	AccountKey        string `xml:"AccountKey"`
	OAuthDeviceID     string `xml:"OAuthDeviceID"`
	AuthorizationCode string `xml:"AuthorizationCode"`
	RedirectURI       string `xml:"RedirectURI"`
	UserIdHashCode    string `xml:"UserIdHashCode"`
	AccountTier       uint32 `xml:"AccountTier"`
}
type AddOAuthAccountXResponse struct {
	AccountUDN      string `xml:"AccountUDN"`
	AccountNickname string `xml:"AccountNickname"`
}

func (s *Service) AddOAuthAccountX(httpClient *http.Client, args *AddOAuthAccountXArgs) (*AddOAuthAccountXResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`AddOAuthAccountX`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{AddOAuthAccountX: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.AddOAuthAccountX == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.AddOAuthAccountX()`)
	}

	return r.Body.AddOAuthAccountX, nil
}

type RemoveAccountArgs struct {
	Xmlns       string `xml:"xmlns:u,attr"`
	AccountType uint32 `xml:"AccountType"`
	AccountID   string `xml:"AccountID"`
}
type RemoveAccountResponse struct {
}

func (s *Service) RemoveAccount(httpClient *http.Client, args *RemoveAccountArgs) (*RemoveAccountResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RemoveAccount`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RemoveAccount: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RemoveAccount == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.RemoveAccount()`)
	}

	return r.Body.RemoveAccount, nil
}

type EditAccountPasswordXArgs struct {
	Xmlns              string `xml:"xmlns:u,attr"`
	AccountType        uint32 `xml:"AccountType"`
	AccountID          string `xml:"AccountID"`
	NewAccountPassword string `xml:"NewAccountPassword"`
}
type EditAccountPasswordXResponse struct {
}

func (s *Service) EditAccountPasswordX(httpClient *http.Client, args *EditAccountPasswordXArgs) (*EditAccountPasswordXResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`EditAccountPasswordX`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{EditAccountPasswordX: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EditAccountPasswordX == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.EditAccountPasswordX()`)
	}

	return r.Body.EditAccountPasswordX, nil
}

type SetAccountNicknameXArgs struct {
	Xmlns           string `xml:"xmlns:u,attr"`
	AccountUDN      string `xml:"AccountUDN"`
	AccountNickname string `xml:"AccountNickname"`
}
type SetAccountNicknameXResponse struct {
}

func (s *Service) SetAccountNicknameX(httpClient *http.Client, args *SetAccountNicknameXArgs) (*SetAccountNicknameXResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`SetAccountNicknameX`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{SetAccountNicknameX: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.SetAccountNicknameX == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.SetAccountNicknameX()`)
	}

	return r.Body.SetAccountNicknameX, nil
}

type RefreshAccountCredentialsXArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	AccountType  uint32 `xml:"AccountType"`
	AccountUID   uint32 `xml:"AccountUID"`
	AccountToken string `xml:"AccountToken"`
	AccountKey   string `xml:"AccountKey"`
}
type RefreshAccountCredentialsXResponse struct {
}

func (s *Service) RefreshAccountCredentialsX(httpClient *http.Client, args *RefreshAccountCredentialsXArgs) (*RefreshAccountCredentialsXResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`RefreshAccountCredentialsX`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{RefreshAccountCredentialsX: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.RefreshAccountCredentialsX == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.RefreshAccountCredentialsX()`)
	}

	return r.Body.RefreshAccountCredentialsX, nil
}

type EditAccountMdArgs struct {
	Xmlns        string `xml:"xmlns:u,attr"`
	AccountType  uint32 `xml:"AccountType"`
	AccountID    string `xml:"AccountID"`
	NewAccountMd string `xml:"NewAccountMd"`
}
type EditAccountMdResponse struct {
}

func (s *Service) EditAccountMd(httpClient *http.Client, args *EditAccountMdArgs) (*EditAccountMdResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`EditAccountMd`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{EditAccountMd: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EditAccountMd == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.EditAccountMd()`)
	}

	return r.Body.EditAccountMd, nil
}

type DoPostUpdateTasksArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type DoPostUpdateTasksResponse struct {
}

func (s *Service) DoPostUpdateTasks(httpClient *http.Client, args *DoPostUpdateTasksArgs) (*DoPostUpdateTasksResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`DoPostUpdateTasks`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{DoPostUpdateTasks: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.DoPostUpdateTasks == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.DoPostUpdateTasks()`)
	}

	return r.Body.DoPostUpdateTasks, nil
}

type ResetThirdPartyCredentialsArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type ResetThirdPartyCredentialsResponse struct {
}

func (s *Service) ResetThirdPartyCredentials(httpClient *http.Client, args *ResetThirdPartyCredentialsArgs) (*ResetThirdPartyCredentialsResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ResetThirdPartyCredentials`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ResetThirdPartyCredentials: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ResetThirdPartyCredentials == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.ResetThirdPartyCredentials()`)
	}

	return r.Body.ResetThirdPartyCredentials, nil
}

type EnableRDMArgs struct {
	Xmlns    string `xml:"xmlns:u,attr"`
	RDMValue bool   `xml:"RDMValue"`
}
type EnableRDMResponse struct {
}

func (s *Service) EnableRDM(httpClient *http.Client, args *EnableRDMArgs) (*EnableRDMResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`EnableRDM`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{EnableRDM: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.EnableRDM == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.EnableRDM()`)
	}

	return r.Body.EnableRDM, nil
}

type GetRDMArgs struct {
	Xmlns string `xml:"xmlns:u,attr"`
}
type GetRDMResponse struct {
	RDMValue bool `xml:"RDMValue"`
}

func (s *Service) GetRDM(httpClient *http.Client, args *GetRDMArgs) (*GetRDMResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`GetRDM`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{GetRDM: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.GetRDM == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.GetRDM()`)
	}

	return r.Body.GetRDM, nil
}

type ReplaceAccountXArgs struct {
	Xmlns              string `xml:"xmlns:u,attr"`
	AccountUDN         string `xml:"AccountUDN"`
	NewAccountID       string `xml:"NewAccountID"`
	NewAccountPassword string `xml:"NewAccountPassword"`
	AccountToken       string `xml:"AccountToken"`
	AccountKey         string `xml:"AccountKey"`
	OAuthDeviceID      string `xml:"OAuthDeviceID"`
}
type ReplaceAccountXResponse struct {
	NewAccountUDN string `xml:"NewAccountUDN"`
}

func (s *Service) ReplaceAccountX(httpClient *http.Client, args *ReplaceAccountXArgs) (*ReplaceAccountXResponse, error) {
	args.Xmlns = _ServiceURN
	r, err := s.exec(`ReplaceAccountX`, httpClient,
		&Envelope{
			EncodingStyle: _EncodingSchema,
			Xmlns:         _EnvelopeSchema,
			Body:          Body{ReplaceAccountX: args},
		})
	if err != nil {
		return nil, err
	}
	if r.Body.ReplaceAccountX == nil {
		return nil, errors.New(`unexpected respose from service calling systemproperties.ReplaceAccountX()`)
	}

	return r.Body.ReplaceAccountX, nil
}
