package svc

// Service Define a service interface
type Service interface {
	Get(uid int) int
}

type UserService struct {
}

// Add implement Add method
func (s UserService) Get(uid int) int {
	return uid
}
