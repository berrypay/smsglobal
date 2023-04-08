/*
 * Project: SMSGlobal API SDK
 * Filename: /path.go
 * Created Date: Friday April 7th 2023 09:26:52 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Friday April 7th 2023 14:52:39 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

const (
	SmsAPI               string = "/v2/sms"
	UserBasePath         string = "/v2/user"
	UserCreditBalanceAPI string = UserBasePath + "/credit-balance"
)

func GetFullPath(api string) string {
	return GetBaseUrl() + api
}
