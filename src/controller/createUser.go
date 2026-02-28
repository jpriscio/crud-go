package controller

import (
	resterr "crud-go/src/configuration/rest_err"
	"crud-go/src/controller/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {

	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestErr(fmt.Sprintf("There are some incorrects fields: %s", err.Error()))

		ctx.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}
