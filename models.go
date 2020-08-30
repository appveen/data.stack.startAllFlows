package main

// Init - initial details from user
var Init = struct {
	URL      string `survey:"url"`
	Username string `survey:"username" json:"username"`
	Password string `survey:"password" json:"password"`
}{}

// LoginResponse - Response structure after logging in
var LoginResponse = struct {
	ID    string `json:"_id"`
	Token string `json:"token"`
}{}

// AppList is the list of apps
var AppList = []struct {
	Name string `json:"_id"`
}{}

// Apps is the name of apps in an array of strings
var Apps = []string{}

// PartnerList is the list of apps
var PartnerList = []struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}{}

// Partners is the name of apps in an array of strings
var Partners = []string{}

// FlowList is the list of apps
var FlowList = []struct {
	ID      string `json:"_id"`
	App     string `json:"app"`
	Name    string `json:"name"`
	Partner string `json:"partner"`
	Status  string `json:"status"`
}{}

// Flow is the list of apps
var Flow = struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	Partner string `json:"partner"`
	Status  string `json:"status"`
}{}

// Flows is the name of apps in an array of strings
var Flows = []string{}
