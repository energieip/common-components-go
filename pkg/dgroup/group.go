package dgroup

import (
	"encoding/json"
)

const (
	// Sensor rules when there are several sensors in the same group
	SensorAverage = "average"
	SensorMin     = "min"
	SensorMax     = "max"
)

//GroupConfig representation
type GroupConfig struct {
	Group                   int      `json:"group"` //groupID
	SensorRule              *string  `json:"sensorRule,omitempty"`
	Auto                    *bool    `json:"auto,omitempty"`
	SlopeStartManual        *int     `json:"slopeStartManual,omitempty"`
	SlopeStopManual         *int     `json:"slopeStopManual,omitempty"`
	SlopeStartAuto          *int     `json:"slopeStartAuto,omitempty"`
	SlopeStopAuto           *int     `json:"slopeStopAuto,omitempty"`
	Watchdog                *int     `json:"watchdog,omitempty"`
	CorrectionInterval      *int     `json:"correctionInterval,omitempty"`
	RuleBrightness          *int     `json:"ruleBrightness,omitempty"`
	RulePresence            *int     `json:"rulePresence,omitempty"`
	SetpointLeds            *int     `json:"setpointLeds,omitempty"`
	SetpointBlinds          *int     `json:"setpointBlinds,omitempty"`
	SetpointTempOffset      *int     `json:"setpointTempOffset,omitempty"` //temperature offset (in 1/10°C)
	SetpointSlatBlinds      *int     `json:"setpointSlatBlinds,omitempty"`
	FriendlyName            *string  `json:"friendlyName,omitempty"`
	Leds                    []string `json:"leds"`       //Mac address list
	Blinds                  []string `json:"blinds"`     //Mac address list
	Hvacs                   []string `json:"hvacs"`      //Mac address list
	Nanosenses              []string `json:"nanosenses"` //MAC + offset
	Sensors                 []string `json:"sensors"`
	FirstDay                []string `json:"firstDay"` //LED Mac address list in first day position
	FirstDayOffset          *int     `json:"firstDayOffset,omitempty"`
	SetpointOccupiedCool1   *int     `json:"setpointOccupiedCool1,omitempty"`   // 1/10°C
	SetpointOccupiedHeat1   *int     `json:"setpointOccupiedHeat1,omitempty"`   // 1/10°C
	SetpointUnoccupiedCool1 *int     `json:"setpointUnoccupiedCool1,omitempty"` // 1/10°C
	SetpointUnoccupiedHeat1 *int     `json:"setpointUnoccupiedHeat1,omitempty"` // 1/10°C
	SetpointStandbyCool1    *int     `json:"setpointStandbyCool1,omitempty"`    // 1/10°C
	SetpointStandbyHeat1    *int     `json:"setpointStandbyHeat1,omitempty"`    // 1/10°C
	HvacsTargetMode         *int     `json:"hvacsTargetMode,omitempty"`
	EipDriversReset         *bool    `json:"eipDriversReset,omitempty"` // reset all eip drivers
}

//GroupStatus status dump to the server
type GroupStatus struct {
	Group                   int      `json:"group"` //groupID
	SensorRule              string   `json:"sensorRule"`
	Auto                    bool     `json:"auto"`
	Watchdog                int      `json:"watchdog"`
	SlopeStartManual        int      `json:"slopeStartManual,omitempty"`
	SlopeStopManual         int      `json:"slopeStopManual,omitempty"`
	SlopeStartAuto          int      `json:"slopeStartAuto,omitempty"`
	SlopeStopAuto           int      `json:"slopeStopAuto,omitempty"`
	CorrectionInterval      int      `json:"correctionInterval"`
	RuleBrightness          *int     `json:"ruleBrightness,omitempty"`
	RulePresence            *int     `json:"rulePresence,omitempty"`
	Error                   int      `json:"error"`
	TimeToAuto              int      `json:"timeToAuto"`
	SetpointLeds            int      `json:"setpointLeds"`
	SetpointTempOffset      int      `json:"setpointTempOffset"` //temperature offset (in 1/10°C)
	Presence                bool     `json:"presence"`
	Temperature             int      `json:"temperature"`        //from nanosense used by HVAC (in 1/10°C)
	Hygrometry              int      `json:"hygrometry"`         //from nanosense used by HVAC (in %)
	CO2                     int      `json:"co2"`                //from nanosense used by HVAC (in ppm)
	COV                     int      `json:"cov"`                //from nanosense used by HVAC (in ppm)
	CeilingTemperature      int      `json:"ceilingTemperature"` //from sensor (in 1/10°C)
	CeilingHumidity         int      `json:"ceilingHumidity"`    //from sensor (in 1/10°C)
	Brightness              int      `json:"brightness"`         // from sensor (in Lux)
	WindowsOpened           bool     `json:"windowsOpened"`      //at least one window is opened
	TimeToLeave             int      `json:"timeToLeave"`
	Leds                    []string `json:"leds"` //Mac address list
	Blinds                  []string `json:"blinds"`
	Hvacs                   []string `json:"hvacs"`
	Sensors                 []string `json:"sensors"`
	Nanosenses              []string `json:"nanosenses"` //MAC + offset
	FriendlyName            string   `json:"friendlyName"`
	FirstDay                []string `json:"firstDay"` //LED Mac address list in first day position
	FirstDayOffset          *int     `json:"firstDayOffset,omitempty"`
	SetpointLedsFirstDay    int      `json:"setpointLedsFirstDay"`
	SetpointOccupiedCool1   int      `json:"setpointOccupiedCool1"`   // 1/10°C
	SetpointOccupiedHeat1   int      `json:"setpointOccupiedHeat1"`   // 1/10°C
	SetpointUnoccupiedCool1 int      `json:"setpointUnoccupiedCool1"` // 1/10°C
	SetpointUnoccupiedHeat1 int      `json:"setpointUnoccupiedHeat1"` // 1/10°C
	SetpointStandbyCool1    int      `json:"setpointStandbyCool1"`    // 1/10°C
	SetpointStandbyHeat1    int      `json:"setpointStandbyHeat1"`    // 1/10°C
	HvacsTargetMode         int      `json:"hvacsTargetMode"`
	HvacsEffectMode         int      `json:"hvacsEffectMode"`
}

// ToMapInterface convert group struct in Map[string] interface{}
func (group GroupConfig) ToMapInterface() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(group)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

// ToJSON dump group struct
func (group GroupConfig) ToJSON() (string, error) {
	inrec, err := json.Marshal(group)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToGroupConfig convert interface to group config object
func ToGroupConfig(val interface{}) (*GroupConfig, error) {
	var group GroupConfig
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &group)
	return &group, err
}

// ToMapInterface convert GroupStatus struct in Map[string] interface{}
func (group GroupStatus) ToMapInterface() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(group)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

// ToJSON dump group GroupStatus
func (group GroupStatus) ToJSON() (string, error) {
	inrec, err := json.Marshal(group)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToGroupStatus convert interface to status object
func ToGroupStatus(val interface{}) (*GroupStatus, error) {
	var group GroupStatus
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &group)
	return &group, err
}
