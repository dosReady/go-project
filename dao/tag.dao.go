package dao

import "github.com/dlog/dto"

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

//DelTagMap: TagMap 정보 삭제
func (session *Session) DelTagMap(tagkey string) {
	db := session.Db
	db.Delete(TbTagMap{}, "tag_key = ?", tagkey)
}

// SrchNrmlTag : 태그  일반 조회
func (session *Session) SrchNrmlTag(tagname string) (data dto.TagInDTO) {
	db := session.Db

	db.Raw(`
		SELECT 
			  tag_key
			, tag_name
		FROM tb_tags
		WHERE tag_name = ?
	`, tagname).Find(&data)
	return data
}

// SrchTagMapByName: 태그 Map 태그이름으로 조회
func (session *Session) SrchTagMapByName(tagname string) (string, string) {
	db := session.Db

	var data struct {
		TagKey  string
		PostKey string
	}

	db.Raw(`
	SELECT 
	  T1.TAG_KEY 
	, T2.POST_KEY
	FROM TB_TAGS T1
	LEFT OUTER JOIN TB_TAG_MAPS T2 ON T1.TAG_KEY  = T2.TAG_KEY 
	WHERE 1=1
	AND T1.TAG_NAME  = ?
	`, tagname).Find(&data)

	return data.TagKey, data.PostKey
}

// SrchPostBindTags: Post에 바인드된 Tag목록 조회
func (session *Session) SrchPostBindTags(postkey string) interface{} {
	db := session.Db

	var data []struct {
		TagKey  string
		TagName string
	}

	db.Raw(`
	SELECT 
		 T1.TAG_KEY 
		,T1.TAG_NAME 
	FROM TB_TAGS T1
	INNER JOIN TB_TAG_MAPS T2 ON T1.TAG_KEY  = T2.TAG_KEY 
	WHERE 1=1
	AND T2.POST_KEY = ?
	`, postkey).Find(&data)
	return data
}
