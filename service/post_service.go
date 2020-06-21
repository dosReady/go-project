package service

import (
	"log"
	"strings"
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

// DelPost : Post 삭제
func DelPost(p core.PostDTO) {
	db := dao.Setup().Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			db.Rollback()
		} else {
			db.Commit()
		}

		db.Close()
	}()

	var models []core.TbTagMap
	db.Where(core.TbTagMap{PostID: p.PostID}).Find(&models)

	var mstIDarr []string
	for i := 0; i < len(models); i++ {
		mstIDarr = append(mstIDarr, models[i].TagMstID)
	}

	if len(mstIDarr) > 0 {
		db.Where("tag_mst_id in(?)", mstIDarr).Delete(core.TbTagMap{})
	}

	db.Where("post_id = ?", p.PostID).Delete(core.TbPost{})
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
func GetPostList(p core.PostDTO) interface{} {
	db := dao.Setup()
	defer db.Close()

	var post []struct {
		PostID    string
		MainTitle string
		SubTitle  string
		Content   string
		TagsJSON  string
		UpdatedAt time.Time
	}
	db = db.Select(`
			t1.post_id
		,	t1.main_title
		,	t1.sub_title
		,	t1.content
		,	t1.updated_at
		,	array_to_json(array_agg(coalesce(t3.tag_name, ''))) as tags_json
	`).Table("tb_posts t1").
		Joins("left outer join tb_tag_maps t2 on t1.post_id = t2.post_id").
		Joins("left outer join tb_tag_msts t3 on t2.tag_mst_id = t3.tag_mst_id").
		Group(`t1.post_id
	,	t1.main_title
	,	t1.sub_title
	,	t1.content
	,	t1.updated_at`).
		Order("t1.created_at desc")

	if len(p.MainTitle) > 0 {
		if strings.Index(p.MainTitle, "#") == 0 {
			db = db.Where(`
				t3.tag_name = ? 
			`, strings.Replace(p.MainTitle, "#", "", 1))
		} else {
			db = db.Where(`t1.main_title like '%'||?||'%' 
			or t1.sub_title like '%'||?||'%'
			or t1.content like '%'||?||'%'
			`, p.MainTitle, p.SubTitle, p.Content)
		}

	}
	db.Find(&post)
	return post
}
