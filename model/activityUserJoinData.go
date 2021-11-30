package model

import (
	"gorm.io/gorm"
	"jlb_shop_go/global"
	"time"
)

// ActivityUserJoinData 会员参与的活动
type ActivityUserJoinData struct {
	Id           uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`  // 编号
	MemId        int       `gorm:"column:mem_id;type:int(11);NOT NULL" json:"mem_id"`                     // 会员编号
	ActivityId   int       `gorm:"column:activity_id;type:int(11);NOT NULL" json:"activity_id"`           // 活动编号
	ActivityType int       `gorm:"column:activity_type;type:tinyint(4);NOT NULL" json:"activity_type"`    // 活动类型，1：晒单活动，2：扫赠活动，3：红包活动，4：买赠活动，5：蒲公英计划
	JoinTime     time.Time `gorm:"column:join_time;type:datetime;NOT NULL" json:"join_time"`              // 参与时间
	RedUrl       string    `gorm:"column:red_url;type:varchar(255)" json:"red_url"`                       // 跳转地址
	IfFinish     int       `gorm:"column:if_finish;type:tinyint(4);default:0;NOT NULL" json:"if_finish"`  // 是否完成，0：未完成，1：已完成
	FinishTime   time.Time `gorm:"column:finish_time;type:datetime" json:"finish_time"`                   // 完成时间
	RetailerId   int       `gorm:"column:retailer_id;type:int(11);default:0;NOT NULL" json:"retailer_id"` // 门店编号
	StaffId      int       `gorm:"column:staff_id;type:int(11);default:0;NOT NULL" json:"staff_id"`       // 店员编号
}

func (m *ActivityUserJoinData) TableName() string {
	return "activity_user_join_data"
}

type ActivityUserJoinDataQuery struct {
	db *gorm.DB
}

func NewActivityUserJoinDataQuery(db ...*gorm.DB) *ActivityUserJoinDataQuery {
	if len(db) > 0 {
		return &ActivityUserJoinDataQuery{db[0]}
	}
	return &ActivityUserJoinDataQuery{global.Db}
}

func (q *ActivityUserJoinDataQuery) BtJoinTime(time []string) *ActivityUserJoinDataQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("join_time between ? and ?", time[0], time[1])
	})

	return q
}

func (q *ActivityUserJoinDataQuery) Count() (*ActivityUserJoinDataQuery, int64) {
	var count int64
	q.db = q.db.Model(new(ActivityUserJoinData)).Count(&count)

	return q, count
}

func (q *ActivityUserJoinDataQuery) FindOne(column ...string) (*ActivityUserJoinData, error) {
	if len(column) > 0 {
		q.db.Select(column)
	}
	data := ActivityUserJoinData{}
	err := q.db.First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (q *ActivityUserJoinDataQuery) Id(id ...int) *ActivityUserJoinDataQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	})

	return q
}

func (q *ActivityUserJoinDataQuery) StaffId(id ...int) *ActivityUserJoinDataQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("staff_id in (?)", ids)
	})

	return q
}

func (q *ActivityUserJoinDataQuery) RetailerId(id ...int) *ActivityUserJoinDataQuery {
	ids := make([]int, 0)
	for _, v := range id {
		ids = append(ids, v)
	}

	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("retailer_id in (?)", ids)
	})

	return q
}

func (q *ActivityUserJoinDataQuery) IsFinish(status int) *ActivityUserJoinDataQuery {
	q.db = q.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("if_finish in (?)", status)
	})

	return q
}
