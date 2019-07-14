package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type identMeResp struct {
	Address string `json:"address"`
}

func getPublicIP() (string, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get("https://v4.ident.me/.json")
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	resp := &identMeResp{}

	json.NewDecoder(r.Body).Decode(resp)

	return resp.Address, nil
}
