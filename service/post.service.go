package service

import (
	"github.com/dlog/dao"
	"github.com/dlog/dto"
)

//GetPostList export
func GetPostList(category string) interface{} {
	session := dao.Setup(false)
	defer session.Close()
	return session.GetPostList(category)
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

	for _, item := range post.Tags {
		if len(item.TagKey) > 0 && item.IsDel == "Y" {
			session.DelTagMap(item.TagKey)
		} else {
			chkTagKey, chkPostKey := session.SrchTagMapByName(item.TagName)
			if len(chkTagKey) == 0 {
				chkTagKey = session.AddTag(item.TagName)
			}

			if len(chkPostKey) == 0 && len(chkTagKey) > 0 {
				session.AddTagMap(postkey, chkTagKey)
			}
		}
	}
}
