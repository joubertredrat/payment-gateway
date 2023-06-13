package infra

import (
	"fmt"
	"time"
)

func DatetimeCanonical(date *time.Time) *string {
	if date == nil {
		return nil
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	d := date.In(loc).Format("2006-01-02 15:04:05")
	return &d
}

func CardExpireTime(year, month string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-01", year, month))
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
