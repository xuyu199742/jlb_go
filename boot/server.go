package boot

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/xinzf/gokit/logger"
	"jlb_shop_go/global"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	//if global.Config.System.UseMultipoint {
	// 初始化redis服务
	Redis()
	//}

	// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}

	Router := Routers()

	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	//global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	logger.DefaultLogger.Error("ListenAndServe", s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
