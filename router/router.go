package router

import (
	"github.com/gin-gonic/gin"
)

var R *gin.Engine 

func InitRouter(){
	R = gin.Default()

	R.GET("/", help)
	user := R.Group("/user")
	{
		user.POST("/add"	, add)
		user.POST("/search"	, search)
		user.POST("/modify"	, modify)
		user.POST("/delete"	, delete)
	}
}