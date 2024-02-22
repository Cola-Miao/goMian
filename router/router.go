package router

import (
	"github.com/gin-gonic/gin"
	"goMian/router/api"
	"goMian/router/middleware"
)

func Init() *gin.Engine {
	r := gin.Default()
	public := r.Group("/")
	{
		public.GET("/health", api.Health)
		public.POST("/signup", api.SignUp)
		public.POST("/login", api.Login)
	}
	private := r.Group("/", middleware.Auth())
	{
		private.GET("/private", api.Private)
		private.GET("/interviewDel", api.DeleteInterview)
		private.GET("/interviewList", api.InterviewList)
		private.GET("/refresh_interview", api.RefreshInterview)
		private.POST("/interview", api.CreateInterview)
		private.POST("/interviewDetail", api.AddInterviewDetail)
	}
	return r
}
