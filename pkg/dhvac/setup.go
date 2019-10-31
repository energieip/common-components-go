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

	if cfg.SetpointCoolOccupied == nil {
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

	if cfg.TemperatureSelection == nil {
		defaultTS := 1
		cfg.TemperatureSelection = &defaultTS
	}
	if cfg.RegulationType == nil {
		defaultRT := 9
		cfg.RegulationType = &defaultRT
	}
	if cfg.LoopUsed == nil {
		loopUsed := 1
		cfg.LoopUsed = &loopUsed
	}
	if cfg.FanOffDelay == nil {
		fanDelay := 0
		cfg.FanOffDelay = &fanDelay
	}
	if cfg.FanConfig == nil {
		fanConfig := 0
		cfg.FanConfig = &fanConfig
	}
	if cfg.FanMode == nil {
		fanMode := 1
		cfg.FanMode = &fanMode
	}
	if cfg.FanOverride == nil {
		fanOverride := 0
		cfg.FanOverride = &fanOverride
	}

	if cfg.OaDamperMode == nil {
		oadamperMode := 5
		cfg.OaDamperMode = &oadamperMode
	}

	if cfg.CO2Mode == nil {
		co2Mode := 1
		cfg.CO2Mode = &co2Mode
	}

	if cfg.CO2Max == nil {
		co2Max := 5000
		cfg.CO2Max = &co2Max
	}
	defaultV := 0
	if cfg.Valve6WayCoolMin == nil {
		cfg.Valve6WayCoolMin = &defaultV
	}
	if cfg.Valve6WayCoolMax == nil {
		cfg.Valve6WayCoolMax = &defaultV
	}
	if cfg.Valve6WayHeatMin == nil {
		cfg.Valve6WayHeatMin = &defaultV
	}
	if cfg.Valve6WayHeatMax == nil {
		cfg.Valve6WayHeatMax = &defaultV
	}
	if cfg.Valve6WayRefPoint == nil {
		cfg.Valve6WayRefPoint = &defaultV
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
	if new.TemperatureSelection != nil {
		setup.TemperatureSelection = new.TemperatureSelection
	}
	if new.RegulationType != nil {
		setup.RegulationType = new.RegulationType
	}
	if new.LoopUsed != nil {
		setup.LoopUsed = new.LoopUsed
	}
	if new.FanOffDelay != nil {
		setup.FanOffDelay = new.FanOffDelay
	}
	if new.FanConfig != nil {
		setup.FanConfig = new.FanConfig
	}
	if new.FanMode != nil {
		setup.FanMode = new.FanMode
	}
	if new.FanOverride != nil {
		setup.FanOverride = new.FanOverride
	}

	if new.OaDamperMode != nil {
		setup.OaDamperMode = new.OaDamperMode
	}

	if new.CO2Mode != nil {
		setup.CO2Mode = new.CO2Mode
	}

	if new.CO2Max != nil {
		setup.CO2Max = new.CO2Max
	}
	if new.Valve6WayCoolMin != nil {
		setup.Valve6WayCoolMin = new.Valve6WayCoolMin
	}
	if new.Valve6WayCoolMax != nil {
		setup.Valve6WayCoolMax = new.Valve6WayCoolMax
	}
	if new.Valve6WayHeatMin != nil {
		setup.Valve6WayHeatMin = new.Valve6WayHeatMin
	}
	if new.Valve6WayHeatMax != nil {
		setup.Valve6WayHeatMax = new.Valve6WayHeatMax
	}
	if new.Valve6WayRefPoint != nil {
		setup.Valve6WayRefPoint = new.Valve6WayRefPoint
	}

	if new.InputE1 != nil {
		setup.InputE1 = new.InputE1
	}
	if new.InputE2 != nil {
		setup.InputE2 = new.InputE2
	}
	if new.InputE3 != nil {
		setup.InputE3 = new.InputE3
	}
	if new.InputE4 != nil {
		setup.InputE4 = new.InputE4
	}
	if new.InputE5 != nil {
		setup.InputE5 = new.InputE5
	}
	if new.InputE6 != nil {
		setup.InputE6 = new.InputE6
	}
	if new.InputC1 != nil {
		setup.InputC1 = new.InputC1
	}
	if new.InputC2 != nil {
		setup.InputC2 = new.InputC2
	}
	if new.OutputY5 != nil {
		setup.OutputY5 = new.OutputY5
	}
	if new.OutputY6 != nil {
		setup.OutputY6 = new.OutputY6
	}
	if new.OutputY7 != nil {
		setup.OutputY7 = new.OutputY7
	}
	if new.OutputY8 != nil {
		setup.OutputY8 = new.OutputY8
	}
	if new.OutputYa != nil {
		setup.OutputYa = new.OutputYa
	}
	if new.OutputYb != nil {
		setup.OutputYb = new.OutputYb
	}

	if new.SwitchMac != "" {
		setup.SwitchMac = new.SwitchMac
	}

	return setup
}
