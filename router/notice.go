package router

import (
	"github.com/gin-gonic/gin"
	"jlb_shop_go/api/controller"
)

type NoticeRouter struct {
}

func (h *NoticeRouter) InitRouter(r *gin.RouterGroup) gin.IRouter {
	noticeRouter := r.Group("notice")

	var noticeApi = controller.NoticeApi
	{
		noticeRouter.GET("list", noticeApi.List)
		noticeRouter.GET("view", noticeApi.Detail)
	}

	return noticeRouter
}
