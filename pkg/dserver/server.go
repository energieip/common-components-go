package dserver

import (
	"encoding/json"

	"github.com/energieip/common-components-go/pkg/dblind"
	"github.com/energieip/common-components-go/pkg/dgroup"
	"github.com/energieip/common-components-go/pkg/dhvac"
	"github.com/energieip/common-components-go/pkg/dled"
	"github.com/energieip/common-components-go/pkg/dnanosense"
	"github.com/energieip/common-components-go/pkg/dsensor"
	"github.com/energieip/common-components-go/pkg/dswitch"
	"github.com/energieip/common-components-go/pkg/dwago"
)

// Server description
type Server struct {
	Mac      string `json:"mac"`
	Protocol string `json:"protocol"`
	IP       string `json:"ip"`
}

//ServerConfig content
type ServerConfig struct {
	Server
}

//IfcInfo ifc component description
type IfcInfo struct {
	Label          string `json:"label"` //cable label
	ModelName      string `json:"modelName"`
	Mac            string `json:"mac"` //device Mac address
	Vendor         string `json:"vendor"`
	URL            string `json:"url"`
	ProductionYear string `json:"productionYear"`
	DeviceType     string `json:"deviceType"`
	ModbusID       *int   `json:"modbusID,omitempty"`
	SlaveID        *int   `json:"slaveID,omitempty"`
}

//Status
type Status struct {
	Leds       []dled.Led             `json:"leds"`
	Sensors    []dsensor.Sensor       `json:"sensors"`
	Blinds     []dblind.Blind         `json:"blinds"`
	Hvacs      []dhvac.Hvac           `json:"hvacs"`
	Wagos      []dwago.Wago           `json:"wagos"`
	Nanosenses []dnanosense.Nanosense `json:"nanosenses"`
}

//DumpBlind
type DumpBlind struct {
	Ifc    IfcInfo           `json:"ifc"`
	Status dblind.Blind      `json:"status"`
	Config dblind.BlindSetup `json:"config"`
}

//DumpHvac
type DumpHvac struct {
	Ifc    IfcInfo         `json:"ifc"`
	Status dhvac.Hvac      `json:"status"`
	Config dhvac.HvacSetup `json:"config"`
}

//DumpLed
type DumpLed struct {
	Ifc    IfcInfo       `json:"ifc"`
	Status dled.Led      `json:"status"`
	Config dled.LedSetup `json:"config"`
}

//DumpSensor
type DumpSensor struct {
	Ifc    IfcInfo             `json:"ifc"`
	Status dsensor.Sensor      `json:"status"`
	Config dsensor.SensorSetup `json:"config"`
}

//DumpSwitch
type DumpSwitch struct {
	Ifc    IfcInfo      `json:"ifc"`
	Status SwitchDump   `json:"status"`
	Config SwitchConfig `json:"config"`
}

//DumpFrame
type DumpFrame struct {
	Ifc    IfcInfo     `json:"ifc"`
	Config Frame       `json:"config"`
	Status FrameStatus `json:"status"`
}

//DumpWago
type DumpWago struct {
	Ifc    IfcInfo         `json:"ifc"`
	Status dwago.Wago      `json:"status"`
	Config dwago.WagoSetup `json:"config"`
}

//DumpNano
type DumpNanosense struct {
	Ifc    IfcInfo                   `json:"ifc"`
	Status dnanosense.Nanosense      `json:"status"`
	Config dnanosense.NanosenseSetup `json:"config"`
}

//DumpSwitch
type DumpGroup struct {
	Status dgroup.GroupStatus `json:"status"`
	Config dgroup.GroupConfig `json:"config"`
}

//Dump
type Dump struct {
	Leds       []DumpLed       `json:"leds"`
	Sensors    []DumpSensor    `json:"sensors"`
	Blinds     []DumpBlind     `json:"blinds"`
	Hvacs      []DumpHvac      `json:"hvacs"`
	Wagos      []DumpWago      `json:"wagos"`
	Switchs    []DumpSwitch    `json:"switchs"`
	Groups     []DumpGroup     `json:"groups"`
	Frames     []DumpFrame     `json:"frames"`
	Nanosenses []DumpNanosense `json:"nanos"`
}

