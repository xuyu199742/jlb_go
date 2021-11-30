package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"jlb_shop_go/api/protocols"
	"jlb_shop_go/global"
	"jlb_shop_go/model"
	"jlb_shop_go/service"
)

type Staff struct {
}

var StaffApi = new(Staff)

func (s *Staff) Login(c *gin.Context) {
	var m map[string]string
	req := &protocols.StaffLoginReq{}

	b, _ := c.GetRawData()
	_ = json.Unmarshal(b, &m)
	if _, ok := m["params"]; !ok {
		protocols.FailWithDetailed(nil, "缺少请求参数", c, 200)
		return
	}

	if err := json.Unmarshal([]byte(m["params"]), req); err != nil {
		protocols.FailWithDetailed(nil, err.Error(), c, 200)
		return
	}

	res, err := service.Staff.Login(req)
	if err != nil {
		protocols.FailWithDetailed(nil, err.Error(), c, 200)
		return
	}

	protocols.OkWithData(res, c, 200)
}

func (s *Staff) GetProfile(c *gin.Context) {
	req := &protocols.StaffProfileReq{}
	res := protocols.StaffProfileRes{}
	retailer := model.RetailerInfo{}
	shopList := make([]map[string]interface{}, 0)
	_ = c.BindQuery(req)

	staff, err := model.NewStaffInfoQuery().Id(global.StaffID).FindOne("id, staff_code")
	if err != nil {
		protocols.FailWithDetailed(nil, "账户信息缺失", c, 200)
		return
	}

	//门店信息
	relations, err := model.NewRetailerStaffRelatedQuery().StaffId(staff.Id).PreloadRetailer(func(db *gorm.DB) *gorm.DB {
		return db.Select("id, shop_name")
	}).All()

	if retailer = relations[0].Retailer; len(relations) <= 0 {
		protocols.FailWithDetailed(nil, "门店与该账户绑定关系异常！", c, 200)
		return
	}
	for _, relation := range relations {
		shopList = append(shopList, map[string]interface{}{
			"shop_name":   relation.Retailer.ShopName,
			"retailer_id": relation.Retailer.Id,
		})
		//切换当前门店
		if req.Params.RetailerId > 0 {
			if relation.RetailerId == req.Params.RetailerId {
				retailer = relation.Retailer
			}
		}
	}
	res.Profile.Id = global.StaffID
	res.Profile.StaffCode = staff.StaffCode
	res.Profile.SelectShopId = cast.ToString(retailer.Id)
	res.Profile.SelectShopName = retailer.ShopName
	res.Shops = shopList

	protocols.OkWithData(res, c, 200)
}
