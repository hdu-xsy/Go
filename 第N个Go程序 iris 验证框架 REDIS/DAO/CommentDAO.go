package DAO
import(
	"../Entity"
)

type Comment struct {

	}
func (d *Comment)FindAll(article string) []Entity.Comment {
	var commnetList []Entity.Comment
	orm.Where("article=?",article).Find(&commnetList)
	return commnetList
}
func (d *Comment)Insert(comment Entity.Comment) (int64,error) {
	i,err := orm.Insert(&comment)
	return i,err
}
