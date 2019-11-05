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

	return setup
}
