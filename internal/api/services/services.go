package services

import (
	"victory-day-bot/internal/api/requests"
)

type Service struct {
	req *requests.Requests
}

func New() *Service {
	return &Service{
		req: requests.New(),
	}
}
