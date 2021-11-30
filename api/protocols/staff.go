package protocols

type StaffLoginReq struct {
	LoginAccount  string `json:"login_account" form:"login_account"`
	Password      string `json:"password" form:"password"`
	WxCode        string `json:"wx_code" form:"wx_code"`
	OpenId        string `json:"open_id" form:"open_id"`
	UnionId       string `json:"union_id" form:"union_id"`
	IV            string `json:"iv" form:"iv"`
	EncryptedData string `json:"encryptedData" form:"encryptedData"`
}

type StaffLoginRes struct {
	AccessToken string `json:"access_token"`
}

type StaffProfileReq struct {
	Params struct {
		RetailerId int `json:"retailerId" form:"retailerId"`
	} `json:"params" form:"params"`
}

type StaffProfileRes struct {
	Profile struct {
		Id             int    `json:"id"`
		StaffCode      string `json:"staffCode"`
		SelectShopId   string `json:"selectShopId"`
		SelectShopName string `json:"selectShopName"`
	} `json:"profile"`
	Shops []map[string]interface{} `json:"shops"`
}
