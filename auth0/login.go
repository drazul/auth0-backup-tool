package auth0

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type loginRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int32  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (c client) login() string {
	b := loginRequest{
		ClientId:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Audience:     "https://" + c.Domain + "/api/v2/",
		GrantType:    "client_credentials",
	}
	jsonData, _ := json.Marshal(b)

	resp, err := http.Post(
		"https://"+c.Domain+"/oauth/token",
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	authResponse := loginResponse{}
	err = json.Unmarshal(responseBody, &authResponse)
	if err != nil {
		panic(err)
	}

	return authResponse.AccessToken
}
