package dnanosense

func UpdateConfig(new NanosenseConf, old NanosenseSetup) NanosenseSetup {
	setup := old
	if new.FriendlyName != nil {
		setup.FriendlyName = new.FriendlyName
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
	return setup
}
