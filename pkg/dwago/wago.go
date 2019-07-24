package dwago

import (
	"encoding/json"

	"github.com/energieip/common-components-go/pkg/dnanosense"
)

//Wago wago driver representation
type Wago struct {
	Mac             string  `json:"mac"`
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
	Mac           string            `json:"mac"`
	Cluster       int               `json:"cluster"`
	FriendlyName  *string           `json:"friendlyName"`
	IsConfigured  *bool             `json:"isConfigured,omitempty"`
	Label         *string           `json:"label,omitempty"`
	DumpFrequency *int              `json:"dumpFrequency,omitempty"`
	API           map[string]string `json:"api"`
}

//WagoConf customizable configuration by the server
type WagoConf struct {
	Mac           string  `json:"mac"`
	Cluster       *int    `json:"cluster,omitempty"`
	IsConfigured  *bool   `json:"isConfigured,omitempty"`
	FriendlyName  *string `json:"friendlyName,omitempty"`
	DumpFrequency *int    `json:"dumpFrequency,omitempty"`
	Label         *string `json:"label,omitempty"`
}

//WagoDef configuration by the switch
type WagoDef struct {
	Mac           string                    `json:"mac"`
	Cluster       *int                      `json:"cluster,omitempty"`
	IsConfigured  *bool                     `json:"isConfigured,omitempty"`
	FriendlyName  *string                   `json:"friendlyName,omitempty"`
	IP            *string                   `json:"ip"`
	Nanosenses    []dnanosense.NanosenseDef `json:"nanosenses"`
	CronJobs      []CronJobDef              `json:"cronJobs"`
	DumpFrequency *int                      `json:"dumpFrequency,omitempty"`
	Label         *string                   `json:"label,omitempty"`
}

type CronJobDef struct {
	Group    int    `json:"group"`
	ModbusID int    `json:"modbusID"`
	Action   string `json:"action"`
}

//ToWago convert map interface to wago object
func ToWago(val interface{}) (*Wago, error) {
	var driver Wago
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToWagoDef convert map interface to Wago object
func ToWagoDef(val interface{}) (*WagoDef, error) {
	var driver WagoDef
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToWagoConf convert map interface to wago object
func ToWagoConf(val interface{}) (*WagoConf, error) {
	var driver WagoConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToWagoSetup convert map interface to wago object
func ToWagoSetup(val interface{}) (*WagoSetup, error) {
	var driver WagoSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

// ToJSON dump wago struct
func (driver Wago) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

// ToJSON dump wago setup struct
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

//ToJSON dump struct in json
func (driver WagoDef) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}
