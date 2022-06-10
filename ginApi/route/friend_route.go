package route

import "github.com/gin-gonic/gin"

func Friend(engine *gin.Engine) {
	friend := engine.Group("/friend")
	{
		friend.POST("/friadd")
	}
}
