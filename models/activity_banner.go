package models

import (
	"github.com/Sirupsen/logrus"
)

type ActivityBanner struct {
	ID         int64  `xorm:"pk autoincr"`
	ActivityId int64  `xorm:"not null default 0 int index"`
	ImgUrl     string `xorm:"not null default '' varchar(128)"`
	LinkUrl    string `xorm:"not null default '' varchar(128)"`
	CreatedAt  int64  `xorm:"not null default 0 int"`
}

func CreateActivityBannerList(list []ActivityBanner) error {
	if len(list) == 0 {
		return nil
	}
	for i := 0; i < len(list); i++ {
		_, err := x.Insert(&list[i])
		if err != nil {
			logrus.Errorf("create activity banner error: %v", err)
			return err
		}
	}
	return nil
}

func GetActivityBannerList(activityId int64) ([]ActivityBanner, error) {
	var bannerList []ActivityBanner
	err := x.Where("activity_id = ?", activityId).Find(&bannerList)
	if err != nil {
		logrus.Errorf("activity_id[%d] get activity banner list error: %v", activityId, err)
		return nil, err
	}
	return bannerList, nil
}

func DelActivityBanner(abId int64) error {
	ab := &ActivityBanner{ID: abId}
	_, err := x.Where("id = ?", abId).Delete(ab)
	if err != nil {
		logrus.Errorf("id[%d] banner delete error: %v", abId, err)
		return err
	}
	return nil
}
