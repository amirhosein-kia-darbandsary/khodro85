package services

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/dto"
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/database"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/security"
	"gorm.io/gorm"
)

type UserService struct {
	cfg        *config.Config
	logger     logging.Logger
	optService OtpUsecase
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) UserService {
	return UserService{
		cfg:        cfg,
		logger:     logging.NewLogger(cfg),
		optService: *NewOtpUsecase(cfg),
		database:   database.GetPostgresConnection(),
	}
}

func (s *UserService) SendOtp(req dto.GetOtpRequest) error {
	otp := security.GenerateOtp()
	err := s.optService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
