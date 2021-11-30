package model

import "time"

// ActivityBuyMemberRecord 活动用户参与记录
type ActivityBuyMemberRecord struct {
	Id                 int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`                     // 编号
	ActivityId         uint      `gorm:"column:activity_id;type:int(11) unsigned;NOT NULL" json:"activity_id"`            // 活动编号
	FirstBuyRetailerId uint      `gorm:"column:first_buy_retailer_id;type:int(11) unsigned" json:"first_buy_retailer_id"` // 首购门店编号
	FirstBuyStaffId    int       `gorm:"column:first_buy_staff_id;type:int(11)" json:"first_buy_staff_id"`                // 首购员工编号
	FirstBuyOrderCode  string    `gorm:"column:first_buy_order_code;type:varchar(255)" json:"first_buy_order_code"`       // 首购单号
	BuyGoodsId         int       `gorm:"column:buy_goods_id;type:int(11)" json:"buy_goods_id"`                            // 购买商品编号
	GiveGoodsId        int       `gorm:"column:give_goods_id;type:int(11)" json:"give_goods_id"`                          // 赠送商品编号
	LatestRetailerId   int       `gorm:"column:latest_retailer_id;type:int(11)" json:"latest_retailer_id"`                // 最新连接的门店编号
	LatestStaffId      int       `gorm:"column:latest_staff_id;type:int(11)" json:"latest_staff_id"`                      // 最新连接的员工编号
	CreateTime         time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`                    // 参与时间
	MemberId           int       `gorm:"column:member_id;type:int(11);NOT NULL" json:"member_id"`                         // 会员编号
	IsFinish           uint      `gorm:"column:is_finish;type:tinyint(1) unsigned;default:0;NOT NULL" json:"is_finish"`   // 是否完成 (0:未完成，1:已完成)
	IsReceive          uint      `gorm:"column:is_receive;type:tinyint(1) unsigned;default:0;NOT NULL" json:"is_receive"` // 是否领取(0:未领取，1:已领取)
	ReceiveTime        time.Time `gorm:"column:receive_time;type:datetime" json:"receive_time"`                           // 领取时间
	CompleteTime       time.Time `gorm:"column:complete_time;type:datetime" json:"complete_time"`                         // 完成时间
}

func (m *ActivityBuyMemberRecord) TableName() string {
	return "activity_buy_member_record"
}
