package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Check - display error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ifErrorExit(e error) {
	if e != nil {
		os.Exit(0)
	}
}

func appListToArrayOfStrings() {
	for _, v := range AppList {
		Apps = append(Apps, v.Name)
	}
}

func partnerListToArrayOfStrings() {
	for _, v := range PartnerList {
		Partners = append(Partners, v.Name)
	}
}

func getPartnerName(_id string) string {
	for _, v := range PartnerList {
		if _id == v.ID {
			return v.Name
		}
	}
	return ""
}

func getFlowName(_id string) string {
	for _, v := range FlowList {
		if _id == v.ID {
			return v.Name
		}
	}
	return ""
}

func deployFlow(_wg *sync.WaitGroup, _flowID string) {
	defer _wg.Done()
	responseBody := Get("/api/a/pm/flow/" + _flowID + "?select=name,partner,status")
	err := json.Unmarshal(responseBody, &Flow)
	check(err)
	fmt.Print("Flow " + color.YellowString(getPartnerName(Flow.Partner)+":"+getFlowName(_flowID)) + " :: ")
	// Deploy
	Put("/api/a/pm/flow/"+_flowID+"/start", "{}")
	// Fetch
	fmt.Print(color.BlueString("Pending"))
	for {
		time.Sleep(1000 * time.Millisecond)
		responseBody := Get("/api/a/pm/flow/" + _flowID + "?select=name,partner,status")
		err := json.Unmarshal(responseBody, &Flow)
		check(err)
		if Flow.Status != "Pending" {
			break
		}
	}
	fmt.Print("\b\b\b\b\b\b\b")
	// fmt.Print("Flow " + color.YellowString(getFlowName(_flowID)) + " :: ")
	if Flow.Status == "Active" {
		fmt.Println(color.GreenString(Flow.Status) + " ")
	}
	if Flow.Status == "Stopped" {
		fmt.Println(color.RedString(Flow.Status))
	}
}
