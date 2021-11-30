package model

import (
	"gorm.io/gorm"
	"jlb_shop_go/global"
	"jlb_shop_go/model/enums"
	"time"
)

// RetailerStaffRelated 门店员工关联表
type RetailerStaffRelated struct {
	Id         uint              `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	RetailerId int               `gorm:"column:retailer_id;type:int(11);NOT NULL" json:"retailer_id"`          // 零售商编号
	StaffId    int               `gorm:"column:staff_id;type:int(11);NOT NULL" json:"staff_id"`                // 员工编号
	BindStatus enums.BaseStatus  `gorm:"column:bind_status;type:tinyint(4);NOT NULL" json:"bind_status"`       // 绑定状态，0：禁用，1：可用
	BindTime   time.Time         `gorm:"column:bind_time;type:datetime;NOT NULL" json:"bind_time"`             // 关联时间
	Staff      RetailerStaffInfo `gorm:"foreignKey:StaffId;references:Id"`
	Retailer   RetailerInfo      `gorm:"foreignKey:RetailerId;references:Id"`
}

func (m *RetailerStaffRelated) TableName() string {
	return "retailer_staff_related"
}

type RetailerStaffRelatedQuery struct {
	db *gorm.DB
}

func NewRetailerStaffRelatedQuery(db ...*gorm.DB) *RetailerStaffRelatedQuery {
	if len(db) > 0 {
		return &RetailerStaffRelatedQuery{db[0]}
	}

	return &RetailerStaffRelatedQuery{global.Db}
}

func (q *RetailerStaffRelatedQuery) Status(status int) *RetailerStaffRelatedQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("bind_status", status)
	})

	return q
}

func (q *RetailerStaffRelatedQuery) Clone() *RetailerStaffRelatedQuery {
	clone := *q
	return &clone
}

func (q *RetailerStaffRelatedQuery) StaffId(id ...int) *RetailerStaffRelatedQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("staff_id in (?)", ids)
	})

	return q
}

func (q *RetailerStaffRelatedQuery) PreloadStaff(scopes ...func(db *gorm.DB) *gorm.DB) *RetailerStaffRelatedQuery {
	q.db = q.db.Preload("Staff", func(db *gorm.DB) *gorm.DB {
		for _, s := range scopes {
			db = db.Scopes(s)
		}
		return db
	})

	return q
}

func (q *RetailerStaffRelatedQuery) PreloadRetailer(scopes ...func(db *gorm.DB) *gorm.DB) *RetailerStaffRelatedQuery {
	q.db = q.db.Preload("Retailer", func(db *gorm.DB) *gorm.DB {
		for _, v := range scopes {
			db = db.Scopes(v)
		}
		return db
	})

	return q
}

func (q *RetailerStaffRelatedQuery) RetailerId(id ...int) *RetailerStaffRelatedQuery {
	ids := make([]int, 0)
	for _, v := range id {
		if v > 0 {
			ids = append(ids, v)
		}
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("retailer_id in (?)", ids)
	})

	return q
}

func (q *RetailerStaffRelatedQuery) All(column ...string) (list []*RetailerStaffRelated, err error) {
	if len(column) > 0 {
		q.db.Select(column)
	}

	if err = q.db.Find(&list).Order("bind_time DESC").Error; err != nil {
		return list, err
	}

	return list, nil
}
