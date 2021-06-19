package SucuriAPI

import "net/url"

// Sucuri represents the endpoint and credentials to submit a SucuriRequest to the API
type Sucuri struct {
	Url       string
	ApiKey    string
	ApiSecret string
}

// UpdateSetting generates a SucuriRequest that will overwrite the specified setting and value
func (s Sucuri) UpdateSetting(setting string, value string) SucuriRequest {
	request := SucuriRequest{
		prefix: "Updating setting '" + setting + "': " + value,
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", "update_setting")
	request.params.Add(setting, value)
	return request
}

// WhitelistIP generates a SucuriRequest that adds the specified IP address to the whitelisted IPs list.
// If remove is set to true the specified IP address will be removed from the whitelisted IP addresses
func (s Sucuri) WhitelistIP(ip string, remove bool) SucuriRequest {
	action := "allowlist_ip"
	prefix := "Whitelisting IP"
	if remove {
		action = "delete_allowlist_ip"
		prefix = "Removing whitelisted IP"
	}
	request := SucuriRequest{
		prefix: prefix + ip + "; ",
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", action)
	request.params.Add("ip", ip)
	return request
}

// WhitelistIP generates a SucuriRequest that adds the specified Path to the whitelisted paths list.
func (s Sucuri) WhitelistDir(path string, pattern string) SucuriRequest {
	request := SucuriRequest{
		prefix: "Whitelisting Path '" + path + "'",
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", "update_setting")
	request.params.Add("allowlist_dir", path)
	request.params.Add("allowlist_dir_pattern", pattern)
	return request
}
