package service

import "github.com/ijkbytes/mega/model"

type articleService struct {
}

var Article = &articleService{}

func (srv *articleService) ExistArticleById(id int) bool {
	var article model.Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.Id > 0 {
		return true
	}

	return false
}

func (srv *articleService) GetArticleTotal(maps interface{}) (count int) {
	db.Model(&model.Article{}).Where(maps).Count(&count)
	return
}

func (srv *articleService) GetArticles(offset int, limit int, maps interface{}) (articles []model.Article) {
	db.Preload("Tag").Where(maps).Offset(offset).Limit(limit).Find(&articles)
	return
}

func (srv *articleService) GetArticle(id int) (article *model.Article) {
	db.Where("id = ?", id).First(article)
	db.Model(&article).Related(article.Tag)

	return
}

func (srv *articleService) EditArticle(id int, data interface{}) bool {
	db.Model(&model.Article{}).Where("id = ?", id).Updates(data)
	return true
}

func (srv *articleService) AddArticle(data map[string]interface{}) bool {
	db.Create(&model.Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		ContentMD: data["content_md"].(string),
		State:     data["state"].(int),
	})

	return true
}

func (srv *articleService) DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(model.Article{})

	return true
}
