package models

import (
	"github.com/Sirupsen/logrus"
)

type ActivityNavigation struct {
	ID         int64  `xorm:"pk autoincr"`
	ActivityID int64  `xorm:"not null default 0 int index"`
	Navigation string `xorm:"not null default '' varchar(64)"`
	Weight     int64  `xorm:"not null default 0 int index"`
	StartTime  int64  `xorm:"not null default 0 bigint"`
	CreatedAt  int64  `xorm:"not null default 0 int"`
}

func CreateActivityNavigationList(list []ActivityNavigation) error {
	if len(list) == 0 {
		return nil
	}
	for i := 0; i < len(list); i++ {
		_, err := x.Insert(&list[i])
		if err != nil {
			logrus.Errorf("create activity navigation error: %v", err)
			return err
		}
	}
	return nil
}

func GetActivityNavigationList(activityId int64) ([]ActivityNavigation, error) {
	var navigationList []ActivityNavigation
	err := x.Where("activity_id = ?", activityId).OrderBy("start_time").Desc("weight").Find(&navigationList)
	if err != nil {
		logrus.Errorf("activity_id[%d] get activity navigation list error: %v", activityId, err)
		return nil, err
	}
	return navigationList, nil
}

func DelActivityNavigation(anId int64) error {
	an := &ActivityNavigation{ID: anId}
	_, err := x.Where("id = ?", anId).Delete(an)
	if err != nil {
		logrus.Errorf("id[%d] navigation delete error: %v", anId, err)
		return err
	}
	return nil
}

func UpdateActivityNavigation(info *ActivityNavigation) error {
	_, err := x.Cols("navigation", "weight").Update(info, &ActivityNavigation{ID: info.ID})
	if err != nil {
		logrus.Errorf("navigation[%v] navigation update error: %v", info, err)
		return err
	}
	return nil
}
