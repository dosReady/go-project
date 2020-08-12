package dao

import (
	"github.com/dlog/dto"
)

//GetPostList export
func GetPostList() interface{} {
	db := Setup()
	defer db.Close()

	var list []struct {
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
	`).Find(&list)

	return list
}

//AddPost export
func AddPost(post dto.PostInDTO) {
	db := Setup()
	defer db.Close()

	data := TbPost{
		PostKey:     post.PostKey,
		PostTitle:   post.PostTitle,
		PostContent: post.PostContent,
	}
	db.Create(&data)
	db.NewRecord(&data)
}
