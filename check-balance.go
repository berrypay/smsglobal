/*
 * Project: SMSGlobal API SDK
 * Filename: /check-balance.go
 * Created Date: Friday April 7th 2023 15:04:22 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday April 9th 2023 11:30:27 +0800
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
	Balance  float64 `json:"balance,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

func GetAccountBalance(timeout int) (*CreditBalanceResponse, error) {
	req, err := http.NewRequest("GET", GetFullPath(UserCreditBalanceAPI), nil)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, err.Error())
		}
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Request URL: %s\n", req.URL)
		fmt.Printf("Request URI: %s\n", req.RequestURI)
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

	if resp.StatusCode != 200 {
		retErr := NewFailedCallError(resp)
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, retErr.Error())
		}

		return nil, retErr
	}

	retVal, err := decodeAccountBalanceResponse(resp)
	if err != nil {
		if err != nil {
			fmt.Printf("Error decoding response: %s\n", err.Error())
		}
		return nil, NewSmsGlobalPayloadDecodeError(err.Error())
	}

	return retVal, nil
}

func decodeAccountBalanceResponse(resp *http.Response) (*CreditBalanceResponse, error) {
	var creditBalanceResponse CreditBalanceResponse
	err := json.NewDecoder(resp.Body).Decode(&creditBalanceResponse)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			bodyBytes, readErr := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading response body: %s\n", readErr.Error())
			} else {
				fmt.Printf("Response Body: %s\n", string(bodyBytes))
			}
			fmt.Printf(ErrorOutputTemplate, err.Error())
		}
		return nil, &SmsGlobalError{Message: err.Error()}
	}

	return &creditBalanceResponse, nil
}
