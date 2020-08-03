package sonos

import (
	"fmt"
	"os"
)

func dispatchAlarmClockTimeZone(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchAlarmClockTimeZone %v\n", value)
}

func dispatchAlarmClockTimeServer(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchAlarmClockTimeServer %v\n", value)
}

func dispatchAlarmClockTimeGeneration(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchAlarmClockTimeGeneration %v\n", value)
}

func dispatchAlarmClockAlarmListVersion(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchAlarmClockAlarmListVersion %v\n", value)
}

func dispatchAlarmClockDailyIndexRefreshTime(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchAlarmClockDailyIndexRefreshTime %v\n", value)
}

func dispatchAlarmClockTimeFormat(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchAlarmClockTimeFormat %v\n", value)
}

func dispatchAlarmClockDateFormat(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchAlarmClockDateFormat %v\n", value)
}

func dispatchAVTransportLastChange(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchAVTransportLastChange %v\n", value)
}

func dispatchConnectionManagerSourceProtocolInfo(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchConnectionManagerSourceProtocolInfo %v\n", value)
}

func dispatchConnectionManagerSinkProtocolInfo(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchConnectionManagerSinkProtocolInfo %v\n", value)
}

func dispatchConnectionManagerCurrentConnectionIDs(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchConnectionManagerCurrentConnectionIDs %v\n", value)
}

func dispatchContentDirectorySystemUpdateID(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectorySystemUpdateID %v\n", value)
}

func dispatchContentDirectoryContainerUpdateIDs(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryContainerUpdateIDs %v\n", value)
}

func dispatchContentDirectoryShareIndexInProgress(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryShareIndexInProgress %v\n", value)
}

func dispatchContentDirectoryShareIndexLastError(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryShareIndexLastError %v\n", value)
}

func dispatchContentDirectoryUserRadioUpdateID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryUserRadioUpdateID %v\n", value)
}

func dispatchContentDirectorySavedQueuesUpdateID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectorySavedQueuesUpdateID %v\n", value)
}

func dispatchContentDirectoryShareListUpdateID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryShareListUpdateID %v\n", value)
}

func dispatchContentDirectoryRecentlyPlayedUpdateID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryRecentlyPlayedUpdateID %v\n", value)
}

func dispatchContentDirectoryBrowseable(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryBrowseable %v\n", value)
}

func dispatchContentDirectoryRadioFavoritesUpdateID(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryRadioFavoritesUpdateID %v\n", value)
}

func dispatchContentDirectoryRadioLocationUpdateID(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryRadioLocationUpdateID %v\n", value)
}

func dispatchContentDirectoryFavoritesUpdateID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryFavoritesUpdateID %v\n", value)
}

func dispatchContentDirectoryFavoritePresetsUpdateID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchContentDirectoryFavoritePresetsUpdateID %v\n", value)
}

func dispatchDevicePropertiesSettingsReplicationState(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesSettingsReplicationState %v\n", value)
}

func dispatchDevicePropertiesZoneName(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesZoneName %v\n", value)
}

func dispatchDevicePropertiesIcon(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesIcon %v\n", value)
}

func dispatchDevicePropertiesConfiguration(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesConfiguration %v\n", value)
}

func dispatchDevicePropertiesInvisible(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesInvisible %v\n", value)
}

func dispatchDevicePropertiesIsZoneBridge(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesIsZoneBridge %v\n", value)
}

func dispatchDevicePropertiesAirPlayEnabled(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesAirPlayEnabled %v\n", value)
}

func dispatchDevicePropertiesSupportsAudioIn(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesSupportsAudioIn %v\n", value)
}

func dispatchDevicePropertiesSupportsAudioClip(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesSupportsAudioClip %v\n", value)
}

func dispatchDevicePropertiesIsIdle(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesIsIdle %v\n", value)
}

func dispatchDevicePropertiesMoreInfo(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesMoreInfo %v\n", value)
}

func dispatchDevicePropertiesChannelMapSet(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesChannelMapSet %v\n", value)
}

func dispatchDevicePropertiesHTSatChanMapSet(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesHTSatChanMapSet %v\n", value)
}

func dispatchDevicePropertiesHTFreq(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesHTFreq %v\n", value)
}

func dispatchDevicePropertiesHTBondedZoneCommitState(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesHTBondedZoneCommitState %v\n", value)
}

func dispatchDevicePropertiesOrientation(zp *ZonePlayer, value int32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesOrientation %v\n", value)
}

func dispatchDevicePropertiesLastChangedPlayState(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesLastChangedPlayState %v\n", value)
}

func dispatchDevicePropertiesRoomCalibrationState(zp *ZonePlayer, value int32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesRoomCalibrationState %v\n", value)
}

func dispatchDevicePropertiesAvailableRoomCalibration(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesAvailableRoomCalibration %v\n", value)
}

func dispatchDevicePropertiesTVConfigurationError(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesTVConfigurationError %v\n", value)
}

func dispatchDevicePropertiesHdmiCecAvailable(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesHdmiCecAvailable %v\n", value)
}

