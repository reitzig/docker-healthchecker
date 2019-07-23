package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type HealthCheck struct {
	Description string
	Command     string
	Timeout     time.Duration
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	checksFile := os.Args[1]
	checksInterval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	check(err)

	checksJSON, err := ioutil.ReadFile(checksFile)
	check(err)

	var checks []HealthCheck
	json.Unmarshal([]byte(checksJSON), &checks)

	for {
		fmt.Println("Computing health")

		for _, check := range checks {
			fmt.Printf(" - %+v?\n", check.Description)
		}

		time.Sleep(time.Duration(checksInterval) * time.Second)
	}
}
