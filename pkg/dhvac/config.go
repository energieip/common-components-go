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

	if new.Mac != "" {
		setup.Mac = new.Mac
	}

	if new.TemperatureOffsetStep != nil {
		setup.TemperatureOffsetStep = new.TemperatureOffsetStep
	}
	if new.SetpointHeatInoccupied != nil {
		setup.SetpointHeatInoccupied = new.SetpointHeatInoccupied
	}
	if new.SetpointHeatOccupied != nil {
		setup.SetpointHeatOccupied = new.SetpointHeatOccupied
	}
	if new.SetpointCoolOccupied != nil {
		setup.SetpointCoolOccupied = new.SetpointCoolOccupied
	}
	if new.SetpointCoolInoccupied != nil {
		setup.SetpointCoolInoccupied = new.SetpointCoolInoccupied
	}
	if new.SetpointCoolStandby != nil {
		setup.SetpointCoolStandby = new.SetpointCoolStandby
	}
	if new.SetpointHeatStandby != nil {
		setup.SetpointHeatStandby = new.SetpointHeatStandby
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

	return setup
}
