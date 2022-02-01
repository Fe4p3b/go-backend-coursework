package models

type URL struct {
	URL      string `json:"original_url,omitempty"`
	ShortURL string `json:"short_url,omitempty"`
}
