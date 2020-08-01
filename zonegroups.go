package sonos

import "encoding/xml"

// type Device struct {
// 	XMLName                 xml.Name `xml:"Device"`
// 	UUID                    string   `xml:"UUID"`
// 	Location                string   `xml:"Location"`
// 	ZoneName                string   `xml:"ZoneName"`
// 	Icon                    string   `xml:"Icon"`
// 	Configuration           string   `xml:"Configuration"`
// 	SoftwareVersion         string   `xml:"SoftwareVersion"`
// 	SWGen                   string   `xml:"SWGen"`
// 	MinCompatibleVersion    string   `xml:"MinCompatibleVersion"`
// 	LegacyCompatibleVersion string   `xml:"LegacyCompatibleVersion"`
// 	BootSeq                 string   `xml:"BootSeq"`
// 	TVConfigurationError    string   `xml:"TVConfigurationError"`
// 	HdmiCecAvailable        string   `xml:"HdmiCecAvailable"`
// 	WirelessMode            string   `xml:"WirelessMode"`
// 	WirelessLeafOnly        string   `xml:"WirelessLeafOnly"`
// 	HasConfiguredSSID       string   `xml:"HasConfiguredSSID"`
// 	ChannelFreq             string   `xml:"ChannelFreq"`
// 	BehindWifiExtender      string   `xml:"BehindWifiExtender"`
// 	WifiEnabled             string   `xml:"WifiEnabled"`
// 	Orientation             string   `xml:"Orientation"`
// 	RoomCalibrationState    string   `xml:"RoomCalibrationState"`
// 	SecureRegState          string   `xml:"SecureRegState"`
// 	VoiceConfigState        string   `xml:"VoiceConfigState"`
// 	MicEnabled              string   `xml:"MicEnabled"`
// 	AirPlayEnabled          string   `xml:"AirPlayEnabled"`
// 	IdleState               string   `xml:"IdleState"`
// 	MoreInfo                string   `xml:"MoreInfo"`
// }

// type Satellite struct {
// 	XMLName                 xml.Name `xml:"Satellite"`
// 	UUID                    string   `xml:"UUID"`
// 	Location                string   `xml:"Location"`
// 	ZoneName                string   `xml:"ZoneName"`
// 	Icon                    string   `xml:"Icon"`
// 	Configuration           string   `xml:"Configuration"`
// 	SoftwareVersion         string   `xml:"SoftwareVersion"`
// 	SWGen                   string   `xml:"SWGen"`
// 	MinCompatibleVersion    string   `xml:"MinCompatibleVersion"`
// 	LegacyCompatibleVersion string   `xml:"LegacyCompatibleVersion"`
// 	BootSeq                 string   `xml:"BootSeq"`
// 	TVConfigurationError    string   `xml:"TVConfigurationError"`
// 	HdmiCecAvailable        string   `xml:"HdmiCecAvailable"`
// 	WirelessMode            string   `xml:"WirelessMode"`
// 	WirelessLeafOnly        string   `xml:"WirelessLeafOnly"`
// 	HasConfiguredSSID       string   `xml:"HasConfiguredSSID"`
// 	ChannelFreq             string   `xml:"ChannelFreq"`
// 	BehindWifiExtender      string   `xml:"BehindWifiExtender"`
// 	WifiEnabled             string   `xml:"WifiEnabled"`
// 	Orientation             string   `xml:"Orientation"`
// 	RoomCalibrationState    string   `xml:"RoomCalibrationState"`
// 	SecureRegState          string   `xml:"SecureRegState"`
// 	VoiceConfigState        string   `xml:"VoiceConfigState"`
// 	MicEnabled              string   `xml:"MicEnabled"`
// 	AirPlayEnabled          string   `xml:"AirPlayEnabled"`
// 	IdleState               string   `xml:"IdleState"`
// 	MoreInfo                string   `xml:"MoreInfo"`
// }

type ZoneGroupMember struct {
	XMLName                 xml.Name `xml:"ZoneGroupMember"`
	UUID                    string   `xml:"UUID"`
	Location                string   `xml:"Location"`
	ZoneName                string   `xml:"ZoneName"`
	Icon                    string   `xml:"Icon"`
	Configuration           string   `xml:"Configuration"`
	SoftwareVersion         string   `xml:"SoftwareVersion"`
	SWGen                   string   `xml:"SWGen"`
	MinCompatibleVersion    string   `xml:"MinCompatibleVersion"`
	LegacyCompatibleVersion string   `xml:"LegacyCompatibleVersion"`
	BootSeq                 string   `xml:"BootSeq"`
	TVConfigurationError    string   `xml:"TVConfigurationError"`
	HdmiCecAvailable        string   `xml:"HdmiCecAvailable"`
	WirelessMode            string   `xml:"WirelessMode"`
	WirelessLeafOnly        string   `xml:"WirelessLeafOnly"`
	HasConfiguredSSID       string   `xml:"HasConfiguredSSID"`
	ChannelFreq             string   `xml:"ChannelFreq"`
	BehindWifiExtender      string   `xml:"BehindWifiExtender"`
	WifiEnabled             string   `xml:"WifiEnabled"`
	Orientation             string   `xml:"Orientation"`
	RoomCalibrationState    string   `xml:"RoomCalibrationState"`
	SecureRegState          string   `xml:"SecureRegState"`
	VoiceConfigState        string   `xml:"VoiceConfigState"`
	MicEnabled              string   `xml:"MicEnabled"`
	AirPlayEnabled          string   `xml:"AirPlayEnabled"`
	IdleState               string   `xml:"IdleState"`
	MoreInfo                string   `xml:"MoreInfo"`
	Satellite               []Device `xml:"Satellite>Device"`
	VanishedDevice          []Device `xml:"VanishedDevices>Device"`
}

type ZoneGroup struct {
	XMLName         xml.Name          `xml:"ZoneGroup"`
	Coordinator     string            `xml:"Coordinator,attr"`
	ID              string            `xml:"ID,attr"`
	ZoneGroupMember []ZoneGroupMember `xml:"ZoneGroupMember"`
}

type ZoneGroupState struct {
	XMLName    xml.Name    `xml:"ZoneGroupState"`
	ZoneGroups []ZoneGroup `xml:"ZoneGroups>ZoneGroup"`
}
