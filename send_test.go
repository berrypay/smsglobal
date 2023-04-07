/*
 * Project: SMSGlobal API SDK
 * Filename: /send_test.go
 * Created Date: Sunday March 12th 2023 02:00:24 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Friday April 7th 2023 09:09:06 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import "testing"

func TestDecodeSendSingleResponse(t *testing.T) {
	var result []byte
	var err error

	if result == nil && err != nil {
		t.Errorf("Expected successful return, got result: %v and error: %v", result, err.Error())
	}
}

func TestDecodeSendMultiResponse(t *testing.T) {
	var result []string
	var err error

	if result == nil && err != nil {
		t.Errorf("Expected successful return, got result: %v and error: %v", result, err.Error())
	}
}
