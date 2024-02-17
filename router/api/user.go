package api

import (
	"github.com/gin-gonic/gin"
	"offerBook/logic"
	"offerBook/model"
	"offerBook/router/service"
)

func SignUp(c *gin.Context) {
	var u *model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		service.ErrorMessage(c, err, "form parse failed", true)
		return
	}
	if err := logic.NewUser(u); err != nil {
		service.ErrorMessage(c, err, "", false)
		return
	}
	service.Message(c, "create user successful")
}

func Login(c *gin.Context) {
	var u *model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		service.ErrorMessage(c, err, "form parse failed", true)
		return
	}
	token, err := logic.Login(u)
	if err != nil {
		service.ErrorMessage(c, err, "", false)
		return
	}
	c.SetCookie("jwt", token, 0, "/", "/", false, false)
	service.Message(c, "login successful")
}
