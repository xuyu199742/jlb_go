package middleware

//
//import (
//	"github.com/gin-gonic/gin"
//	"jlb_shop_go/utils"
//)
//
//func JWTAuth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
//		token := c.Request.Header.Get("x-token")
//		if token == "" {
//			//response.FailWithMessage("未登录或者非法访问", c, http.StatusUnauthorized)
//			c.Abort()
//			return
//		}
//		j := utils.NewJwt()
//		claims, err := j.ParseToken(token)
//		if err != nil {
//			//response.FailWithMessage(err.Error(), c, http.StatusUnauthorized)
//			c.Abort()
//			return
//		}
//		c.Set("claims", claims)
//		c.Next()
//	}
//}
