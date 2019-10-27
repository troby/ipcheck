package ipcheck

import (
	"log"
	"fmt"
	"regexp"
	"testing"
)

func TestWTFMyIP(t *testing.T) {
	expected := "^[0-9.]+$"
	wtf, err := WTFMyIP()
	if err != nil {
		t.Error(err)
	}
	if match, _ := regexp.MatchString(expected, wtf.YourFuckingIPAddress); match != true {
		t.Errorf("%s does not look like an ip address", wtf.YourFuckingIPAddress)
	}
	log.Println(wtf.YourFuckingIPAddress)
	log.Println(wtf.YourFuckingLocation)
	log.Println(wtf.YourFuckingHostname)
	log.Println(wtf.YourFuckingISP)
	log.Println(fmt.Sprintf("%t", wtf.YourFuckingTorExit))
	log.Println(wtf.YourFuckingCountryCode)
}
