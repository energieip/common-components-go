package dsensor

import (
	"encoding/json"
)

//Sensor driver representation
type Sensor struct {
	IP                         string  `json:"ip"`
	Mac                        string  `json:"mac"`
	Group                      int     `json:"group"`
	Protocol                   string  `json:"protocol"`
	SwitchMac                  string  `json:"switchMac"`
	IsConfigured               bool    `json:"isConfigured"`
	SoftwareVersion            float32 `json:"softwareVersion"`
	HardwareVersion            float32 `json:"hardwareVersion"`
	IsBleEnabled               bool    `json:"isBleEnabled"`
	Temperature                int     `json:"temperature"`
	Humidity                   int     `json:"humidity"`
	Error                      int     `json:"error"`
	Brightness                 int     `json:"brightness"`
	Presence                   bool    `json:"presence"`
	BrightnessCorrectionFactor float32 `json:"brightnessCorrectionFactor"`
	BrightnessCorrectionOffset float32 `json:"brightnessCorrectionOffset"`
	ThresholdPresence          int     `json:"thresholdPresence"`
	TemperatureOffset          int     `json:"temperatureOffset"`
	BrightnessRaw              int     `json:"brightnessRaw"`
	LastMovement               int     `json:"lastMovement"`
	VoltageInput               int     `json:"voltageInput"`
	TemperatureRaw             int     `json:"temperatureRaw"`
	FriendlyName               string  `json:"friendlyName"`
	DumpFrequency              int     `json:"dumpFrequency"`
	BleMode                    string  `json:"bleMode"` //bleMode could be: remote/iBeacon/ptm
	PtmMac                     string  `json:"mac_ptm"` //In ptm mode it corresponds to the enOcean switch
	IBeaconUUID                string  `json:"iBeaconUUID"`
	IBeaconMajor               int     `json:"iBeaconMajor"`
	IBeaconMinor               int     `json:"iBeaconMinor"`
	IBeaconTxPower             int     `json:"iBeaconTxPower"`
	Label                      *string `json:"label,omitempty"`
}

//SensorSetup initial setup send by the server when the driver is authorized
type SensorSetup struct {
	Mac                        string   `json:"mac"`
	Group                      *int     `json:"group,omitempty"`
	BrightnessCorrectionFactor *float32 `json:"brightnessCorrectionFactor,omitempty"`
	BrightnessCorrectionOffset *float32 `json:"brightnessCorrectionOffset,omitempty"`
	ThresholdPresence          *int     `json:"thresholdPresence,omitempty"`
	TemperatureOffset          *int     `json:"temperatureOffset,omitempty"`
	IsBleEnabled               *bool    `json:"isBleEnabled,omitempty"`
	FriendlyName               *string  `json:"friendlyName,omitempty"`
	SwitchMac                  string   `json:"switchMac"`
	IsConfigured               *bool    `json:"isConfigured,omitempty"`
	DumpFrequency              int      `json:"dumpFrequency"`
	Label                      *string  `json:"label,omitempty"`
	BleMode                    *string  `json:"bleMode,omitempty"` //bleMode could be: remote/iBeacon/ptm
	PtmMac                     *string  `json:"mac_ptm,omitempty"` //In ptm mode it corresponds to the enOcean switch
	IBeaconUUID                *string  `json:"iBeaconUUID,omitempty"`
	IBeaconMajor               *int     `json:"iBeaconMajor,omitempty"`
	IBeaconMinor               *int     `json:"iBeaconMinor,omitempty"`
	IBeaconTxPower             *int     `json:"iBeaconTxPower,omitempty"`
}

//SensorConf customizable configuration by the server
type SensorConf struct {
	Mac                        string   `json:"mac"`
	Group                      *int     `json:"group,omitempty"`
	BrightnessCorrectionFactor *float32 `json:"brightnessCorrectionFactor,omitempty"`
	BrightnessCorrectionOffset *float32 `json:"brightnessCorrectionOffset,omitempty"`
	IsConfigured               *bool    `json:"isConfigured,omitempty"`
	ThresholdPresence          *int     `json:"thresholdPresence,omitempty"`
	TemperatureOffset          *int     `json:"temperatureOffset,omitempty"`
	IsBleEnabled               *bool    `json:"isBleEnabled,omitempty"`
	FriendlyName               *string  `json:"friendlyName,omitempty"`
	DumpFrequency              *int     `json:"dumpFrequency,omitempty"`
	Label                      *string  `json:"label,omitempty"`
	BleMode                    *string  `json:"bleMode,omitempty"` //bleMode could be: remote/iBeacon/ptm
	IBeaconUUID                *string  `json:"iBeaconUUID,omitempty"`
	IBeaconMajor               *int     `json:"iBeaconMajor,omitempty"`
	IBeaconMinor               *int     `json:"iBeaconMinor,omitempty"`
	IBeaconTxPower             *int     `json:"iBeaconTxPower,omitempty"`
	PtmMac                     *string  `json:"mac_ptm,omitempty"` //In ptm mode it corresponds to the enOcean switch
}

//ToSensor convert interface to Sensor object
func ToSensor(val interface{}) (*Sensor, error) {
	var cell Sensor
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &cell)
	return &cell, err
}

//ToSensorSetup convert interface to SensorSetup object
func ToSensorSetup(val interface{}) (*SensorSetup, error) {
	var cell SensorSetup
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &cell)
	return &cell, err
}

//ToSensorConf convert map interface to Sensor object
func ToSensorConf(val interface{}) (*SensorConf, error) {
	var light SensorConf
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &light)
	return &light, err
}

// ToJSON dump sensor struct
func (sensor Sensor) ToJSON() (string, error) {
	inrec, err := json.Marshal(sensor)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

// ToJSON dump sensor struct
func (sensor SensorSetup) ToJSON() (string, error) {
	inrec, err := json.Marshal(sensor)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToJSON dump struct in json
func (sensor SensorConf) ToJSON() (string, error) {
	inrec, err := json.Marshal(sensor)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}
