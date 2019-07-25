package dwago

//FillDefaultValue fill default parameter for setup
func FillDefaultValue(cfg WagoSetup) WagoSetup {
	if cfg.FriendlyName == nil && cfg.Label != nil {
		name := *cfg.Label
		cfg.FriendlyName = &name
	}
	if cfg.DumpFrequency == nil {
		freq := 1000 //1s
		cfg.DumpFrequency = &freq
	}
	if cfg.ModbusOffset == nil {
		offset := 0
		cfg.ModbusOffset = &offset
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

	if new.DumpFrequency != nil {
		setup.DumpFrequency = new.DumpFrequency
	}

	if new.IP != nil {
		setup.IP = new.IP
	}
	if new.ModbusOffset != nil {
		setup.ModbusOffset = new.ModbusOffset
	}
	setup.Cluster = new.Cluster
	return setup
}
