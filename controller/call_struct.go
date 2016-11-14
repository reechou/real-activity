package controller

const (
	RESPONSE_OK = iota
	RESPONSE_ERR
)

type AddActivityReq struct {
	Title string `json:"title"`
}

type AddActivityHeaderReq struct {
	ActivityId int64                `json:"activityId"`
	Navigation []string             `json:"navigation"`
	Banner     []ActivityBannerInfo `json:"banner"`
}

type AddActivityItemsReq struct {
	NavigationId int64    `json:"navigationId"`
	ItemList     []string `json:"itemList"`
}

type GetActivityHeaderReq struct {
	ActivityId int64 `json:"activityId"`
}

type GetActivityItemsReq struct {
	NavigationId int64 `json:"navigationId"`
	Offset       int64 `json:"offset"`
	Num          int64 `json:"num"`
}

type ActivityNavigationInfo struct {
	NavigationName string `json:"navigationName"`
	NavigationId   int64  `json:"navigationId"`
}

type ActivityBannerInfo struct {
	BannerImgUrl  string `json:"bannerImgUrl"`
	BannerLinkUrl string `json:"bannerLinkUrl"`
	BannerId      int64  `json:"bannerId"`
}

type ActivityHeader struct {
	Title      string                   `json:"title"`
	Navigation []ActivityNavigationInfo `json:"navigation"`
	Banner     []ActivityBannerInfo     `json:"banner"`
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
