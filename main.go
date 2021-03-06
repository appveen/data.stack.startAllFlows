package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	fileName := "runLog_" + strconv.FormatInt(t.Unix(), 10) + ".log"
	loggingFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0664)
	check(err)

	logger.SetOutput(loggingFile)
	logger.SetFlags(log.LstdFlags)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	logger.Println("Disabling TLS")

	initialQuestions()
	Login()

	logger.Println("Getting the list of apps")
	responseBody := Get("/api/a/rbac/app?select=_id&count=-1")
	err = json.Unmarshal(responseBody, &AppList)
	check(err)

	selectedApp := selectAppQuestion()
	fmt.Println(selectedApp)

	responseBody = Get("/api/a/pm/partner?page=1&count=-1&filter=%7B%22app%22:%22" + selectedApp + "%22%7D&select=name&sort=name")
	err = json.Unmarshal(responseBody, &PartnerList)
	check(err)
	selectedPartners := selectPartnerQuestion()

	filterForFlows := "filter=%7B%22app%22:%22" + selectedApp + "%22"
	filterForFlows += ",%20%22partner%22:%20%7B%22$in%22:%5B" + selectedPartners + "%5D%7D%7D"

	responseBody = Get("/api/a/pm/flow?page=1&count=-1&select=name,app,partner,status&sort=_id&" + filterForFlows)
	err = json.Unmarshal(responseBody, &FlowList)
	check(err)

	selectedFlows := selectFlowsQuestion()

	fmt.Println("")
	var wg sync.WaitGroup
	for i := 0; i < len(selectedFlows); {
		wg.Add(1)
		deployFlow(&wg, selectedFlows[i])
		Logout()
		Login()
		i++
	}
	fmt.Println("")
}
