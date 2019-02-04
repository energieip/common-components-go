package dblind

import (
	"encoding/json"
)

const (
	DbStatus  = "status"
	DbConfig  = "config"
	TableName = "blinds"

	UrlHello   = "setup/hello"
	UrlStatus  = "status/dump"
	UrlSetup   = "setup/config"
	UrlSetting = "update/settings"
)

//Blind driver representation
type Blind struct {
	IP                string  `json:"ip"`
	Mac               string  `json:"mac"`
	Group             int     `json:"group"`
	Protocol          string  `json:"protocol"`
	SwitchMac         string  `json:"switchMac"`
	IsConfigured      bool    `json:"isConfigured"`
	SoftwareVersion   float32 `json:"softwareVersion"`
	HardwareVersion   float32 `json:"hardwareVersion"`
	IsBleEnabled      bool    `json:"isBleEnabled"`
	Error             int     `json:"error"`
	FriendlyName      string  `json:"friendlyName"`
	DumpFrequency     int     `json:"dumpFrequency"`
	WindowsOpen1      bool    `json:"windowOpen1"`
	WindowsOpen2      bool    `json:"windowOpen2"`
	VoltageInput      int     `json:"voltageInput"`
	DaisyChainEnabled bool    `json:"daisyChainEnabled"`
	DaisyChainPos     int     `json:"daisyChainPos"`
	LinePower         int     `json:"linePower"`
}

//ActionProfile
type ActionProfile struct {
	Action   int  `json:"action"`
	Duration *int `json:"duration,omitempty"`
}

//BlindProfile constructor recommendation
type BlindProfile struct {
	UP           []ActionProfile `json:"up"`
	Down         []ActionProfile `json:"down"`
	Stop         []ActionProfile `json:"stop"`
	WindowOpen   bool            `json:"windowOpen,omitempty"`
	SlatDuration int             `json:"slatDuration,omitempty"`
}

//BlindSetup initial setup send by the server when the driver is authorized
type BlindSetup struct {
	Mac           string       `json:"mac"`
	Group         *int         `json:"group,omitempty"`
	IsBleEnabled  *bool        `json:"isBleEnabled,omitempty"`
	FriendlyName  *string      `json:"friendlyName,omitempty"`
	SwitchMac     string       `json:"switchMac"`
	IsConfigured  *bool        `json:"isConfigured,omitempty"`
	DumpFrequency int          `json:"dumpFrequency"`
	Blind1        *int         `json:"blind1,omitempty"`
	Blind2        *int         `json:"blind2,omitempty"`
	Slat1         *int         `json:"slat1,omitempty"`
	Slat2         *int         `json:"slat2,omitempty"`
	Profile       BlindProfile `json:"profile,omitempty"`
}

//BlindConf customizable configuration by the server
type BlindConf struct {
	Mac           string  `json:"mac"`
	Group         *int    `json:"group,omitempty"`
	IsConfigured  *bool   `json:"isConfigured,omitempty"`
	IsBleEnabled  *bool   `json:"isBleEnabled,omitempty"`
	FriendlyName  *string `json:"friendlyName,omitempty"`
	DumpFrequency *int    `json:"dumpFrequency,omitempty"`
	Blind1        *int    `json:"blind1,omitempty"`
	Blind2        *int    `json:"blind2,omitempty"`
	Slat1         *int    `json:"slat1,omitempty"`
	Slat2         *int    `json:"slat2,omitempty"`
}

//ToBlind convert interface to Sensor object
func ToBlind(val interface{}) (*Blind, error) {
	var driver Blind
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToBlindSetup convert interface to blindsetup object
func ToBlindSetup(val interface{}) (*BlindSetup, error) {
	var driver BlindSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

//ToBlindConf convert map interface to BlindConf object
func ToBlindConf(val interface{}) (*BlindConf, error) {
	var driver BlindConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &driver)
	return &driver, err
}

// ToJSON dump blind struct
func (driver Blind) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

// ToJSON dump blindSetup struct
func (blind BlindSetup) ToJSON() (string, error) {
	inrec, err := json.Marshal(blind)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

//ToJSON dump struct in json
func (driver BlindConf) ToJSON() (string, error) {
	inrec, err := json.Marshal(driver)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}
