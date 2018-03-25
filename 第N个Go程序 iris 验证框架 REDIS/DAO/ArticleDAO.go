package DAO

import(
	"../Entity"
)

type Article struct {

}

func (d *Article)Get(article Entity.Article) (bool,error,Entity.Article) {
	bo,err := orm.Get(&article)
	return bo,err,article
}
func (d *Article)FindAll(id string) []Entity.Article {
	var articleList []Entity.Article
	orm.Where("menu=?",id).Find(&articleList)
	return articleList
}
func (d *Article)Insert(article Entity.Article) (int64,error) {
	i,err := orm.Insert(&article)
	return i,err
}
func (d *Article)OrderByTime() []Entity.Article{
	var articleList []Entity.Article
	orm.Desc("time").Limit(20,0).Find(&articleList)
	return articleList
}
func (d *Article)FindAllA() []Entity.Article{
	var articleList []Entity.Article
	orm.Find(&articleList)
	return articleList
}
func (d *Article)Update(article Entity.Article) (int64,error) {
	i,err := orm.Id(article.Id).Update(&article)
	return i,err
}