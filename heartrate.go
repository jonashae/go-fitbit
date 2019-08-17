package fitbit

import (
	"fmt"
	"time"
)

// DataResolution determines the granularity of the heart rate data
type DataResolution = string

// Available resolution levels
const (
	OneSecondResolution DataResolution = "1sec"
	OneMinuteResolution DataResolution = "1min"
)

// HeartRateMeasurement represents the heartrate at a point in time
type HeartRateMeasurement struct {
	Time  LocalTime `json:"time"`
	Value int       `json:"value"`
}

// GetHeartRateLogs returns all heartrate time series for a given date
func (client Client) GetHeartRateLogs(date time.Time, resolution DataResolution) ([]HeartRateMeasurement, error) {
	var result struct {
		Root struct {
			DataSet []HeartRateMeasurement `json:"dataset"`
		} `json:"activities-heart-intraday"`
	}

	path := fmt.Sprintf("/1/user/-/activities/heart/date/%s/1d/%s", date.Format(dateLayout), resolution)
	err := client.getResource(path, nil, &result)
	if err != nil {
		return nil, err
	}

	dataset := result.Root.DataSet

	// Adds the missing date part in the time
	for i, entry := range dataset {
		adjusted := entry.Time.AddDate(date.Year(), int(date.Month())-1, int(date.Day())-1)
		dataset[i].Time = LocalTime{adjusted}
	}

	return dataset, nil
}
