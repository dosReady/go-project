package service

import (
	"github.com/dlog/dao"
	"github.com/dlog/dto"
)

//GetPostList export
func GetPostList(category string) interface{} {
	return dao.GetPostList(category)
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

//InputPost export
func InputPost(post dto.PostInDTO) {
	if len(post.PostKey) > 0 {
		dao.UpdPost(post)
	} else {
		dao.AddPost(post)
	}
}