func dispatchDevicePropertiesWirelessMode(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesWirelessMode %v\n", value)
}

func dispatchDevicePropertiesWirelessLeafOnly(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesWirelessLeafOnly %v\n", value)
}

func dispatchDevicePropertiesHasConfiguredSSID(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesHasConfiguredSSID %v\n", value)
}

func dispatchDevicePropertiesChannelFreq(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesChannelFreq %v\n", value)
}

func dispatchDevicePropertiesBehindWifiExtender(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesBehindWifiExtender %v\n", value)
}

func dispatchDevicePropertiesWifiEnabled(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesWifiEnabled %v\n", value)
}

func dispatchDevicePropertiesConfigMode(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesConfigMode %v\n", value)
}

func dispatchDevicePropertiesSecureRegState(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesSecureRegState %v\n", value)
}

func dispatchDevicePropertiesVoiceConfigState(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesVoiceConfigState %v\n", value)
}

func dispatchDevicePropertiesMicEnabled(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchDevicePropertiesMicEnabled %v\n", value)
}

func dispatchGroupManagementGroupCoordinatorIsLocal(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchGroupManagementGroupCoordinatorIsLocal %v\n", value)
}

func dispatchGroupManagementLocalGroupUUID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchGroupManagementLocalGroupUUID %v\n", value)
}

func dispatchGroupManagementVirtualLineInGroupID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchGroupManagementVirtualLineInGroupID %v\n", value)
}

func dispatchGroupManagementResetVolumeAfter(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchGroupManagementResetVolumeAfter %v\n", value)
}

func dispatchGroupManagementVolumeAVTransportURI(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchGroupManagementVolumeAVTransportURI %v\n", value)
}

func dispatchGroupRenderingControlGroupMute(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchGroupRenderingControlGroupMute %v\n", value)
}

func dispatchGroupRenderingControlGroupVolume(zp *ZonePlayer, value uint16) {
	fmt.Fprintf(os.Stderr, "dispatchGroupRenderingControlGroupVolume %v\n", value)
}

func dispatchGroupRenderingControlGroupVolumeChangeable(zp *ZonePlayer, value bool) {
	fmt.Fprintf(os.Stderr, "dispatchGroupRenderingControlGroupVolumeChangeable %v\n", value)
}

func dispatchMusicServicesServiceListVersion(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchMusicServicesServiceListVersion %v\n", value)
}

func dispatchQueueLastChange(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchQueueLastChange %v\n", value)
}

func dispatchRenderingControlLastChange(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchRenderingControlLastChange %v\n", value)
}

func dispatchSystemPropertiesCustomerID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchSystemPropertiesCustomerID %v\n", value)
}

func dispatchSystemPropertiesUpdateID(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchSystemPropertiesUpdateID %v\n", value)
}

func dispatchSystemPropertiesUpdateIDX(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchSystemPropertiesUpdateIDX %v\n", value)
}

func dispatchSystemPropertiesVoiceUpdateID(zp *ZonePlayer, value uint32) {
	fmt.Fprintf(os.Stderr, "dispatchSystemPropertiesVoiceUpdateID %v\n", value)
}

func dispatchSystemPropertiesThirdPartyHash(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchSystemPropertiesThirdPartyHash %v\n", value)
}

func dispatchVirtualLineInLastChange(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchVirtualLineInLastChange %v\n", value)
}

func dispatchZoneGroupTopologyAvailableSoftwareUpdate(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchZoneGroupTopologyAvailableSoftwareUpdate %v\n", value)
}

func dispatchZoneGroupTopologyThirdPartyMediaServersX(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchZoneGroupTopologyThirdPartyMediaServersX %v\n", value)
}

func dispatchZoneGroupTopologyAlarmRunSequence(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchZoneGroupTopologyAlarmRunSequence %v\n", value)
}

func dispatchZoneGroupTopologyMuseHouseholdId(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchZoneGroupTopologyMuseHouseholdId %v\n", value)
}

func dispatchZoneGroupTopologyZoneGroupName(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchZoneGroupTopologyZoneGroupName %v\n", value)
}

func dispatchZoneGroupTopologyZoneGroupID(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchZoneGroupTopologyZoneGroupID %v\n", value)
}

func dispatchZoneGroupTopologyZonePlayerUUIDsInGroup(zp *ZonePlayer, value string) {
	fmt.Fprintf(os.Stderr, "dispatchZoneGroupTopologyZonePlayerUUIDsInGroup %v\n", value)
}

func dispatchZoneGroupTopologyAreasUpdateID(zp *ZonePlayer, value string) {
	if zp.AreasUpdateID == nil {
		return
	}
	zp.AreasUpdateID(value)
}

func dispatchZoneGroupTopologySourceAreasUpdateID(zp *ZonePlayer, value string) {
	if zp.SourceAreasUpdateID == nil {
		return
	}
	zp.SourceAreasUpdateID(value)
}

func dispatchZoneGroupTopologyNetsettingsUpdateID(zp *ZonePlayer, value string) {
	if zp.NetsettingsUpdateID == nil {
		return
	}
	zp.NetsettingsUpdateID(value)
}
