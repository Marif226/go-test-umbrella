package service

import "time"

type Service struct {
	
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int {
	days := time.Until(time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24
	return int(days)
}