package api

import "time"

// JSONInfo is JSON of info
type JSONInfo struct {
	ServerUpTime time.Time `json:"server_up_time"`
}

// GetJSON is convert Info to JSONInfo
func (info *Info) GetJSON() JSONInfo {
	return JSONInfo{
		ServerUpTime: info.serverUpTime,
	}
}

// JSONToken is JSON of token
type JSONToken struct {
	GenerateTime time.Time `json:"generate_time"`
	SpecificCode string    `json:"specific_code"`
}

// GetJSON is convert Token to JSONToken
func (token *Token) GetJSON() JSONToken {
	return JSONToken{
		GenerateTime: token.generateTime,
		SpecificCode: token.specificCode,
	}
}
