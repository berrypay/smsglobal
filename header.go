/*
 * Project: SMSGlobal API SDK
 * Filename: /header.go
 * Created Date: Friday April 7th 2023 14:31:28 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday April 9th 2023 13:36:56 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-http-utils/headers"
)

const (
	AuthHeaderTemplate string = `MAC id="%s", ts="%d", nonce="%s", mac="%s"`
)

func NewAuthHeader(ts int64, nonce string, method string, uri string, extraData string) string {
	hmac := GetSignature(ts, nonce, method, uri, extraData)
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Calculated HMAC256: %s\n", hmac)
	}
	newAuthHeader := fmt.Sprintf(AuthHeaderTemplate, Settings.Credential.ApiKey, ts, nonce, hmac)
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Setting Authorization Header: %s\n", newAuthHeader)
	}
	return fmt.Sprintf(AuthHeaderTemplate, Settings.Credential.ApiKey, ts, nonce, hmac)
}

func setHeader(req *http.Request, ts int64, nonce string, method string, uri string, extraData string) {
	authHeader := NewAuthHeader(ts, nonce, method, uri, extraData)
	req.Header.Set(headers.Authorization, authHeader)
	req.Header.Set(headers.ContentType, "application/json")
	req.Header.Set(headers.Accept, "application/json")
}
