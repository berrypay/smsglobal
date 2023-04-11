/*
 * Project: SMSGlobal API SDK
 * Filename: /error.go
 * Created Date: Sunday March 12th 2023 00:03:25 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday April 11th 2023 10:15:54 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"fmt"
	"io"
	"net/http"
)

const (
	ErrorOutputTemplate   string = "Error: %s\n"
	WarningOutputTemplate string = "Warning: %s\n"
)

type SmsGlobalPayloadDecodeError struct {
	Message string
}

func (m *SmsGlobalPayloadDecodeError) Error() string {
	return fmt.Sprintf(ErrorOutputTemplate, m.Message)
}

func NewSmsGlobalPayloadDecodeError(message string) *SmsGlobalPayloadDecodeError {
	return &SmsGlobalPayloadDecodeError{
		Message: message,
	}
}

type SmsGlobalError struct {
	Code    int
	Message string
}

func (m *SmsGlobalError) Error() string {
	return fmt.Sprintf("[%d] %s", m.Code, m.Message)
}

func NewSmsGlobalError(code int, message string) *SmsGlobalError {
	return &SmsGlobalError{
		Code:    code,
		Message: message,
	}
}

func NewFailedCallError(resp *http.Response) *SmsGlobalError {
	code := resp.StatusCode
	var message string

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		message = "error details unavailable due to response body read error"
	}
	message = string(bodyBytes)

	return &SmsGlobalError{
		Code:    code,
		Message: message,
	}
}
