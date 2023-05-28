package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// getRelationships returns a response object full of discord account relationships
func getRelationships(authorization string) (Relationships, error) {

	client := http.Client{Timeout: 10 * time.Second}
	url := "https://discord.com/api/v9/users/@me/relationships"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return Relationships{}, errors.New("failed to create new http request: " + err.Error())
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Add("authorization", authorization)
	req.Header.Add("x-debug-options", "bugReporterEnabled")
	req.Header.Add("x-discord-locale", "en-US")
	req.Header.Add("x-discord-timezone", "America/New_York")
	req.Header.Add("x-super-properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzExMy4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTEzLjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjE5OTkzMywiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbCwiZGVzaWduX2lkIjowfQ==")

	resp, err := client.Do(req)
	if err != nil {
		return Relationships{}, errors.New("failed to do http request: " + err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Relationships{}, errors.New("failed reading response body: " + err.Error())
	}

	var relationships Relationships
	err = json.Unmarshal(body, &relationships)
	if err != nil {
		return Relationships{}, errors.New("failed unmarshalling json into relationships object: " + err.Error())
	}

	return relationships, nil
}

type (
	Relationships []struct {
		ID       string    `json:"id"`
		Type     int       `json:"type"`
		Nickname any       `json:"nickname"`
		User     User      `json:"user"`
		Since    time.Time `json:"since"`
	}

	User struct {
		ID               string `json:"id"`
		Username         string `json:"username"`
		GlobalName       any    `json:"global_name"`
		Avatar           string `json:"avatar"`
		Discriminator    string `json:"discriminator"`
		PublicFlags      int    `json:"public_flags"`
		AvatarDecoration string `json:"avatar_decoration"`
	}
)
