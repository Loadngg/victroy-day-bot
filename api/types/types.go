package types

type Path struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Region int    `json:"region,omitempty"`
	Places []int  `json:"places"`
}

type Place struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Coordinates    string `json:"coordinates"`
	TextRiddle     string `json:"text_riddle,omitempty"`
	VideoRiddleUrl string `json:"video_riddle_url"`
	VideoStoryUrl  string `json:"video_story_url"`
	TaskFile       string `json:"task_file"`
}

type Region struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Team struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Contacts string `json:"contacts"`
	Path     int    `json:"path"`
}

type TeamPlaceAnswer struct {
	Id            int    `json:"id"`
	StartDatetime string `json:"start_datetime,omitempty"`
	EndDatetime   string `json:"end_datetime,omitempty"`
	Photo         string `json:"photo,omitempty"`
	TaskAnswer    string `json:"task_answer,omitempty"`
	Team          int    `json:"team"`
	Place         int    `json:"place"`
}
