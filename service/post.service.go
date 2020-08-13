package service

import (
	"github.com/dlog/dao"
	"github.com/dlog/dto"
)

//GetPostList export
func GetPostList() interface{} {
	return dao.GetPostList()
}

//GetPost export
func GetPost(postkey string) interface{} {
	return dao.GetPost(postkey)
}

//AddPost export
func AddPost(post dto.PostInDTO) {
	dao.AddPost(post)
}

//AddPost export
func RemovePost(postkey string) {
	dao.RemovePost(postkey)
}
