package SucuriAPI

import (
	"fmt"
	"net/http"
	"net/url"
)

type Sucuri struct {
	url       string
	apiKey    string
	apiSecret string
}

func (s Sucuri) setUrl(newUrl string) {
	s.url = newUrl
}

func (s Sucuri) setApiKey(newApiKey string) {
	s.apiKey = newApiKey
}

func (s Sucuri) setApiSecret(newApiSecret string) {
	s.apiSecret = newApiSecret
}

func (s Sucuri) updateSetting(setting string, value string) SRequest {
	request := SRequest{
		prefix: "Updating setting '" + setting + "': " + value,
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", "update_setting")
	return request
}

func (s Sucuri) whitelistIP(ip string, remove bool) SRequest {
	action := "allowlist_ip"
	prefix := "Whitelisting IP "
	if remove {
		action = "delete_allowlist_ip"
		prefix = "Removing whitelisted IP "
	}
	request := SRequest{
		prefix: prefix + ip + "; ",
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", action)
	request.params.Add("ip", ip)
	return request
}

type SRequest struct {
	prefix string
	sucuri Sucuri
	params url.Values
}

func (r SRequest) submit() {
	r.params.Add("k", r.sucuri.apiKey)
	r.params.Add("s", r.sucuri.apiSecret)
	requestURL, err := url.Parse(r.sucuri.url + "?" + r.params.Encode())
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.Get(requestURL.String())
	if err != nil {
		fmt.Printf("Error during request: %s", err)
	}
	defer resp.Body.Close()
}
