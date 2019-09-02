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
	return setup
}
