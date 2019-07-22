package dled

//FillDefaultValue fill default parameter for setup
func FillDefaultValue(config LedSetup) LedSetup {
	if config.IsBleEnabled == nil {
		enabled := false
		config.IsBleEnabled = &enabled
	}
	if config.DumpFrequency == 0 {
		config.DumpFrequency = 1000
	}
	if config.Auto == nil {
		auto := true
		config.Auto = &auto
	}
	if config.ThresholdHigh == nil {
		high := 100
		config.ThresholdHigh = &high
	}
	if config.ThresholdLow == nil {
		low := 0
		config.ThresholdLow = &low
	}
	if config.Group == nil {
		group := 0
		config.Group = &group
	}
	if config.Watchdog == nil {
		watchdog := 600
		config.Watchdog = &watchdog
	}
	slope := 10000
	if config.SlopeStartManual == nil {
		config.SlopeStartManual = &slope
	}
	if config.SlopeStopManual == nil {
		config.SlopeStopManual = &slope
	}
	if config.DefaultSetpoint == nil {
		defaultValue := 0
		config.DefaultSetpoint = &defaultValue
	}
	if config.BleMode == nil {
		defaultMode := "remote"
		config.BleMode = &defaultMode
	}
	if config.PMax == 0 {
		config.PMax = 5
	}
	if config.FriendlyName == nil && config.Label != nil {
		name := *config.Label
		config.FriendlyName = &name
	}

	if config.FirstDay == nil {
		firstDay := false
		config.FirstDay = &firstDay
	}
	return config
}

//UpdateSetup update setup struct
func UpdateSetup(new LedSetup, old LedSetup) LedSetup {
	setup := old
	if new.ThresholdHigh != nil {
		setup.ThresholdHigh = new.ThresholdHigh
	}

	if new.ThresholdLow != nil {
		setup.ThresholdLow = new.ThresholdLow
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

	if new.SlopeStartAuto != nil {
		setup.SlopeStartAuto = new.SlopeStartAuto
	}

	if new.SlopeStartManual != nil {
		setup.SlopeStartManual = new.SlopeStartManual
	}

	if new.SlopeStopAuto != nil {
		setup.SlopeStopAuto = new.SlopeStopAuto
	}

	if new.SlopeStopManual != nil {
		setup.SlopeStopManual = new.SlopeStopManual
	}

	if new.BleMode != nil {
		setup.BleMode = new.BleMode
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

	if new.DumpFrequency != 0 {
		setup.DumpFrequency = new.DumpFrequency
	}

	if new.Watchdog != nil {
		setup.Watchdog = new.Watchdog
	}

	if new.PMax != 0 {
		setup.PMax = new.PMax
	}

	if new.Label != nil {
		setup.Label = new.Label
	}

	if new.Mac != "" {
		setup.Mac = new.Mac
	}

	if new.FirstDay != nil {
		setup.FirstDay = new.FirstDay
	}
	return setup
}
