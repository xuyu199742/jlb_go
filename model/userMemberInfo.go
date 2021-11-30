package model

import (
	"gorm.io/gorm"
	"jlb_shop_go/global"
	"time"
)

// UserMemberInfo  会员信息表
type UserMemberInfo struct {
	Id              uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`              // 编号
	Name            string    `gorm:"column:name;type:varchar(30)" json:"name"`                                          // 姓名
	Sex             int       `gorm:"column:sex;type:tinyint(4);default:0;NOT NULL" json:"sex"`                          // 性别，0：未知，1：男，2：女
	Birth           time.Time `gorm:"column:birth;type:date" json:"birth"`                                               // 生日
	Phone           string    `gorm:"column:phone;type:varchar(20);NOT NULL" json:"phone"`                               // 手机号
	Nickname        string    `gorm:"column:nickname;type:varchar(64)" json:"nickname"`                                  // 微信昵称
	Avatar          string    `gorm:"column:avatar;type:varchar(255)" json:"avatar"`                                     // 微信头像
	OpenId          string    `gorm:"column:open_id;type:varchar(255);NOT NULL" json:"open_id"`                          // 微信open_id
	UnionId         string    `gorm:"column:union_id;type:varchar(255)" json:"union_id"`                                 // 微信union_id
	CrmCode         string    `gorm:"column:crm_code;type:varchar(255)" json:"crm_code"`                                 // crm唯一标识
	UserCode        string    `gorm:"column:user_code;type:varchar(255);NOT NULL" json:"user_code"`                      // 会员编码
	Status          int       `gorm:"column:status;type:tinyint(4);default:1;NOT NULL" json:"status"`                    // 状态，0：禁用，1：可用
	CreateTime      time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`                      // 注册时间
	MemStage        int       `gorm:"column:mem_stage;type:tinyint(4);default:1;NOT NULL" json:"mem_stage"`              // 会员阶段，1：新会员，2：新客，3：老客
	ChannelId       int       `gorm:"column:channel_id;type:int(11);default:0;NOT NULL" json:"channel_id"`               // 隶属渠道
	GoodsId         int       `gorm:"column:goods_id;type:int(11);default:0;NOT NULL" json:"goods_id"`                   // 最近交易的商品编号
	ComeFrom        int       `gorm:"column:come_from;type:tinyint(4);NOT NULL" json:"come_from"`                        // 来源，1：扫码注册，2：直接注册，3：晒单注册，4：红包分享注册，5：买一赠一注册，6：导购邀请注册
	SupMemId        int       `gorm:"column:sup_mem_id;type:int(11);default:0;NOT NULL" json:"sup_mem_id"`               // 上级会员ID
	CrmShopCode     string    `gorm:"column:crm_shop_code;type:varchar(50);NOT NULL" json:"crm_shop_code"`               // CRM隶属门店
	FirstRegShop    int       `gorm:"column:first_reg_shop;type:int(11);default:0;NOT NULL" json:"first_reg_shop"`       // 首次注册门店
	FirstRegStaff   int       `gorm:"column:first_reg_staff;type:int(11);default:0;NOT NULL" json:"first_reg_staff"`     // 首次注册店员
	FirstActivityId int       `gorm:"column:first_activity_id;type:int(11);default:0;NOT NULL" json:"first_activity_id"` // 首次注册活动编号
}

func (m *UserMemberInfo) TableName() string {
	return "user_member_info"
}

type UserMemberQuery struct {
	db *gorm.DB
}

func NewUserMemberQuery(db ...*gorm.DB) *UserMemberQuery {
	if len(db) > 0 {
		return &UserMemberQuery{db[0]}
	}
	return &UserMemberQuery{global.Db}
}

func (q *UserMemberQuery) Status(status int) *UserMemberQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("status", status)
	})

	return q
}

func (q *UserMemberQuery) FirstRegStaff(staffId int) *UserMemberQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("first_reg_staff", staffId)
	})

	return q
}

func (q *UserMemberQuery) FirstRegShop(shopId int) *UserMemberQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("first_reg_staff", shopId)
	})

	return q
}

func (q *UserMemberQuery) BtCreateTime(time []string) *UserMemberQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("create_time between ? and ?", time[0], time[1])
	})

	return q
}

func (q *UserMemberQuery) Count() (*UserMemberQuery, int64) {
	var count int64
	q.db = q.db.Model(new(UserMemberInfo)).Count(&count)

	return q, count
}

func (q *UserMemberQuery) MemStage(state int) *UserMemberQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("mem_stage = ?", state)
	})

	return q
}

func (q *UserMemberQuery) FindOne(column ...string) (*UserMemberInfo, error) {
	if len(column) > 0 {
		q.db.Select(column)
	}
	user := UserMemberInfo{}
	err := q.db.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *UserMemberQuery) Id(id ...int) *UserMemberQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	})

	return q
}

func (q *UserMemberQuery) PageList(page, size int, column ...string) ([]*UserMemberInfo, int64, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	offset := (page - 1) * size

	if len(column) > 0 {
		q.db = q.db.Select(column)
	}

	var total int64
	var list []*UserMemberInfo
	err := q.db.Model(new(UserMemberInfo)).Count(&total).Order("create_time desc").Limit(size).Offset(offset).Find(&list).Error

	return list, total, err
}
