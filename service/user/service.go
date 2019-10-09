package user

type Service interface {
	Get(uid int) int
}

type UserService struct {
}

func (s UserService) Get(uid int) int {
	return uid
}