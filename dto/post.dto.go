package dto

// PostInDTO export
type PostInDTO struct {
	PostKey      string `json:"PostKey"`
	PostTitle    string `json:"PostTitle"`
	PostSubTitle string `json:"PostSubTitle"`
	PostContent  string `json:"PostContent"`
	PostCategory string `json:"PostCategory"`
}
