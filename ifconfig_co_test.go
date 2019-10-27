package ipcheck

import (
	"log"
	"testing"
)

func TestIfconfigCoMyIP(t *testing.T) {
	ip, err := IfconfigCoMyIP()
	if err != nil {
		t.Error(err.Error())
	}
	log.Print("ifconfig.co: " + ip)
}
