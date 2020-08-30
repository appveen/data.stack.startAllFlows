package main

import (
	"os"
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

// func getListOfPartners() {
// 	for _, v := range ListOfFlows {
// 		jsonList, _ := json.Marshal(ListOfPartnerIds)
// 		list := string(jsonList)
// 		if strings.Index(list, v.Partner) == -1 {
// 			ListOfPartnerIds = append(ListOfPartnerIds, v.Partner)
// 		}
// 	}
// }
