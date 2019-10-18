package dnanosense

//FillDefaultValue fill default parameter for setup
func FillDefaultValue(cfg NanosenseSetup) NanosenseSetup {
	if cfg.FriendlyName == nil && cfg.Label != "" {
		name := cfg.Label
		cfg.FriendlyName = &name
	}
	return cfg
}

//UpdateSetup update blind struct
func UpdateSetup(new NanosenseSetup, old NanosenseSetup) NanosenseSetup {
	setup := old
	if new.FriendlyName != nil {
		setup.FriendlyName = new.FriendlyName
	}

	if len(new.API) != 0 {
		setup.API = new.API
	}
	setup.Group = new.Group
	setup.Cluster = new.Cluster
	setup.ModbusID = new.ModbusID
	if new.IP != nil {
		setup.IP = new.IP
	}
	return setup
}
