package models

import (
	"time"

	"github.com/Sirupsen/logrus"
	"fmt"
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

func GetActivityInfo(info *Activity) (bool, error) {
	has, err := x.Where("id = ?", info.ID).Get(info)
	if err != nil {
		logrus.Errorf("get activity id[%d] error: %v", info.ID, err)
		return false, fmt.Errorf("get activity id[%d] error: %v", info.ID, err)
	}
	if !has {
		logrus.Errorf("get activity id[%d] has no this order.", info.ID)
		return false, nil
	}
	return true, nil
}

func GetActivityList() ([]Activity, error) {
	var activityList []Activity
	err := x.Find(&activityList)
	if err != nil {
		logrus.Errorf("get activity list error: %v", err)
		return nil, err
	}
	return activityList, nil
}

func DelActivity(aId int64) error {
	a := &Activity{ID: aId}
	_, err := x.Where("id = ?", aId).Delete(a)
	if err != nil {
		logrus.Errorf("id[%d] activity delete error: %v", aId, err)
		return err
	}
	return nil
}

func UpdateActivity(info *Activity) error {
	_, err := x.Cols("title").Update(info, &Activity{ID: info.ID})
	if err != nil {
		logrus.Errorf("activity[%v] activity update error: %v", info, err)
		return err
	}
	return nil
}
