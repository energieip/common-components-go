package dhvac

import (
	"encoding/json"
)

const (
	//Regulation Mode (heatCool)
	HVAC_MODE_AUTO           = 0
	HVAC_MODE_HEAT           = 1
	HVAC_MODE_COOL           = 3
	HVAC_MODE_OFF            = 6
	HVAC_MODE_TEST           = 7
	HVAC_MODE_EMERGENCY_HEAT = 8

	//CVC Occupation mode
	OCCUPANCY_AUTO                = 0
	OCCUPANCY_COMFORT             = 1
	OCCUPANCY_STANDBY             = 2
	OCCUPANCY_ECONOMY             = 3
	OCCUPANCY_BUILDING_PROTECTION = 4
)

//Hvac hvac driver representation
type Hvac struct {
	Mac                     string  `json:"mac"`
	IP                      string  `json:"ip"`
	Group                   int     `json:"group"`
	Protocol                string  `json:"protocol"`
	SwitchMac               string  `json:"switchMac"`
	IsConfigured            bool    `json:"isConfigured"`
	SoftwareVersion         string  `json:"softwareVersion"`
	Error                   int     `json:"error"`
	LinePower               int     `json:"linePower"`
	HoldOff1                int     `json:"holdOff1"`   //window open in the group
	SpaceTemp1              int     `json:"spaceTemp1"` //nanosense temperature
	OccManCmd1              int     `json:"occManCmd1"` //occupied/inoccupied from sensor
	Shift                   int     `json:"shift"`      //temperature shift (in 1/10°C)
	HeatCool1               int     `json:"heatCool1"`  //pcVue control for heat/cool mode
	TargetMode              int     `json:"targetMode"` //TargetMode
	DewSensor1              int     `json:"dewSensor1"` //humidity contact
	SpaceCO2                int     `json:"spaceCO2"`
	TemperatureSelect       int     `json:"temperatureSelect"`       //target temperature
	SetpointOccupiedCool1   int     `json:"setpointOccupiedCool1"`   // 1/10°C
	SetpointOccupiedHeat1   int     `json:"setpointOccupiedHeat1"`   // 1/10°C
	SetpointUnoccupiedCool1 int     `json:"setpointUnoccupiedCool1"` // 1/10°C
	SetpointUnoccupiedHeat1 int     `json:"setpointUnoccupiedHeat1"` // 1/10°C
	SetpointStandbyCool1    int     `json:"setpointStandbyCool1"`    // 1/10°C
	SetpointStandbyHeat1    int     `json:"setpointStandbyHeat1"`    // 1/10°C
	EffectSetPoint1         int     `json:"effectSetPoint1"`         //opening valve percentage
	HeatOutput1             int     `json:"heatOutput1"`             //?
	CoolOutput1             int     `json:"coolOutput1"`             //?
	OADamper                int     `json:"oaDamper"`                //opening damper percentage
	InputE1                 int     `json:"inputE1"`
	InputE2                 int     `json:"inputE2"`
	InputE3                 int     `json:"inputE3"`
	InputE4                 int     `json:"inputE4"`
	InputE5                 int     `json:"inputE5"`
	InputE6                 int     `json:"inputE6"`
	InputC1                 int     `json:"inputC1"`
	InputC2                 int     `json:"inputC2"`
	OutputY5                int     `json:"outputY5"`
	OutputY6                int     `json:"outputY6"`
	OutputY7                int     `json:"outputY7"`
	OutputY8                int     `json:"outputY8"`
	OutputYa                int     `json:"outputYa"`
	OutputYb                int     `json:"outputYb"`
	TemperatureOffsetStep   int     `json:"temperatureOffsetStep"`
	DumpFrequency           int     `json:"dumpFrequency"`
	SpaceCOV                int     `json:"spaceCOV"`
	SpaceHygro              int     `json:"spaceHygro"`
	FriendlyName            string  `json:"friendlyName"`
	Label                   *string `json:"label,omitempty"`
}

