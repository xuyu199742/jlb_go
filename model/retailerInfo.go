package model

import (
	"gorm.io/gorm"
	"jlb_shop_go/global"
	"time"
)

// RetailerInfo 零售商信息表
type RetailerInfo struct {
	Id         int       `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	ShopName   string    `gorm:"column:shop_name;type:varchar(255);NOT NULL" json:"shop_name"`         // 门店名称
	ParentId   int       `gorm:"column:parent_id;type:int(11);default:0;NOT NULL" json:"parent_id"`    // 父级门店编号
	Contacts   string    `gorm:"column:contacts;type:varchar(20);NOT NULL" json:"contacts"`            // 联系人
	Phone      string    `gorm:"column:phone;type:varchar(20);NOT NULL" json:"phone"`                  // 联系人手机号
	Longitude  string    `gorm:"column:longitude;type:varchar(255)" json:"longitude"`                  // 经度
	Latitude   string    `gorm:"column:latitude;type:varchar(255)" json:"latitude"`                    // 纬度
	ShopCode   string    `gorm:"column:shop_code;type:varchar(100);NOT NULL" json:"shop_code"`         // 精准营销标识
	Province   int       `gorm:"column:province;type:int(11);default:0" json:"province"`               // 省份
	City       int       `gorm:"column:city;type:int(11);default:0" json:"city"`                       // 城市
	District   int       `gorm:"column:district;type:int(11);default:0" json:"district"`               // 区县
	Address    string    `gorm:"column:address;type:varchar(255)" json:"address"`                      // 详细地址
	DoorPhoto  string    `gorm:"column:door_photo;type:varchar(255)" json:"door_photo"`                // 门头照图片
	LicenseNo  string    `gorm:"column:license_no;type:varchar(50)" json:"license_no"`                 // 营业执照号
	LicenseImg string    `gorm:"column:license_img;type:varchar(255)" json:"license_img"`              // 营业执照图片
	IsUse      int       `gorm:"column:is_use;type:tinyint(4);default:1;NOT NULL" json:"is_use"`       // 是否可用，0：禁用，1：可用
	RegTime    time.Time `gorm:"column:reg_time;type:datetime;NOT NULL" json:"reg_time"`               // 注册时间
	Status     int       `gorm:"column:status;type:tinyint(2);default:1;NOT NULL" json:"status"`       // 精准营销状态，0：无效，1：有效
}

func (m *RetailerInfo) TableName() string {
	return "retailer_info"
}

type RetailerInfoQuery struct {
	db *gorm.DB
}

func NewRetailerInfoQuery(db ...*gorm.DB) *RetailerInfoQuery {
	if len(db) > 0 {
		return &RetailerInfoQuery{db[0]}
	}

	return &RetailerInfoQuery{global.Db}
}

func (q *RetailerInfoQuery) Id(id ...int) *RetailerInfoQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	})

	return q
}

func (q *RetailerInfoQuery) FindOne(column ...string) (*RetailerInfo, error) {
	if len(column) > 0 {
		q.db.Select(column)
	}
	retailer := RetailerInfo{}
	err := q.db.First(&retailer).Error
	if err != nil {
		return nil, err
	}

	return &retailer, nil
}
