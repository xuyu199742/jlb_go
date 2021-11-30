package model

import (
	"gorm.io/gorm"
	"jlb_shop_go/global"
	"time"
)

// RetailerStaffInfo 零售商员工信息表
type RetailerStaffInfo struct {
	Id           int       `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	StaffCode    string    `gorm:"column:staff_code;type:varchar(100);NOT NULL" json:"staff_code"`       // 员工编码
	LoginAccount string    `gorm:"column:login_account;type:varchar(100);NOT NULL" json:"login_account"` // 登录账号
	Username     string    `gorm:"column:username;type:varchar(100)" json:"username"`                    // 姓名
	Gender       int       `gorm:"column:gender;type:tinyint(2);default:0;NOT NULL" json:"gender"`       // 性别，0：未知，1：男，2：女
	Phone        string    `gorm:"column:phone;type:varchar(20)" json:"phone"`                           // 精准营销手机号
	Position     int       `gorm:"column:position;type:tinyint(2);NOT NULL" json:"position"`             // 职位，1：店长，2：店员
	StaffType    string    `gorm:"column:staff_type;type:varchar(20)" json:"staff_type"`                 // 店员类型
	IsUse        int       `gorm:"column:is_use;type:tinyint(2);default:1;NOT NULL" json:"is_use"`       // 是否可用，0：禁用，1：可用
	Status       int       `gorm:"column:status;type:tinyint(2);default:1;NOT NULL" json:"status"`       // 精准营销状态，0：无效，1：有效
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`         // 创建时间
	OpenId       string    `gorm:"column:open_id;type:varchar(255)" json:"open_id"`                      // 微信open_id
	UnionId      string    `gorm:"column:union_id;type:varchar(255)" json:"union_id"`                    // 微信union_id
	Nickname     string    `gorm:"column:nickname;type:varchar(64)" json:"nickname"`                     // 微信昵称
	Avatar       string    `gorm:"column:avatar;type:varchar(255)" json:"avatar"`                        // 微信头像
	WxPhone      string    `gorm:"column:wx_phone;type:varchar(20)" json:"wx_phone"`                     // 微信授权手机号
}

func (m *RetailerStaffInfo) TableName() string {
	return "retailer_staff_info"
}

// ============= executor ==============//

type StaffInfoExecutor struct {
	db *gorm.DB
}

func NewStaffInfoExecutor(db ...*gorm.DB) *StaffInfoExecutor {
	if len(db) > 0 {
		return &StaffInfoExecutor{db[0]}
	}
	return &StaffInfoExecutor{global.Db}
}

func (e *StaffInfoExecutor) Save(staff *RetailerStaffInfo) error {
	return e.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(staff).Error; err != nil {
			return err
		}
		return nil
	})
}

// ============= query ==============//

type StaffInfoQuery struct {
	db *gorm.DB
}

func NewStaffInfoQuery(db ...*gorm.DB) *StaffInfoQuery {
	if len(db) > 0 {
		return &StaffInfoQuery{db[0]}
	}

	return &StaffInfoQuery{global.Db}
}

func (q *StaffInfoQuery) Status(status int) *StaffInfoQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("status", status)
	})

	return q
}

func (q *StaffInfoQuery) Id(id ...int) *StaffInfoQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	})

	return q
}

func (q *StaffInfoQuery) Account(account string) *StaffInfoQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("login_account", account)
	})

	return q
}

func (q *StaffInfoQuery) FindOne(column ...string) (*RetailerStaffInfo, error) {
	if len(column) > 0 {
		q.db.Select(column)
	}
	staff := RetailerStaffInfo{}
	err := q.db.First(&staff).Error
	if err != nil {
		return nil, err
	}

	return &staff, nil
}
