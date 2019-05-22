package dserver

import (
	"encoding/json"
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

// ToJSON dump status struct
func (i ServerConfig) ToJSON() (string, error) {
	inrec, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
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
