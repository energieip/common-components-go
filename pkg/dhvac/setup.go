package dhvac

//FillDefaultValue fill default parameter for setup
func FillDefaultValue(cfg HvacSetup) HvacSetup {
	if cfg.FriendlyName == nil && cfg.Label != nil {
		name := *cfg.Label
		cfg.FriendlyName = &name
	}
	if cfg.Group == nil {
		group := 0
		cfg.Group = &group
	}
	if cfg.DumpFrequency == 0 {
		cfg.DumpFrequency = 1000
	}
	if cfg.TemperatureOffsetStep == nil {
		defaultStep := 5 //It corresponds to 0.5°C
		cfg.TemperatureOffsetStep = &defaultStep
	}
	return cfg
}

//UpdateSetup update blind struct
func UpdateSetup(new HvacSetup, old HvacSetup) HvacSetup {
	setup := old
	if new.FriendlyName != nil {
		setup.FriendlyName = new.FriendlyName
	}
	if new.Group != nil {
		setup.Group = new.Group
	}

	if new.DumpFrequency != 0 {
		setup.DumpFrequency = new.DumpFrequency
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
