package protocols

type UserListReq struct {
	Params struct {
		Page       int    `json:"page" form:"page"`
		RetailerId string `json:"retailerId" form:"retailerId"`
		Stage      string `json:"stage" form:"stage"`
		BeginTime  string `json:"begin_time" form:"begin_time"`
		EndTime    string `json:"end_time" form:"end_time"`
	} `json:"params" form:"params"`
}

type UserListRes struct {
	Total  int64                    `json:"total"`
	Result []map[string]interface{} `json:"result"`
}
