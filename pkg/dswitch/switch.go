package dswitch

import (
	"encoding/json"

	"github.com/energieip/common-components-go/pkg/dgroup"
	"github.com/energieip/common-components-go/pkg/dled"
	"github.com/energieip/common-components-go/pkg/dsensor"
	"github.com/energieip/common-components-go/pkg/service"
)

// Switch description
type Switch struct {
	Mac                   string `json:"mac"`
	Protocol              string `json:"protocol"`
	IP                    string `json:"ip"`
	Topic                 string `json:"topic"`
	LastSystemUpgradeDate string `json:"lastUpgradeDate"`
	IsConfigured          *bool  `json:"isConfigured"`
	FriendlyName          string `json:"friendlyName"`
	DumpFrequency         int    `json:"dumpFrequency"`
}

//SwitchConfig content
type SwitchConfig struct {
	Switch
	Services      map[string]service.Service     `json:"services"`
	Groups        map[int]dgroup.GroupConfig     `json:"groups"`
	LedsSetup     map[string]dled.LedSetup       `json:"ledsSetup"`
	LedsConfig    map[string]dled.LedConf        `json:"ledsConfig"`
	SensorsSetup  map[string]dsensor.SensorSetup `json:"sensorsSetup"`
	SensorsConfig map[string]dsensor.SensorConf  `json:"sensorsConfig"`
}

//SwitchStatus description
type SwitchStatus struct {
	Switch
	ErrorCode *int                             `json:"errorCode"`
	Services  map[string]service.ServiceStatus `json:"services"`
	Leds      map[string]dled.Led              `json:"leds"`
	Sensors   map[string]dsensor.Sensor        `json:"sensors"`
	Groups    map[int]dgroup.GroupStatus       `json:"groups"`
}

// ToJSON dump status struct
func (status SwitchStatus) ToJSON() (string, error) {
	inrec, err := json.Marshal(status)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

// ToJSON dump switch struct
func (status Switch) ToJSON() (string, error) {
	inrec, err := json.Marshal(status)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

// ToJSON dump switch config struct
func (config SwitchConfig) ToJSON() (string, error) {
	inrec, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}
