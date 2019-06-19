package dblind

func UpdateConfig(new BlindConf, old BlindSetup) BlindSetup {
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

	if new.DumpFrequency != nil {
		setup.DumpFrequency = *new.DumpFrequency
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

	if new.FullMac != nil {
		setup.FullMac = new.FullMac
	}
	return setup
}
