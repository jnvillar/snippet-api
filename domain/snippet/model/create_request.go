package snippetmodel

type CreateRequest struct {
	Name      string `json:"name"`
	ExpiresIn int    `json:"expires_in"`
	Content   string `json:"snippet"`
}
