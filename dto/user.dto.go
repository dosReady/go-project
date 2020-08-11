package dto

// UserInDTO export
type UserInDTO struct {
	User struct {
		LoginID     string `json:"LoginID"`
		Password    string `json:"Password"`
		Role        string `json:"Role"`
		AccessToken string `json:"AccessToken"`
	} `json:"user"`
}

// UserOutDTO export
type UserOutDTO struct {
	LoginID     string `json:"LoginID"`
	Role        string `json:"Role"`
	AccessToken string `json:"AccessToken"`
}

// UserRsDTO export
type UserRsDTO struct {
	LoginID     string
	Role        string
	AccessToken string
}
