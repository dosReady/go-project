package service

import (
	"log"

	"github.com/dlog/core"
	"github.com/dlog/dao"
)

// GetCategoryList export : PostID로 해당 카테고리를 조회한다.
func GetCategoryList(target string) []core.CategoryJSON {
	db := dao.Setup()
	defer db.Close()

	var categoryList []core.CategoryJSON
	db.Select("t1.ctg_id, t1.ctg_title, count(t2.post_id) as ctg_cnt").
		Table("tb_categories t1").
		Joins("inner join tb_posts t2 on t1.ctg_id = t2.ctg_id").
		Group("t1.ctg_id, t1.ctg_title").
		Order("count(t2.post_id) desc").Find(&categoryList).Limit(10)

	return categoryList
}

// GetCategory export : PostID로 해당 카테고리를 조회한다.
func GetCategory(postID uint32) core.TbCategory {
	db := dao.Setup()
	defer db.Close()

	result := core.TbCategory{}

	db.Table("tb_posts").
		Joins("LEFT OUTER JOIN tb_categories ON tb_posts.CtgID = tb_categories.CtgID").
		Where("tb_posts.post_id = ?", postID).First(&result)

	return result
}

// GetCategoryForTitle export : 카테로고리 이름으로로 해당 카테고리를 조회한다.
func GetCategoryForTitle(title string) core.TbCategory {
	db := dao.Setup()
	defer db.Close()

	result := core.TbCategory{}
	db.Table("tb_categories").Where("ctg_title = TRIM(?)", title).First(&result)

	return result
}

// InsertCategory export : 카테고리 정보 입력
func InsertCategory(title string) uint32 {
	db := dao.Setup()
	defer db.Close()

	category := core.TbCategory{
		CtgTitle: title,
	}

	db.Create(&category)
	db.NewRecord(category)

	log.Println(category.CtgID)
	return category.CtgID
}
