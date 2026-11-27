package models

import (
	"github.com/goravel/framework/database/orm"
)

type Link struct {
	orm.Model
	Code   string `gorm:"column:code" json:"code"`
	URL    string `gorm:"column:url" json:"url"`
	Clicks int    `gorm:"column:clicks" json:"clicks"`
}
