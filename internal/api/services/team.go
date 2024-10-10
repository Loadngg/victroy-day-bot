package services

import (
	"encoding/json"

	"victory-day-bot/api/routes"
	"victory-day-bot/api/types"
)

func (s *Service) CreateTeam(newTeam *types.Team) (*types.Team, error) {
	body, err := s.req.Post(routes.TeamApiUrl, newTeam)
	if err != nil {
		return nil, err
	}

	var createdTeam types.Team
	err = json.Unmarshal(body, &createdTeam)
	if err != nil {
		return nil, err
	}

	return &createdTeam, nil
}
