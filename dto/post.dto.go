package dto

// PostInDTO export
type PostInDTO struct {
	Value       string `json:"LoginID"`
	Password    string `json:"Password"`
	Role        string `json:"Role"`
	AccessToken string `json:"AccessToken"`
}

// PostOutDTO export
type PostOutDTO struct {
	LoginID     string `json:"LoginID"`
	Role        string `json:"Role"`
	AccessToken string `json:"AccessToken"`
}

// PostRsDto export
type PostRsDto struct {
	LoginID     string
	Role        string
	AccessToken string
}
