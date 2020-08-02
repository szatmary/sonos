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
	ZoneGroupState      func(*ZoneGroupState)
	NetsettingsUpdateID func(string)
	SourceAreasUpdateID func(string)
	AreasUpdateID       func(string)
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
		Scheme:   "http",
		Host:     host,
		RawQuery: "sn=" + zp.SerialNum(),
		Path:     service.EventEndpoint().Path,
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
	// fmt.Printf("%v\n", req)
	fmt.Fprintf(conn, req+"\r\n")
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	// fmt.Printf("%v\n", body)
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

// func (zp *ZonePlayer) EventCallback(evt interface{}) {
// 	switch e := evt.(type) {
// 	default:
// 		fmt.Printf("Unhandeld event %T\n", evt)
// 	case ZoneGroupTopologyZoneGroupState:
// 		var zoneGroupState ZoneGroupState
// 		err := xml.Unmarshal([]byte(e), &zoneGroupState)
// 		if err == nil && zp.ZoneGroupStateCallback != nil {
// 			zp.ZoneGroupStateCallback(&zoneGroupState)
// 		}
// 		// type AlarmClockTimeZone string
// 		// type AlarmClockTimeServer string
// 		// type AlarmClockTimeGeneration uint32
// 		// type AlarmClockAlarmListVersion string
// 		// type AlarmClockDailyIndexRefreshTime string
// 		// type AlarmClockTimeFormat string
// 		// type AlarmClockDateFormat string

// 		// type AVTransportLastChange string

// 		// type ConnectionManagerSourceProtocolInfo string
// 		// type ConnectionManagerSinkProtocolInfo string
// 		// type ConnectionManagerCurrentConnectionIDs string

// 		// type ContentDirectorySystemUpdateID uint32
// 		// type ContentDirectoryContainerUpdateIDs string
// 		// type ContentDirectoryShareIndexInProgress bool
// 		// type ContentDirectoryShareIndexLastError string
// 		// type ContentDirectoryUserRadioUpdateID string
// 		// type ContentDirectorySavedQueuesUpdateID string
// 		// type ContentDirectoryShareListUpdateID string
// 		// type ContentDirectoryRecentlyPlayedUpdateID string
// 		// type ContentDirectoryBrowseable bool
// 		// type ContentDirectoryRadioFavoritesUpdateID uint32
// 		// type ContentDirectoryRadioLocationUpdateID uint32
// 		// type ContentDirectoryFavoritesUpdateID string
// 		// type ContentDirectoryFavoritePresetsUpdateID string

// 		// type ZoneGroupTopologyAvailableSoftwareUpdate string
// 		// type ZoneGroupTopologyZoneGroupState string
// 		// type ZoneGroupTopologyThirdPartyMediaServersX string
// 		// type ZoneGroupTopologyAlarmRunSequence string
// 		// type ZoneGroupTopologyMuseHouseholdId string
// 		// type ZoneGroupTopologyZoneGroupName string
// 		// type ZoneGroupTopologyZoneGroupID string
// 		// type ZoneGroupTopologyZonePlayerUUIDsInGroup string
// 		// type ZoneGroupTopologyAreasUpdateID string
// 		// type ZoneGroupTopologySourceAreasUpdateID string
// 		// type ZoneGroupTopologyNetsettingsUpdateID string
// 	}
// }
