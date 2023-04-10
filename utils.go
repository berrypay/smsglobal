/*
 * Project: SMSGlobal API SDK
 * Filename: /utils.go
 * Created Date: Monday April 10th 2023 10:17:04 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Monday April 10th 2023 10:18:15 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package smsglobal

import (
	"fmt"
	"os"
)

func printDebug(template string, data ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf(template, data...)
	}
}
