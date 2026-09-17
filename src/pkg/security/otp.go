package security

import (
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateOtp() string {
	rand.Seed(time.Now().Unix())
	var min = 1000
	var max = 10000
	RandomIntegerwithinRange := rand.Intn(max-min) + min

	return strconv.Itoa(RandomIntegerwithinRange)
}

func GeneratePasswordHash(password string) (hashedpassword string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
