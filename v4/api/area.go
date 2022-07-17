package api

type Areas struct {
	BaseResponse
	Areas []*Area `json:"areas"`
}

type Area struct {
	Id           uint64  `json:"id" dynamodbav:"id"`
	Name         string  `json:"name" dynamodbav:"name"`
	Code         string  `json:"code" dynamodbav:"code"`
	CountryCode  string  `json:"countryCode" dynamodbav:"countryCode"`
	Flag         string  `json:"flag,omitempty" dynamodbav:"flag"`
	ParentAreaId *uint64 `json:"parentAreaId,omitempty" dynamodbav:"parentAreaId"`
	ParentArea   string  `json:"parentArea,omitempty" dynamodbav:"parentArea"`
	ChildAreas   []*Area `json:"childAreas,omitempty" dynamodbav:"childAreas"`
}
