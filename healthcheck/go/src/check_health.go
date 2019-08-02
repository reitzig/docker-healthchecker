package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	targetFolder := os.Getenv("SUMMARY_DIR")

	summaryJson, err := ioutil.ReadFile(targetFolder + "/all.json")
	check(err)

	var summary map[string]*HealthSummary
	err = json.Unmarshal([]byte(summaryJson), &summary)
	check(err)

	os.Exit(summary["all"].Status)
}
