package handlers

import (
	"net/http"

	"github.com/amirhosein-kia-darbandsary/khodro85/api/base"
	_ "github.com/amirhosein-kia-darbandsary/khodro85/docs"
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

// TestHandler godoc
// @Summary      Test handler
// @Description  Test API endpoint
// @Tags         test
// @Produce      json
// @Success      200 {object} interface{}
// @Router       /test/ [get]
func (t *Test) TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, base.GenerateBaseResponse("handler is good", true, 200))
}

// TestBindingHandler godoc
// @Summary      Test JSON binding
// @Description  Test JSON request body binding
// @Tags         test
// @Accept       json
// @Produce      json
// @Param        user body User true "User data"
// @Success      200 {object} User
// @Failure      400 {object} interface{}
// @Router       /test/ [post]
func (t *Test) TestBindingHandler(ctx *gin.Context) {
	p := User{}

	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			base.GenerateBaseResponseWithValidationError(
				nil,
				false,
				0,
				err,
			),
		)
	}

	ctx.JSON(
		http.StatusOK,
		base.GenerateBaseResponse(p, true, 200),
	)
}
