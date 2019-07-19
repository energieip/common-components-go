package dblind

//FillDefaultValue fill default parameter for setup
func FillDefaultValue(cfg BlindSetup) BlindSetup {
	if cfg.FriendlyName != nil && cfg.Label != nil {
		name := *cfg.Label
		cfg.FriendlyName = &name
	}
	if cfg.DumpFrequency == 0 {
		cfg.DumpFrequency = 1000
	}
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
	return cfg
}

//UpdateSetup update blind struct
func UpdateSetup(new BlindSetup, old BlindSetup) BlindSetup {
	setup := old
	if new.FriendlyName != nil {
		setup.FriendlyName = new.FriendlyName
	}

	if new.Group != nil {
		setup.Group = new.Group
	}

	if new.IsBleEnabled != nil {
		setup.IsBleEnabled = new.IsBleEnabled
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
