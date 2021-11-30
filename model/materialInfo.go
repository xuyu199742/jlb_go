package model

import "time"

// MaterialInfo 素材信息表
type MaterialInfo struct {
	Id            uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`               // 编号
	Title         string    `gorm:"column:title;type:varchar(50);NOT NULL" json:"title"`                                // 素材标题
	MaterialType  int       `gorm:"column:material_type;type:int(11);NOT NULL" json:"material_type"`                    // 素材类型
	Content       string    `gorm:"column:content;type:varchar(500);NOT NULL" json:"content"`                           // 素材内容
	ShowType      int       `gorm:"column:show_type;type:tinyint(2);NOT NULL" json:"show_type"`                         // 门店可见类型1：部分可见，2：全部可见
	Type          int       `gorm:"column:type;type:tinyint(4);NOT NULL" json:"type"`                                   // 文件类型，1：图片，2：视频
	ImgUrl        string    `gorm:"column:img_url;type:varchar(800)" json:"img_url"`                                    // 图片地址
	VideoUrl      string    `gorm:"column:video_url;type:varchar(500)" json:"video_url"`                                // 视频地址
	VideoImgCover string    `gorm:"column:video_img_cover;type:varchar(255)" json:"video_img_cover"`                    // 视频封面地址
	BaseSendNum   uint      `gorm:"column:base_send_num;type:int(11) unsigned;default:0;NOT NULL" json:"base_send_num"` // 基础转发次数
	ActSendNum    uint      `gorm:"column:act_send_num;type:int(11) unsigned;default:0;NOT NULL" json:"act_send_num"`   // 实际转发次数
	BeginTime     time.Time `gorm:"column:begin_time;type:datetime;NOT NULL" json:"begin_time"`                         // 生效时间
	EndTime       time.Time `gorm:"column:end_time;type:datetime;NOT NULL" json:"end_time"`                             // 失效时间
	Status        int       `gorm:"column:status;type:tinyint(4);default:1;NOT NULL" json:"status"`                     // 状态，0：禁用，1：启用
	IfPublish     int       `gorm:"column:if_publish;type:tinyint(4);default:0;NOT NULL" json:"if_publish"`             // 是否发布，0：未发布，1：已发布
	IsShow        int       `gorm:"column:is_show;type:tinyint(2);default:1;NOT NULL" json:"is_show"`                   // 前端显示
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`                       // 创建时间
	Creator       string    `gorm:"column:creator;type:varchar(255);NOT NULL" json:"creator"`                           // 创建者
}

func (m *MaterialInfo) TableName() string {
	return "material_info"
}
