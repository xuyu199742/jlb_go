package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"jlb_shop_go/api/protocols"
	"jlb_shop_go/global"
	"jlb_shop_go/model"
)

type User struct {
}

var UserApi = new(User)

func (u *User) List(c *gin.Context) {
	req := &protocols.UserListReq{}
	res := &protocols.UserListRes{}

	_ = c.BindQuery(req)
	retailer, err := model.NewStaffInfoQuery().Id(cast.ToInt(req.Params.RetailerId)).FindOne()
	if err != nil || retailer.Status == 0 || retailer.IsUse == 0 {
		protocols.FailWithDetailed(nil, "门店信息异常", c, 500)
		return
	}

	staff, err := model.NewStaffInfoQuery().Id(global.StaffID).FindOne()
	if err != nil || staff.Status == 0 || staff.IsUse == 0 {
		protocols.FailWithDetailed(nil, "账号信息异常", c, 500)
		return
	}

	userQuery := model.NewUserMemberQuery().Status(1).FirstRegShop(retailer.Id)
	if staff.Position == 2 {
		userQuery = userQuery.FirstRegStaff(staff.Id)
	}

	if cast.ToInt(req.Params.Stage) > 0 {
		userQuery = userQuery.MemStage(cast.ToInt(req.Params.Stage))
	}

	if req.Params.BeginTime != "" && req.Params.EndTime != "" {
		dateRange := []string{req.Params.BeginTime, req.Params.EndTime}
		userQuery = userQuery.BtCreateTime(dateRange)
	}

	users, total, err := userQuery.PageList(req.Params.Page, 20)
	if err != nil {
		protocols.FailWithDetailed(nil, err.Error(), c, 500)
		return
	}
	result := map[string]interface{}{}
	for _, v := range users {
		result["memStageCode"] = v.MemStage
		result["memStage"] = v.MemStage
		result["comeFrom"] = v.ComeFrom
		res.Result = append(res.Result, result)
	}

	res.Total = total

	protocols.OkWithData(res, c, 200)
}
