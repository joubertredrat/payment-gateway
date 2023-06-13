package infra_test

import (
	"joubertredrat/transaction-ms/internal/infra"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatetimeCanonical(t *testing.T) {
	tests := []struct {
		name             string
		time             *time.Time
		datetimeExpected *string
	}{
		{
			name: "test datetime canonical with valid datetime",
			time: func() *time.Time {
				date, _ := time.Parse("2006-01-02 15:04:05", "2029-06-13 17:27:51")
				return &date
			}(),
			datetimeExpected: func() *string {
				str := "2029-06-13 14:27:51"
				return &str
			}(),
		},
		{
			name:             "test datetime canonical with no datetime",
			time:             nil,
			datetimeExpected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			datetimeGot := infra.DatetimeCanonical(test.time)
			assert.Equal(t, test.datetimeExpected, datetimeGot)
		})
	}
}

func TestCardExpireTime(t *testing.T) {
	tests := []struct {
		name         string
		year         string
		month        string
		timeExpected time.Time
		errExpected  error
	}{
		{
			name:  "test card expire time with valid data",
			year:  "2029",
			month: "05",
			timeExpected: func() time.Time {
				date, _ := time.Parse("2006-01-02", "2029-05-01")
				return date
			}(),
			errExpected: nil,
		},
		{
			name:         "test card expire time with invalid year",
			year:         "2k29",
			month:        "05",
			timeExpected: time.Time{},
			errExpected: &time.ParseError{
				Layout:     "2006-01-02",
				Value:      "2k29-05-01",
				LayoutElem: "2006",
				ValueElem:  "-05-01",
				Message:    "",
			},
		},
		{
			name:         "test card expire time with invalid month",
			year:         "2029",
			month:        "16",
			timeExpected: time.Time{},
			errExpected: &time.ParseError{
				Layout:     "2006-01-02",
				Value:      "2029-16-01",
				LayoutElem: "01",
				ValueElem:  "-01",
				Message:    ": month out of range",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			timeGot, errGot := infra.CardExpireTime(test.year, test.month)

			assert.Equal(t, test.timeExpected, timeGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
