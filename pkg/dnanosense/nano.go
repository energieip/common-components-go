package dnanosense

import (
	"encoding/json"
)

//Nanosense driver representation
type Nanosense struct {
	Mac           string  `json:"mac"`
	IP            string  `json:"ip"`
	Protocol      string  `json:"protocol"`
	Group         int     `json:"group"`
	Cluster       int     `json:"cluster"`
	IsConfigured  bool    `json:"isConfigured"`
	Temperature   int     `json:"temperature"` //in 1/10°C
	CO2           int     `json:"c02"`
	CoV           int     `json:"cov"`
	Error         int     `json:"error"`
	DumpFrequency int     `json:"dumpFrequency"`
	FriendlyName  string  `json:"friendlyName"`
	Label         *string `json:"label,omitempty"`
}

//NanosenseSetup initial setup send by the server when the driver is authorized
type NanosenseSetup struct {
	Mac          string            `json:"mac"`
	Cluster      int               `json:"cluster"`
	Group        int               `json:"group"`
	FriendlyName *string           `json:"friendlyName"`
	IsConfigured *bool             `json:"isConfigured,omitempty"`
	Label        *string           `json:"label,omitempty"`
	APIType      string            `json:"apiType"`
	API          map[string]string `json:"api"`
	Protocol     string            `json:"protocol"`
	ModbusOffset string            `json:"modbusOffset"`
}

//NanosenseConf customizable configuration by the server
type NanosenseConf struct {
	Mac          string  `json:"mac"`
	Cluster      *int    `json:"cluster,omitempty"`
	Group        *int    `json:"group,omitempty"`
	IsConfigured *bool   `json:"isConfigured,omitempty"`
	FriendlyName *string `json:"friendlyName,omitempty"`
	Label        *string `json:"label,omitempty"`
	ModbusOffset *string `json:"modbusOffset,omitempty"`
}

//ToNanosense convert map interface to driver object
func ToNanosense(val interface{}) (*Nanosense, error) {
	var driver Nanosense
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToNanosenseConf convert map interface to driver object
func ToNanosenseConf(val interface{}) (*NanosenseConf, error) {
	var driver NanosenseConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToNanosenseSetup convert map interface to driver object
func ToNanosenseSetup(val interface{}) (*NanosenseSetup, error) {
	var driver NanosenseSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

// ToJSON dump hvac struct
func (driver Nanosense) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

// ToJSON dump hvac setup struct
func (driver NanosenseSetup) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToJSON dump struct in json
func (driver NanosenseConf) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}
