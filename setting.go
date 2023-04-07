/*
 * Project: SMSGlobal API SDK
 * Filename: /setting.go
 * Created Date: Saturday March 11th 2023 21:04:51 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Friday April 7th 2023 09:30:38 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

type SmsGlobalCredential struct {
	MasterUser string `json:"masterUser"`
	MasterPass string `json:"masterPass"`
	ApiKey     string `json:"apiKey"`
	ApiSecret  string `json:"apiSecret"`
}

type SmsGlobalSettings struct {
	BaseUrl     string               `json:"baseUrl"`
	Credential  *SmsGlobalCredential `json:"credential"`
	DefaultFrom string               `json:"defaultFrom"`
	SmsPath     string               `json:"smsPath"`
	BalancePath string               `json:"balancePath"`
}

var Settings *SmsGlobalSettings

func init() {
	Settings = &SmsGlobalSettings{
		BaseUrl: "https://api.smsglobal.com",
		Credential: &SmsGlobalCredential{
			MasterUser: "TEST000",
			MasterPass: "",
			ApiKey:     "A_Super_Secret_Key",
			ApiSecret:  "A_Super_Secret_Phrase",
		},
		DefaultFrom: "Private Sender",
	}
}

func SetCredential(credential *SmsGlobalCredential) {
	Settings.Credential = credential
}

func GetCredential() *SmsGlobalCredential {
	return Settings.Credential
}

func SetDefaultFrom(from string) {
	Settings.DefaultFrom = from
}

func GetDefaultFrom() string {
	return Settings.DefaultFrom
}

func SetBaseUrl(baseUrl string) {
	Settings.BaseUrl = baseUrl
}

func GetBaseUrl() string {
	return Settings.BaseUrl
}
