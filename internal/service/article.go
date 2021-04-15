package service

import "my-blog-service-go/internal/model"

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
	Title string `form:"title" binding:"required, min=2, max=100"`
	Desc string `form:"desc" binding:"required, min=2, max=255"`
	Content string `form:"content" binding:"required, min=2, max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required, url"`
	CreatedBy string `form:"created_by" binding:"required, min=2, max=100"`
	State uint8 `form:"state, default=1" binding:"oneof=0 1"`
}


type UpdateArticleRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
}

type Article struct {
	ID uint32 `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
	Tag *model.Tag `json:"tag"`
}

func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTag, err := svc.dao.GetArticleTagByAId(article.ID)
	if err != nil {
		return nil, err	
	}
	tag, err := svc.dao.GetTag(articleTag.TagID, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}	

	return &Article{
		ID: article.ID,
		Title: article.Title,
		Desc: article.Desc,
		Content: article.Content,
		CoverImageUrl: article.CoverImageUrl,
		Tag: &tag,
	}, nil
}

// 通过标号删除article 和它的tag
func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		return err	
	}

	err = svc.dao.DeleteArticleTag(param.ID) 
	if err != nil {
		return err
	}
	return nil
}