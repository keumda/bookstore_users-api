package date

import "time"

// Z means standard timezone, notation -07:00 specifies timezone
// or time.Now().UTC()
const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
