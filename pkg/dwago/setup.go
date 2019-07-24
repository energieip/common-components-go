package dwago

//FillDefaultValue fill default parameter for setup
func FillDefaultValue(cfg WagoSetup) WagoSetup {
	if cfg.FriendlyName == nil && cfg.Label != nil {
		name := *cfg.Label
		cfg.FriendlyName = &name
	}
	return cfg
}

//UpdateSetup update blind struct
func UpdateSetup(new WagoSetup, old WagoSetup) WagoSetup {
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

	if len(new.API) != 0 {
		setup.API = new.API
	}

	if new.IP != nil {
		setup.IP = new.IP
	}
	setup.Cluster = new.Cluster
	return setup
}
