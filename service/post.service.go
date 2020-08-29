package service

import (
	"github.com/dlog/dao"
	"github.com/dlog/dto"
)

//GetPostList export
func GetPostList() interface{} {
	session := dao.Setup(false)
	defer session.Close()
	return session.GetPostList()
}

//GetPost export
func GetPost(postkey string) (interface{}, interface{}) {
	session := dao.Setup(false)
	defer session.Close()
	return session.GetPost(postkey), session.SrchPostBindTags(postkey)
}

//RemovePost export
func RemovePost(postkey string) {
	session := dao.Setup(false)
	defer session.Close()
	session.RemovePost(postkey)
}

//InputPost export
func InputPost(post dto.PostInDTO) {
	session := dao.Setup(true)
	defer session.Close()

	var postkey string
	if len(post.PostKey) > 0 {
		postkey = post.PostKey
		session.UpdPost(post)
	} else {
		postkey = session.AddPost(post)
	}

	var tagkey string
	session.DelTagMapByPostKey(postkey)
	for _, tagName := range post.Tags {
		tagkey = session.GetTagKey(tagName)
		if len(tagkey) == 0 {
			tagkey = session.AddTag(tagName)
		}

		session.AddTagMap(postkey, tagkey)
	}
}
