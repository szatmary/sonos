package gono6

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SpecVersion struct {
	XMLName xml.Name `xml:"specVersion"`
	Major   int      `xml:"major"`
	Minor   int      `xml:"minor"`
}

type Service struct {
	XMLName     xml.Name `xml:"service"`
	ServiceType string   `xml:"serviceType"`
	ServiceId   string   `xml:"serviceId"`
	ControlURL  string   `xml:"controlURL"`
	EventSubURL string   `xml:"eventSubURL"`
	SCPDURL     string   `xml:"SCPDURL"`
}

type Icon struct {
	XMLName  xml.Name `xml:"icon"`
	Id       string   `xml:"id"`
	Mimetype string   `xml:"mimetype"`
	Width    int      `xml:"width"`
	Height   int      `xml:"height"`
	Depth    int      `xml:"depth"`
	Url      url.URL  `xml:"url"`
}

type Device struct {
	XMLName                 xml.Name  `xml:"device"`
	DeviceType              string    `xml:"deviceType"`
	FriendlyName            string    `xml:"friendlyName"`
	Manufacturer            string    `xml:"manufacturer"`
	ManufacturerURL         string    `xml:"manufacturerURL"`
	ModelNumber             string    `xml:"modelNumber"`
	ModelDescription        string    `xml:"modelDescription"`
	ModelName               string    `xml:"modelName"`
	ModelURL                string    `xml:"modelURL"`
	SoftwareVersion         string    `xml:"softwareVersion"`
	SwGen                   string    `xml:"swGen"`
	HardwareVersion         string    `xml:"hardwareVersion"`
	SerialNum               string    `xml:"serialNum"`
	MACAddress              string    `xml:"MACAddress"`
	UDN                     string    `xml:"UDN"`
	Icons                   []Icon    `xml:"iconList>icon"`
	MinCompatibleVersion    string    `xml:"minCompatibleVersion"`
	LegacyCompatibleVersion string    `xml:"legacyCompatibleVersion"`
	ApiVersion              string    `xml:"apiVersion"`
	MinApiVersion           string    `xml:"minApiVersion"`
	DisplayVersion          string    `xml:"displayVersion"`
	ExtraVersion            string    `xml:"extraVersion"`
	RoomName                string    `xml:"roomName"`
	DisplayName             string    `xml:"displayName"`
	ZoneType                int       `xml:"zoneType"`
	Feature1                string    `xml:"feature1"`
	Feature2                string    `xml:"feature2"`
	Feature3                string    `xml:"feature3"`
	Seriesid                string    `xml:"seriesid"`
	Variant                 int       `xml:"variant"`
	InternalSpeakerSize     float32   `xml:"internalSpeakerSize"`
	BassExtension           float32   `xml:"bassExtension"`
	SatGainOffset           float32   `xml:"satGainOffset"`
	Memory                  int       `xml:"memory"`
	Flash                   int       `xml:"flash"`
	FlashRepartitioned      int       `xml:"flashRepartitioned"`
	AmpOnTime               int       `xml:"ampOnTime"`
	RetailMode              int       `xml:"retailMode"`
	Services                []Service `xml:"serviceList>service"`
	Devices                 []Device  `xml:"deviceList>device"`
}

type Root struct {
	XMLName     xml.Name    `xml:"root"`
	Xmlns       string      `xml:"xmlns,attr"`
	SpecVersion SpecVersion `xml:"specVersion"`
	Device      Device      `xml:"device"`
}

type ZonePlayer struct {
	Root                 *Root
	HttpClient           *http.Client
	DeviceDescriptionURL *url.URL
	// services
	AlarmClock            *AlarmClockService
	AVTransport           *AVTransportService
	ConnectionManager     *ConnectionManagerService
	ContentDirectory      *ContentDirectoryService
	DeviceProperties      *DevicePropertiesService
	GroupManagement       *GroupManagementService
	GroupRenderingControl *GroupRenderingControlService
	MusicServices         *MusicServicesService
	QPlay                 *QPlayService
	Queue                 *QueueService
	RenderingControl      *RenderingControlService
	SystemProperties      *SystemPropertiesService
	VirtualLineIn         *VirtualLineInService
	ZoneGroupTopology     *ZoneGroupTopologyService
	// Callbaacks
	ZoneGroupStateCallback func(*ZoneGroupState)
}

