package protocols

type NoticeListReq struct {
	Params struct {
		Page int `json:"page" form:"page"`
	} `json:"params" form:"params"`
}

type NoticeListRes struct {
	NoticeList []Notice `json:"notice_list"`
	Page       int      `json:"page"`
	PageSize   int      `json:"pageSize"`
	Total      int64    `json:"total"`
}

type Notice struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	CreateTime string `json:"create_time"`
}

type NoticeDetailRes struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type NoticeDetailReq struct {
	Params struct {
		NoticeId string `json:"notice_id" form:"notice_id"`
	} `json:"params" form:"params"`
}
