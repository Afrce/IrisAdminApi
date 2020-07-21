package models

import (
	"IrisAdminApi/sysinit"
	"time"
)

type ApiLog struct {
	ID     uint64 `gorm:"primary_key"`
	Url    string
	UserId uint
	Params string
	Time   time.Time
}

func (a ApiLog) TableName() string {
	return "apiLogs"
}

func RecordLogs(log ApiLog) {
	sysinit.DB.Create(&log)
}
