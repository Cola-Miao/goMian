package api

import (
	"github.com/gin-gonic/gin"
	"goMian/logic"
	"goMian/model"
	"goMian/router/service"
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
	service.Message(c, "create interview successful")
}

func RefreshInterview(c *gin.Context) {
	id := c.GetInt("id")
	if err := logic.RefreshInterview(id); err != nil {
		service.ErrorMessage(c, err, "refresh interview failed", true)
		return
	}
	service.Message(c, "refresh interview successful")
}

func InterviewList(c *gin.Context) {
	id := c.GetInt("id")
	its, err := logic.InterviewList(id)
	if err != nil {
		service.ErrorMessage(c, err, "get interview list failed", true)
		return
	}
	service.Message(c, "get interview list successful", its)
}

func DeleteInterview(c *gin.Context) {
	id := c.GetInt("id")
	itID := c.Query("id")
	if err := logic.DeleteInterview(id, itID); err != nil {
		service.ErrorMessage(c, err, "delete interview failed", true)
		return
	}
	service.Message(c, "delete interview successful")
}

func AddInterviewDetail(c *gin.Context) {
	id := c.GetInt("id")
	relevance := c.Query("rel")
	var detail model.InterviewDetail
	if err := c.ShouldBindJSON(&detail); err != nil {
		service.ErrorMessage(c, err, "parse form failed", true)
		return
	}
	if err := logic.AddInterviewDetail(id, relevance, &detail); err != nil {
		service.ErrorMessage(c, err, "add interview failed", true)
		return
	}
	service.Message(c, "add details successful")
}
