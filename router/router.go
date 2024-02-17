package router

import (
	"github.com/gin-gonic/gin"
	"offerBook/router/api"
	"offerBook/router/middleware"
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
		private.POST("/interview", api.CreateInterview)
	}
	return r
}