type Conf struct {
	Leds    []dled.LedConf       `json:"leds"`
	Sensors []dsensor.SensorConf `json:"sensors"`
	Blinds  []dblind.BlindConf   `json:"blinds"`
	Hvacs   []dhvac.HvacConf     `json:"hvacs"`
	Wagos   []dwago.WagoConf     `json:"wagos"`
	Groups  []dgroup.GroupConfig `json:"groups"`
	Switchs []SwitchConfig       `json:"switchs"`
}

//SwitchConfig content
type SwitchConfig struct {
	Mac           *string `json:"mac,omitempty"`
	FriendlyName  *string `json:"friendlyName,omitempty"`
	IP            *string `json:"ip,omitempty"`
	Cluster       *int    `json:"cluster,omitempty"`
	Profil        string  `json:"profil"` // I/O card configuration none/pulse
	IsConfigured  *bool   `json:"isConfigured,omitempty"`
	DumpFrequency *int    `json:"dumpFrequency,omitempty"`
	Label         *string `json:"label,omitempty"`
}

//SwitchCmd content
type SwitchCmd struct {
	Mac          string `json:"mac"`
	IsConfigured *bool  `json:"isConfigured,omitempty"`
}

type SwitchDump struct {
	dswitch.Switch
	Cluster      int   `json:"cluster"`
	StatePuls1   int   `json:"statePuls1"`
	StatePuls2   int   `json:"statePuls2"`
	StatePuls3   int   `json:"statePuls3"`
	StatePuls4   int   `json:"statePuls4"`
	StatePuls5   int   `json:"statePuls5"`
	StateBaes    int   `json:"stateBaes"`
	LedsPower    int64 `json:"ledsPower"`
	BlindsPower  int64 `json:"blindsPower"`
	HvacsPower   int64 `json:"hvacsPower"`
	TotalPower   int64 `json:"totalPower"`
	HvacsEnergy  int64 `json:"hvacsEnergy"`
	LedsEnergy   int64 `json:"ledsEnergy"`
	BlindsEnergy int64 `json:"blindsEnergy"`
	TotalEnergy  int64 `json:"totalEnergy"`
	ErrorCode    *int  `json:"errorCode"`
}

//ToSwitchDump convert map interface to SwitchDump object
func ToSwitchDump(val interface{}) (*SwitchDump, error) {
	var sw SwitchDump
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &sw)
	return &sw, err
}

//ToSwitchConfig convert map interface to SwitchConfig object
func ToSwitchConfig(val interface{}) (*SwitchConfig, error) {
	var sw SwitchConfig
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &sw)
	return &sw, err
}

// ToJSON dump status struct
func (i ServerConfig) ToJSON() (string, error) {
	inrec, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToServer convert map interface to ServerConfig object
func ToServer(val interface{}) (*ServerConfig, error) {
	var device ServerConfig
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &device)
	return &device, err
}

//Frame represent device object in the building map
type Frame struct {
	Label        string `json:"label"`
	FriendlyName string `json:"friendlyName"`
	Cluster      int    `json:"cluster"`
}

type FrameStatus struct {
	Label        string `json:"label"`
	FriendlyName string `json:"friendlyName"`
	Cluster      int    `json:"cluster"`
	StatePuls1   int    `json:"statePuls1"`
	StatePuls2   int    `json:"statePuls2"`
	StatePuls3   int    `json:"statePuls3"`
	StatePuls4   int    `json:"statePuls4"`
	StatePuls5   int    `json:"statePuls5"`
	StateBaes    int    `json:"stateBaes"`
	LedsPower    int64  `json:"ledsPower"`
	BlindsPower  int64  `json:"blindsPower"`
	HvacsPower   int64  `json:"hvacsPower"`
	TotalPower   int64  `json:"totalPower"`
	HvacsEnergy  int64  `json:"hvacsEnergy"`
	LedsEnergy   int64  `json:"ledsEnergy"`
	BlindsEnergy int64  `json:"blindsEnergy"`
	TotalEnergy  int64  `json:"totalEnergy"`
	Profil       string `json:"profil"`
	Error        int    `json:"error"`
}

// ToJSON dump Model struct
func (m Frame) ToJSON() (string, error) {
	inrec, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToFrame convert map interface to Model object
func ToFrame(val interface{}) (*Frame, error) {
	var m Frame
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &m)
	return &m, err
}
