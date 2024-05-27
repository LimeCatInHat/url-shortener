package models

type ShortenURLRequest struct {
	URL string `json:"url"`
}

type ShortenURLResponse struct {
	ShortLink string `json:"result"`
}
