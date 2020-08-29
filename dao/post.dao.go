package dao

import (
	"github.com/dlog/dto"
)

//GetPostList export
func (session *Session) GetPostList() interface{} {
	db := session.Db
	var list []struct {
		PostKey     string
		PostTitle   string
		PostContent string
		Tags        string
		CreatedAt   string
		UpdatedAt   string
	}

	db.Raw(`
	SELECT 
		T1.POST_KEY
		, T1.POST_TITLE
		, T1.POST_CONTENT
		, STRING_AGG('#'||T3.TAG_NAME, ' ' ORDER BY T3.TAG_NAME) AS TAGS
		, TO_CHAR(T1.CREATED_AT ,'YYYYMMDD') AS CREATED_AT
	FROM TB_POSTS T1
	LEFT OUTER JOIN TB_TAG_MAPS T2 ON T1.POST_KEY = T2.POST_KEY 
	LEFT OUTER JOIN TB_TAGS T3 ON T2.TAG_KEY  = T3.TAG_KEY 
	WHERE 1=1
	GROUP BY T1.POST_KEY, T1.POST_TITLE, T1.POST_CONTENT, T1.CREATED_AT 
	ORDER BY T1.CREATED_AT DESC
	`).Find(&list)
	return list
}

//GetPost export
func (session *Session) GetPost(postkey string) interface{} {
	db := session.Db
	var post struct {
		PostKey     string
		PostTitle   string
		PostContent string
		CreatedAt   string
	}

	db.Raw(`
		SELECT 
			POST_KEY
			, POST_TITLE
			, POST_CONTENT
			, TO_CHAR(CREATED_AT ,'YYYYMMDD') AS CREATED_AT
		FROM TB_POSTS
		WHERE POST_KEY = ?
	`, postkey).Find(&post)

	return post
}

//AddPost export
func (session *Session) AddPost(post dto.PostInDTO) string {
	db := session.Db
	data := TbPost{
		PostKey:     post.PostKey,
		PostTitle:   post.PostTitle,
		PostContent: post.PostContent,
	}
	db.Create(&data)
	db.NewRecord(&data)

	return data.PostKey
}

//UpdPost export
func (session *Session) UpdPost(post dto.PostInDTO) {
	db := session.Db
	data := TbPost{
		PostTitle:   post.PostTitle,
		PostContent: post.PostContent,
	}
	db.Model(TbPost{PostKey: post.PostKey}).Updates(data)
}

//RemovePost export
func (session *Session) RemovePost(postkey string) {
	db := session.Db
	db.Delete(TbPost{PostKey: postkey})
}
