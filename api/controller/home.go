package controller

import (
	"github.com/gin-gonic/gin"
	"jlb_shop_go/api/protocols"
	"jlb_shop_go/service"
)

type Home struct {
}

var HomeApi = new(Home)

func (h *Home) Data(c *gin.Context) {
	req := &protocols.HomeDataReq{}
	_ = c.BindQuery(req)
	res, err := service.Home.GetData(req)

	if err != nil {
		protocols.FailWithDetailed(nil, err.Error(), c, 200)
		return
	}

	protocols.OkWithData(res, c, 200)
}
