// Copyright 2021 The SucuriAPI-Go AUTHORS. All rights reserved.
//
// Use of this source code is governed by an MIT License
// license that can be found in the LICENSE file.

// TODO Audit trail
// TODO Get protected pages /request.params.Add(setting, value)/
// TODO Add/Remove site

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

// ClearCache generates a SucuriRequest that will clear the site cache
func (s Sucuri) ClearCache() SucuriRequest {
	request := SucuriRequest{
		prefix: "Clearing Cache",
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", "clear_cache")
	return request
}

// ClearFileFromCache generates a SucuriRequest that will clear the file specified by the path from cache
func (s Sucuri) ClearFileFromCache(filepath string) SucuriRequest {
	request := s.ClearCache()
	request.params.Add("file", filepath)
	return request
}

// WhitelistIP generates a SucuriRequest that adds the specified IP address to the whitelisted IPs list.
// If remove is set to true the specified IP address will be removed from the whitelisted IP addresses
func (s Sucuri) WhitelistIP(ip string, remove bool) SucuriRequest {
	action := "allowlist_ip"
	prefix := "Whitelisting IP "
	if remove {
		action = "delete_allowlist_ip"
		prefix = "Removing whitelisted IP "
	}
	request := SucuriRequest{
		prefix: prefix + ip,
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", action)
	request.params.Add("ip", ip)
	return request
}

// WhitelistIP generates a SucuriRequest that adds the specified IP address to the whitelisted IPs list.
// If remove is set to true the specified IP address will be removed from the whitelisted IP addresses
func (s Sucuri) BlacklistIP(ip string, remove bool) SucuriRequest {
	action := "blacklist_ip"
	prefix := "Blacklisting IP "
	if remove {
		action = "delete_blacklist_ip"
		prefix = "Removing blacklisted IP "
	}
	request := SucuriRequest{
		prefix: prefix + ip,
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", action)
	request.params.Add("ip", ip)
	return request
}

// WhitelistPath generates a SucuriRequest that adds the specified Path to the whitelisted paths list, with the specified pattern (matches|begins_with|ends_with|equals).
func (s Sucuri) WhitelistPath(path string, pattern string) SucuriRequest {
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

// BlacklistPath generates a SucuriRequest that adds the specified Path to the blacklisted paths list, with the specified pattern (matches|begins_with|ends_with|equals).
func (s Sucuri) BlacklistPath(path string, pattern string) SucuriRequest {
	request := SucuriRequest{
		prefix: "Blacklisting Path '" + path + "'",
		sucuri: s,
		params: url.Values{},
	}
	request.params.Add("a", "update_setting")
	request.params.Add("blacklist_dir", path)
	request.params.Add("blacklist_dir_pattern", pattern)
	return request
}
