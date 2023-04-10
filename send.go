/*
 * Project: SMSGlobal API SDK
 * Filename: /send.go
 * Created Date: Saturday March 11th 2023 19:06:05 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Monday April 10th 2023 12:59:54 +0800
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
	"strings"
	"time"

	"github.com/berrypay/runegen"
)

type SmsGlobalSendSMSRequest struct {
	Destination       string     `json:"destination,omitempty"`
	Message           string     `json:"message,omitempty"`
	Origin            string     `json:"origin,omitempty"`
	ScheduledDateTime *time.Time `json:"scheduledDateTime,omitempty"`
	Campaign          string     `json:"campaign,omitempty"`
	SharedPool        string     `json:"sharedPool,omitempty"`
	NotifyUrl         string     `json:"notifyUrl,omitempty"`
	IncomingUrl       string     `json:"incomingUrl,omitempty"`
	ExpiryDateTime    *time.Time `json:"expiryDateTime,omitempty"`
}
type SmsGlobalSendSMSResponse struct {
	Messages []SmsGlobalSendSMSResponseMessageItem `json:"messages"`
}

type SmsGlobalSendSMSResponseMessageItem struct {
	Id          int64             `json:"id"`
	OutgoingId  int64             `json:"outgoingId"`
	Origin      string            `json:"origin"`
	Destination string            `json:"destination"`
	Message     string            `json:"message"`
	Status      string            `json:"status"`
	DateTime    SmsGlobalDateTime `json:"dateTime"`
}

type SmsGlobalDateTime time.Time

const customTimeLayout = "2006-01-02 15:04:05 -0700"

func (ct *SmsGlobalDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(customTimeLayout, s)
	if err != nil {
		return err
	}
	*ct = SmsGlobalDateTime(t)
	return nil
}

func SendSingle(to string, from string, title string, message string, timeout int) (*SmsGlobalSendSMSResponseMessageItem, error) {
	payload := SmsGlobalSendSMSRequest{
		Destination: to,
		Origin:      from,
	}
	payload.Message = fmt.Sprintf("%s: %s", title, message)
	payloadByteArray, err := json.Marshal(payload)
	if err != nil {
		printDebug(ErrorOutputTemplate, err.Error())
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Send Payload: %s\n", string(payloadByteArray))
	}

	body := bytes.NewBuffer(payloadByteArray)
	req, err := http.NewRequest("POST", GetFullPath(SmsAPI), body)
	if err != nil {
		printDebug(ErrorOutputTemplate, err.Error())
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		printDebug("Sending sms by calling API endpoint at: %s\n", req.URL)
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
		printDebug(ErrorOutputTemplate, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		retErr := NewFailedCallError(resp)
		printDebug(ErrorOutputTemplate, retErr.Error())

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

func decodeSendSMSResponse(resp *http.Response) (*SmsGlobalSendSMSResponseMessageItem, error) {
	// TODO: implement decodeSendSMSResponse decoding
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		printDebug(ErrorOutputTemplate, err.Error())
		return nil, err
	}
	printDebug("Observed response: [%d] %s", resp.StatusCode, string(bodyBytes))

	// Create a bytes.Reader from the byte slice
	bodyReader := bytes.NewReader(bodyBytes)

	var sendSmsResponse SmsGlobalSendSMSResponse
	jsonDecodeErr := json.NewDecoder(bodyReader).Decode(&sendSmsResponse)
	if jsonDecodeErr != nil {
		printDebug(ErrorOutputTemplate, jsonDecodeErr.Error())
		return nil, jsonDecodeErr
	}

	// check sanity of response
	if len(sendSmsResponse.Messages) < 1 {
		return nil, NewSmsGlobalPayloadDecodeError("No result returned")
	}

	// TODO: better checking of response. Response may return more that 1 result which it should not. For now just assume taking the first result
	return &sendSmsResponse.Messages[0], nil
}
