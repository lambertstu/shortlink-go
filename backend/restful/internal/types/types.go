// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type CreateGroupRequest struct {
	Name     string `form:"name"`
	Username string `form:"username"`
}

type CreateGroupResponse struct {
	Gid      string `json:"gid"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type DeleteGroupRequest struct {
	Gid      string `form:"gid"`
	Username string `form:"username"`
}

type DeleteUserRequest struct {
	Username string `form:"username"`
}

type GetGroupRequest struct {
	Username string `form:"username"`
}

type GetGroupResponse struct {
	Data []GroupData `json:"data"`
}

type GetUserRequest struct {
	Username string `form:"username"`
}

type GetUserResponse struct {
	Username string `json:"username"`
	RealName string `json:"realName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type GroupData struct {
	Gid      string `json:"gid"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type IsExistUserRequest struct {
	Username string `form:"username"`
}

type IsUserLoginRequest struct {
	Token    string `form:"token"`
	Username string `form:"username"`
}

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LogoutRequest struct {
	Token    string `form:"token"`
	Username string `form:"username"`
}

type NilResponse struct {
	Nil string `json:"code"`
}

type RegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	RealName string `form:"realName"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
}

type ShortLinkCreateRequest struct {
	OriginUrl string `form:"origin_url"`
	Gid       string `form:"gid"`
	Describe  string `form:"describe"`
}

type ShortLinkCreateResponse struct {
	OriginUrl string `json:"origin_url"`
}

type ShortLinkDeleteRequest struct {
	ShortUri  string `form:"shortUri"`
	OriginUrl string `form:"origin_url"`
}

type ShortLinkPageData struct {
	ShortUri     string `json:"short_uri"`
	FullShortUrl string `json:"full_short_url"`
	OriginUrl    string `json:"origin_url"`
	Gid          string `json:"gid"`
	EnableStatus int32  `json:"enable_status"`
	CreateTime   int64  `json:"create_time"`
	Describe     string `json:"describe"`
	ClickNum     string `json:"click_num"`
	Favicon      string `json:"favicon"`
	TotalPv      int32  `json:"total_pv"`
	TodayPv      int32  `json:"today_pv"`
	TotalUv      int32  `json:"total_uv"`
	TodayUv      int32  `json:"today_uv"`
	TotalUip     int32  `json:"total_uip"`
	TodayUip     int32  `json:"today_uip"`
}

type ShortLinkPageRequest struct {
	Gid      string `form:"gid"`
	Page     int64  `form:"page"`
	Size     int64  `form:"size"`
	OrderTag int64  `form:"orderTag"`
}

type ShortLinkPageResponse struct {
	List    []ShortLinkPageData `json:"list"`
	Page    int64               `json:"page"`
	MaxPage int64               `json:"maxPage"`
	Total   int64               `json:"total"`
}

type ShortLinkRestoreRequest struct {
	ShortUri string `path:"shortUri"`
}

type ShortLinkRestoreResponse struct {
	OriginUrl string `json:"originUrl"`
}

type ShortLinkUpdateRequest struct {
	Gid       string `form:"gid"`
	OriginUrl string `form:"origin_url"`
	ShortUri  string `form:"short_uri"`
	Describe  string `form:"describe"`
	Favicon   string `form:"favicon"`
	ClickNum  int64  `form:"clickNum"`
	TotalPv   int64  `form:"totalPv"`
	TotalUv   int64  `form:"totalUv"`
	TotalUip  int64  `form:"totalUip"`
	TodayPv   int64  `form:"todayPv"`
	TodayUv   int64  `form:"todayUv"`
	TodayUip  int64  `form:"todayUip"`
}

type UpdateGroupRequest struct {
	Gid      string `form:"gid"`
	Name     string `form:"name"`
	Username string `form:"username"`
}

type UpsertUserInfoRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	RealName string `form:"realName"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
}
