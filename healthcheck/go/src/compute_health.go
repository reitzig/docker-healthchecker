package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type HealthCheck struct {
	Description string
	Command     []string
	Timeout     time.Duration
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func expandEnv(slice []string) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = os.ExpandEnv(v)
	}
	return result
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

			// TODO: Use timout
			// TODO: capture stdout
			parameters := expandEnv(check.Command[1:])
			cmd := exec.Command(check.Command[0], parameters...)
			err := cmd.Run()

			if err != nil {
				fmt.Printf("   -> ERROR: %+v\n", err)
			} else {
				fmt.Println("   -> OKAY")
			}

			// TODO: Collect result
		}

		// TODO: Write summary
		time.Sleep(time.Duration(checksInterval) * time.Second)
	}
}
