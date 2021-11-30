package model

// ActivityRetailerJoinData 门店活动参数数据
type ActivityRetailerJoinData struct {
	Id           uint `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	ActivityId   int  `gorm:"column:activity_id;type:int(11);NOT NULL" json:"activity_id"`          // 活动编号
	ActivityType int  `gorm:"column:activity_type;type:tinyint(4);NOT NULL" json:"activity_type"`   // 活动类型，1：买一赠一，2：蒲公英计划
	RetailerId   int  `gorm:"column:retailer_id;type:int(11);NOT NULL" json:"retailer_id"`          // 门店编号
	StaffId      int  `gorm:"column:staff_id;type:int(11);NOT NULL" json:"staff_id"`                // 员工编号
	JoinNum      int  `gorm:"column:join_num;type:int(11);NOT NULL" json:"join_num"`                // 参与人数
	FinishNum    int  `gorm:"column:finish_num;type:int(11);default:0;NOT NULL" json:"finish_num"`  // 完成人数
	RegNum       int  `gorm:"column:reg_num;type:int(11);default:0;NOT NULL" json:"reg_num"`        // 注册人数
}

func (m *ActivityRetailerJoinData) TableName() string {
	return "activity_retailer_join_data"
}
