package lib

import (
	"bytes"
	"net/http"
)

func IfconfigCoMyIP() (string, error) {
	resp, err := http.Get("https://ifconfig.co/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := new(bytes.Buffer)
	if _, err = body.ReadFrom(resp.Body); err != nil {
		return "", err
	}
	return body.String(), nil
}
