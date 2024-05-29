package util

import (
	"fmt"
	"strconv"
	"time"
)

func GetExpiryTime(expiryStr string) time.Duration {
	expiry, err := strconv.Atoi(expiryStr)
	if err != nil {
		fmt.Printf("Error converting expiry string to integer: %v\n", err)
	}
	expirationDuration := time.Minute * time.Duration(expiry)
	return expirationDuration
}
