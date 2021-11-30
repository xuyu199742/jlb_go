package protocols

type HomeDataReq struct {
	Params struct {
		RetailerId int `json:"retailer_id" form:"retailer_id"`
	} `json:"params" form:"params"`
}

type HomeDataRes struct {
	ShopId          int                      `json:"shopId"`
	ShopName        string                   `json:"shopName"`
	ShopList        []map[string]interface{} `json:"shopList"`
	Banner          []BannerRes              `json:"banner"`
	Notice          NoticeRes                `json:"notice"`
	NewMemberNum    int64                    `json:"newMemberNum"`
	AddActivityNum  int64                    `json:"addActivityNum"`
	DoneActivityNum int64                    `json:"doneActivityNum"`
}

type BannerRes struct {
	ImgUrl      string `json:"img_url"`
	NavigateUrl string `json:"navigate_url"`
}

type NoticeRes struct {
	Title string `json:"title"`
}
