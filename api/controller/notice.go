package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"jlb_shop_go/api/protocols"
	"jlb_shop_go/model"
	"strings"
)

type Notice struct {
}

var NoticeApi = new(Notice)

func (n *Notice) List(c *gin.Context) {
	var (
		req    = protocols.NoticeListReq{}
		res    = protocols.NoticeListRes{}
		notice = protocols.Notice{}
	)
	err := c.ShouldBind(&req)
	if err != nil {
		protocols.FailWithDetailed(nil, "请求参数异常", c, 200)
		return
	}

	items, total, err := model.NewNoticeQuery().ListPage(req.Params.Page, 20, "id,title,create_time")
	if err != nil {
		protocols.FailWithDetailed(nil, err.Error(), c, 200)
		return
	}

	for _, v := range items {
		notice.CreateTime = v.CreateTime.String()
		notice.Id = v.Id
		notice.Title = v.Title
		res.NoticeList = append(res.NoticeList, notice)
	}
	res.Total = total
	res.Page = req.Params.Page

	protocols.OkWithData(res, c, 200)
}

func (n *Notice) Detail(c *gin.Context) {
	var (
		req = protocols.NoticeDetailReq{}
		res = protocols.NoticeDetailRes{}
	)
	_ = c.BindQuery(&req)

	notice, err := model.NewNoticeQuery().Id(cast.ToInt(req.Params.NoticeId)).First()
	if err != nil {
		protocols.FailWithDetailed(nil, "该公告不存在", c, 200)
		return
	}

	res.CreateTime = notice.CreateTime.String()
	res.Content = strings.Replace(notice.Content, " /<img/", "< img style=\"max-width:100%;height:auto\" ", -1)
	res.Title = notice.Title

	protocols.OkWithData(res, c, 200)
}
