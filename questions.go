package main

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func initialQuestions() {
	questions := []*survey.Question{
		{
			Name:     "url",
			Prompt:   &survey.Input{Message: "Server URL (e.g. https://staging.appveen.com):"},
			Validate: survey.Required,
		},
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: "Username:"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Password:"},
			Validate: survey.Required,
		},
	}
	err := survey.Ask(questions, &Init)
	check(err)
}

func selectAppQuestion() string {
	appListToArrayOfStrings()
	questions := []*survey.Question{
		{
			Name: "app",
			Prompt: &survey.Select{
				Message: "Select app:",
				Options: Apps,
			},
			Validate: survey.Required,
		},
	}
	var app string
	err := survey.Ask(questions, &app)
	check(err)
	logger.Println("Selected app " + app)
	return app
}

func selectPartnerQuestion() string {
	partnerListToArrayOfStrings()
	questions := []*survey.Question{
		{
			Name: "partner",
			Prompt: &survey.MultiSelect{
				Message: "Select partners:",
				Options: Partners,
			},
			Validate: survey.Required,
		},
	}
	var selectedPartners []string
	err := survey.Ask(questions, &selectedPartners)
	check(err)

	logger.Println("Selected Partners")
	selectedPartnerIds := []string{}
	for _, v := range selectedPartners {
		for _, p := range PartnerList {
			if v == p.Name {
				selectedPartnerIds = append(selectedPartnerIds, p.ID)
				logger.Println(" Partners " + p.ID + ":" + p.Name)
			}
		}
	}
	concatenatedStringOfPartnerIds := []string{}
	for _, v := range selectedPartnerIds {
		concatenatedStringOfPartnerIds = append(concatenatedStringOfPartnerIds, "%22"+v+"%22")
	}
	return strings.Join(concatenatedStringOfPartnerIds, ",")
}

func selectFlowsQuestion() []string {
	partnerNameFormat := color.New(color.FgYellow, color.Italic).SprintfFunc()
	statusFormat := color.New(color.FgBlue, color.Italic).SprintfFunc()
	flowList := []string{}
	for _, v := range FlowList {
		menuItem := partnerNameFormat(getPartnerName(v.Partner) + ":")
		menuItem += v.Name + " "
		menuItem += "(" + v.ID + ") "
		menuItem += statusFormat(v.Status)
		flowList = append(flowList, menuItem)
	}
	questions := []*survey.Question{
		{
			Name: "flow",
			Prompt: &survey.MultiSelect{
				Message:  "Select flows:",
				Options:  flowList,
				PageSize: 15,
			},
			Validate: survey.Required,
		},
	}
	var selectedFlows []string
	err := survey.Ask(questions, &selectedFlows)
	check(err)

	logger.Println("Selected Flows")
	selectedFlowIds := []string{}
	for _, v := range selectedFlows {
		flowid := strings.Split(v, "(")[1]
		flowid = strings.Split(flowid, ")")[0]

		flowName := strings.Split(v, "(")[0]
		flowName = strings.Split(flowName, ":")[1]
		flowName = strings.Split(flowName, "[0m")[1]

		partnerName := strings.Split(v, ":")[0]
		partnerName = strings.Split(partnerName, "[33;3m")[1]

		logger.Println(" Partner " + partnerName + "/Flow " + flowid + ":" + flowName)
		selectedFlowIds = append(selectedFlowIds, flowid)
	}
	logger.Println("Selected flow ids: " + strings.Join(selectedFlowIds, ","))
	return selectedFlowIds
}
