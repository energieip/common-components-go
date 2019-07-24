package dwago

func UpdateConfig(new WagoConf, old WagoSetup) WagoSetup {
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

	if new.IsConfigured != nil {
		setup.IsConfigured = new.IsConfigured
	}

	if new.IP != nil {
		setup.IP = new.IP
	}
	return setup
}
