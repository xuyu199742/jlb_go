package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/xinzf/gokit/logger"
)

//GetUserId 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		logger.DefaultLogger.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.ID
	}
}

//GetUserInfo  从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		logger.DefaultLogger.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}
