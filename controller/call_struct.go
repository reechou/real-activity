package controller

const (
	RESPONSE_OK = iota
	RESPONSE_ERR
)

type AddActivityReq struct {
	Title string `json:"title"`
}

type AddActivityHeaderReq struct {
	ActivityId int64                    `json:"activityId"`
	Navigation []ActivityNavigationInfo `json:"navigation"`
	Banner     []ActivityBannerInfo     `json:"banner"`
}

type AddActivityItemsReq struct {
	NavigationId int64              `json:"navigationId"`
	ItemList     []ActivityItemInfo `json:"itemList"`
}

type GetActivityHeaderReq struct {
	ActivityId int64 `json:"activityId"`
}

type GetActivityItemsReq struct {
	NavigationId int64  `json:"navigationId"`
	TaobaoPid    string `json:"taobaoPid"`
	Offset       int64  `json:"offset"`
	Num          int64  `json:"num"`
}

type ActivityNavigationInfo struct {
	NavigationName      string `json:"navigationName"`
	NavigationId        int64  `json:"navigationId"`
	NavigationWeight    int64  `json:"navigationWeight"`
	NavigationStartTime int64  `json:"navigationStartTime"`
}

type ActivityBannerInfo struct {
	BannerImgUrl  string `json:"bannerImgUrl"`
	BannerLinkUrl string `json:"bannerLinkUrl"`
	BannerId      int64  `json:"bannerId"`
}

type ActivityItemInfo struct {
	Item string `json:"item"`
	Pid  string `json:"pid"`
}

type ActivityHeader struct {
	Title      string                   `json:"title"`
	Navigation []ActivityNavigationInfo `json:"navigation"`
	Banner     []ActivityBannerInfo     `json:"banner"`
}

type DelActivityReq struct {
	ActivityId int64 `json:"activityId"`
}

type DelActivityItemReq struct {
	ItemId int64 `json:"itemId"`
}

type DelNavigationReq struct {
	NavigationId int64 `json:"navigationId"`
}

type DelBannerReq struct {
	BannerId int64 `json:"bannerId"`
}

type UpdateActivityReq struct {
	ActivityId int64  `json:"activityId"`
	Title      string `json:"title"`
}

type UpdateNavigationReq struct {
	NavigationId int64  `json:"navigationId"`
	Navigation   string `json:"navigation"`
	Weight       int64  `json:"weight"`
}

type UpdateBannerReq struct {
	BannerId int64  `json:"bannerId"`
	ImgUrl   string `json:"imgUrl"`
	LinkUrl  string `json:"linkUrl"`
}

type ActivityResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
