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
	return setup
}
