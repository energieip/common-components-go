package dswitch

import (
	"encoding/json"

	"github.com/energieip/common-components-go/pkg/dblind"
	"github.com/energieip/common-components-go/pkg/dhvac"
	"github.com/energieip/common-components-go/pkg/duser"
	"github.com/energieip/common-components-go/pkg/dwago"

	"github.com/energieip/common-components-go/pkg/dgroup"
	"github.com/energieip/common-components-go/pkg/dled"
	"github.com/energieip/common-components-go/pkg/dsensor"
	"github.com/energieip/common-components-go/pkg/service"
)

// Switch description
type Switch struct {
	Mac                   string  `json:"mac"`
	FullMac               string  `json:"fullMac"`
	Protocol              string  `json:"protocol"`
	IP                    string  `json:"ip"`
	LastSystemUpgradeDate string  `json:"lastUpgradeDate"`
	IsConfigured          *bool   `json:"isConfigured"`
	FriendlyName          string  `json:"friendlyName"`
	DumpFrequency         int     `json:"dumpFrequency"`
	Label                 *string `json:"label,omitempty"`
}

//SwitchConfig content
type SwitchConfig struct {
	Switch
	ClusterBroker map[string]SwitchCluster       `json:"clusterBroker"`
	Services      map[string]service.Service     `json:"services"`
	Groups        map[int]dgroup.GroupConfig     `json:"groups"`
	LedsSetup     map[string]dled.LedSetup       `json:"ledsSetup"`
	LedsConfig    map[string]dled.LedConf        `json:"ledsConfig"`
	BlindsSetup   map[string]dblind.BlindSetup   `json:"blindsSetup"`
	BlindsConfig  map[string]dblind.BlindConf    `json:"blindsConfig"`
	SensorsSetup  map[string]dsensor.SensorSetup `json:"sensorsSetup"`
	SensorsConfig map[string]dsensor.SensorConf  `json:"sensorsConfig"`
	HvacsSetup    map[string]dhvac.HvacSetup     `json:"hvacsSetup"`
	HvacsConfig   map[string]dhvac.HvacConf      `json:"hvacsConfig"`
	WagosSetup    map[string]dwago.WagoSetup     `json:"wagosSetup"`
	WagosConfig   map[string]dwago.WagoConf      `json:"wagosConfig"`
	Users         map[string]duser.UserAccess    `json:"users"`
	Label         *string                        `json:"label,omitempty"`
}

//SwitchStatus description
type SwitchStatus struct {
	Switch
	ErrorCode     *int                             `json:"errorCode"`
	Services      map[string]service.ServiceStatus `json:"services"`
	Leds          map[string]dled.Led              `json:"leds"`
	Sensors       map[string]dsensor.Sensor        `json:"sensors"`
	Blinds        map[string]dblind.Blind          `json:"blinds"`
	Hvacs         map[string]dhvac.Hvac            `json:"hvacs"`
	Wagos         map[string]dwago.Wago            `json:"wagos"`
	Groups        map[int]dgroup.GroupStatus       `json:"groups"`
	ClusterBroker map[string]SwitchCluster         `json:"clusterBroker"`
	Users         map[string]duser.UserAccess      `json:"users"`
	Label         *string                          `json:"label,omitempty"`
}

type SwitchCluster struct {
	Mac string `json:"mac"`
	IP  string `json:"ip"`
}

// ToJSON dump status struct
func (i SwitchCluster) ToJSON() (string, error) {
	inrec, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToSwitchCluster convert map interface to SwitchCluster object
func ToSwitchCluster(val interface{}) (*SwitchCluster, error) {
	var cluster SwitchCluster
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &cluster)
	return &cluster, err
}

// ToJSON dump status struct
func (status SwitchStatus) ToJSON() (string, error) {
	inrec, err := json.Marshal(status)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

// ToJSON dump switch struct
func (status Switch) ToJSON() (string, error) {
	inrec, err := json.Marshal(status)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

// ToJSON dump switch config struct
func (config SwitchConfig) ToJSON() (string, error) {
	inrec, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}
