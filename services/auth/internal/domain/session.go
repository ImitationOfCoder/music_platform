package domain

type Session struct {
	Id     int64
	UserId int64
	Token  string
}

func NewSession(id int64, userId int64, token string) Session {
	return Session{
		Id:     id,
		UserId: userId,
		Token:  token,
	}
}
