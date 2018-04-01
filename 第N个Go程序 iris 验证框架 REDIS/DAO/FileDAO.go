package DAO

import "../Entity"
type FileDAO struct {

}

func (d *FileDAO) GetAll() []Entity.File {
	var fileList []Entity.File
	orm.Find(&fileList)
	return fileList
}
func (d *FileDAO) Insert(file Entity.File) (int64,error) {
	i,err :=orm.Insert(&file)
	return i,err
}

