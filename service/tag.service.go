package service

import (
	"github.com/dlog/dao"
	"github.com/dlog/dto"
)

//Service - GetTagList  export
func GetTagList(param dto.TagDTO) interface{} {
	session := dao.Setup(false)
	defer session.Close()

	return session.GetTagList(param)
}