//HvacSetup initial setup send by the server when the driver is authorized
type HvacSetup struct {
	Mac                    string  `json:"mac"`
	SwitchMac              string  `json:"switchMac"`
	Group                  *int    `json:"group,omitempty"`
	DumpFrequency          int     `json:"dumpFrequency"`
	SetpointCoolOccupied   *int    `json:"setpointCoolOccupied,omitempty"`   // 1/10°C
	SetpointHeatOccupied   *int    `json:"setpointHeatOccupied,omitempty"`   // 1/10°C
	SetpointCoolInoccupied *int    `json:"setpointCoolInoccupied,omitempty"` // 1/10°C
	SetpointHeatInoccupied *int    `json:"setpointHeatInoccupied,omitempty"` // 1/10°C
	SetpointCoolStandby    *int    `json:"setpointCoolStandby,omitempty"`    // 1/10°C
	SetpointHeatStandby    *int    `json:"setpointHeatStandby,omitempty"`    // 1/10°C
	InputE1                *int    `json:"inputE1,omitempty"`
	InputE2                *int    `json:"inputE2,omitempty"`
	InputE3                *int    `json:"inputE3,omitempty"`
	InputE4                *int    `json:"inputE4,omitempty"`
	InputE5                *int    `json:"inputE5,omitempty"`
	InputE6                *int    `json:"inputE6,omitempty"`
	InputC1                *int    `json:"inputC1,omitempty"`
	InputC2                *int    `json:"inputC2,omitempty"`
	OutputY5               *int    `json:"outputY5,omitempty"`
	OutputY6               *int    `json:"outputY6,omitempty"`
	OutputY7               *int    `json:"outputY7,omitempty"`
	OutputY8               *int    `json:"outputY8,omitempty"`
	OutputYa               *int    `json:"outputYa,omitempty"`
	OutputYb               *int    `json:"outputYb,omitempty"`
	TemperatureOffsetStep  *int    `json:"temperatureOffsetStep,omitempty"`
	TemperatureSelection   *int    `json:"temperatureSelection,omitempty"` //reception mode of the temperature: sensor/network
	RegulationType         *int    `json:"regulationType,omitempty"`       //regulation type
	LoopUsed               *int    `json:"loopUsed,omitempty"`             //loop used
	FanOffDelay            *int    `json:"fanOffDelay,omitempty"`          //ventilation delay on battery after stop
	FanConfig              *int    `json:"fanConfig,omitempty"`            // Ventilation mode
	FanMode                *int    `json:"fanMode,omitempty"`              //ventilation speed
	FanOverride            *int    `json:"fanOverride,omitempty"`          //blowing when regulation is 0
	OaDamperMode           *int    `json:"oaDamperMode,omitempty"`         // 0-10V register detection
	CO2Mode                *int    `json:"co2Mode,omitempty"`              //reception mode of the co2: sensor/network
	CO2Max                 *int    `json:"co2Max,omitempty"`               //max limit in ppm
	TargetMode             *int    `json:"targetMode,omitempty"`           //TargetMode
	FriendlyName           *string `json:"friendlyName,omitempty"`
	IsConfigured           *bool   `json:"isConfigured,omitempty"`
	Label                  *string `json:"label,omitempty"`
}

