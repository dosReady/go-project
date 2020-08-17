package dao

import (
	"github.com/dlog/dto"
)

//GetPostList export
func GetPostList(category string) interface{} {
	db := Setup()
	defer db.Close()

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
			post_key
			, post_title
			, post_sub_title
			, post_content
			, TO_CHAR(created_at ,'YYYYMMDD') AS created_at
			, TO_CHAR(updated_at ,'YYYYMMDD') AS updated_at
		FROM tb_posts
		WHERE post_category = ?
	`, category).Find(&list)

	return list
}

//GetPost export
func GetPost(postkey string) interface{} {
	db := Setup()
	defer db.Close()

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
			post_key
			, post_title
			, post_sub_title
			, post_content
			, TO_CHAR(created_at ,'YYYYMMDD') AS created_at
			, TO_CHAR(updated_at ,'YYYYMMDD') AS updated_at
		FROM tb_posts
		WHERE post_key = ?
	`, postkey).Find(&post)

	return post
}

//AddPost export
func AddPost(post dto.PostInDTO) {
	db := Setup()
	defer db.Close()

	data := TbPost{
		PostKey:      post.PostKey,
		PostTitle:    post.PostTitle,
		PostContent:  post.PostContent,
		PostCategory: post.PostCategory,
	}
	db.Create(&data)
	db.NewRecord(&data)
}

//RemovePost export
func RemovePost(postkey string) {
	db := Setup()
	defer db.Close()

	db.Delete(TbPost{PostKey: postkey})
}
