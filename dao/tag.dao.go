package dao

import (
	"log"
)

// AddTag export
func (session *Session) AddTag(param string) string {
	db := session.Db

	tag := TbTag{
		TagName: param,
	}

	db.Create(&tag)
	db.NewRecord(&tag)

	return tag.TagKey
}

// AddTagMap : TagMaps 테이블 입력
func (session *Session) AddTagMap(postkey string, tagkey string) {
	db := session.Db

	tagMap := TbTagMap{
		TagKey:  tagkey,
		PostKey: postkey,
	}

	db.Create(&tagMap)
	db.NewRecord(&tagMap)
}

//DelTagMapByPostKey: PostKey로 조회하여 TagMap 정보 삭제
func (session *Session) DelTagMapByPostKey(postkey string) {
	db := session.Db
	db.Delete(TbTagMap{}, "post_key = ?", postkey)
}

// GetTagKey : TagKey 조회
func (session *Session) GetTagKey(tagname string) (tagkey string) {
	db := session.Db

	var result struct {
		TagKey string
	}

	db.Raw(`
		SELECT 
			TAG_KEY
		FROM TB_TAGS T1
		WHERE TAG_NAME = ?
	`, tagname).Find(&result)
	return result.TagKey
}

// SrchPostBindTags: Post에 바인드된 Tag목록 조회
func (session *Session) SrchPostBindTags(postkey string) (datas []string) {
	db := session.Db

	rows, _ := db.Raw(`
	SELECT 
		T1.TAG_NAME 
	FROM TB_TAGS T1
	INNER JOIN TB_TAG_MAPS T2 ON T1.TAG_KEY  = T2.TAG_KEY 
	WHERE 1=1
	AND T2.POST_KEY = ?
	`, postkey).Rows()

	var rs string

	for rows.Next() {
		if err := rows.Scan(&rs); err != nil {
			log.Panicln(err)
		}
		datas = append(datas, rs)
	}

	return datas
}
