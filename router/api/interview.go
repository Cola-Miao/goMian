package api

import (
	"github.com/gin-gonic/gin"
	"offerBook/logic"
	"offerBook/model"
	"offerBook/router/service"
)

func CreateInterview(c *gin.Context) {
	var it model.Interview
	if err := c.ShouldBindJSON(&it); err != nil {
		service.ErrorMessage(c, err, "parse form failed", true)
		return
	}
	it.Owner = c.GetInt("id")
	if err := logic.CreateInterview(&it); err != nil {
		service.ErrorMessage(c, err, "create interview failed", true)
	}
	service.Message(c, "create interview success")
}
