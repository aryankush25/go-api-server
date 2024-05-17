package services

import (
	"math/rand"
	"time"
)

func GenerateOTP() string {
    rand.Seed(time.Now().UnixNano())
    otp := ""
    for i := 0; i < 6; i++ {
        otp += string(rand.Intn(10) + '0')
    }
    return otp
}
