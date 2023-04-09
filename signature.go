/*
 * Project: SMSGlobal API SDK
 * Filename: /signature.go
 * Created Date: Friday April 7th 2023 14:38:47 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday April 9th 2023 11:40:28 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

const (
	SignatureTemplate string = "%d\n%s\n%s\n%s\n%s\n%d\n%s\n"
)

func GetSignature(ts int64, nonce string, method string, uri string, extraData string) string {
	sigRaw := fmt.Sprintf(SignatureTemplate, ts, nonce, method, uri, Settings.Host, Settings.Port, extraData)

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Signature input: \n%s", sigRaw)
	}

	return computeHMAC256(sigRaw, Settings.Credential.ApiSecret)
}

func computeHMAC256(raw string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(raw))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
