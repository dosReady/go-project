package service

import (
	"time"

	"github.com/dlog/core"
	"github.com/dlog/dao"
)

// InstPost export
func InstPost(p core.INPostInfo) {
	db := dao.Setup()
	defer db.Close()

	var ctgID uint32
	rs := GetCategoryForTitle(p.CategoryJSON.CtgTitle)
	ctgID = rs.CtgID
	if ctgID == 0 {
		category := core.TbCategory{
			CtgTitle: p.CategoryJSON.CtgTitle,
		}

		db.Create(&category)
		db.NewRecord(category)

		ctgID = category.CtgID
	}

	post := core.TbPost{
		MainTitle: p.PostJSON.MainTitle,
		Content:   p.PostJSON.Content,
		CtgID:     ctgID,
	}

	db.Create(&post)
	db.NewRecord(post)

}

// UpdPost export
func UpdPost(p core.INPostInfo) {
	db := dao.Setup()
	defer db.Close()

	var ctgID uint32
	if rs := GetCategoryForTitle(p.CategoryJSON.CtgTitle); rs.CtgID > 0 {
		ctgID = rs.CtgID
	} else {
		ctgID = InsertCategory(p.CategoryJSON.CtgTitle)
	}

	db.Table("tb_posts").
		Where("tb_posts.post_id = ?", p.PostJSON.PostID).
		Updates(core.TbPost{MainTitle: p.PostJSON.MainTitle,
			Content: p.PostJSON.Content, CtgID: ctgID, CommonModel: core.CommonModel{UpdatedAt: time.Now()}})

}

// GetPost export
func GetPost(p core.INPostInfo) core.OUTPostInfo {
	db := dao.Setup()
	defer db.Close()

	var rs struct {
		PostID    uint32
		MainTitle string
		Content   string
		CtgID     uint32
		CtgTitle  string
	}

	db.Select(`t1.post_id,
	t1.main_title,
	t1."content",
	t1.ctg_id,
	t2.ctg_title`).
		Table("tb_posts t1").
		Joins("LEFT OUTER JOIN tb_categories as t2 ON t1.ctg_id = t2.ctg_id").
		Where("t1.post_id = ?", p.PostJSON.PostID).Scan(&rs)

	postInfo := core.OUTPostInfo{
		TbPost: core.TbPost{
			PostID:    rs.PostID,
			MainTitle: rs.MainTitle,
			Content:   rs.Content,
		},
		TbCategory: core.TbCategory{
			CtgID:    rs.CtgID,
			CtgTitle: rs.CtgTitle,
		},
	}

	return postInfo
}

// GetPostList export
func GetPostList(p core.INPostInfo) []core.OUTPostInfo {
	db := dao.Setup()
	defer db.Close()

	db = db.Select(`t1.post_id, t1.main_title, t1."content", t1.ctg_id, t2.ctg_title, t1.updated_at`).
		Table("tb_posts t1").
		Joins("LEFT OUTER JOIN tb_categories as t2 ON t1.ctg_id = t2.ctg_id").
		Order("updated_at DESC")

	if p.CategoryJSON.CtgID > 0 {
		db = db.Where("t1.ctg_id = ?", p.CategoryJSON.CtgID)
	}

	rs, _ := db.Rows()

	var list []core.OUTPostInfo
	for rs.Next() {

		var object struct {
			PostID    uint32
			MainTitle string
			Content   string
			CtgID     uint32
			CtgTitle  string
			UpdatedAt time.Time
		}

		if err := db.ScanRows(rs, &object); err != nil {
			panic(err)
		}

		item := core.OUTPostInfo{
			TbPost: core.TbPost{
				PostID:      object.PostID,
				MainTitle:   object.MainTitle,
				Content:     object.Content,
				CommonModel: core.CommonModel{UpdatedAt: object.UpdatedAt},
			},
			TbCategory: core.TbCategory{
				CtgID:    object.CtgID,
				CtgTitle: object.CtgTitle,
			},
		}
		list = append(list, item)
	}

	return list
}
