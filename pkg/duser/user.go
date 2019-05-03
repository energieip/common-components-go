package duser

import "encoding/json"

//UserAcess
type UserAccess struct {
	UserHash    string   `json:"userHash"`
	Priviledges []string `json:"priviledges"`
	AccessGroup []int    `json:"accessGroups"`
}

// ToJSON dump User struct
func (u UserAccess) ToJSON() (string, error) {
	inrec, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(inrec), err
}

//ToUserAccess convert map interface to UserAuthorization object
func ToUserAccess(val interface{}) (*UserAccess, error) {
	var m UserAccess
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &m)
	return &m, err
}
