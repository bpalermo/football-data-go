package api

import (
	"encoding/json"
	"time"
)

const (
	ISO8601 = "2006-01-02"
)

type ISODate struct {
	Format string
	time.Time
}

// UnmarshalJSON ISODate method
func (d *ISODate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	d.Format = ISO8601
	t, err := time.Parse(d.Format, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalJSON ISODate method
func (d ISODate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(d.Format))
}
