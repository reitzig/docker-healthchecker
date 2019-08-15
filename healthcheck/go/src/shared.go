package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type HealthCheck struct {
	Description string
	Command     []string
	Timeout     Duration
}

type HealthSummary struct {
	Status   int
	Output   string
	Duration Duration
}

type Duration time.Duration

func (d *Duration) UnmarshalJSON(data []byte) error {
	stringLiteral := string(data)
	durationString := strings.Trim(stringLiteral, "\"")
	duration, err := time.ParseDuration(durationString)
	*d = Duration(duration)
	return err
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s", time.Duration(d)))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func exitStatus(err error) int {
	exitError := err.(*exec.ExitError)
	ws := exitError.Sys().(syscall.WaitStatus)
	return ws.ExitStatus()
}

func expandEnv(slice []string) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = os.ExpandEnv(v)
	}
	return result
}
