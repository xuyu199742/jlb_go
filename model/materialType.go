package model

import "time"

// MaterialType 素材类型
type MaterialType struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	TypeName   string    `gorm:"column:type_name;type:varchar(255);NOT NULL" json:"type_name"` // 素材类型名称
	Sort       int       `gorm:"column:sort;type:int(11);default:100" json:"sort"`             // 排序
	Status     int       `gorm:"column:status;type:tinyint(2);NOT NULL" json:"status"`         // 状态
	Creator    string    `gorm:"column:creator;type:varchar(255);NOT NULL" json:"creator"`     // 创建人
	CreateTime time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"` // 创建时间
}

func (m *MaterialType) TableName() string {
	return "material_type"
}
