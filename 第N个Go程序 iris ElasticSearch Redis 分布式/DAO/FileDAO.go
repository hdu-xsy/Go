package DAO

import "../Entity"
type File struct {

}

func (d *File) GetAll() []Entity.File {
	var fileList []Entity.File
	orm.Find(&fileList)
	return fileList
}
func (d *File) Insert(file Entity.File) (int64,error) {
	i,err :=orm.Insert(&file)
	return i,err
}
func (d *File) GetName(id int64) string {
	var file Entity.File
	orm.Id(id).Get(&file)
	return file.Name
}
