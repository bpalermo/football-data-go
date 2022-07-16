package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	year  = 2022
	month = time.Month(12)
	day   = 31
)

func TestDate_UnmarshallJSON(t *testing.T) {
	expected := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	payload := []byte(fmt.Sprintf("\"%04d-%02d-%02d\"", year, month, day))
	result := &ISODate{}
	if err := result.UnmarshalJSON(payload); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ISO8601, result.Format)
	assert.Equal(t, expected, result.Time)

}

func TestDate_MarshallJSON(t *testing.T) {
	expectedTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	d := &ISODate{
		Format: ISO8601,
		Time:   expectedTime,
	}

	result, err := d.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, fmt.Sprintf("\"%04d-%02d-%02d\"", year, month, day), string(result))
}
