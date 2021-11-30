package model

import (
	"gorm.io/gorm"
	"jlb_shop_go/global"
	"jlb_shop_go/model/enums"
	"jlb_shop_go/utils"
	"time"
)

//RetailerHomeBanner 零售商轮播图
type RetailerHomeBanner struct {
	Id          uint             `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	Name        string           `gorm:"column:name;type:varchar(80);NOT NULL" json:"name"`                    // 轮播名称
	ImgUrl      string           `gorm:"column:img_url;type:varchar(255);NOT NULL" json:"img_url"`             // 轮播图
	Sort        int              `gorm:"column:sort;type:int(11);default:100;NOT NULL" json:"sort"`            // 排序
	BeginTime   time.Time        `gorm:"column:begin_time;type:datetime;NOT NULL" json:"begin_time"`           // 开始时间
	EndTime     time.Time        `gorm:"column:end_time;type:datetime;NOT NULL" json:"end_time"`               // 结束时间
	NavigateUrl string           `gorm:"column:navigate_url;type:varchar(100);NOT NULL" json:"navigate_url"`   // 跳转地址
	IsUse       enums.BaseStatus `gorm:"column:is_use;type:tinyint(2);default:1;NOT NULL" json:"is_use"`       // 是否可用，0：禁用，1：可用
	CreateTime  time.Time        `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`         // 创建时间
	Creator     int              `gorm:"column:creator;type:int(11);NOT NULL" json:"creator"`                  // 创建人
}

func (m *RetailerHomeBanner) TableName() string {
	return "retailer_home_banner"
}

type BannerQuery struct {
	db *gorm.DB
}

func NewBannerQuery(db ...*gorm.DB) *BannerQuery {
	if len(db) > 0 {
		return &BannerQuery{db[0]}
	}

	return &BannerQuery{global.Db}
}

func (q *BannerQuery) IsUse(use int) *BannerQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("is_use = ?", use)
	})

	return q
}

func (q *BannerQuery) BeginTime(time ...string) *BannerQuery {

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		if len(time) > 0 {
			return db.Where("begin_time <= ?", time[0])
		}

		return db.Where("begin_time <= ?", utils.TimeToSting())
	})

	return q
}

func (q *BannerQuery) EndTime(time ...string) *BannerQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		if len(time) > 0 {
			return db.Where("end_time >= ?", time[0])
		}

		return db.Where("end_time >= ?", utils.TimeToSting())
	})

	return q
}

func (q *BannerQuery) All(column ...string) ([]*RetailerHomeBanner, error) {
	banners := make([]*RetailerHomeBanner, 0)

	if len(column) > 0 {
		q.db.Select(column)
	}
	if err := q.db.Find(&banners).Order("sort DESC").Error; err != nil {
		return nil, err
	}

	return banners, nil
}
