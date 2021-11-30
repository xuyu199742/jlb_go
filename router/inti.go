package router

type Group struct {
	HomeRouter
	StaffRouter
	NoticeRouter
	UserRouter
}

var ApiGroup = new(Group)
