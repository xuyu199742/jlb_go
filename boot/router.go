package boot

import (
	"jlb_shop_go/global"
	"jlb_shop_go/middleware"
	"jlb_shop_go/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	//Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/static", "./dist/static")   // dist里面的静态资源
	//Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	ApiGroup := router.ApiGroup
	PublicGroup := Router.Group("shopapi")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		ApiGroup.StaffRouter.InitRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("shopapi")
	PrivateGroup.Use(middleware.CheckAuth())
	{
		ApiGroup.HomeRouter.InitRouter(PrivateGroup)
		ApiGroup.NoticeRouter.InitRouter(PrivateGroup)
		ApiGroup.UserRouter.InitRouter(PrivateGroup)
	}

	return Router
}
