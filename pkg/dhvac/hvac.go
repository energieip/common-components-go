package dhvac

import (
	"encoding/json"
)

//Hvac hvac driver representation
type Hvac struct {
	Mac                         string  `json:"mac"`
	FullMac                     *string `json:"fullMac,omitempty"`
	IP                          string  `json:"ip"`
	Group                       int     `json:"group"`
	Protocol                    string  `json:"protocol"`
	SwitchMac                   string  `json:"switchMac"`
	IsConfigured                bool    `json:"isConfigured"`
	SoftwareVersion             float32 `json:"softwareVersion"`
	HardwareVersion             float32 `json:"hardwareVersion"`
	Error                       int     `json:"error"`
	HoldOff1                    int     `json:"holdOff1"`   //window open in the group
	SpaceTemp1                  int     `json:"spaceTemp1"` //nanosense temperature
	OccManCmd1                  int     `json:"occManCmd1"` //occupied/inoccupied from sensor
	Shift                       int     `json:"shift"`      //temperature shift
	HeatCool1                   int     `json:"heatCool1"`  //pcVue control for heat/cool mode
	DewSensor1                  int     `json:"dewSensor1"` //humidity contact
	SpaceCO2                    int     `json:"spaceCO2"`
	TemperatureSelect           int     `json:"temperatureSelect"`           //target temperature
	SetpointOccupiedCool1       int     `json:"setpointOccupiedCool1"`       // 1/10°C
	SetpointOccupiedHeat1       int     `json:"setpointOccupiedHeat1"`       // 1/10°C
	SetpointInoccupiedCool1     int     `json:"setpointInoccupiedCool1"`     // 1/10°C
	SetpointHeatInoccupiedHeat1 int     `json:"setpointHeatInoccupiedHeat1"` // 1/10°C
	SetpointStandbyCool1        int     `json:"setpointStandbyCool1"`        // 1/10°C
	SetpointStandbyHeat1        int     `json:"setpointStandbyHeat1"`        // 1/10°C
	EffectSetPoint1             int     `json:"effectSetPoint1"`             //opening valve percentage
	HeatOutput1                 int     `json:"heatOutput1"`                 //?
	CoolOutput1                 int     `json:"coolOutput1"`                 //?
	OADamper                    int     `json:"oaDamper"`                    //opening damper percentage
	DumpFrequency               int     `json:"dumpFrequency"`
	SpaceCOV                    int     `json:"spaceCOV"`
	FriendlyName                string  `json:"friendlyName"`
	Label                       *string `json:"label,omitempty"`
}

//HvacSetup initial setup send by the server when the driver is authorized
type HvacSetup struct {
	Mac                    string  `json:"mac"`
	SwitchMac              string  `json:"switchMac"`
	FullMac                *string `json:"fullMac,omitempty"`
	Group                  *int    `json:"group,omitempty"`
	DumpFrequency          int     `json:"dumpFrequency"`
	SetpointCoolOccupied   *int    `json:"setpointCoolOccupied,omitempty"`   // 1/10°C
	SetpointHeatOccupied   *int    `json:"setpointHeatOccupied,omitempty"`   // 1/10°C
	SetpointCoolInoccupied *int    `json:"setpointCoolInoccupied,omitempty"` // 1/10°C
	SetpointHeatInoccupied *int    `json:"setpointHeatInoccupied,omitempty"` // 1/10°C
	SetpointCoolStandby    *int    `json:"setpointCoolStandby,omitempty"`    // 1/10°C
	SetpointHeatStandby    *int    `json:"setpointHeatStandby,omitempty"`    // 1/10°C
	FriendlyName           *string `json:"friendlyName"`
	IsConfigured           *bool   `json:"isConfigured,omitempty"`
	Label                  *string `json:"label,omitempty"`
}

//HvacConf customizable configuration by the server
type HvacConf struct {
	Mac                    string  `json:"mac"`
	FullMac                *string `json:"fullMac,omitempty"`
	Group                  *int    `json:"group,omitempty"`
	Shift                  int     `json:"shift"`                            //temperature shift (+6/-6)
	WindowStatus           *bool   `json:"windowStatus,omitempty"`           //corresponds to holdOff1
	Temperature            *int    `json:"temperature,omitempty"`            // 1/10°C from nanosense device
	Presence               *bool   `json:"presence,omitempty"`               //corresponds to oCCManCmd1
	Mode                   *string `json:"mode,omitempty"`                   //heat/cold corresponds to HeatCool1
	CO2                    *int    `json:"co2,omitempty"`                    //CO2 from nanosense
	COV                    *int    `json:"cov,omitempty"`                    //COV from nanosense
	SetpointCoolOccupied   *int    `json:"setpointCoolOccupied,omitempty"`   // 1/10°C
	SetpointHeatOccupied   *int    `json:"setpointHeatOccupied,omitempty"`   // 1/10°C
	SetpointCoolInoccupied *int    `json:"setpointCoolInoccupied,omitempty"` // 1/10°C
	SetpointHeatInoccupied *int    `json:"setpointHeatInoccupied,omitempty"` // 1/10°C
	SetpointCoolStandby    *int    `json:"setpointCoolStandby,omitempty"`    // 1/10°C
	SetpointHeatStandby    *int    `json:"setpointHeatStandby,omitempty"`    // 1/10°C
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
