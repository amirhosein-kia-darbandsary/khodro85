package services

import (
	"crypto/rsa"
	"errors"
	"os"
	"time"

	"github.com/amirhosein-kia-darbandsary/khodro85/config"
	"github.com/amirhosein-kia-darbandsary/khodro85/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	cfg        *config.Config
	logger     logging.Logger
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type TokenDto struct {
	UserId    string
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Roles     []string
}

func NewTokenService(cfg *config.Config) (TokenService, error) {
	logger := logging.NewLogger(cfg)

	privateKeyBytes, err := os.ReadFile(cfg.JWT.PrivateKeyPath)
	if err != nil {
		return TokenService{}, err
	}

	publicKeyBytes, err := os.ReadFile(cfg.JWT.PublicKeyPath)
	if err != nil {
		return TokenService{}, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return TokenService{}, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return TokenService{}, err
	}

	return TokenService{
		cfg:        cfg,
		logger:     logger,
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (t *TokenService) CreateToken(req TokenDto) (string, error) {

	now := time.Now()

	claims := jwt.MapClaims{
		"username":  req.UserName,
		"userid":    req.UserId,
		"firstname": req.FirstName,
		"lastname":  req.LastName,
		"email":     req.Email,
		"roles":     req.Roles,

		"iat": now.Unix(),
		"exp": now.Add(
			time.Duration(t.cfg.JWT.AccessTokenExpireMinutes) * time.Minute,
		).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		claims,
	)

	tokenString, err := token.SignedString(t.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *TokenService) ValidateToken(tokenString string) (bool, error) {

	parsedToken, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {

			if token.Method != jwt.SigningMethodRS256 {
				return nil, errors.New("unexpected signing method")
			}

			return t.publicKey, nil
		},
	)

	if err != nil {
		t.logger.Error(
			logging.JwtToken,
			logging.InvalidToken,
			err.Error(),
			nil,
		)

		return false, err
	}

	if !parsedToken.Valid {
		err := errors.New("invalid token")

		t.logger.Error(
			logging.JwtToken,
			logging.InvalidToken,
			err.Error(),
			nil,
		)
 
		return false, err
	}

	return true, nil
}
