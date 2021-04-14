package service

type ArticleRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"oneof=0 1"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}
type CreateArticleResult struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	
}


