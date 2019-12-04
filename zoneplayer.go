package sonos

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"

	avt "github.com/szatmary/sonos/AVTransport"
	clk "github.com/szatmary/sonos/AlarmClock"
	con "github.com/szatmary/sonos/ConnectionManager"
	dir "github.com/szatmary/sonos/ContentDirectory"
	dev "github.com/szatmary/sonos/DeviceProperties"
	gmn "github.com/szatmary/sonos/GroupManagement"
	rcg "github.com/szatmary/sonos/GroupRenderingControl"
	mus "github.com/szatmary/sonos/MusicServices"
	ply "github.com/szatmary/sonos/QPlay"
	que "github.com/szatmary/sonos/Queue"
	ren "github.com/szatmary/sonos/RenderingControl"
	sys "github.com/szatmary/sonos/SystemProperties"
	vli "github.com/szatmary/sonos/VirtualLineIn"
	zgt "github.com/szatmary/sonos/ZoneGroupTopology"
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
	AlarmClock            *clk.Service
	AVTransport           *avt.Service
	ConnectionManager     *con.Service
	ContentDirectory      *dir.Service
	DeviceProperties      *dev.Service
	GroupManagement       *gmn.Service
	GroupRenderingControl *rcg.Service
	MusicServices         *mus.Service
	QPlay                 *ply.Service
	Queue                 *que.Service
	RenderingControl      *ren.Service
	SystemProperties      *sys.Service
	VirtualLineIn         *vli.Service
	ZoneGroupTopology     *zgt.Service
}

func NewZonePlayer(deviceDescriptionURL *url.URL) (*ZonePlayer, error) {
	zp := ZonePlayer{
		Root:                  &Root{},
		HttpClient:            &http.Client{},
		DeviceDescriptionURL:  deviceDescriptionURL,
		AlarmClock:            clk.NewService(deviceDescriptionURL),
		AVTransport:           avt.NewService(deviceDescriptionURL),
		ConnectionManager:     con.NewService(deviceDescriptionURL),
		ContentDirectory:      dir.NewService(deviceDescriptionURL),
		DeviceProperties:      dev.NewService(deviceDescriptionURL),
		GroupManagement:       gmn.NewService(deviceDescriptionURL),
		GroupRenderingControl: rcg.NewService(deviceDescriptionURL),
		MusicServices:         mus.NewService(deviceDescriptionURL),
		QPlay:                 ply.NewService(deviceDescriptionURL),
		Queue:                 que.NewService(deviceDescriptionURL),
		RenderingControl:      ren.NewService(deviceDescriptionURL),
		SystemProperties:      sys.NewService(deviceDescriptionURL),
		VirtualLineIn:         vli.NewService(deviceDescriptionURL),
		ZoneGroupTopology:     zgt.NewService(deviceDescriptionURL),
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

func (z *ZonePlayer) GetZoneGroupState() (*ZoneGroupState, error) {
	zoneGroupStateResponse, err := z.ZoneGroupTopology.GetZoneGroupState(z.HttpClient, &zgt.GetZoneGroupStateArgs{})
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
	res, err := z.RenderingControl.GetVolume(z.HttpClient, &ren.GetVolumeArgs{Channel: "Master"})
	if err != nil {
		return 0, err
	}

	return int(res.CurrentVolume), err
}

func (z *ZonePlayer) SetVolume(desiredVolume int) error {
	_, err := z.RenderingControl.SetVolume(z.HttpClient, &ren.SetVolumeArgs{
		Channel:       "Master",
		DesiredVolume: uint16(desiredVolume),
	})
	return err
}
