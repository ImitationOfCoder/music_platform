package domain

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CurrentUser struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