func NewZonePlayer(deviceDescriptionURL *url.URL) (*ZonePlayer, error) {
	zp := ZonePlayer{
		Root:                  &Root{},
		HttpClient:            &http.Client{},
		DeviceDescriptionURL:  deviceDescriptionURL,
		AlarmClock:            NewAlarmClockService(deviceDescriptionURL),
		AVTransport:           NewAVTransportService(deviceDescriptionURL),
		ConnectionManager:     NewConnectionManagerService(deviceDescriptionURL),
		ContentDirectory:      NewContentDirectoryService(deviceDescriptionURL),
		DeviceProperties:      NewDevicePropertiesService(deviceDescriptionURL),
		GroupManagement:       NewGroupManagementService(deviceDescriptionURL),
		GroupRenderingControl: NewGroupRenderingControlService(deviceDescriptionURL),
		MusicServices:         NewMusicServicesService(deviceDescriptionURL),
		QPlay:                 NewQPlayService(deviceDescriptionURL),
		Queue:                 NewQueueService(deviceDescriptionURL),
		RenderingControl:      NewRenderingControlService(deviceDescriptionURL),
		SystemProperties:      NewSystemPropertiesService(deviceDescriptionURL),
		VirtualLineIn:         NewVirtualLineInService(deviceDescriptionURL),
		ZoneGroupTopology:     NewZoneGroupTopologyService(deviceDescriptionURL),
	}

	resp, err := zp.HttpClient.Get(zp.DeviceDescriptionURL.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = xml.Unmarshal(body, zp.Root)
	if err != nil {
		return nil, err
	}

	return &zp, nil
}

// go http library does sone stuff that make upnp scbscriptions difficult
// hecce here we just open a tcp connection and fire off the request manually
func (s *Sonos) subscribe(zp *ZonePlayer) error {
	// sub := func(serialNum string, eventEndpoint *url.URL) error {
	// 	host := eventEndpoint.Hostname() + ":" + eventEndpoint.Port()
	// 	conn, err := net.Dial("tcp", host)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// fmt.Sprintf("%s:%d", GetLocalAddress(eventEndpoint.Hostname()), s.httpListener.Addr().Port),
	// 	callback := url.URL{
	// 		Scheme:   "http",
	// 		Host:     s.httpListener.Addr().String(),
	// 		Fragment: serialNum,
	// 	}

	// 	var req string
	// 	req += fmt.Sprintf("SUBSCRIBE %s HTTP/1.0\r\n", eventEndpoint.Path)
	// 	req += fmt.Sprintf("HOST: %s\r\n", host)
	// 	req += fmt.Sprintf("USER-AGENT: Unknown UPnP/1.0 Gonos/1.0\r\n")
	// 	req += fmt.Sprintf("CALLBACK: <%s>\r\n", callback.String())
	// 	req += fmt.Sprintf("NT: upnp:event\r\n")
	// 	req += fmt.Sprintf("TIMEOUT: Second-300\r\n")
	// 	fmt.Fprintf(conn, req+"\r\n")
	// 	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer res.Body.Close()
	// 	body, err := ioutil.ReadAll(res.Body)
	// 	if 200 != res.StatusCode {
	// 		fmt.Printf("%v\n", res)
	// 		return errors.New(string(body))
	// 	}

	// 	return nil
	// }

	callbackURL := url.URL{
		Scheme:   "http",
		Fragment: zp.SerialNum(),
		Host:     s.httpListener.Addr().String(),
	}
	callbackURL.Path = zp.AlarmClock.EventEndpoint.Path
	zp.AlarmClock.AlarmClockSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.AVTransport.AVTransportSubscribe(callbackURL)
	callbackURL.Path = zp.ConnectionManager.EventEndpoint.Path
	zp.ConnectionManager.ConnectionManagerSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.ContentDirectory.ContentDirectorySubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.DeviceProperties.DevicePropertiesSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.GroupManagement.GroupManagementSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.GroupRenderingControl.GroupRenderingControlSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.MusicServices.MusicServicesSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.Queue.QueueSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.RenderingControl.RenderingControlSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.SystemProperties.SystemPropertiesSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.VirtualLineIn.VirtualLineInSubscribe(callbackURL)
	callbackURL.Path = zp.AVTransport.EventEndpoint.Path
	zp.ZoneGroupTopology.ZoneGroupTopologySubscribe(callbackURL)
	return nil
}

func (zp *ZonePlayer) unsubscribe() error {
	return nil
}

func (z *ZonePlayer) RoomName() string {
	return z.Root.Device.RoomName
}

func (z *ZonePlayer) ModelName() string {
	return z.Root.Device.ModelName
}

func (z *ZonePlayer) HardwareVersion() string {
	return z.Root.Device.HardwareVersion
}

func (z *ZonePlayer) SerialNum() string {
	return z.Root.Device.SerialNum
}

func (z *ZonePlayer) IsCoordinator() bool {
	zoneGroupState, err := z.GetZoneGroupState()
	if err != nil {
		return false
	}
	for _, group := range zoneGroupState.ZoneGroups {
		if "uuid:"+group.Coordinator == z.Root.Device.UDN {
			return true
		}
	}

	return false
}

// Convience functions

func (z *ZonePlayer) GetZoneGroupState() (*ZoneGroupState, error) {
	zoneGroupStateResponse, err := z.ZoneGroupTopology.GetZoneGroupState(z.HttpClient, &ZoneGroupTopologyGetZoneGroupStateArgs{})
	if err != nil {
		return nil, err
	}
	var zoneGroupState ZoneGroupState
	err = xml.Unmarshal([]byte(zoneGroupStateResponse.ZoneGroupState), &zoneGroupState)
	if err != nil {
		return nil, err
	}

	return &zoneGroupState, nil
}

func (z *ZonePlayer) GetVolume() (int, error) {
	res, err := z.RenderingControl.GetVolume(z.HttpClient, &RenderingControlGetVolumeArgs{Channel: "Master"})
	if err != nil {
		return 0, err
	}

	return int(res.CurrentVolume), err
}

func (z *ZonePlayer) SetVolume(desiredVolume int) error {
	_, err := z.RenderingControl.SetVolume(z.HttpClient, &RenderingControlSetVolumeArgs{
		Channel:       "Master",
		DesiredVolume: uint16(desiredVolume),
	})
	return err
}

func (z *ZonePlayer) Play() error {
	_, err := z.AVTransport.Play(z.HttpClient, &AVTransportPlayArgs{
		Speed: "1",
	})
	return err
}

func (z *ZonePlayer) SetAVTransportURI(url string) error {
	_, err := z.AVTransport.SetAVTransportURI(z.HttpClient, &AVTransportSetAVTransportURIArgs{
		CurrentURI: url,
	})
	return err
}

func (zp *ZonePlayer) EventCallback(evt interface{}) {
	switch e := evt.(type) {
	default:
		fmt.Sprintf("Unhandeld type %T", evt)
	case ZoneGroupTopologyZoneGroupState:
		var zoneGroupState ZoneGroupState
		err := xml.Unmarshal([]byte(e), &zoneGroupState)
		if err == nil && zp.ZoneGroupStateCallback != nil {
			zp.ZoneGroupStateCallback(&zoneGroupState)
		}
	case DevicePropertiesSettingsReplicationState:
	}
}

// Event handlers
func (z *ZonePlayer) AVTransportLastChangeEvent(evt string) {
}
func (z *ZonePlayer) AlarmClockTimeZoneEvent(evt string) {
}
func (z *ZonePlayer) AlarmClockTimeServerEvent(evt string) {
}
func (z *ZonePlayer) AlarmClockTimeGenerationEvent(evt uint32) {
}
func (z *ZonePlayer) AlarmClockAlarmListVersionEvent(evt string) {
}
func (z *ZonePlayer) AlarmClockDailyIndexRefreshTimeEvent(evt string) {
}
func (z *ZonePlayer) AlarmClockTimeFormatEvent(evt string) {
}
func (z *ZonePlayer) AlarmClockDateFormatEvent(evt string) {
}
func (z *ZonePlayer) ConnectionManagerSourceProtocolInfoEvent(evt string) {
}
func (z *ZonePlayer) ConnectionManagerSinkProtocolInfoEvent(evt string) {
}
func (z *ZonePlayer) ConnectionManagerCurrentConnectionIDsEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectorySystemUpdateIDEvent(evt uint32) {
}
func (z *ZonePlayer) ContentDirectoryContainerUpdateIDsEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectoryShareIndexInProgressEvent(evt bool) {
}
func (z *ZonePlayer) ContentDirectoryShareIndexLastErrorEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectoryUserRadioUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectorySavedQueuesUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectoryShareListUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectoryRecentlyPlayedUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectoryBrowseableEvent(evt bool) {
}
func (z *ZonePlayer) ContentDirectoryRadioFavoritesUpdateIDEvent(evt uint32) {
}
func (z *ZonePlayer) ContentDirectoryRadioLocationUpdateIDEvent(evt uint32) {
}
func (z *ZonePlayer) ContentDirectoryFavoritesUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) ContentDirectoryFavoritePresetsUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesSettingsReplicationStateEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesZoneNameEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesIconEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesConfigurationEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesInvisibleEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesIsZoneBridgeEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesAirPlayEnabledEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesSupportsAudioInEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesSupportsAudioClipEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesIsIdleEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesMoreInfoEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesChannelMapSetEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesHTSatChanMapSetEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesHTFreqEvent(evt uint32) {
}
func (z *ZonePlayer) DevicePropertiesHTBondedZoneCommitStateEvent(evt uint32) {
}
func (z *ZonePlayer) DevicePropertiesOrientationEvent(evt int32) {
}
func (z *ZonePlayer) DevicePropertiesLastChangedPlayStateEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesRoomCalibrationStateEvent(evt int32) {
}
func (z *ZonePlayer) DevicePropertiesAvailableRoomCalibrationEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesTVConfigurationErrorEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesHdmiCecAvailableEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesWirelessModeEvent(evt uint32) {
}
func (z *ZonePlayer) DevicePropertiesWirelessLeafOnlyEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesHasConfiguredSSIDEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesChannelFreqEvent(evt uint32) {
}
func (z *ZonePlayer) DevicePropertiesBehindWifiExtenderEvent(evt uint32) {
}
func (z *ZonePlayer) DevicePropertiesWifiEnabledEvent(evt bool) {
}
func (z *ZonePlayer) DevicePropertiesConfigModeEvent(evt string) {
}
func (z *ZonePlayer) DevicePropertiesSecureRegStateEvent(evt uint32) {
}
func (z *ZonePlayer) DevicePropertiesVoiceConfigStateEvent(evt uint32) {
}
func (z *ZonePlayer) DevicePropertiesMicEnabledEvent(evt uint32) {
}
func (z *ZonePlayer) GroupRenderingControlGroupMuteEvent(evt bool) {
}
func (z *ZonePlayer) GroupRenderingControlGroupVolumeEvent(evt uint16) {
}
func (z *ZonePlayer) GroupRenderingControlGroupVolumeChangeableEvent(evt bool) {
}
func (z *ZonePlayer) GroupManagementGroupCoordinatorIsLocalEvent(evt bool) {
}
func (z *ZonePlayer) GroupManagementLocalGroupUUIDEvent(evt string) {
}
func (z *ZonePlayer) GroupManagementVirtualLineInGroupIDEvent(evt string) {
}
func (z *ZonePlayer) GroupManagementResetVolumeAfterEvent(evt bool) {
}
func (z *ZonePlayer) GroupManagementVolumeAVTransportURIEvent(evt string) {
}
func (z *ZonePlayer) MusicServicesServiceListVersionEvent(evt string) {
}
func (z *ZonePlayer) QueueLastChangeEvent(evt string) {
}
func (z *ZonePlayer) RenderingControlLastChangeEvent(evt string) {
}
func (z *ZonePlayer) SystemPropertiesCustomerIDEvent(evt string) {
}
func (z *ZonePlayer) SystemPropertiesUpdateIDEvent(evt uint32) {
}
func (z *ZonePlayer) SystemPropertiesUpdateIDXEvent(evt uint32) {
}
func (z *ZonePlayer) SystemPropertiesVoiceUpdateIDEvent(evt uint32) {
}
func (z *ZonePlayer) SystemPropertiesThirdPartyHashEvent(evt string) {
}
func (z *ZonePlayer) VirtualLineInLastChangeEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyAvailableSoftwareUpdateEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyZoneGroupStateEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyThirdPartyMediaServersXEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyAlarmRunSequenceEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyMuseHouseholdIdEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyZoneGroupNameEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyZoneGroupIDEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyZonePlayerUUIDsInGroupEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyAreasUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologySourceAreasUpdateIDEvent(evt string) {
}
func (z *ZonePlayer) ZoneGroupTopologyNetsettingsUpdateIDEvent(evt string) {
}
