package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
)

type Loadavg struct {
	LoadAverage1     float64
	LoadAverage5     float64
	LoadAverage10    float64
	RunningProcesses int
	TotalProcesses   int
	LastProcessId    int
}

func ParseLoadAvg() (*Loadavg, error) {
	switch runtime.GOOS {
	case "linux":
		return parse_linux()
	default:
		return nil, errors.New("loadavg unimplemented on " + runtime.GOOS)
	}
}

func parse_linux() (*Loadavg, error) {
	self := new(Loadavg)

	raw, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return self, err
	}

	fmt.Sscanf(string(raw), "%f %f %f %d/%d %d",
		&self.LoadAverage1, &self.LoadAverage5, &self.LoadAverage10,
		&self.RunningProcesses, &self.TotalProcesses,
		&self.LastProcessId)

	return self, nil
}
