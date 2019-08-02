package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

type HealthCheck struct {
	Description string
	Command     []string
	//Timeout     time.Duration
}

type HealthSummary struct {
	Status   int
	Output   string
	Duration time.Duration
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
