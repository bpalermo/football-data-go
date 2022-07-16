package api

import "time"

type Competitions struct {
	BaseResponse
	Competitions []*Competition `json:"competitions,omitempty"`
}

type Competition struct {
	Id                       uint64          `json:"id"`
	Area                     *Area           `json:"area,omitempty"`
	Name                     string          `json:"name,omitempty"`
	Code                     string          `json:"code,omitempty"`
	Type                     CompetitionType `json:"type,omitempty"`
	Emblem                   string          `json:"emblem,omitempty"`
	Plan                     Plan            `json:"plan,omitempty"`
	CurrentSeason            *Season         `json:"currentSeason,omitempty"`
	Seasons                  []*Season       `json:"seasons,omitempty"`
	NumberOfAvailableSeasons uint64          `json:"numberOfAvailableSeasons,omitempty"`
	LastUpdated              time.Time       `json:"lastUpdated,omitempty"`
}
