// Copyright 2021 The SucuriAPI-Go AUTHORS. All rights reserved.
//
// Use of this source code is governed by an MIT License
// license that can be found in the LICENSE file.

package SucuriAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// TODO Show Settings response

type SucuriRequest struct {
	prefix string
	sucuri Sucuri
	params url.Values
}

type sucuriResponse struct {
	Status      int               `json:"status,omitempty"`
	Messages    []string          `json:"messages,omitempty"`
	Action      string            `json:"action,omitempty"`
	RequestTime int               `json:"request_time,omitempty"`
	Verbose     int               `json:"verbose,omitempty"`
	Output      map[string]string `json:"output,omitempty"`
}

func (r SucuriRequest) Submit() {
	r.params.Add("k", r.sucuri.ApiKey)
	r.params.Add("s", r.sucuri.ApiSecret)
	requestURL, err := url.Parse(r.sucuri.Url + "&" + r.params.Encode())
	if err != nil {
		fmt.Println(err)
	}
	response := sucuriResponse{Output: make(map[string]string)}
	resp, err := http.Get(requestURL.String())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(body))
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Error during request: %s\n", err)
		fmt.Println(string(body))
	} else {
		if response.Status == 0 {
			fmt.Println("No Change - " + r.prefix + "; " + response.Messages[0])
		} else if response.Status == 1 {
			fmt.Println(r.prefix + "; " + response.Messages[0])
		}
	}
}
