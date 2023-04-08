/*
 * Project: SMSGlobal API SDK
 * Filename: /check-balance.go
 * Created Date: Friday April 7th 2023 15:04:22 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Saturday April 8th 2023 13:07:43 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/berrypay/runegen"
)

type CreditBalanceResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

func GetAccountBalance(timeout int) (*CreditBalanceResponse, error) {
	req, err := http.NewRequest("GET", GetFullPath(UserCreditBalanceAPI), nil)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, err.Error())
		}
		return nil, err
	}

	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}

	nonce := runegen.GetRandom(7, 32)
	// as per API documentation, ts must be a Unix timestamp
	ts := time.Now().Unix()
	req.Header.Set("Authorization", NewAuthHeader(ts, nonce, req.Method, req.RequestURI, ""))
	resp, err := client.Do(req)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, err.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403:
		retErr := NewFailedCallError(resp.StatusCode)
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, retErr.Error())
		}
		return nil, NewFailedCallError(resp.StatusCode)
	case 404:
		retErr := NewFailedCallError(resp.StatusCode)
		retErr.Message = "balance not available for postpaid user"
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, retErr.Error())
		}
		return nil, retErr
	case 405:
		retErr := NewFailedCallError(resp.StatusCode)
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, retErr.Error())
		}
		return nil, retErr
	}

	retVal, err := decodeAccountBalanceResponse(resp.Body)
	if err != nil {
		return nil, NewSmsGlobalPayloadDecodeError(err.Error())
	}

	return retVal, nil
}

func decodeAccountBalanceResponse(body io.ReadCloser) (*CreditBalanceResponse, error) {
	var creditBalanceResponse CreditBalanceResponse
	err := json.NewDecoder(body).Decode(&creditBalanceResponse)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, err.Error())
		}
		return nil, &SmsGlobalError{Message: err.Error()}
	}

	return &creditBalanceResponse, nil
}
