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
func (d Article)Insert(article Entity.Article) (int64,error) {
	i,err := orm.Insert(&article)
	return i,err
}