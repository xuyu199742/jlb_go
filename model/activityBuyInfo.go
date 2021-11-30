package model

import "time"

// ActivityBuyInfo 活动信息表
type ActivityBuyInfo struct {
	Id                uint      `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`         // 编号
	Name              string    `gorm:"column:name;type:varchar(80);NOT NULL" json:"name"`                            // 活动名称
	BeginTime         time.Time `gorm:"column:begin_time;type:datetime;NOT NULL" json:"begin_time"`                   // 活动开始时间
	EndTime           time.Time `gorm:"column:end_time;type:datetime;NOT NULL" json:"end_time"`                       // 活动结束时间
	JoinRange         int       `gorm:"column:join_range;type:tinyint(1);NOT NULL" json:"join_range"`                 // 门店参与范围1:全国2:部分区域3:指定门店
	ReceiveMethod     int       `gorm:"column:receive_method;type:tinyint(1);NOT NULL" json:"receive_method"`         // 礼品领取方式1:线上邮寄2:门店核销
	GoodsLineId       uint      `gorm:"column:goods_line_id;type:int(10) unsigned;NOT NULL" json:"goods_line_id"`     // 活动对应系列编号
	Status            int       `gorm:"column:status;type:tinyint(1);default:0;NOT NULL" json:"status"`               // 状态0:禁用1:启用
	IsPublish         int       `gorm:"column:is_publish;type:tinyint(1);default:0;NOT NULL" json:"is_publish"`       // 发布状态0:未发布1:已发布
	IsPostage         int       `gorm:"column:is_postage;type:tinyint(1)" json:"is_postage"`                          // 是否包邮0:否1:是
	Freight           float64   `gorm:"column:freight;type:decimal(16,2)" json:"freight"`                             // 运费
	CRules            string    `gorm:"column:c_rules;type:text;NOT NULL" json:"c_rules"`                             // C端活动规则
	RetailerRules     string    `gorm:"column:retailer_rules;type:text;NOT NULL" json:"retailer_rules"`               // 门店活动规则
	BirthCertAgeLimit int       `gorm:"column:birth_cert_age_limit;type:int(10)" json:"birth_cert_age_limit"`         // 出生证年龄限制-几周岁以下
	CBackgroundImg    string    `gorm:"column:c_background_img;type:varchar(255);NOT NULL" json:"c_background_img"`   // C端活动背景图
	BBackgroundImg    string    `gorm:"column:b_background_img;type:varchar(255);NOT NULL" json:"b_background_img"`   // B端活动背景图
	ListImg           string    `gorm:"column:list_img;type:varchar(255);NOT NULL" json:"list_img"`                   // 活动列表图
	BSharePostImg     string    `gorm:"column:b_share_post_img;type:varchar(255);NOT NULL" json:"b_share_post_img"`   // B端分享海报图
	CSharePostImg     string    `gorm:"column:c_share_post_img;type:varchar(255);NOT NULL" json:"c_share_post_img"`   // C端分享海报图
	WelfareEnterImg   string    `gorm:"column:welfare_enter_img;type:varchar(255);NOT NULL" json:"welfare_enter_img"` // 福利页入口图
	HelpPageImg       string    `gorm:"column:help_page_img;type:varchar(255);NOT NULL" json:"help_page_img"`         // 助力页面图
	CShareCardImg     string    `gorm:"column:c_share_card_img;type:varchar(255);NOT NULL" json:"c_share_card_img"`   // C端小程序分享卡片
	BShareCardImg     string    `gorm:"column:b_share_card_img;type:varchar(255);NOT NULL" json:"b_share_card_img"`   // B端小程序分享卡片
	CShareDesc        string    `gorm:"column:c_share_desc;type:varchar(50);NOT NULL" json:"c_share_desc"`            // C端小程序分享描述
	BShareDesc        string    `gorm:"column:b_share_desc;type:varchar(50);NOT NULL" json:"b_share_desc"`            // B端小程序分享描述
	CreateTime        time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`                 // 创建时间
	Creator           uint      `gorm:"column:creator;type:int(10) unsigned;NOT NULL" json:"creator"`                 // 创建人
}

func (m *ActivityBuyInfo) TableName() string {
	return "activity_buy_info"
}
