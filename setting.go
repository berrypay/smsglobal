/*
 * Project: SMSGlobal API SDK
 * Filename: /setting.go
 * Created Date: Saturday March 11th 2023 21:04:51 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday April 9th 2023 11:15:51 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import "fmt"

type SmsGlobalCredential struct {
	MasterUser string `json:"masterUser"`
	MasterPass string `json:"masterPass"`
	ApiKey     string `json:"apiKey"`
	ApiSecret  string `json:"apiSecret"`
}

type SmsGlobalSettings struct {
	Host        string              `json:"host"`
	Port        int                 `json:"port"`
	TLS         bool                `json:"tls"`
	Credential  SmsGlobalCredential `json:"credential"`
	DefaultFrom string              `json:"defaultFrom"`
}

var Settings *SmsGlobalSettings

func init() {
	Settings = &SmsGlobalSettings{
		Host: "api.smsglobal.com",
		Port: 443,
		TLS:  true,
		Credential: SmsGlobalCredential{
			MasterUser: "TEST000",
			MasterPass: "",
			ApiKey:     "A_Super_Secret_Key",
			ApiSecret:  "A_Super_Secret_Phrase",
		},
		DefaultFrom: "Private Sender",
	}
}

func SetCredential(credential SmsGlobalCredential) {
	Settings.Credential = credential
}

func GetCredential() SmsGlobalCredential {
	return Settings.Credential
}

func SetHost(host string) {
	Settings.Host = host
}

func GetHost() string {
	return Settings.Host
}

func SetTLS(tls bool) {
	Settings.TLS = tls
}

func GetTLS() bool {
	return Settings.TLS
}

func SetPort(port int) {
	Settings.Port = port
}

func GetPort() int {
	return Settings.Port
}

func GetBaseUrl() string {
	baseUrl := "https://api.smsglobal.com"
	if Settings.TLS {
		if Settings.Port != 443 {
			baseUrl = fmt.Sprintf("https://%s:%d", Settings.Host, Settings.Port)
		} else {
			baseUrl = fmt.Sprintf("https://%s", Settings.Host)
		}
	} else {
		if Settings.Port != 80 {
			baseUrl = fmt.Sprintf("http://%s:%d", Settings.Host, Settings.Port)
		} else {
			baseUrl = fmt.Sprintf("http://%s", Settings.Host)
		}
	}

	return baseUrl
}

func SetDefaultFrom(from string) {
	Settings.DefaultFrom = from
}

func GetDefaultFrom() string {
	return Settings.DefaultFrom
}
