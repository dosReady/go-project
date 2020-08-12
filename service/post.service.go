package service

import (
	"github.com/dlog/dao"
	"github.com/dlog/dto"
)

//GetPostList export
func GetPostList() interface{} {
	return dao.GetPostList()
}

//AddPost export
func AddPost(post dto.PostInDTO) {
	dao.AddPost(post)
}