//HvacConf customizable configuration by the server
type HvacConf struct {
	Mac                    string  `json:"mac"`
	Group                  *int    `json:"group,omitempty"`
	Shift                  *int    `json:"shift,omitempty"`                  //temperature shift (+6/-6) in (1/10°C)
	WindowStatus           *bool   `json:"windowStatus,omitempty"`           //corresponds to holdOff1
	Temperature            *int    `json:"temperature,omitempty"`            // 1/10°C from nanosense device
	Hygrometry             *int    `json:"hygrometry,omitempty"`             // % from nanosense device
	Presence               *bool   `json:"presence,omitempty"`               //corresponds to oCCManCmd1
	HeatCool               *int    `json:"heatCool,omitempty"`               //Regulation Mode (heatCool)
	TargetMode             *int    `json:"targetMode,omitempty"`             //TargetMode
	CO2                    *int    `json:"co2,omitempty"`                    //CO2 from nanosense
	COV                    *int    `json:"cov,omitempty"`                    //COV from nanosense
	SetpointCoolOccupied   *int    `json:"setpointCoolOccupied,omitempty"`   // 1/10°C
	SetpointHeatOccupied   *int    `json:"setpointHeatOccupied,omitempty"`   // 1/10°C
	SetpointCoolInoccupied *int    `json:"setpointCoolInoccupied,omitempty"` // 1/10°C
	SetpointHeatInoccupied *int    `json:"setpointHeatInoccupied,omitempty"` // 1/10°C
	SetpointCoolStandby    *int    `json:"setpointCoolStandby,omitempty"`    // 1/10°C
	SetpointHeatStandby    *int    `json:"setpointHeatStandby,omitempty"`    // 1/10°C
	InputE1                *int    `json:"inputE1,omitempty"`
	InputE2                *int    `json:"inputE2,omitempty"`
	InputE3                *int    `json:"inputE3,omitempty"`
	InputE4                *int    `json:"inputE4,omitempty"`
	InputE5                *int    `json:"inputE5,omitempty"`
	InputE6                *int    `json:"inputE6,omitempty"`
	InputC1                *int    `json:"inputC1,omitempty"`
	InputC2                *int    `json:"inputC2,omitempty"`
	OutputY5               *int    `json:"outputY5,omitempty"`
	OutputY6               *int    `json:"outputY6,omitempty"`
	OutputY7               *int    `json:"outputY7,omitempty"`
	OutputY8               *int    `json:"outputY8,omitempty"`
	OutputYa               *int    `json:"outputYa,omitempty"`
	OutputYb               *int    `json:"outputYb,omitempty"`
	TemperatureOffsetStep  *int    `json:"temperatureOffsetStep,omitempty"`
	TemperatureSelection   *int    `json:"temperatureSelection,omitempty"` //reception mode of the temperature: sensor/network
	RegulationType         *int    `json:"regulationType,omitempty"`       //regulation type
	LoopUsed               *int    `json:"loopUsed,omitempty"`             //loop used
	FanOffDelay            *int    `json:"fanOffDelay,omitempty"`          //ventilation delay on battery after stop
	FanConfig              *int    `json:"fanConfig,omitempty"`            // Ventilation mode
	FanMode                *int    `json:"fanMode,omitempty"`              //ventilation speed
	FanOverride            *int    `json:"fanOverride,omitempty"`          //blowing when regulation is 0
	OaDamperMode           *int    `json:"oaDamperMode,omitempty"`         // 0-10V register detection
	CO2Mode                *int    `json:"co2Mode,omitempty"`              //reception mode of the co2: sensor/network
	CO2Max                 *int    `json:"co2Max,omitempty"`               //max limit in ppm
	FriendlyName           *string `json:"friendlyName"`
	IsConfigured           *bool   `json:"isConfigured,omitempty"`
	DumpFrequency          *int    `json:"dumpFrequency"`
	Label                  *string `json:"label,omitempty"`
}

//ToHvac convert map interface to Hvac object
func ToHvac(val interface{}) (*Hvac, error) {
	var driver Hvac
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToHvacConf convert map interface to Hvac object
func ToHvacConf(val interface{}) (*HvacConf, error) {
	var driver HvacConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToHvacSetup convert map interface to Hvac object
func ToHvacSetup(val interface{}) (*HvacSetup, error) {
	var driver HvacSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

// ToJSON dump hvac struct
func (driver Hvac) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

// ToJSON dump hvac setup struct
func (driver HvacSetup) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToJSON dump struct in json
func (driver HvacConf) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}
