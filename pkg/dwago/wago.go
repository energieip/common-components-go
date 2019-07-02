package dwago

import (
	"encoding/json"
)

//Wago wago driver representation
type Wago struct {
	Mac             string  `json:"mac"`
	FullMac         *string `json:"fullMac,omitempty"`
	IP              string  `json:"ip"`
	Protocol        string  `json:"protocol"`
	Cluster         int     `json:"cluster"`
	IsConfigured    bool    `json:"isConfigured"`
	SoftwareVersion float32 `json:"softwareVersion"`
	HardwareVersion float32 `json:"hardwareVersion"`
	Error           int     `json:"error"`
	FriendlyName    string  `json:"friendlyName"`
	Label           *string `json:"label,omitempty"`
}

//WagoSetup initial setup send by the server when the driver is authorized
type WagoSetup struct {
	Mac          string  `json:"mac"`
	Cluster      int     `json:"cluster"`
	FullMac      *string `json:"fullMac,omitempty"`
	FriendlyName *string `json:"friendlyName"`
	IsConfigured *bool   `json:"isConfigured,omitempty"`
	Label        *string `json:"label,omitempty"`
}

//WagoConf customizable configuration by the server
type WagoConf struct {
	Mac          string  `json:"mac"`
	FullMac      *string `json:"fullMac,omitempty"`
	Group        *int    `json:"group,omitempty"`
	IsConfigured *bool   `json:"isConfigured,omitempty"`
	FriendlyName *string `json:"friendlyName"`
	Label        *string `json:"label,omitempty"`
}

//ToWago convert map interface to Hvac object
func ToWago(val interface{}) (*Wago, error) {
	var driver Wago
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToWagoConf convert map interface to Hvac object
func ToWagoConf(val interface{}) (*WagoConf, error) {
	var driver WagoConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToWagoSetup convert map interface to Hvac object
func ToWagoSetup(val interface{}) (*WagoSetup, error) {
	var driver WagoSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

// ToJSON dump hvac struct
func (driver Wago) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

// ToJSON dump hvac setup struct
func (driver WagoSetup) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToJSON dump struct in json
func (driver WagoConf) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}
