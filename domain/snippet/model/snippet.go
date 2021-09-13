package snippetmodel

import "time"

type Snippet struct {
	Name      string    `json:"name"`
	ExpiresAt time.Time `json:"expires_at"`
	Content   string    `json:"snippet"`
	URL       string    `json:"url"`
	Likes     int       `json:"likes"`
}

func NewSnippet(name, content, url string, expiresAt time.Time) *Snippet {
	return &Snippet{
		Name:      name,
		ExpiresAt: expiresAt,
		Content:   content,
		URL:       url,
	}
}
