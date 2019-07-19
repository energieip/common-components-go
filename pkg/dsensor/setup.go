package dsensor

//FillDefaultValue fill default parameter for setup
func FillDefaultValue(cfg SensorSetup) SensorSetup {
	if cfg.BleMode == nil {
		ble := "remote"
		cfg.BleMode = &ble
	}
	if cfg.IsBleEnabled == nil {
		bleEnable := false
		cfg.IsBleEnabled = &bleEnable
	}
	if cfg.Group == nil {
		group := 0
		cfg.Group = &group
	}
	if cfg.FriendlyName == nil && cfg.Label != nil {
		name := *cfg.Label
		cfg.FriendlyName = &name
	}

	if cfg.BrightnessCorrectionFactor == nil {
		defaultValue := 1
		cfg.BrightnessCorrectionFactor = &defaultValue
	}
	if cfg.DumpFrequency == 0 {
		cfg.DumpFrequency = 1000
	}
	if cfg.TemperatureOffset == nil {
		defaultValue := 0
		cfg.TemperatureOffset = &defaultValue
	}
	if cfg.ThresholdPresence == nil {
		presence := 10
		cfg.ThresholdPresence = &presence
	}
	return cfg
}

//UpdateSetup update setup struct
func UpdateSetup(new SensorSetup, old SensorSetup) SensorSetup {
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

	if new.DumpFrequency != 0 {
		setup.DumpFrequency = new.DumpFrequency
	}

	if new.Label != nil {
		setup.Label = new.Label
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

	if new.Mac != "" {
		setup.Mac = new.Mac
	}

	return setup
}
