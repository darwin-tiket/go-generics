package main

// START OMIT
func IsOrderTypeIn(orderTypes []string, expectedOrderType string) bool {
	for _, orderType := range orderTypes {
		if orderType == expectedOrderType {
			return true
		}
	}
	return false
}

func IsStatusCodeIn(statusCodes []int, expectedStatusCode int) bool {
	for _, code := range statusCodes {
		if code == expectedStatusCode {
			return true
		}
	}
	return false
}

// END OMIT
