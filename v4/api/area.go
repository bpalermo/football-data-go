package api

type Areas struct {
	BaseResponse
	Areas []*Area `json:"areas"`
}

type Area struct {
	Id           uint64  `json:"id"`
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	CountryCode  string  `json:"countryCode"`
	Flag         string  `json:"flag,omitempty"`
	ParentAreaId *uint64 `json:"parentAreaId,omitempty"`
	ParentArea   string  `json:"parentArea,omitempty"`
	ChildAreas   []*Area `json:"childAreas,omitempty"`
}
