package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/reechou/real-activity/models"
)

func (al *ActLogic) addActivityHandler(w http.ResponseWriter, r *http.Request) {
	req := &AddActivityReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	activity := &models.Activity{
		Title: req.Title,
	}
	err := models.CreateActivity(activity)
	if err != nil {
		logrus.Errorf("Error add activity error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error add activity error: %v", err)
	}
	rsp.Data = activity.ID

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) addActivityHeaderHandler(w http.ResponseWriter, r *http.Request) {
	req := &AddActivityHeaderReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	now := time.Now().Unix()
	var navigationList []models.ActivityNavigation
	for _, v := range req.Navigation {
		n := models.ActivityNavigation{
			ActivityID: req.ActivityId,
			Navigation: v,
			CreatedAt:  now,
		}
		navigationList = append(navigationList, n)
	}
	err := models.CreateActivityNavigationList(navigationList)
	if err != nil {
		logrus.Errorf("Error add activity navigation list error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error add activity navigation list error: %v", err)
	} else {
		var bannerList []models.ActivityBanner
		for _, v := range req.Banner {
			b := models.ActivityBanner{
				ActivityId: req.ActivityId,
				ImgUrl:     v.BannerImgUrl,
				LinkUrl:    v.BannerLinkUrl,
				CreatedAt:  now,
			}
			bannerList = append(bannerList, b)
		}
		err := models.CreateActivityBannerList(bannerList)
		if err != nil {
			logrus.Errorf("Error add activity banner list error: %v", err)
			rsp.Code = RESPONSE_ERR
			rsp.Msg = fmt.Sprintf("Error add activity banner list error: %v", err)
		} else {
			var nInfo []ActivityNavigationInfo
			for _, v := range navigationList {
				ani := ActivityNavigationInfo{
					NavigationName: v.Navigation,
					NavigationId:   v.ID,
				}
				nInfo = append(nInfo, ani)
			}
			var bInfo []ActivityBannerInfo
			for _, v := range bannerList {
				abi := ActivityBannerInfo{
					BannerImgUrl: v.ImgUrl,
					BannerLinkUrl: v.LinkUrl,
					BannerId: v.ID,
				}
				bInfo = append(bInfo, abi)
			}
			header := &ActivityHeader{
				Navigation: nInfo,
				Banner:     bInfo,
			}
			rsp.Data = header
		}
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) addActivityItemsHandler(w http.ResponseWriter, r *http.Request) {
	req := &AddActivityItemsReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	now := time.Now().Unix()
	var itemList []models.ActivityItem
	for _, v := range req.ItemList {
		ai := models.ActivityItem{
			NavigationId: req.NavigationId,
			Item:         v,
			CreatedAt:    now,
		}
		itemList = append(itemList, ai)
	}
	err := models.CreateActivityItemsList(itemList)
	if err != nil {
		logrus.Errorf("Error add activity items error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error add activity items error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) getActivityHeaderHandler(w http.ResponseWriter, r *http.Request) {
	req := &GetActivityHeaderReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	nList, err := models.GetActivityNavigationList(req.ActivityId)
	if err != nil {
		logrus.Errorf("Error get activity navigation list error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error get activity navigation list error: %v", err)
	} else {
		bList, err := models.GetActivityBannerList(req.ActivityId)
		if err != nil {
			logrus.Errorf("Error get activity banner list error: %v", err)
			rsp.Code = RESPONSE_ERR
			rsp.Msg = fmt.Sprintf("Error get activity banner list error: %v", err)
		} else {
			var aviList []ActivityNavigationInfo
			for _, v := range nList {
				avi := ActivityNavigationInfo{
					NavigationName: v.Navigation,
					NavigationId:   v.ID,
				}
				aviList = append(aviList, avi)
			}
			var bannerList []ActivityBannerInfo
			for _, v := range bList {
				abi := ActivityBannerInfo{
					BannerImgUrl: v.ImgUrl,
					BannerLinkUrl: v.LinkUrl,
					BannerId:   v.ID,
				}
				bannerList = append(bannerList, abi)
			}
			rsp.Data = &ActivityHeader{
				Navigation: aviList,
				Banner:     bannerList,
			}
		}
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) getActivityItemsHandler(w http.ResponseWriter, r *http.Request) {
	req := &GetActivityItemsReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	list, err := models.GetActivityItemList(req.NavigationId)
	if err != nil {
		logrus.Errorf("Error get activity items error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error get activity items error: %v", err)
	} else {
		var itemList []string
		for _, v := range list {
			itemList = append(itemList, v.Item)
		}
		rsp.Data = itemList
	}

	WriteJSON(w, http.StatusOK, rsp)
}