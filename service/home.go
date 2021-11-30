package service

import (
	"errors"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
	"jlb_shop_go/api/protocols"
	"jlb_shop_go/global"
	"jlb_shop_go/model"
	"jlb_shop_go/model/enums"
)

type HomeService struct {
}

var Home = new(HomeService)

func (h *HomeService) GetData(req *protocols.HomeDataReq) (res protocols.HomeDataRes, err error) {
	var (
		banner   = protocols.BannerRes{}
		notice   = &model.RetailerNoticeInfo{}
		retailer = model.RetailerInfo{}
	)
	relations := make([]*model.RetailerStaffRelated, 0)
	banners := make([]*model.RetailerHomeBanner, 0)
	shopList := make([]map[string]interface{}, 0)

	staff, err := model.NewStaffInfoQuery().Id(global.StaffID).FindOne("id, position")
	if err != nil {
		return res, errors.New("账户信息缺失！")
	}

	//门店信息
	relations, err = model.NewRetailerStaffRelatedQuery().StaffId(staff.Id).PreloadRetailer(func(db *gorm.DB) *gorm.DB {
		return db.Select("id, shop_name")
	}).All()

	if retailer = relations[0].Retailer; len(relations) <= 0 {
		return res, errors.New("门店与该账户绑定关系异常！")
	}

	for _, relation := range relations {
		shopList = append(shopList, map[string]interface{}{
			"label": relation.Retailer.ShopName,
			"value": relation.Retailer.Id,
		})
		//切换当前门店
		if req.Params.RetailerId > 0 {
			if relation.RetailerId == req.Params.RetailerId {
				retailer = relation.Retailer
			}
		}
	}
	res.ShopId = retailer.Id
	res.ShopName = retailer.ShopName
	res.ShopList = shopList

	//最新公告
	notice, err = model.NewNoticeQuery().First("title")
	if notice.Title != "" {
		res.Notice.Title = notice.Title
	}

	// banner 列表
	banners, err = model.NewBannerQuery().BeginTime().EndTime().IsUse(enums.StatusTrue).All("img_url", "navigate_url")
	for _, v := range banners {
		banner.NavigateUrl = v.NavigateUrl
		banner.ImgUrl = global.Config.Local.ImgUrl + "/" + v.ImgUrl
		res.Banner = append(res.Banner, banner)
	}

	// 本月新增会员
	monthRange := []string{carbon.Now().StartOfMonth().ToDateTimeString(), carbon.Now().EndOfMonth().ToDateTimeString()}
	userQuery := model.NewUserMemberQuery().BtCreateTime(monthRange).Status(1)
	if staff.Position == 2 {
		userQuery = userQuery.FirstRegShop(retailer.Id).FirstRegStaff(global.StaffID)
	} else {
		userQuery = userQuery.FirstRegShop(retailer.Id)
	}
	_, res.NewMemberNum = userQuery.Count()

	//活动参与人数
	userJoinQuery := model.NewActivityUserJoinDataQuery().BtJoinTime(monthRange)
	if staff.Position == 2 {
		_, res.AddActivityNum = userJoinQuery.RetailerId(retailer.Id).StaffId(global.StaffID).Count()
	} else {
		_, res.AddActivityNum = userJoinQuery.RetailerId(retailer.Id).Count()
	}

	//活动完成人数
	_, res.DoneActivityNum = userJoinQuery.IsFinish(1).Count()

	if err != nil {
		return res, err
	}

	return res, err
}
