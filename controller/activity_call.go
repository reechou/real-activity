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
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &AddActivityReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("addActivityHandler json decode error: %v", err)
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
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &AddActivityHeaderReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("addActivityHeaderHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	now := time.Now().Unix()
	var navigationList []models.ActivityNavigation
	for _, v := range req.Navigation {
		logrus.Debugf("add navigation: %v", v)
		n := models.ActivityNavigation{
			ActivityID: req.ActivityId,
			Navigation: v.NavigationName,
			Weight:     v.NavigationWeight,
			StartTime:  v.NavigationStartTime,
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
					BannerImgUrl:  v.ImgUrl,
					BannerLinkUrl: v.LinkUrl,
					BannerId:      v.ID,
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
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &AddActivityItemsReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("addActivityItemsHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	now := time.Now().Unix()
	var itemList []models.ActivityItem
	for _, v := range req.ItemList {
		ai := models.ActivityItem{
			NavigationId: req.NavigationId,
			Item:         v.Item,
			TaobaoPid:    v.Pid,
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

func (al *ActLogic) getActivityListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	list, err := models.GetActivityList()
	if err != nil {
		logrus.Errorf("Error get activity list error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error get activity list error: %v", err)
	} else {
		rsp.Data = list
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) getActivityHeaderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &GetActivityHeaderReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("getActivityHeaderHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	aInfo := &models.Activity{
		ID: req.ActivityId,
	}
	has, err := models.GetActivityInfo(aInfo)
	if err != nil || !has {
		logrus.Errorf("Error get activity info error: %v or not found", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error get activity info error: %v", err)
	} else {
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
						NavigationName:      v.Navigation,
						NavigationId:        v.ID,
						NavigationWeight:    v.Weight,
						NavigationStartTime: v.StartTime,
					}
					aviList = append(aviList, avi)
				}
				var bannerList []ActivityBannerInfo
				for _, v := range bList {
					abi := ActivityBannerInfo{
						BannerImgUrl:  v.ImgUrl,
						BannerLinkUrl: v.LinkUrl,
						BannerId:      v.ID,
					}
					bannerList = append(bannerList, abi)
				}
				rsp.Data = &ActivityHeader{
					Title:      aInfo.Title,
					Navigation: aviList,
					Banner:     bannerList,
				}
			}
		}
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) getActivityItemsHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		logrus.Debugf("get activity item list call use_time[%v]", time.Now().Sub(start))
	}()
	
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &GetActivityItemsReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("getActivityItemsHandler json decode error: %v", err)
		return
	}
	logrus.Debugf("get items req: %v", req)

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	count, err := models.GetActivityItemCount(req.NavigationId, req.TaobaoPid)
	if err != nil {
		logrus.Errorf("Error get activity items count error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error get activity items count error: %v", err)
	} else {
		list, err := models.GetActivityItemList(req.NavigationId, req.Offset, req.Num, req.TaobaoPid)
		if err != nil {
			logrus.Errorf("Error get activity items error: %v", err)
			rsp.Code = RESPONSE_ERR
			rsp.Msg = fmt.Sprintf("Error get activity items error: %v", err)
		} else {
			type ItemList struct {
				Count int64
				List  []models.ActivityItem
			}
			l := &ItemList{
				Count: count,
				List:  list,
			}
			rsp.Data = l
		}
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) delActivityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &DelActivityReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("delActivityHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	err := models.DelActivity(req.ActivityId)
	if err != nil {
		logrus.Errorf("Error del activity error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error del activity error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) delActivityItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &DelActivityItemReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("delActivityItemHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	err := models.DelActivityItem(req.ItemId)
	if err != nil {
		logrus.Errorf("Error del activity item error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error del activity item error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) delActivityNavigationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &DelNavigationReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("delActivityNavigationHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	err := models.DelActivityNavigation(req.NavigationId)
	if err != nil {
		logrus.Errorf("Error del activity navigation error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error del activity navigation error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) delActivityBannerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &DelBannerReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("delActivityBannerHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	err := models.DelActivityBanner(req.BannerId)
	if err != nil {
		logrus.Errorf("Error del activity banner error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error del activity banner error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) updateActivityNavigationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &UpdateNavigationReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("updateActivityNavigationHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	info := &models.ActivityNavigation{
		ID:         req.NavigationId,
		Navigation: req.Navigation,
		Weight:     req.Weight,
	}
	err := models.UpdateActivityNavigation(info)
	if err != nil {
		logrus.Errorf("Error update activity navigation error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error update activity navigation error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) updateActivityBannerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &UpdateBannerReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("updateActivityBannerHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	info := &models.ActivityBanner{
		ID:      req.BannerId,
		ImgUrl:  req.ImgUrl,
		LinkUrl: req.LinkUrl,
	}
	err := models.UpdateActivityBanner(info)
	if err != nil {
		logrus.Errorf("Error update activity banner error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error update activity banner error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}

func (al *ActLogic) updateActivityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WriteJSON(w, http.StatusOK, nil)
		return
	}
	req := &UpdateActivityReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logrus.Errorf("updateActivityHandler json decode error: %v", err)
		return
	}

	rsp := &ActivityResponse{Code: RESPONSE_OK}

	info := &models.Activity{
		ID:    req.ActivityId,
		Title: req.Title,
	}
	err := models.UpdateActivity(info)
	if err != nil {
		logrus.Errorf("Error update activity error: %v", err)
		rsp.Code = RESPONSE_ERR
		rsp.Msg = fmt.Sprintf("Error update activity error: %v", err)
	}

	WriteJSON(w, http.StatusOK, rsp)
}
