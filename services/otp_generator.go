package services

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
    src := rand.NewSource(time.Now().UnixNano())
    r := rand.New(src)
    otp := ""
    for i := 0; i < 6; i++ {
        otp += strconv.Itoa(r.Intn(10))
    }
    return otp
}
