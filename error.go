/*
 * Project: Macrokiosk SMS Gateway API SDK
 * Filename: /error.go
 * Created Date: Sunday March 12th 2023 00:03:25 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Friday April 7th 2023 09:04:46 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import "fmt"

type SmsGlobalResponseBodyError struct {
	ViolationPart []byte
	Message       string
}

func (m *SmsGlobalResponseBodyError) Error() string {
	return fmt.Sprintf("Unexpected response body structure found. Violation part: %s, Message: %s", m.ViolationPart, m.Message)
}

type SmsGlobalError struct {
	Code    string
	Message string
}

func (m *SmsGlobalError) Error() string {
	return fmt.Sprintf("API Error: %s %s", m.Code, m.Message)
}
