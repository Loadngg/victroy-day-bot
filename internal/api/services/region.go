package services

import (
	"encoding/json"

	"victory-day-bot/api/routes"
	"victory-day-bot/api/types"
)

func (s *Service) GetRegions() (*[]types.Region, error) {
	body, err := s.req.Get(routes.RegionApiUrl)
	if err != nil {
		return nil, err
	}

	var regions []types.Region
	err = json.Unmarshal(body, &regions)
	if err != nil {
		return nil, err
	}

	return &regions, nil
}
