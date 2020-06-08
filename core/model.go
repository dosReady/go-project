package core

import "time"

// CommonModel : 공통 테이블 모델
type CommonModel struct {
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

// TbPost : Post 테이블 모델
type TbPost struct {
	PostID    string `gorm:"type:bigserial;primary_key;auto_increment"`
	MainTitle string `gorm:"type:varchar(100);not null;index:tb_post_idx1"`
	SubTitle  string `gorm:"type:varchar(100);not null;index:tb_post_idx2"`
	Content   string `gorm:"type:text;not null;"`
	CommonModel
}

// TbTagMst : Tag 마스터 테이블 모델
type TbTagMst struct {
	TagMstID string `gorm:"type:bigserial;primary_key;auto_increment"`
	TagName  string `gorm:"type:varchar(100);not null;index:tb_tag_mst_idx1"`
	CommonModel
}

// TbTagMap : Tag Post Mapping 테이블 모델
type TbTagMap struct {
	PostID   string `gorm:"type:bigserial;not null"`
	TagMstID string `gorm:"type:bigserial;not null"`
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

// UserJSON :
type UserJSON struct {
	LoginID      string `json:"LoginID"`
	Password     string `json:"Password"`
	Role         string `json:"Role"`
	RefreshToken string `json:"RefreshToken"`
	AccessToken  string `json:"AccessToken"`
}

// UserInDTO :
type UserInDTO struct {
	UserJSON `json:"user"`
}

// UserOutDTO :
type UserOutDTO struct {
	LoginID      string `json:"LoginID"`
	Role         string `json:"Role"`
	RefreshToken string `json:"RefreshToken"`
	AccessToken  string `json:"AccessToken"`
}

// PostJSON export
type PostJSON struct {
	PostID    string `json:"PostID"`
	MainTitle string `json:"MainTitle"`
	SubTitle  string `json:"SubTitle"`
	Content   string `json:"Content"`
}

// TagJSON :
type TagJSON struct {
	TagMstID string `json:"TagMstID"`
	TagName  string `json:"TagName"`
}

/* ========================= DTO  ===========================*/

// PostDTO : Post 입력 DTO
type PostDTO struct {
	PostJSON `json:"post"`
	TagJSON  []TagJSON `json:"tags"`
}
