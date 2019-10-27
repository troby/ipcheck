package lib

import (
	"log"
	"testing"
)

func TestIfconfigCoMyIP(t *testing.T) {
	ip, err := IfconfigCoMyIP()
	if err != nil {
		t.Error(err.Error())
	}
	log.Println("ifconfig.co: " + ip)
}
