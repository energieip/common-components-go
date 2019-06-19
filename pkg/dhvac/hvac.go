package dhvac

import (
	"encoding/json"
)

//Hvac hvac driver representation
type Hvac struct {
	Mac             string  `json:"mac"`
	FullMac         string  `json:"fullMac"`
	IP              string  `json:"ip"`
	Group           int     `json:"group"`
	Protocol        string  `json:"protocol"`
	SwitchMac       string  `json:"switchMac"`
	IsConfigured    bool    `json:"isConfigured"`
	SoftwareVersion float32 `json:"softwareVersion"`
	HardwareVersion float32 `json:"hardwareVersion"`
	Error           int     `json:"error"`
	Setpoint        int     `json:"setpoint"`
	Heat            int     `json:"heat"`   //percentage heat setpoint
	Cold            int     `json:"cold"`   //percentage cold setpoint
	Damper          int     `json:"damper"` //percentage open
	DumpFrequency   int     `json:"dumpFrequency"`
	FriendlyName    string  `json:"friendlyName"`
	Label           *string `json:"label,omitempty"`
}

//HvacSetup initial setup send by the server when the driver is authorized
type HvacSetup struct {
	Mac                    string  `json:"mac"`
	SwitchMac              string  `json:"switchMac"`
	FullMac                string  `json:"fullMac"`
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
	FullMac                string  `json:"fullMac"`
	Group                  *int    `json:"group,omitempty"`
	WindowStatus           *bool   `json:"windowStatus,omitempty"`
	Temperature            *int    `json:"temperature,omitempty"` // 1/10°C
	Presence               *bool   `json:"presence,omitempty"`
	Mode                   *string `json:"mode,omitempty"` //heat/cold
	DewPoint               *bool   `json:"dewPoint,omitempty"`
	C02                    *int    `json:"c02,omitempty"`
	TemperatureSelect      *int    `json:"temperatureSelect,omitempty"`      //????
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
