/*
 * Project: SMSGlobal API SDK
 * Filename: /error.go
 * Created Date: Sunday March 12th 2023 00:03:25 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Saturday April 8th 2023 13:07:26 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"fmt"
	"strconv"
)

const (
	ErrorOutputTemplate string = "Error: %s\n"
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
	Code    string
	Message string
}

func (m *SmsGlobalError) Error() string {
	return fmt.Sprintf("API Error: %s %s", m.Code, m.Message)
}

func NewSmsGlobalError(code string, message string) *SmsGlobalError {
	return &SmsGlobalError{
		Code:    code,
		Message: message,
	}
}

func NewFailedCallError(statusCode int) *SmsGlobalError {
	code := strconv.Itoa(statusCode)
	var message string
	switch statusCode {
	case 400:
		message = "request contained invalid or missing data"
	case 401:
		message = "authentication failed or the authenticate header was not provided"
	case 403:
		message = "user not authorized"
	case 404:
		message = "URI does not match any of the recognized resources or resource does not exist"
	case 405:
		message = "method ot allowed, make OPTIONS request for allowed methods"
	case 406:
		message = "content type not supported"
	case 415:
		message = "content-type header not supported"
	default:
		message = "unexpected return code"
	}

	return &SmsGlobalError{
		Code:    code,
		Message: message,
	}
}
