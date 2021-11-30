package model

type ActivityBuyType struct {
	Id   uint   `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	Name string `gorm:"column:name;type:varchar(20);NOT NULL" json:"name"`                    // 活动类型名称
	Img  string `gorm:"column:img;type:varchar(255);NOT NULL" json:"img"`                     // 活动类型图片
}

func (m *ActivityBuyType) TableName() string {
	return "activity_buy_type"
}
