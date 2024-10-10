package services

import (
	"encoding/json"
	"fmt"

	"victory-day-bot/api/routes"
	"victory-day-bot/api/types"
)

func (s *Service) GetPlaceById(id int) (*types.Place, error) {
	body, err := s.req.Get(fmt.Sprintf("%s/%d", routes.PlaceApiUrl, id))
	if err != nil {
		return nil, err
	}

	var place types.Place
	err = json.Unmarshal(body, &place)
	if err != nil {
		return nil, err
	}

	return &place, nil
}
