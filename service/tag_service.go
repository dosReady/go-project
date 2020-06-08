package service

import (
	"strconv"

	"github.com/dlog/core"
	"github.com/dlog/dao"
	"github.com/jinzhu/gorm"
)

// GetTagsMap : Tag Map 정보 가져오기
func GetTagsMap(postID string) (tags []core.TbTagMst) {
	db := dao.Setup()
	defer db.Close()

	db.Select("t1.*").Table("tb_tag_msts t1").Joins("inner join tb_tag_maps t2 on t1.tag_mst_id = t2.tag_mst_id ").
		Where("t2.post_id =? ", postID).Find(&tags)
	return tags
}

// InputTag : 태그 입력
func InputTag(postID string, tags []core.TagJSON, db *gorm.DB) {
	for i := 0; i < len(tags); i++ {
		tagMst := core.TbTagMst{}
		if len(tags[i].TagName) > 0 {
			db.Where(core.TbTagMst{TagName: tags[i].TagName}).First(&tagMst)
			tagMstID, _ := strconv.ParseInt(tagMst.TagMstID, 10, 32)
			if tagMstID <= 0 {
				tagMst = core.TbTagMst{
					TagName: tags[i].TagName,
				}
				db.Create(&tagMst)
				db.NewRecord(tagMst)
			}

			tbMap := core.TbTagMap{
				PostID:   postID,
				TagMstID: tagMst.TagMstID,
			}

			var count int = 0
			db.Model(tbMap).Where(tbMap).Count(&count)
			if count == 0 {
				db.Create(&tbMap)
				db.NewRecord(tbMap)
			}
		}
	}
}

// DelTagMaps :
func DelTagMaps(paramPostID string) {
	db := dao.Setup()
	defer db.Close()

	tagMaps := core.TbTagMap{
		PostID: paramPostID,
	}
	db.Where(tagMaps).Delete(tagMaps)
}
