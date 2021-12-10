package handler

import (
	"9gag-api/contract"
	"9gag-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	authService service.AuthService
}

func NewUserHandler(authService service.AuthService) UserHandler {
	return UserHandler{
		authService: authService,
	}
}

func (u UserHandler) Signup(c *gin.Context) {
	var req contract.SignupRequest
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	res, err := u.authService.Signup(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "success": false, "message": err.Error() })
	}

	c.JSON(http.StatusOK, gin.H{ "success": true, "data": res })
}

func (u UserHandler) Signin(c *gin.Context) {
	var req contract.SigninRequest
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	res, err := u.authService.Signin(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "success": false, "message": err.Error() })
		return
	}
	c.JSON(http.StatusOK, gin.H{ "success": true, "data": res })
}
