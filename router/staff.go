package router

import (
	"github.com/gin-gonic/gin"
	"jlb_shop_go/api/controller"
)

type StaffRouter struct {
}

func (h *StaffRouter) InitRouter(r *gin.RouterGroup) gin.IRouter {
	staffRouter := r.Group("user-staff-info")

	var StaffApi = controller.StaffApi
	{
		staffRouter.POST("login", StaffApi.Login)
		staffRouter.GET("query-user-profile", StaffApi.GetProfile)
	}

	return staffRouter
}
