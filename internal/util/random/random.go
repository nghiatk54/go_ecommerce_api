package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigitOtp() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := random.Intn(900000) + 100000
	return otp
}
