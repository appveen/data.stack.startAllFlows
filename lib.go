package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Check - display error
func check(e error) {
	if e != nil {
		logger.Panic(e)
	}
}

func ifErrorExit(e error) {
	if e != nil {
		logger.Panic(e)
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
	fmt.Print("Flow " + getPartnerName(Flow.Partner) + ":" + getFlowName(_flowID) + " -> ")
	// Deploy
	Put("/api/a/pm/flow/"+_flowID+"/start", "{}")
	// Fetch
	fmt.Print("Pending")
	for {
		time.Sleep(10000 * time.Millisecond)
		responseBody := Get("/api/a/pm/flow/" + _flowID + "?select=status")
		err := json.Unmarshal(responseBody, &Flow)
		check(err)
		if Flow.Status != "Pending" {
			break
		}
	}
	fmt.Print("\b\b\b\b\b\b\b")
	fmt.Println(Flow.Status + " ")
}
