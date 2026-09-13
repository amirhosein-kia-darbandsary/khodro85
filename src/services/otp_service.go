package services

import (
	"fmt"

	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/constants"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/cache"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/errors"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/security"
	"github.com/redis/go-redis/v9"
)

type OtpUsecase struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type otpDto struct {
	Value string
	Used  bool
}

func NewOtpUsecase(cfg *config.Config) *OtpUsecase {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedsi()
	return &OtpUsecase{logger: logger, cfg: cfg, redisClient: redis}
}

func (u *OtpUsecase) SendOtp(mobileNumber string) error {
	otp := security.GenerateOtp()
	err := u.SetOtp(mobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

func (u *OtpUsecase) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &otpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[otpDto](u.redisClient, key)
	if err == nil && !res.Used {
		return &errors.ServiceError{EndUserMessage: errors.OptExists}
	} else if err == nil && res.Used {
		return &errors.ServiceError{EndUserMessage: errors.OtpUsed}
	}
	err = cache.Set(u.redisClient, key, val, 120)
	if err != nil {
		return err
	}
	return nil
}

func (u *OtpUsecase) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[otpDto](u.redisClient, key)
	if err != nil {
		return err
	} else if res.Used {
		return &errors.ServiceError{EndUserMessage: errors.OtpUsed}
	} else if !res.Used && res.Value != otp {
		return &errors.ServiceError{EndUserMessage: errors.OtpNotValid}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(u.redisClient, key, res, 120)
		if err != nil {
			return err
		}
	}
	return nil
}
