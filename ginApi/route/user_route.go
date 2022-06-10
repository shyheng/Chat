package route

import "github.com/gin-gonic/gin"
import "ginApi/controller"

func User(engine *gin.Engine) {
	login := engine.Group("/")
	{
		login.POST("/login", controller.VueUser{}.LoginUser)
		login.POST("/add", controller.VueUser{}.AddUser)
	}
	fri := engine.Group("/fri")
	{
		fri.POST("/sel", controller.VueUser{}.SelFriend)
		fri.POST("/add", controller.VueUser{}.AddFriend)
	}
}
