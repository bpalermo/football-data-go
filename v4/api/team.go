package api

import "time"

type Team struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name,omitempty"`
	ShortName   string    `json:"shortName,omitempty"`
	Tla         string    `json:"tla,omitempty"`
	Crest       string    `json:"crest,omitempty"`
	Address     string    `json:"address,omitempty"`
	Website     string    `json:"website,omitempty"`
	Founded     uint16    `json:"founded,omitempty"`
	ClubColors  string    `json:"ClubColors,omitempty"`
	Venue       string    `json:"venue,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
}
