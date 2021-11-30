package service

import (
	"errors"
	"jlb_shop_go/api/protocols"
	"jlb_shop_go/global"
	"jlb_shop_go/model"
	"jlb_shop_go/package/wechat/wxMini"
	"jlb_shop_go/utils"
	"strconv"
	"time"
)

type StaffService struct {
}

var Staff = new(StaffService)

func (s *StaffService) Login(req *protocols.StaffLoginReq) (*protocols.StaffLoginRes, error) {
	res := &protocols.StaffLoginRes{}
	//todo params valid
	staff, err := model.NewStaffInfoQuery().Account(req.LoginAccount).FindOne()
	if err != nil || staff == nil {
		return res, errors.New("账户信息不存在，请稍后再试！")
	}

	if staff.Status != 1 || staff.IsUse != 1 {
		return res, errors.New("账户状态异常，请联系客服！")
	}

	wxServer := wxMini.NewWxMiniServer(global.Config.Wx.SfaAppid, global.Config.Wx.SfaSecret)
	sessionRes, err := wxServer.GetCode2Session(req.WxCode)
	if err != nil {
		return res, err
	}
	if sessionRes.Openid == "" || sessionRes.UnionId == "" {
		return res, errors.New("系统异常，微信openid获取失败！")
	}

	userMobile, err := wxMini.MobileDecrypt(req.EncryptedData, sessionRes.SessionKey, req.IV)
	staff.UnionId = sessionRes.UnionId
	staff.OpenId = sessionRes.Openid
	staff.Phone = userMobile.PurePhoneNumber
	err = model.NewStaffInfoExecutor().Save(staff)

	key := global.SfaLoginKey + utils.RandStr(30)
	err = global.Redis.Set(global.Ctx, key, strconv.Itoa(int(staff.Id)), time.Duration(global.SfaLoginExpired)*time.Second).Err()
	res.AccessToken = key
	if err != nil {
		return res, err
	}

	return res, nil
}
