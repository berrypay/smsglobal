/*
 * Project: SMSGlobal API SDK
 * Filename: /send-multi.go
 * Created Date: Saturday March 11th 2023 19:06:05 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Sunday April 9th 2023 13:42:12 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/berrypay/runegen"
)

type SendMultiSMSResponse struct {
}

func SendMulti(to string, from string, message string, title string, timeout int) (*SendMultiSMSResponse, error) {
	req, err := http.NewRequest("POST", GetFullPath(SmsAPI), nil)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, err.Error())
		}
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Sending sms by calling API endpoint at: %s\n", req.URL)
	}

	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}

	nonce := runegen.GetRandom(7, 32)
	// as per API documentation, ts must be a Unix timestamp
	ts := time.Now().Unix()
	setHeader(req, ts, nonce, req.Method, SmsAPI, "")
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

	retVal, err := decodeSendMultiSMSResponse(resp)
	if err != nil {
		if err != nil {
			fmt.Printf("Error decoding response: %s\n", err.Error())
		}
		return nil, NewSmsGlobalPayloadDecodeError(err.Error())
	}

	return retVal, nil
}

func decodeSendMultiSMSResponse(resp *http.Response) (*SendMultiSMSResponse, error) {
	// TODO: implement SendMultiSMSResponse decoding
	return nil, nil
}
