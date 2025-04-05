package util

import "time"

// TimeUnit represents the unit of time
type TimeUnit string

const (
	Seconds TimeUnit = "seconds"
	Minutes TimeUnit = "minutes"
	Hours   TimeUnit = "hours"
)

func GetCurrentTimeString(unit TimeUnit) string {
	now := time.Now()
	switch unit {
    case Hours:
		return now.Format("2025-04-05T15")
	case Minutes:
		return now.Format("2025-04-05T15:04")
    case Seconds:
        return now.Format("2025-04-05T15:04:05")
	default:
		return now.Format("2025-04-05T15:04:05")
	}
}
