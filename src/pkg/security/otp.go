package security

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOtp() string {
	rand.Seed(time.Now().Unix())
	var min = 1000
	var max = 10000
	RandomIntegerwithinRange := rand.Intn(max-min) + min

	return strconv.Itoa(RandomIntegerwithinRange)
}
