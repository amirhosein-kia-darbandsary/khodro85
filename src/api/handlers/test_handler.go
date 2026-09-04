package handlers

import (
	"net/http"

	"github.com/amirhosein-kia-darbandsary/khodro85/api/base"
	"github.com/gin-gonic/gin"
)

type Test struct{}

type User struct {
	FirstName       string `json:"first_name" binding:"required,alpha,min=2,max=80"`
	LastName        string `json:"last_name" binding:"required,alpha,min=2,max=80"`
	EnteredPassword string `json:"entered_password" binding:"password,min=8,max=64"`
	MobileNumber    string `json:"mobile_number" binding:"iranian_mobile,min=11,max=11"`
}

func NewTest() *Test {
	return &Test{}
}

func (t *Test) TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, base.GenerateBaseResponse("handler is good", true, 200))
}

func (t *Test) TestBindingHandler(ctx *gin.Context) {
	p := User{}
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.GenerateBaseResponseWithValidationError(nil, false, 0, err))
	}

	ctx.JSON(http.StatusOK, base.GenerateBaseResponse(p, true, 200))

}
