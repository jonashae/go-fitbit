package fitbit

import (
	"strconv"
	"time"
)

const dateTimeLayout = "\"2006-01-02T15:04:05.000\""
const timeLayout = "\"15:04:05\""
const dateLayout = "2006-01-02"

// LocalDate is used to parse the date fields in the API
type LocalDate struct {
	time.Time
}

// UnmarshalJSON parses the date string from the JSON response
func (date *LocalDate) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(dateTimeLayout, string(data), time.Local)
	if err != nil {
		return err
	}

	date.Time = parsed
	return nil
}

// LocalTime is used to parse time fields in the API
type LocalTime struct {
	time.Time
}

// UnmarshalJSON parses the time string from the JSON response
func (date *LocalTime) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(timeLayout, string(data), time.Local)
	if err != nil {
		return err
	}

	date.Time = parsed
	return nil
}

// MillisecondDuration is used to parse millisecond durations in the API
type MillisecondDuration struct {
	time.Duration
}

// UnmarshalJSON parses the milliseconds int from the JSON response
func (duration *MillisecondDuration) UnmarshalJSON(data []byte) error {
	millis, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}

	duration.Duration = time.Duration(millis) * time.Millisecond
	return nil
}

// SecondDuration is used to parse second durations in the API
type SecondDuration struct {
	time.Duration
}

// UnmarshalJSON parses the seconds int from the JSON response
func (duration *SecondDuration) UnmarshalJSON(data []byte) error {
	millis, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}

	duration.Duration = time.Duration(millis) * time.Second
	return nil
}

// MinuteDuration is used to parse minute durations in the API
type MinuteDuration struct {
	time.Duration
}

// UnmarshalJSON parses the minutes int from the JSON response
func (duration *MinuteDuration) UnmarshalJSON(data []byte) error {
	millis, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}

	duration.Duration = time.Duration(millis) * time.Minute
	return nil
}
