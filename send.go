/*
 * Project: SMSGlobal API SDK
 * Filename: /send.go
 * Created Date: Saturday March 11th 2023 19:06:05 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Monday April 10th 2023 09:51:48 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/berrypay/runegen"
)

type SendSMSRequest struct {
	Destination       string    `json:"destination,omitempty"`
	Message           string    `json:"message,omitempty"`
	Origin            string    `json:"origin,omitempty"`
	ScheduledDateTime time.Time `json:"scheduledDateTime,omitempty"`
	Campaign          string    `json:"campaign,omitempty"`
	SharedPool        string    `json:"sharedPool,omitempty"`
	NotifyUrl         string    `json:"notifyUrl,omitempty"`
	IncomingUrl       string    `json:"incomingUrl,omitempty"`
	ExpiryDateTime    time.Time `json:"expiryDateTime,omitempty"`
}
type SendSMSResponse struct {
}

func SendSingle(to string, from string, title string, message string, timeout int) (*SendSMSResponse, error) {
	payload := SendSMSRequest{
		Destination: to,
		Origin:      from,
	}
	payload.Message = fmt.Sprintf("%s: %s", title, message)
	payloadByteArray, err := json.Marshal(payload)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Printf(ErrorOutputTemplate, err.Error())
		}
		return nil, err
	}

	body := bytes.NewBuffer(payloadByteArray)
	req, err := http.NewRequest("POST", GetFullPath(SmsAPI), body)
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

	retVal, err := decodeSendSMSResponse(resp)
	if err != nil {
		if err != nil {
			fmt.Printf("Error decoding response: %s\n", err.Error())
		}
		return nil, NewSmsGlobalPayloadDecodeError(err.Error())
	}

	return retVal, nil
}

func decodeSendSMSResponse(resp *http.Response) (*SendSMSResponse, error) {
	// TODO: implement decodeSendSMSResponse decoding
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Observed response: [%d] ", resp.StatusCode)
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Response Body Read Error: %s\n", err.Error())
		} else {
			fmt.Printf("%s\n", string(bodyBytes))
		}
	}

	return nil, nil
}
