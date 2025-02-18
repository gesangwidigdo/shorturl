package dto

// UrlCreateRequest is a struct to store request body for creating short URL
type UrlCreateRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,url"`
	ShortUrl    string `json:"short_url"`
}

type UrlCreateResponse struct {
	ShortUrl string `json:"short_url"`
}