package dao

import (
	"github.com/dlog/dto"
)

//GetPostList export
func (session *Session) GetPostList(category string) interface{} {
	db := session.Db
	var list []struct {
		PostKey      string
		PostTitle    string
		PostSubTitle string
		PostContent  string
		PostCategory string
		CreatedAt    string
		UpdatedAt    string
	}

	db.Raw(`
		SELECT 
			POST_KEY
			, POST_TITLE
			, POST_SUB_TITLE
			, POST_CONTENT
			, TO_CHAR(CREATED_AT ,'YYYYMMDD') AS CREATED_AT
			, TO_CHAR(UPDATED_AT ,'YYYYMMDD') AS UPDATED_AT
		FROM TB_POSTS
		WHERE POST_CATEGORY = ?
		ORDER BY CREATED_AT DESC
	`, category).Find(&list)

	return list
}

//GetPost export
func (session *Session) GetPost(postkey string) interface{} {
	db := session.Db
	var post struct {
		PostKey      string
		PostTitle    string
		PostSubTitle string
		PostContent  string
		CreatedAt    string
		UpdatedAt    string
	}

	db.Raw(`
		SELECT 
			POST_KEY
			, POST_TITLE
			, POST_SUB_TITLE
			, POST_CONTENT
			, TO_CHAR(CREATED_AT ,'YYYYMMDD') AS CREATED_AT
			, TO_CHAR(UPDATED_AT ,'YYYYMMDD') AS UPDATED_AT
		FROM TB_POSTS
		WHERE POST_KEY = ?
	`, postkey).Find(&post)

	return post
}

//AddPost export
func (session *Session) AddPost(post dto.PostInDTO) string {
	db := session.Db
	data := TbPost{
		PostKey:      post.PostKey,
		PostTitle:    post.PostTitle,
		PostContent:  post.PostContent,
		PostCategory: post.PostCategory,
	}
	db.Create(&data)
	db.NewRecord(&data)

	return data.PostKey
}

//UpdPost export
func (session *Session) UpdPost(post dto.PostInDTO) {
	db := session.Db
	data := TbPost{
		PostTitle:    post.PostTitle,
		PostContent:  post.PostContent,
		PostCategory: post.PostCategory,
	}
	db.Model(TbPost{PostKey: post.PostKey}).Updates(data)
}

//RemovePost export
func (session *Session) RemovePost(postkey string) {
	db := session.Db
	db.Delete(TbPost{PostKey: postkey})
}
