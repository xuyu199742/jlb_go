package router

import (
	"github.com/gin-gonic/gin"
	"jlb_shop_go/api/controller"
)

type UserRouter struct {
}

func (h *UserRouter) InitRouter(r *gin.RouterGroup) gin.IRouter {
	userRouter := r.Group("user-member")

	var userApi = controller.UserApi
	{
		userRouter.GET("list", userApi.List)
	}

	return userRouter
}
