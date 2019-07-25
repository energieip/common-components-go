package dnanosense

import (
	"encoding/json"
)

//Nanosense driver representation
type Nanosense struct {
	Group         int    `json:"group"`
	Cluster       int    `json:"cluster"`
	Temperature   int    `json:"temperature"` //in 1/10°C
	CO2           int    `json:"c02"`
	Hygrometry    int    `json:"hygrometry"`
	COV           int    `json:"cov"`
	Error         int    `json:"error"`
	DumpFrequency int    `json:"dumpFrequency"`
	FriendlyName  string `json:"friendlyName"`
	Label         string `json:"label"`
}

//NanosenseSetup initial setup send by the server when the driver is authorized
type NanosenseSetup struct {
	Cluster      int               `json:"cluster"`
	Group        int               `json:"group"`
	FriendlyName *string           `json:"friendlyName"`
	Label        string            `json:"label"`
	APIType      string            `json:"apiType"`
	API          map[string]string `json:"api"`
	ModbusOffset int               `json:"modbusOffset"`
}

//NanosenseConf customizable configuration by the server
type NanosenseConf struct {
	Cluster      *int    `json:"cluster,omitempty"`
	Group        *int    `json:"group,omitempty"`
	FriendlyName *string `json:"friendlyName,omitempty"`
	Label        string  `json:"label"`
	ModbusOffset *int    `json:"modbusOffset,omitempty"`
}

//NanosenseDef customizable configuration by the server
type NanosenseDef struct {
	Group        int    `json:"group"`
	Cluster      int    `json:"cluster"`
	FriendlyName string `json:"friendlyName"`
	Label        string `json:"label"`
	Hygrometry   int    `json:"hygrometry"`  //modbusID
	Temperature  int    `json:"temperature"` //modbusID
	CO2          int    `json:"co2"`         //modbusID
	COV          int    `json:"cov"`         //modbusID
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
