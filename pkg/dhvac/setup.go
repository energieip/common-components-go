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
	if cfg.SetpointCoolInoccupied == nil {
		defaultCI := 300
		cfg.SetpointCoolInoccupied = &defaultCI
	}

	if cfg.SetpointCoolOccupied != nil {
		defaultCO := 190
		cfg.SetpointCoolOccupied = &defaultCO
	}

	if cfg.SetpointCoolStandby == nil {
		defaultCS := 280
		cfg.SetpointCoolStandby = &defaultCS
	}

	if cfg.SetpointHeatStandby == nil {
		defaultHS := 170
		cfg.SetpointHeatStandby = &defaultHS
	}

	if cfg.SetpointHeatInoccupied == nil {
		defaultHI := 150
		cfg.SetpointHeatInoccupied = &defaultHI
	}

	if cfg.SetpointHeatOccupied == nil {
		defaultHO := 260
		cfg.SetpointHeatOccupied = &defaultHO
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

	if new.SetpointCoolInoccupied != nil {
		setup.SetpointCoolInoccupied = new.SetpointCoolInoccupied
	}

	if new.SetpointCoolOccupied != nil {
		setup.SetpointCoolOccupied = new.SetpointCoolOccupied
	}

	if new.SetpointCoolStandby != nil {
		setup.SetpointCoolStandby = new.SetpointCoolStandby
	}

	if new.SetpointHeatStandby != nil {
		setup.SetpointHeatStandby = new.SetpointHeatStandby
	}

	if new.SetpointHeatInoccupied != nil {
		setup.SetpointHeatInoccupied = new.SetpointHeatInoccupied
	}

	if new.SetpointHeatOccupied != nil {
		setup.SetpointHeatOccupied = new.SetpointHeatOccupied
	}

	if new.TemperatureOffsetStep != nil {
		setup.TemperatureOffsetStep = new.TemperatureOffsetStep
	}
	return setup
}
