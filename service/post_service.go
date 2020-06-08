package service

import (
	"time"

	"github.com/dlog/core"
	"github.com/dlog/dao"
)

// InstPost export
func InstPost(p core.PostDTO) {
	db := dao.Setup().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		} else {
			db.Commit()
		}

		db.Close()
	}()

	post := core.TbPost{
		MainTitle: p.PostJSON.MainTitle,
		SubTitle:  p.PostJSON.SubTitle,
		Content:   p.PostJSON.Content,
	}

	db.Create(&post)
	db.NewRecord(post)

	InputTag(post.PostID, p.TagJSON, db)

	if err := db.Error; err != nil {
		panic(err)
	}

}

// UpdPost export
func UpdPost(p core.PostDTO) {
	db := dao.Setup().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		} else {
			db.Commit()
		}

		db.Close()
	}()

	var post = core.TbPost{
		PostID:    p.PostJSON.PostID,
		MainTitle: p.PostJSON.MainTitle,
		SubTitle:  p.PostJSON.SubTitle,
		Content:   p.PostJSON.Content,
		CommonModel: core.CommonModel{
			UpdatedAt: time.Now(),
		},
	}
	db.Save(&post)
	DelTagMaps(post.PostID)
	InputTag(post.PostID, p.TagJSON, db)

	if err := db.Error; err != nil {
		panic(err)
	}
}

// GetPost : Post 상세 가져오기
func GetPost(p core.PostDTO) (post core.TbPost, tags []core.TbTagMst) {
	db := dao.Setup()
	defer db.Close()

	db.Where("post_id = ?", p.PostID).First(&post)

	tags = GetTagsMap(p.PostID)

	return post, tags
}

// GetPostList : Post 목록 가져오기
func GetPostList(p core.PostDTO) (postList []core.TbPost) {
	db := dao.Setup()
	defer db.Close()

	db.Find(&postList)
	return postList
}
