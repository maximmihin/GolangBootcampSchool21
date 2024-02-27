package models

import (
	"ex00/entities"
)

type PlacesModel struct {
	Name     string           `json:"name"`
	Total    int              `json:"total"`
	Places   []entities.Place `json:"places"`
	PrevPage int              `json:"prev_page"`
	NextPage int              `json:"next_page"`
	LastPage int              `json:"last_page"`
}

type TokenModel struct {
	Token string `json:"tokenModel"`
}

type RecommendedPlacesModel struct {
	Name   string           `json:"name"`
	Places []entities.Place `json:"places"`
}
