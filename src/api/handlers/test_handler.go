package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct{}

type User struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=2,max=80"`
	LastName     string `json:"last_name" binding:"required,alpha,min=2,max=80"`
	MobileNumber string `json:"mobile_number" binding:"iranian_mobile",min=11, max=11`
}

func NewTest() *Test {
	return &Test{}
}

func (t *Test) TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "mamaa is good",
	})
}

func (t *Test) TestBindingHandler(ctx *gin.Context) {
	p := User{}
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": p,
	})

}
