package lib

import (
	"bytes"
	"net/http"
	"encoding/json"
)

type WTF struct {
	YourFuckingIPAddress	string	`json:"YourFuckingIPAddress"`
	YourFuckingLocation	string	`json:"YourFuckingLocation"`
	YourFuckingHostname	string	`json:"YourFuckingHostname"`
	YourFuckingISP		string	`json:"YourFuckingISP"`
	YourFuckingTorExit	bool	`json:"YourFuckingTorExit"`
	YourFuckingCountryCode	string	`json:"YourFuckingCountryCode"`
}

func WTFMyIP() (*WTF, error) {
	wtf := new(WTF)
	resp, err := http.Get("https://wtfismyip.com/json")
	if err != nil {
		return wtf, err
	}
	defer resp.Body.Close()
	body := new(bytes.Buffer)
	if _, err = body.ReadFrom(resp.Body); err != nil {
		return wtf, err
	}
	if err := json.Unmarshal(body.Bytes(), wtf); err != nil {
		return wtf, err
	}
	return wtf, nil
}
