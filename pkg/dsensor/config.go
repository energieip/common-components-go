package dsensor

func UpdateConfig(new SensorConf, old SensorSetup) SensorSetup {
	setup := old
	if new.BrightnessCorrectionFactor != nil {
		setup.BrightnessCorrectionFactor = new.BrightnessCorrectionFactor
	}

	if new.FriendlyName != nil {
		setup.FriendlyName = new.FriendlyName
	}

	if new.Group != nil {
		setup.Group = new.Group
	}

	if new.IsBleEnabled != nil {
		setup.IsBleEnabled = new.IsBleEnabled
	}

	if new.TemperatureOffset != nil {
		setup.TemperatureOffset = new.TemperatureOffset
	}

	if new.ThresholdPresence != nil {
		setup.ThresholdPresence = new.ThresholdPresence
	}

	if new.DumpFrequency != nil {
		setup.DumpFrequency = *new.DumpFrequency
	}

	if new.IBeaconMajor != nil {
		setup.IBeaconMajor = new.IBeaconMajor
	}

	if new.IBeaconMinor != nil {
		setup.IBeaconMinor = new.IBeaconMinor
	}

	if new.IBeaconTxPower != nil {
		setup.IBeaconTxPower = new.IBeaconTxPower
	}

	if new.IBeaconUUID != nil {
		setup.IBeaconUUID = new.IBeaconUUID
	}

	if new.BleMode != nil {
		setup.BleMode = new.BleMode
	}
	return setup
}
