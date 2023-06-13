package infra

import "time"

func GetDatetimeCanonical(date *time.Time) *string {
	if date == nil {
		return nil
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	d := date.In(loc).Format("2006-01-02 15:04:05")
	return &d
}
