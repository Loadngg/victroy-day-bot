package services

import (
	"encoding/json"
	"fmt"

	"victory-day-bot/api/routes"
	"victory-day-bot/api/types"
)

func (s *Service) GetPathByRegionId(regionId int) (*types.Path, error) {
	body, err := s.req.Get(fmt.Sprintf("%s/%d", routes.PathApiUrl, regionId))
	if err != nil {
		return nil, err
	}

	var path types.Path
	err = json.Unmarshal(body, &path)
	if err != nil {
		return nil, err
	}

	return &path, nil
}
