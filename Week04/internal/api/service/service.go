package service

import "github.com/B8kingS0ga/Go-000/tree/main/Week04/internal/api/dao"

type Service struct {
	Dao dao.Dao
}

func (s *Service) ServiceSome() string {
	return s.Dao.GetUser()
}

func NewService(d dao.Dao) Service {
	return Service{
		Dao: d,
	}
}
