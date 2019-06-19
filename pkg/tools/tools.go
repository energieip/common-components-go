package tools

import (
	"bytes"
	"encoding/json"
	"net"
	"strings"

	"github.com/energieip/common-components-go/pkg/pconst"
)

//GetNetworkInfo return current ip and mac address
func GetNetworkInfo() (addr, ip string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			if iface.Flags&net.FlagUp != 0 && bytes.Compare(iface.HardwareAddr, nil) != 0 {
				addr = iface.HardwareAddr.String()
				ips, err := iface.Addrs()
				if err != nil {
					return
				}
				if len(ips) > 0 {
					ipValue := ips[0].String()
					ip = strings.Split(ipValue, "/")[0]
				}
				break
			}
		}
	}
	return
}

//GetMac return current mac address
func GetMac() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

//IntInSlice check if integer is present in the list
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//StringInSlice check if integer is present in the list
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//Model2Type convert model in device type
func Model2Type(model string) string {
	if strings.HasPrefix(model, "led") {
		return pconst.LED
	}
	if strings.HasPrefix(model, "bld") {
		return pconst.BLIND
	}
	if strings.HasPrefix(model, "hvac") {
		return pconst.HVAC
	}
	if strings.HasPrefix(model, "mca") {
		return pconst.SENSOR
	}
	if strings.HasPrefix(model, "swh") {
		return pconst.SWITCH
	}
	return ""
}

//ToJSON dump struct in json
func ToJSON(elt interface{}) (string, error) {
	inrec, err := json.Marshal(elt)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}
