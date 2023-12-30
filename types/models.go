package types

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ApiToken struct {
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}
