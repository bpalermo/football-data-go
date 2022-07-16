package client

import (
	"fmt"
	api2 "github.com/bpalermo/football-data-go/v4/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func TestMain(m *testing.M) {
	teardown := setup()
	m.Run()
	teardown()
}

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = New(BaseURL(server.URL))

	return func() {
		server.Close()
	}
}

func TestNew_Defaults(t *testing.T) {
	c, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, defaultBaseURL, c.baseURL)
	assert.Equal(t, defaultAuthToken, c.authToken)
	assert.Equal(t, defaultTimeout, c.httpClient.Timeout)
}

func TestNew_WithOptions(t *testing.T) {
	expectedBaseUrl := "http://localhost"
	expectedAuthToken := "fake"
	expectedTimeout := 5 * time.Second

	c, err := New(BaseURL(expectedBaseUrl), AuthToken(expectedAuthToken), Timeout(expectedTimeout))
	assert.Nil(t, err)
	assert.Equal(t, expectedBaseUrl, c.baseURL)
	assert.Equal(t, expectedAuthToken, c.authToken)
	assert.Equal(t, expectedTimeout, c.httpClient.Timeout)
}

func TestClient_GetAreas(t *testing.T) {
	addEndpoint(areasEndpoint, "areas.json")

	areas, err := client.GetAreas()
	if err != nil {
		t.Fatal(err)
	}

	var afghanistanId uint64 = 2014
	var africaId uint64 = 2267

	expected := &api2.Areas{
		BaseResponse: api2.BaseResponse{
			Count: uint32(2),
		},
		Areas: []*api2.Area{
			{
				Id:           2000,
				Name:         "Afghanistan",
				CountryCode:  "AFG",
				Flag:         "",
				ParentAreaId: &afghanistanId,
				ParentArea:   "Asia",
			},
			{
				Id:           2001,
				Name:         "Africa",
				CountryCode:  "AFR",
				Flag:         "",
				ParentAreaId: &africaId,
				ParentArea:   "World",
			},
		},
	}

	assert.NotNil(t, areas)
	assert.Equal(t, expected, areas)
}

func TestClient_GetArea(t *testing.T) {
	areaId := "INT"
	addEndpoint(fmt.Sprintf("%s/%s", areasEndpoint, areaId), "area.json")

	area, err := client.GetArea(areaId)
	if err != nil {
		t.Fatal(err)
	}

	var expectedWorldId uint64 = 2267

	var expected = &api2.Area{
		Id:   expectedWorldId,
		Code: "INT",
		Name: "World",
		ChildAreas: []*api2.Area{
			{
				Id:           2001,
				Name:         "Africa",
				CountryCode:  "AFR",
				ParentAreaId: &expectedWorldId,
				ParentArea:   "World",
			},
			{
				Id:           2010,
				Name:         "Arameans Suryoye",
				CountryCode:  "ARS",
				ParentAreaId: &expectedWorldId,
				ParentArea:   "World",
			},
		},
	}
	assert.NotNil(t, area)
	assert.Equal(t, expected, area)
}

func TestClient_GetCompetitions(t *testing.T) {
	addEndpoint(competitionsEndpoint, "competitions.json")

	var expectedCount uint32 = 157

	expectedStartDate := time.Date(2019, 9, 4, 0, 0, 0, 0, time.UTC)
	expectedEndDate := time.Date(2021, 11, 16, 0, 0, 0, 0, time.UTC)
	expectedLastUpdated := time.Date(2022, 3, 13, 18, 51, 44, 0, time.UTC)

	expectedCompetition := &api2.Competition{
		Id: 2006,
		Area: &api2.Area{
			Id:   2001,
			Name: "Africa",
			Code: "AFR",
		},
		Name: "WC Qualification CAF",
		Code: "QCAF",
		Type: api2.COMPETITION_TYPE_CUP,
		Plan: api2.PLAN_TIER_FOUR,
		CurrentSeason: &api2.Season{
			Id: 555,
			StartDate: &api2.ISODate{
				Format: api2.ISO8601,
				Time:   expectedStartDate,
			},
			EndDate: &api2.ISODate{
				Format: api2.ISO8601,
				Time:   expectedEndDate,
			},
			CurrentMatchDay: 6,
		},
		NumberOfAvailableSeasons: 2,
		LastUpdated:              expectedLastUpdated,
	}

	competitions, err := client.GetCompetitions()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, competitions)
	assert.Equal(t, expectedCount, competitions.Count)
	assert.Equal(t, expectedCompetition, competitions.Competitions[0])
}

func TestClient_GetCompetition(t *testing.T) {
	competitionCode := "PL"
	addEndpoint(fmt.Sprintf("%s/%s", competitionsEndpoint, competitionCode), "competition.json")

	expectedCurrentSeason := &api2.Season{
		Id:              733,
		CurrentMatchDay: 37,
		StartDate: &api2.ISODate{
			Format: api2.ISO8601,
			Time:   time.Date(2021, 8, 13, 0, 0, 0, 0, time.UTC),
		},
		EndDate: &api2.ISODate{
			Format: api2.ISO8601,
			Time:   time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC),
		},
	}

	expectedWinner := &api2.Team{
		Id:          uint64(65),
		Name:        "Manchester City FC",
		ShortName:   "Man City",
		Tla:         "MCI",
		Crest:       "https://crests.football-data.org/65.png",
		Address:     "SportCity Manchester M11 3FF",
		Website:     "https://www.mancity.com",
		Founded:     1880,
		ClubColors:  "Sky Blue / White",
		Venue:       "Etihad Stadium",
		LastUpdated: time.Date(2022, 2, 10, 19, 48, 37, 0, time.UTC),
	}

	expected := &api2.Competition{
		Id:     uint64(2021),
		Name:   "Premier League",
		Code:   "PL",
		Type:   api2.COMPETITION_TYPE_LEAGUE,
		Emblem: "https://crests.football-data.org/PL.png",
		Area: &api2.Area{
			Id:   2072,
			Name: "England",
			Code: "ENG",
			Flag: "https://crests.football-data.org/770.svg",
		},
		CurrentSeason: expectedCurrentSeason,
		Seasons: []*api2.Season{
			expectedCurrentSeason,
			{
				Id: uint64(619),
				StartDate: &api2.ISODate{
					Format: api2.ISO8601,
					Time:   time.Date(2020, 9, 12, 0, 0, 0, 0, time.UTC),
				},
				EndDate: &api2.ISODate{
					Format: api2.ISO8601,
					Time:   time.Date(2021, 5, 23, 0, 0, 0, 0, time.UTC),
				},
				CurrentMatchDay: 38,
				Winner:          expectedWinner,
			},
		},
		LastUpdated: time.Date(2022, 3, 20, 8, 58, 54, 0, time.UTC),
	}

	competition, err := client.GetCompetition(competitionCode)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, competition)
}

func addEndpoint(endpoint string, filename string) {
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, readFile(filename))
	})
}

func readFile(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
