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
	Navigation []ActivityNavigationInfo `json:"navigation"`
	Banner     []ActivityBannerInfo     `json:"banner"`
}

type ActivityResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
