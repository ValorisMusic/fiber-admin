package models

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Name     string  `gorm:"size:30;not null;column:name;comment:等级名称" json:"name" validate:"required"`
	Icon     string  `gorm:"size:255;column:icon;comment:等级图标" json:"icon" validate:"required"`
	PayPrice float64 `gorm:"type:decimal(10,2);column:pay_price;comment:支付价格" json:"pay_price" validate:"required"`
	Revenue  float64 `gorm:"type:decimal(10,2);column:revenue;comment:任务收入" json:"revenue" validate:"required"`
	Daily    int     `gorm:"column:daily;comment:每日任务数量" json:"daily" validate:"required"`
	IsStatus *int8   `gorm:"type:tinyint(1);column:is_status;comment:1启用 0禁用" json:"is_status"`
	Period   int     `gorm:"column:period;comment:周期" json:"period"`
}

// TableName sets the insert table name for this struct type
func (s *Level) TableName() string {
	return "level"
}
