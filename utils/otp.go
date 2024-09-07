package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

func VerifyOTP(phoneNumber, otp string) bool {
	// In a real application, you would check the OTP against a stored value
	// This is a simplified example
	return otp == "1234" // Replace with actual verification logic
}