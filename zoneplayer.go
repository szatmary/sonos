package sonos

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
)

type SonosService interface {
	ControlEndpoint() *url.URL
	EventEndpoint() *url.URL
}

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
	// Services             []SonosService
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

// go http library does some stuff that make upnp scbscriptions difficult
// hecce here we just open a tcp connection and fire off the request manually
func (s *Sonos) SubscribeAll(zp *ZonePlayer) {
	s.Subscribe(zp, zp.AlarmClock)
	s.Subscribe(zp, zp.AVTransport)
	s.Subscribe(zp, zp.ConnectionManager)
	s.Subscribe(zp, zp.ContentDirectory)
	s.Subscribe(zp, zp.DeviceProperties)
	s.Subscribe(zp, zp.GroupManagement)
	s.Subscribe(zp, zp.GroupRenderingControl)
	s.Subscribe(zp, zp.MusicServices)
	s.Subscribe(zp, zp.Queue)
	s.Subscribe(zp, zp.RenderingControl)
	s.Subscribe(zp, zp.SystemProperties)
	s.Subscribe(zp, zp.VirtualLineIn)
	s.Subscribe(zp, zp.ZoneGroupTopology)
	s.Subscribe(zp, zp.RenderingControl)
}

func (s *Sonos) Subscribe(zp *ZonePlayer, service SonosService) error {

	conn, err := net.Dial("tcp", service.EventEndpoint().Host)
	host := fmt.Sprintf("%s:%d", conn.LocalAddr().(*net.TCPAddr).IP.String(), s.httpListener.Addr().(*net.TCPAddr).Port)
	calbackUrl := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   service.EventEndpoint().Path,
	}
	var req string
	req += fmt.Sprintf("SUBSCRIBE %s HTTP/1.0\r\n", service.EventEndpoint().String())
	req += fmt.Sprintf("HOST: %s\r\n", service.EventEndpoint().Host)
	req += fmt.Sprintf("USER-AGENT: Unknown UPnP/1.0 sonos.szatmary.com.github/2.0\r\n")
	req += fmt.Sprintf("CALLBACK: <%s>\r\n", calbackUrl.String())
	req += fmt.Sprintf("NT: upnp:event\r\n")
	req += fmt.Sprintf("TIMEOUT: Second-300\r\n")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", req)
	fmt.Fprintf(conn, req+"\r\n")
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%v\n", body)
	if 200 != res.StatusCode {
		fmt.Printf("%v\n", res)
		return errors.New(string(body))
	}
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
	// fmt.Printf("GetZoneGroupState %v %v\n", zoneGroupState, err)
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
	// fmt.Printf("z.ZoneGroupTopology.GetZoneGroupState %v %v\n", zoneGroupStateResponse, err)
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
	fmt.Fprintf(os.Stderr, "AVTransportLastChangeEvent: %v\n", evt)
}
func (z *ZonePlayer) AlarmClockTimeZoneEvent(evt string) {
	fmt.Fprintf(os.Stderr, "AlarmClockTimeZoneEvent: %v\n", evt)
}
func (z *ZonePlayer) AlarmClockTimeServerEvent(evt string) {
	fmt.Fprintf(os.Stderr, "AlarmClockTimeServerEvent: %v\n", evt)
}
func (z *ZonePlayer) AlarmClockTimeGenerationEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "AlarmClockTimeGenerationEvent: %v\n", evt)
}
func (z *ZonePlayer) AlarmClockAlarmListVersionEvent(evt string) {
	fmt.Fprintf(os.Stderr, "AlarmClockAlarmListVersionEvent: %v\n", evt)
}
func (z *ZonePlayer) AlarmClockDailyIndexRefreshTimeEvent(evt string) {
	fmt.Fprintf(os.Stderr, "AlarmClockDailyIndexRefreshTimeEvent: %v\n", evt)
}
func (z *ZonePlayer) AlarmClockTimeFormatEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) AlarmClockDateFormatEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ConnectionManagerSourceProtocolInfoEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ConnectionManagerSinkProtocolInfoEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ConnectionManagerCurrentConnectionIDsEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectorySystemUpdateIDEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryContainerUpdateIDsEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryShareIndexInProgressEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryShareIndexLastErrorEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryUserRadioUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectorySavedQueuesUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryShareListUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryRecentlyPlayedUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryBrowseableEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryRadioFavoritesUpdateIDEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryRadioLocationUpdateIDEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryFavoritesUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ContentDirectoryFavoritePresetsUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesSettingsReplicationStateEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesZoneNameEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesIconEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesConfigurationEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesInvisibleEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesIsZoneBridgeEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesAirPlayEnabledEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesSupportsAudioInEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesSupportsAudioClipEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesIsIdleEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesMoreInfoEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesChannelMapSetEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesHTSatChanMapSetEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesHTFreqEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesHTBondedZoneCommitStateEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesOrientationEvent(evt int32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesLastChangedPlayStateEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesRoomCalibrationStateEvent(evt int32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesAvailableRoomCalibrationEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesTVConfigurationErrorEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesHdmiCecAvailableEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesWirelessModeEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesWirelessLeafOnlyEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesHasConfiguredSSIDEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesChannelFreqEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesBehindWifiExtenderEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesWifiEnabledEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesConfigModeEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesSecureRegStateEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesVoiceConfigStateEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) DevicePropertiesMicEnabledEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupRenderingControlGroupMuteEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupRenderingControlGroupVolumeEvent(evt uint16) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupRenderingControlGroupVolumeChangeableEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupManagementGroupCoordinatorIsLocalEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupManagementLocalGroupUUIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupManagementVirtualLineInGroupIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupManagementResetVolumeAfterEvent(evt bool) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) GroupManagementVolumeAVTransportURIEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) MusicServicesServiceListVersionEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) QueueLastChangeEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) RenderingControlLastChangeEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) SystemPropertiesCustomerIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) SystemPropertiesUpdateIDEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) SystemPropertiesUpdateIDXEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) SystemPropertiesVoiceUpdateIDEvent(evt uint32) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) SystemPropertiesThirdPartyHashEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) VirtualLineInLastChangeEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyAvailableSoftwareUpdateEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyZoneGroupStateEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyThirdPartyMediaServersXEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyAlarmRunSequenceEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyMuseHouseholdIdEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyZoneGroupNameEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyZoneGroupIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyZonePlayerUUIDsInGroupEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyAreasUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologySourceAreasUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
func (z *ZonePlayer) ZoneGroupTopologyNetsettingsUpdateIDEvent(evt string) {
	fmt.Fprintf(os.Stderr, "%v\n", evt)
}
