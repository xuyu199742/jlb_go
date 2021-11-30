package model

import (
	"gorm.io/gorm"
	"jlb_shop_go/global"
)

// RetailerNoticeInfo B端公告表
type RetailerNoticeInfo struct {
	Id         int    `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"` // 编号
	Title      string `gorm:"column:title;type:varchar(200);NOT NULL" json:"title"`                 // 标题
	Content    string `gorm:"column:content;type:longtext;NOT NULL" json:"content"`                 // 内容
	IsUse      int    `gorm:"column:is_use;type:tinyint(4);NOT NULL" json:"is_use"`                 // 是否可用，0：禁用，1：可用
	CreateTime Time   `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`         // 创建时间
	Creator    int    `gorm:"column:creator;type:int(11);NOT NULL" json:"creator"`                  // 创建者
}

func (m *RetailerNoticeInfo) TableName() string {
	return "retailer_notice_info"
}

type NoticeQuery struct {
	db *gorm.DB
}

func NewNoticeQuery(db ...*gorm.DB) *NoticeQuery {
	if len(db) > 0 {
		return &NoticeQuery{db[0]}
	}

	return &NoticeQuery{global.Db}
}

func (q *NoticeQuery) Id(id ...int) *NoticeQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	})

	return q
}

func (q *NoticeQuery) IsUse(status int) *NoticeQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("is_use = ?", status)
	})

	return q
}

func (q *NoticeQuery) WhereRaw(sql string, args ...interface{}) *NoticeQuery {
	q.db = q.db.Where(sql, args)

	return q
}

func (q *NoticeQuery) OrderBy(value ...interface{}) *NoticeQuery {
	q.db = q.db.Order(value)

	return q
}

func (q *NoticeQuery) First(column ...string) (*RetailerNoticeInfo, error) {
	if len(column) > 0 {
		q.db.Select(column)
	}
	notice := RetailerNoticeInfo{}
	if err := q.db.Order("id desc").First(&notice).Error; err != nil {
		return nil, err
	}

	return &notice, nil
}

func (q *NoticeQuery) ListPage(page, pageSize int, column ...string) ([]*RetailerNoticeInfo, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	if len(column) > 0 {
		q.db.Select(column)
	}

	var total int64
	var list []*RetailerNoticeInfo
	err := q.db.Model(new(RetailerNoticeInfo)).Count(&total).Limit(pageSize).Offset(offset).Find(&list).Error

	return list, total, err
}
