package api

type Season struct {
	Id              uint64   `json:"id" dynamodbav:"id"`
	StartDate       *ISODate `json:"startDate,omitempty" dynamodbav:"startDate"`
	EndDate         *ISODate `json:"endDate,omitempty" dynamodbav:"endDate"`
	CurrentMatchDay uint8    `json:"currentMatchDay" dynamodbav:"currentMatchDay"`
	Winner          *Team    `json:"winner,omitempty" dynamodbav:"winner"`
}
