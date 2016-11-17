package models

import (
	"github.com/Sirupsen/logrus"
)

type ActivityItem struct {
	ID           int64  `xorm:"pk autoincr"`
	NavigationId int64  `xorm:"not null default 0 int index"`
	TaobaoPid    string `xorm:"not null default '' varchar(128) index"`
	Item         string `xorm:"not null default '' varchar(2048)"`
	CreatedAt    int64  `xorm:"not null default 0 int"`
}

func CreateActivityItemsList(list []ActivityItem) error {
	if len(list) == 0 {
		return nil
	}
	_, err := x.Insert(&list)
	if err != nil {
		logrus.Errorf("create activity item list error: %v", err)
		return err
	}
	return nil
}

func GetActivityItemCount(navigationId int64, taobaoPid string) (int64, error) {
	count, err := x.Where("navigation_id = ?", navigationId).And("taobao_pid = ?", taobaoPid).Count(&ActivityItem{})
	if err != nil {
		logrus.Errorf("navigation_id[%d] get item list count error: %v", navigationId, err)
		return 0, err
	}
	return count, nil
}

func GetActivityItemList(navigationId, offset, num int64, taobaoPid string) ([]ActivityItem, error) {
	var itemList []ActivityItem
	err := x.Where("navigation_id = ?", navigationId).And("taobao_pid = ?", taobaoPid).Limit(int(num), int(offset)).Find(&itemList)
	if err != nil {
		logrus.Errorf("navigation_id[%d] get activity item list error: %v", navigationId, err)
		return nil, err
	}
	return itemList, nil
}

func DelActivityItem(itemId int64) error {
	item := &ActivityItem{ID: itemId}
	_, err := x.Where("id = ?", itemId).Delete(item)
	if err != nil {
		logrus.Errorf("id[%d] item delete error: %v", itemId, err)
		return err
	}
	return nil
}
