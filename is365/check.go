package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	apiURL    = "https://login.microsoftonline.com/common/GetCredentialType"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"
)

func checkEmail(username, proxy string) error {
	resp, err := doPost(username, proxy)
	if err != nil {
		return err
	}
	switch resp.IfExistsResult {
	case 0:
		if resp.ThrottleStatus == 1 {
			return ErrRateLimiting{}
		}
		return nil
	case 1:
		return ErrUsernameDoesNotExist{}
	case 5:
		return nil
	case 6:
		return nil
	default:
		return fmt.Errorf("undefined error")
	}
}

type GetCredentialTypeBody struct {
	Username string `json:"username"`
}

type GetCredentialTypeResponse struct {
	Credentials struct {
		CertAuthParams  interface{} `json:"CertAuthParams"`
		FacebookParams  interface{} `json:"FacebookParams"`
		FidoParams      interface{} `json:"FidoParams"`
		GoogleParams    interface{} `json:"GoogleParams"`
		HasPassword     bool        `json:"HasPassword"`
		PrefCredential  int64       `json:"PrefCredential"`
		RemoteNgcParams interface{} `json:"RemoteNgcParams"`
		SasParams       interface{} `json:"SasParams"`
	} `json:"Credentials"`
	Display        string `json:"Display"`
	EstsProperties struct {
		DomainType         int64       `json:"DomainType"`
		UserTenantBranding interface{} `json:"UserTenantBranding"`
	} `json:"EstsProperties"`
	IfExistsResult     int64  `json:"IfExistsResult"`
	IsSignupDisallowed bool   `json:"IsSignupDisallowed"`
	IsUnmanaged        bool   `json:"IsUnmanaged"`
	ThrottleStatus     int64  `json:"ThrottleStatus"`
	Username           string `json:"Username"`
	APICanary          string `json:"apiCanary"`
}

func doPost(username, proxy string) (GetCredentialTypeResponse, error) {
	body, err := json.Marshal(&GetCredentialTypeBody{
		Username: username,
	})
	if err != nil {
		return GetCredentialTypeResponse{}, err
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	if proxy != "" {
		transport, err := createTransport(proxy)
		if err != nil {
			return GetCredentialTypeResponse{}, fmt.Errorf("error creating transport: %v", err)
		}
		client.Transport = transport
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return GetCredentialTypeResponse{}, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "close")

	var getCredResp GetCredentialTypeResponse
	resp, err := client.Do(req)
	if err != nil {
		return GetCredentialTypeResponse{}, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &getCredResp)
	if err != nil {
		return GetCredentialTypeResponse{}, err
	}
	return getCredResp, nil
}

func createTransport(proxy string) (*http.Transport, error) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return &http.Transport{}, fmt.Errorf("cannot parse proxy url: %v", proxy)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	return transport, nil
}
