package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,iranian_mobile	,min=11,max=11"`
}
