package api

import "time"

type Competitions struct {
	BaseResponse
	Competitions []*Competition `json:"competitions,omitempty"`
}

type Competition struct {
	Id                       uint64          `json:"id" dynamodbav:"id"`
	Area                     *Area           `json:"area,omitempty" dynamodbav:"area"`
	Name                     string          `json:"name,omitempty" dynamodbav:"name"`
	Code                     string          `json:"code,omitempty" dynamodbav:"code"`
	Type                     CompetitionType `json:"type,omitempty" dynamodbav:"type"`
	Emblem                   string          `json:"emblem,omitempty" dynamodbav:"emblem"`
	Plan                     Plan            `json:"plan,omitempty" dynamodbav:"plan"`
	CurrentSeason            *Season         `json:"currentSeason,omitempty" dynamodbav:"currentSeason"`
	Seasons                  []*Season       `json:"seasons,omitempty" dynamodbav:"seasons"`
	NumberOfAvailableSeasons uint64          `json:"numberOfAvailableSeasons,omitempty" dynamodbav:"numberOfAvailableSeasons"`
	LastUpdated              time.Time       `json:"lastUpdated,omitempty" dynamodbav:"lastUpdated"`
}
