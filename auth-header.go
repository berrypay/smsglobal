/*
 * Project: SMSGlobal API SDK
 * Filename: /auth-header.go
 * Created Date: Friday April 7th 2023 14:31:28 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Saturday April 8th 2023 12:43:23 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"fmt"
	"os"
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
