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
