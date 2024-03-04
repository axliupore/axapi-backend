package model

import "time"

// Model 全局的实体属性
type Model struct {
	CreateTime time.Time `db:"create_time" json:"-"`
	UpdateTime time.Time `db:"update_time" json:"-"`
	IsDelete   int8      `db:"is_delete" json:"-"`
}

func (m *Model) Init() {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	m.IsDelete = 0
}
