package core

import "time"

// CommonModel : 공통 테이블 모델
type CommonModel struct {
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

// TbPost : Post 테이블 모델
type TbPost struct {
	PostID    uint32 `gorm:"primary_key;auto_increment"`
	MainTitle string `gorm:"type:varchar(100);not null;index:tb_post_idx1"`
	Content   string `gorm:"type:text;not null;"`
	CtgID     uint32
	CommonModel
}

// TbCategory : 카테고리 테이블 모델
type TbCategory struct {
	CtgID    uint32 `gorm:"primary_key;auto_increment"`
	CtgTitle string `gorm:"type:varchar(100);not null;index:tb_tag_idx1"`
	CtgAlias string `gorm:"type:varchar(255);not null;index:tb_tag_idx2"`
	CommonModel
}

// TbUser : 유저 테이블 모델
type TbUser struct {
	LoginID      string `gorm:"type:varchar(100);primary_key:auto_increment"`
	Password     string `gorm:"type:varchar(100);not null;"`
	Role         string `gorm:"varchar(100);not null;"`
	RefreshToken string `gorm:"text;"`
	CommonModel
}

type UserJSON struct {
	LoginID      string `json:"LoginID"`
	Password     string `json:"Password"`
	Role         string `json:"Role"`
	RefreshToken string `json:"RefreshToken"`
	AccessToken  string `json:"AccessToken"`
}

type UserInDTO struct {
	UserJSON `json:"user"`
}

type UserOutDTO struct {
	LoginID      string `json:"LoginID"`
	Role         string `json:"Role"`
	RefreshToken string `json:"RefreshToken"`
	AccessToken  string `json:"AccessToken"`
}

// PostJSON export
type PostJSON struct {
	PostID    uint32 `json:"PostID"`
	MainTitle string `json:"MainTitle"`
	SubTitle  string `json:"SubTitle"`
	Content   string `json:"Content"`
}

// CategoryJSON export
type CategoryJSON struct {
	CtgID    uint32 `json:"CtgID,string"`
	CtgTitle string `json:"CtgTitle"`
	CtgAlias string `json:"CtgAlias"`
	CtgCnt   uint32 `json:"CtgCnt"`
}

// RsPostInfo export
type RsPostInfo struct {
	PostID    uint32
	MainTitle string
	CtgID     uint32
	CtgTitle  string
}

// OUTPostInfo : QUERY 결과 VO
type OUTPostInfo struct {
	TbPost     `json:"post"`
	TbCategory `json:"ctg"`
}

// INPostInfo : 입력 VO
type INPostInfo struct {
	PostJSON     `json:"info"`
	CategoryJSON `json:"category"`
}
