package services

import (
	"github.com/amirhosein-kia-darbandsary/khodro85/api/dto"
	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/constants"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/database"
	"github.com/amirhosein-kia-darbandsary/khodro85/data/models"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/errors"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/security"
	"gorm.io/gorm"
)

type UserService struct {
	cfg          *config.Config
	logger       logging.Logger
	optService   OtpUsecase
	tokenService TokenService
	database     *gorm.DB
}

func NewUserService(cfg *config.Config) UserService {
	return UserService{
		cfg:          cfg,
		logger:       logging.NewLogger(cfg),
		optService:   *NewOtpUsecase(cfg),
		tokenService: NewTokenService(cfg),
		database:     database.GetPostgresConnection(),
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

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool

	err := s.database.
		Model(&models.User{}).
		Select("1").
		Where("email = ?", email).
		Limit(1).
		Scan(&exists).Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool

	err := s.database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", username).
		Limit(1).
		Scan(&exists).Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByMobilenumber(mobileNumber string) (bool, error) {
	var exists bool

	err := s.database.
		Model(&models.User{}).
		Select("1").
		Where("mobilenumber = ?", mobileNumber).
		Limit(1).
		Scan(&exists).Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {
	var roleIdres int
	err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleIdres).Error
	if err != nil {
		s.logger.Error(logging.Postgres,
			logging.Select,
			err.Error(), nil)
		return 0, err
	}
	return roleIdres, nil

}

func (s *UserService) RegisterUser(req dto.RegisterRequestByUsername) error {

	exists, err := s.existsByUsername(req.UserName)
	if err != nil {

		return err
	}

	if exists {
		return &errors.ServiceError{
			EndUserMessage: "username Exists",
		}
	}

	hashedpassword, err := security.GeneratePasswordHash(req.Password)

	if err != nil {
		s.logger.Error(logging.Validation, logging.HashPassword, err.Error(), nil)
		return err

	}
	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	user := models.User{Username: req.UserName,
		LastName:  req.LastName,
		Email:     req.Email,
		FirstName: req.FirstName,
		Password:  hashedpassword,
	}

	tx := s.database.Begin()
	defer tx.Rollback()
	err = tx.Create(&user).Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}

	err = tx.Create(&models.UserRole{RoleID: uint(roleId), UserID: uint(user.ID)}).Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err

	}
	tx.Commit()
	return nil

}

func (s *UserService) RegisterLoginByMobileNumber(
	req dto.LoginByMobileNumberRequest,
) (res string, err error) {

	// 1. Validate OTP
	if err = s.optService.ValidateOtp(req.MobileNumber, req.Otp); err != nil {
		return "", err
	}

	// 2. Check whether user exists
	exists, err := s.existsByMobilenumber(req.MobileNumber)
	if err != nil {
		return "", err
	}

	// 3. Existing user -> Login
	if exists {
		var user models.User

		err = s.database.
			Where("mobile_number = ?", req.MobileNumber).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			First(&user).Error

		if err != nil {
			return "", err
		}

		tokenDTO := TokenDto{
			UserId:    user.ID,
			FirstName: user.FirstName,
			UserName:  user.Username,
			Email:     user.Email,
			LastName:  user.LastName,
		}

		if user.UserRoles != nil {
			for _, userRole := range *user.UserRoles {
				tokenDTO.Roles = append(
					tokenDTO.Roles,
					userRole.Role.Name,
				)

			}
		}

		return s.tokenService.CreateToken(tokenDTO)
	}

	// 4. New user -> Register
	registerReq := dto.RegisterRequestByUsername{
		UserName: req.MobileNumber,
		Password: req.MobileNumber,
	}

	err = s.RegisterUser(registerReq)
	if err != nil {
		return "", err
	}

	// 5. Get the newly created user
	var user models.User

	err = s.database.
		Where("mobile_number = ?", req.MobileNumber).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		First(&user).Error

	if err != nil {
		return "", err
	}

	// 6. Create token
	tokenDTO := TokenDto{
		UserId:    user.ID,
		FirstName: user.FirstName,
		UserName:  user.Username,
		Email:     user.Email,
		LastName:  user.LastName,
	}

	if user.UserRoles != nil {
		for _, userRole := range *user.UserRoles {
			tokenDTO.Roles = append(
				tokenDTO.Roles,
				userRole.Role.Name,
			)

		}
	}

	return s.tokenService.CreateToken(tokenDTO)
}
