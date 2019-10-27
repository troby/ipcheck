package ipcheck

import (
	"fmt"
	"log"
	"sync"
	"time"
	"strings"
)

type IPCheck struct {
	ResponseChannel		chan string
	Quiet			bool
	Response		string
	Timeout			time.Duration
	Mutex			sync.Mutex
}

func NewIPCheck(q bool) *IPCheck {
	i := new(IPCheck)
	i.ResponseChannel = make(chan string)
	i.Quiet = q
	duration, err := time.ParseDuration("5s")
	if err != nil {
		panic(err.Error())
	}
	i.Timeout = duration
	return i
}

func (i *IPCheck) Wait() {
	timer := time.NewTimer(i.Timeout)
	select {
	case i.Response = <-i.ResponseChannel:
		return
	case <-timer.C:
		i.Logger(`timeout reached: ` + i.Timeout.String())
	}
	fmt.Println(i.Response)
}

func (i *IPCheck) Start() {
	// wtf thread
	go func() {
		wtf, err := WTFMyIP()
		if err != nil {
			i.Logger(`wtf lookup failed: ` + err.Error())
		}
		if !i.Quiet {
			i.Mutex.Lock()
			i.Logger(`wtf: ` + wtf.YourFuckingIPAddress)
			i.Logger(`wtf: ` + wtf.YourFuckingLocation)
			i.Logger(`wtf: ` + wtf.YourFuckingHostname)
			i.Logger(`wtf: ` + wtf.YourFuckingISP)
			i.Logger(`wtf: ` + wtf.YourFuckingCountryCode)
			i.Mutex.Unlock()
		}
		i.ResponseChannel <- strings.TrimSpace(wtf.YourFuckingIPAddress)
	}()
	// ifconfig thread
	go func() {
		addr, err := IfconfigCoMyIP()
		if err != nil {
			i.Logger(`ifconfig lookup failed: ` + err.Error())
		}
		i.ResponseChannel <- strings.TrimSpace(addr)
	}()
}

func (i *IPCheck) Logger(m string) {
	if !i.Quiet {
		log.Printf("%s", m)
	}
}
