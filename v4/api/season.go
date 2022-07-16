package api

type Season struct {
	Id              uint64   `json:"id"`
	StartDate       *ISODate `json:"startDate,omitempty"`
	EndDate         *ISODate `json:"endDate,omitempty"`
	CurrentMatchDay uint8    `json:"currentMatchDay"`
	Winner          *Team    `json:"winner,omitempty"`
}
