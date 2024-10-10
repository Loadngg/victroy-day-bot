package services

import (
	"encoding/json"
	"fmt"

	"victory-day-bot/api/routes"
	"victory-day-bot/api/types"
)

func (s *Service) CreateAnswer(newAnswer *types.TeamPlaceAnswer) (*types.TeamPlaceAnswer, error) {
	body, err := s.req.Post(routes.TeamPlaceAnswerApiUrl, newAnswer)
	if err != nil {
		return nil, err
	}

	var createdAnswer types.TeamPlaceAnswer
	err = json.Unmarshal(body, &createdAnswer)
	if err != nil {
		return nil, err
	}

	return &createdAnswer, nil
}

func (s *Service) UpdateAnswer(updatedAnswer *types.TeamPlaceAnswer) (*types.TeamPlaceAnswer, error) {
	body, err := s.req.Put(fmt.Sprintf("%s/%d", routes.TeamPlaceAnswerApiUrl, updatedAnswer.Id), updatedAnswer)
	if err != nil {
		return nil, err
	}

	var updatedTeamPlaceAnswer types.TeamPlaceAnswer
	err = json.Unmarshal(body, &updatedTeamPlaceAnswer)
	if err != nil {
		return nil, err
	}

	return &updatedTeamPlaceAnswer, nil
}
