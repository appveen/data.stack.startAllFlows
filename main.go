package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	initialQuestions()
	Login()

	responseBody := Get("/api/a/rbac/app?select=_id&count=-1")
	err := json.Unmarshal(responseBody, &AppList)
	check(err)

	selectedApp := selectAppQuestion()
	fmt.Println(selectedApp)

	responseBody = Get("/api/a/pm/partner?page=1&count=-1&filter=%7B%22app%22:%22" + selectedApp + "%22%7D&select=name")
	err = json.Unmarshal(responseBody, &PartnerList)
	check(err)
	selectedPartners := selectPartnerQuestion()

	filterForFlows := "filter=%7B%22app%22:%22" + selectedApp + "%22"
	filterForFlows += ",%20%22partner%22:%20%7B%22$in%22:%5B" + selectedPartners + "%5D%7D%7D"

	responseBody = Get("/api/a/pm/flow?page=1&count=-1&select=name,app,partner,status&" + filterForFlows)
	err = json.Unmarshal(responseBody, &FlowList)
	check(err)

	selectedFlows := selectFlowsQuestion()
	fmt.Println(selectedFlows)
}
