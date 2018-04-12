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
func (d *Article)GetClassify() []Entity.Article {
	var articleList []Entity.Article
	orm.GroupBy("classify").Find(&articleList)
	return articleList
}
func (d *Article)Count() map[string]int64 {
	var articleList []Entity.Article
	var article Entity.Article
	var m map[string]int64
	m = make(map[string]int64)
	orm.GroupBy("classify").Find(&articleList)
	for _,a := range articleList {
		i,_ := orm.Where("classify=?",a.Classify).Count(article)
		m[a.Classify] = i
	}
	return m
}
func (d *Article)FindByClassify(classify string) []Entity.Article {
	var articleList []Entity.Article
	orm.Where("classify=?",classify).Find(&articleList)
	return articleList
}