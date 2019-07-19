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
	Protocol              string  `json:"protocol"`
	IP                    string  `json:"ip"`
	LastSystemUpgradeDate string  `json:"lastUpgradeDate"`
	IsConfigured          *bool   `json:"isConfigured"`
	FriendlyName          string  `json:"friendlyName"`
	DumpFrequency         int     `json:"dumpFrequency"`
	Profil                string  `json:"profil"` //I/O card configuration none/puls
	Label                 *string `json:"label,omitempty"`
}

//SwitchConfig content
type SwitchConfig struct {
	Switch
	Cluster       *int                           `json:"cluster"`
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
}

//SwitchStatus description
type SwitchStatus struct {
	Switch
	Cluster       int                              `json:"cluster"`
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
	StatePuls1    int                              `json:"statePuls1"`
	StatePuls2    int                              `json:"statePuls2"`
	StatePuls3    int                              `json:"statePuls3"`
	StatePuls4    int                              `json:"statePuls4"`
	StatePuls5    int                              `json:"statePuls5"`
	StateBaes     int                              `json:"stateBaes"`
	LedsPower     int64                            `json:"ledsPower"`
	BlindsPower   int64                            `json:"blindsPower"`
	HvacsPower    int64                            `json:"hvacsPower"`
	TotalPower    int64                            `json:"totalPower"`
	HvacsEnergy   int64                            `json:"hvacsEnergy"`
	LedsEnergy    int64                            `json:"ledsEnergy"`
	BlindsEnergy  int64                            `json:"blindsEnergy"`
	TotalEnergy   int64                            `json:"totalEnergy"`
}

type SwitchCluster struct {
	Mac string `json:"mac"`
	IP  string `json:"ip"`
}

// SwitchDefinition description
type SwitchDefinition struct {
	IP                    string  `json:"ip"`
	LastSystemUpgradeDate string  `json:"lastUpgradeDate"`
	FriendlyName          string  `json:"friendlyName"`
	DumpFrequency         int     `json:"dumpFrequency"`
	Profil                string  `json:"profil"` //I/O card configuration none/puls
	Cluster               int     `json:"cluster"`
	Label                 *string `json:"label,omitempty"`
	EnergyLeds            int64   `json:"energyLeds"`
	EnergyBlinds          int64   `json:"energyBlinds"`
	EnergyHvacs           int64   `json:"energyHvacs"`
}

//ToSwitchDefinition convert map interface to Switch object
func ToSwitchDefinition(val interface{}) (*SwitchDefinition, error) {
	var sw SwitchDefinition
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &sw)
	return &sw, err
}

// ToJSON dump status struct
func (i SwitchDefinition) ToJSON() (string, error) {
	inrec, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(inrec), err
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

//ToSwitch convert map interface to Switch object
func ToSwitch(val interface{}) (*Switch, error) {
	var sw Switch
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &sw)
	return &sw, err
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
