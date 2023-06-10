package pkg

import "time"

func TimeFromCanonical(datetime string) *time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", datetime)
	return &t
}
