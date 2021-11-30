package model

import "time"

// ActivityBuyRetailerRelated 活动关联门店表
type ActivityBuyRetailerRelated struct {
	Id         uint      `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	ActivityId uint      `gorm:"column:activity_id;type:int(10) unsigned;NOT NULL" json:"activity_id"` // 活动编号
	RetailerId uint      `gorm:"column:retailer_id;type:int(10) unsigned;NOT NULL" json:"retailer_id"` // 门店编号
	Creator    int       `gorm:"column:creator;type:int(11);NOT NULL" json:"creator"`                  // 创建人编号
	CreateTime time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`         // 创建时间
}

func (m *ActivityBuyRetailerRelated) TableName() string {
	return "activity_buy_retailer_related"
}
