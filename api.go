package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login - perform login
func Login() {
	url := Init.URL + "/api/a/rbac/login"
	logger.Println("Logging URL :: " + url)

	payload, err := json.Marshal(Init)
	check(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	check(err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	err = json.Unmarshal(body, &LoginResponse)
	check(err)
	logger.Println("User " + Init.Username + " logged")
}

// Logout - perform log out
func Logout() {
	url := Init.URL + "/api/a/rbac/logout"
	logger.Println("Logging URL :: " + url)

	req, err := http.NewRequest("DELETE", url, nil)
	check(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JWT "+LoginResponse.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	logger.Println("Response -> " + string(body))
	logger.Println("User " + Init.Username + " logged out")
}

// Get api call
func Get(_url string) []byte {
	logger.Println("GET " + Init.URL + _url)
	req, err := http.NewRequest("GET", Init.URL+_url, nil)
	check(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JWT "+LoginResponse.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	logger.Println("Response -> " + string(body))

	return body
}

// Put api call
func Put(_url string, _data string) []byte {
	logger.Println("PUT " + Init.URL + _url)
	logger.Println("Data ->" + _data)
	req, err := http.NewRequest("PUT", Init.URL+_url, bytes.NewBuffer([]byte(_data)))
	check(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "JWT "+LoginResponse.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	logger.Println("Response -> " + string(body))

	return body
}
