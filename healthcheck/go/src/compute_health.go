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

func main() {
	checksFile := os.Args[1]
	checksInterval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	check(err)

	targetFolder := os.Getenv("SUMMARY_DIR")

	checksJSON, err := ioutil.ReadFile(checksFile)
	check(err)

	var checks []HealthCheck
	err = json.Unmarshal(checksJSON, &checks)
	check(err)

	for {
		fmt.Println("Computing health")

		summary := make(map[string]*HealthSummary)
		allSummary := new(HealthSummary)
		allSummary.Status = 0
		allStartTime := time.Now()
		for _, check := range checks {
			fmt.Printf(" - %+v?\n", check.Description)

			checkSummary := new(HealthSummary)
			startTime := time.Now()

			// TODO: capture stdout --> https://stackoverflow.com/a/40770011/539599
			parameters := expandEnv(check.Command[1:])
			cmd := exec.Command(check.Command[0], parameters...)
			err := cmd.Run() // TODO: Use timout
			// --> https://medium.com/@vCabbage/go-timeout-commands-with-os-exec-commandcontext-ba0c861ed738

			if err != nil {
				fmt.Printf("   -> ERROR: %+v\n", err)
				checkSummary.Status = exitStatus(err)
				checkSummary.Output = "??" // TODO
				allSummary.Status = 1
			} else {
				fmt.Println("   -> OKAY")
				checkSummary.Status = 0
				checkSummary.Output = "??" // TODO
			}

			checkSummary.Duration = Duration(time.Now().Sub(startTime))
			summary[check.Description] = checkSummary
		}

		allSummary.Duration = Duration(time.Now().Sub(allStartTime))
		summary["all"] = allSummary

		summaryFile, err := json.MarshalIndent(summary, "", "    ")
		check(err)

		err = ioutil.WriteFile(targetFolder+"/all.json", summaryFile, os.ModeExclusive)
		check(err)
		err = os.Chmod(targetFolder+"/all.json", 0644)
		check(err)

		time.Sleep(time.Duration(checksInterval) * time.Second)
	}
}
