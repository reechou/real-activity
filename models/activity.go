package models

import (
	"time"

	"github.com/Sirupsen/logrus"
)

type Activity struct {
	ID        int64  `xorm:"pk autoincr"`
	Title     string `xorm:"not null default '' varchar(64)"`
	CreatedAt int64  `xorm:"not null default 0 int"`
	UpdatedAt int64  `xorm:"not null default 0 int"`
}

func CreateActivity(info *Activity) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		logrus.Errorf("create activity error: %v", err)
		return err
	}
	logrus.Infof("create activity[%s] success.", info.Title)
	
	return nil
}
