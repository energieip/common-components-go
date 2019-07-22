package dnanosense

func UpdateConfig(new NanosenseConf, old NanosenseSetup) NanosenseSetup {
	setup := old
	if new.FriendlyName != nil {
		setup.FriendlyName = new.FriendlyName
	}

	if new.Label != nil {
		setup.Label = new.Label
	}

	if new.Mac != "" {
		setup.Mac = new.Mac
	}

	if new.Cluster != nil {
		setup.Cluster = *new.Cluster
	}

	if new.Group != nil {
		setup.Group = *new.Group
	}

	if new.ModbusOffset != nil {
		setup.ModbusOffset = *new.ModbusOffset
	}
	if new.IsConfigured != nil {
		setup.IsConfigured = new.IsConfigured
	}

	return setup
}
