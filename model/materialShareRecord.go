package model

import "time"

// MaterialShareRecord 素材分享转发记录
type MaterialShareRecord struct {
	Id         uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	MaterialId int       `gorm:"column:material_id;type:int(11);NOT NULL" json:"material_id"`          // 素材编号
	RetailerId int       `gorm:"column:retailer_id;type:int(11);NOT NULL" json:"retailer_id"`          // 零售商编号
	IdType     int       `gorm:"column:id_type;type:tinyint(4);NOT NULL" json:"id_type"`               // 转发者身份，1：店主，2：店员
	Operator   int       `gorm:"column:operator;type:int(11);NOT NULL" json:"operator"`                // 转发者编号
	RecordTime time.Time `gorm:"column:record_time;type:datetime;NOT NULL" json:"record_time"`         // 转发记录时间
	RecordDate time.Time `gorm:"column:record_date;type:date;NOT NULL" json:"record_date"`             // 转发记录日期
}

func (m *MaterialShareRecord) TableName() string {
	return "material_share_record"
}
