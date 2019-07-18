package dhvac

func UpdateConfig(new HvacConf, old HvacSetup) HvacSetup {
	setup := old
	if new.FriendlyName != nil {
		setup.FriendlyName = new.FriendlyName
	}

	if new.Group != nil {
		setup.Group = new.Group
	}

	if new.DumpFrequency != nil {
		setup.DumpFrequency = *new.DumpFrequency
	}

	if new.Label != nil {
		setup.Label = new.Label
	}

	if new.Mac != "" {
		setup.Mac = new.Mac
	}

	if new.FullMac != nil {
		setup.FullMac = new.FullMac
	}
	if new.TemperatureOffsetStep != nil {
		setup.TemperatureOffsetStep = new.TemperatureOffsetStep
	}
	return setup
}
