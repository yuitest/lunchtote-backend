package model

import (
	"time"

	"github.com/kaikourok/lunchtote-backend/entity/service"
)

type CharacterOverview struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Mainicon string `json:"mainicon"`
}

type Icon struct {
	Path string `json:"path"`
}

type ProfileImage struct {
	Path string `json:"path"`
}

type ProfileEditData struct {
	Name     string   `json:"name"`
	Nickname string   `json:"nickname"`
	Summary  string   `json:"summary"`
	Profile  string   `json:"profile"`
	Mainicon string   `json:"Mainicon"`
	Tags     []string `json:"tags"`
}

type Profile struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Nickname        string   `json:"nickname"`
	Summary         string   `json:"summary"`
	Profile         string   `json:"profile"`
	ProfileImages   []string `json:"profileImages"`
	Tags            []string `json:"tags"`
	FollowingNumber int      `json:"followingNumber"`
	FollowedNumber  int      `json:"followedNumber"`
	Icons           []Icon   `json:"icons"`
	IsFollowing     bool     `json:"isFollowing"`
	IsFollowed      bool     `json:"isFollowed"`
	IsMuting        bool     `json:"isMuting"`
	IsBlocking      bool     `json:"isBlocking"`
	IsBlocked       bool     `json:"isBlocked"`
}

type CharacterListItem struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Nickname    string   `json:"nickname"`
	Summary     string   `json:"summary"`
	Mainicon    string   `json:"Mainicon"`
	Tags        []string `json:"tags"`
	IsFollowing *bool    `json:"isFollowing,omitempty"`
	IsFollowed  *bool    `json:"isFollowed,omitempty"`
	IsMuting    *bool    `json:"isMuting,omitempty"`
	IsBlocking  *bool    `json:"isBlocking,omitempty"`
}

type UploadedImage struct {
	Id         int       `json:"id"`
	Path       string    `json:"path"`
	UploadedAt time.Time `json:"uploadedAt"`
}

type HomeData struct {
	Nickname string `json:"nickname"`
}

type CharacterSuggestionData struct {
	Id   int
	Name string
}

type CharacterSuggestion struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type CharacterSuggestionsData []CharacterSuggestionData
type CharacterSuggestions []CharacterSuggestion

func (s *CharacterSuggestionData) ToDomain() *CharacterSuggestion {
	return &CharacterSuggestion{
		Id:   s.Id,
		Text: service.ConvertCharacterIdToText(s.Id) + " " + s.Name,
	}
}

func (s *CharacterSuggestionsData) ToDomain() *CharacterSuggestions {
	suggestions := make(CharacterSuggestions, len(*s))
	for i, suggestionData := range *s {
		suggestions[i] = *suggestionData.ToDomain()
	}
	return &suggestions
}

type CharacterEmailRegistratedData struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type CharacterIconLayerItemEditData struct {
	Path string `json:"path"`
}

type CharacterIconLayerItem struct {
	CharacterIconLayerItemEditData
	Id int `json:"id"`
}

type CharacterIconProcessSchema struct {
	Name   string  `json:"name"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Rotate float64 `json:"rotate"` // degree
	Scale  float64 `json:"scale"`  // percent
}

type CharacterIconProcessSchemaEditData struct {
	CharacterIconProcessSchema
	Id int `json:"id"`
}

type CharacterIconLayerGroup struct {
	Id    int                      `json:"id"`
	Name  string                   `json:"name"`
	Items []CharacterIconLayerItem `json:"items"`
}

type CharacterIconLayeringGroupOverview struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CharacterIconLayeringGroup struct {
	CharacterIconLayeringGroupOverview
	Layers    []CharacterIconLayerGroup            `json:"layers"`
	Processes []CharacterIconProcessSchemaEditData `json:"processes"`
}

type CharacterIconLayerGroupOrder struct {
	LayerGroup int `json:"layerGroup"`
	Order      int `json:"order"`
}