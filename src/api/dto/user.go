package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,iranian_mobile	,min=11,max=11"`
}

type TokenResponse struct {
	AccessToken            string `json:"access_token"`
	RefreshTokn            string `json:"refresh_token"`
	AccessTokenExpireTime  int    `json:"accesstokenexpiretime"`
	RefreshTokenExpireTime int    `json:"refreshtokenexpiretime"`
}

type RegisterRequestByUsername struct {
	FirstName string `json:"firstname" binding:"required,min=4,max=120"`
	LastName  string `json:"lastname" binding:"min=10,max=180"`
	UserName  string `json:"username" binding:"required,min=4,max=120"`
	Email     string `json:"email" binding:"min=6,email"`
	Password  string `json:"password" binding:"required,password,min=6"`
}

type LoginByMobileNumberRequest struct {
	MobileNumber string `json:"mobilenumber" binding:"required,iranian_mobile,min=11,max=11"`
	Otp          string `json:"otp" binding:"required,min=6,max=6"`
}

type LoginByUserNameRequest struct {
	UserName string `json:"mobilenumber" binding:"required,min=4,max=120"`
	Password string `json:"password" binding:"required,password,min=6"`
}
