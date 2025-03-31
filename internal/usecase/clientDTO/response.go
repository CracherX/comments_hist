package clientDTO

type ProfileResponse struct {
	ID       int    `json:"id"`
	IsAdmin  bool   `json:"isAdmin"`
	Username string `json:"username"`
}
