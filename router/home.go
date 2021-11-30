package router

import (
	"github.com/gin-gonic/gin"
	"jlb_shop_go/api/controller"
)

type HomeRouter struct {
}

func (h *HomeRouter) InitRouter(r *gin.RouterGroup) gin.IRouter {
	homeRouter := r.Group("home")

	var homeApi = controller.HomeApi
	{
		homeRouter.GET("data", homeApi.Data)
	}

	return homeRouter
}
