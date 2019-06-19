package dled

func UpdateConfig(new LedConf, old LedSetup) LedSetup {
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

	if new.DumpFrequency != nil {
		setup.DumpFrequency = *new.DumpFrequency
	}

	if new.Watchdog != nil {
		setup.Watchdog = new.Watchdog
	}

	if new.Label != nil {
		setup.Label = new.Label
	}
	return setup
}
