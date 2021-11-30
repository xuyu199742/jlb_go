package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"jlb_shop_go/api/protocols"
	"jlb_shop_go/global"
)

type Request struct {
	AccessToken string `json:"access_token" form:"access_token"`
}

func CheckAuth() gin.HandlerFunc {
	//todo 单账号登录
	return func(c *gin.Context) {
		req := Request{}
		_ = c.BindQuery(&req)
		if req.AccessToken == "" {
			protocols.FailWithMessage(20004, "无效的access token", c, 200)
			c.Abort()
			return
		}

		staffId, _ := global.Redis.Get(global.Ctx, req.AccessToken).Result()
		if staffId == "" {
			protocols.FailWithMessage(20004, "未登录或者非法访问", c, 200)
			c.Abort()
			return
		}
		global.StaffID = cast.ToInt(staffId)

		c.Next()
	}
}
