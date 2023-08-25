package models

type ParamVoteData struct {
	// userid从请求中获取
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1)反对票(-1)取消投票(0)
}
