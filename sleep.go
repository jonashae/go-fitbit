package fitbit

import (
	"fmt"
	"time"
)

// SleepPhase represents a sleep phase
type SleepPhase = string

// Enum of all sleep phases
const (
	Awake SleepPhase = "wake"
	Deep  SleepPhase = "deep"
	Light SleepPhase = "light"
	REM   SleepPhase = "rem"
)

// SleepSummary is the summary of a sleep phase
type SleepSummary struct {
	Count    int            `json:"count"`
	Duration MinuteDuration `json:"minutes"`
}

// SleepEvent is a sleep phase in the sleep timeline
type SleepEvent struct {
	Date     LocalDate      `json:"datetime"`
	Phase    SleepPhase     `json:"level"`
	Duration SecondDuration `json:"seconds"`
}

// SleepData contains the summary and details for a sleep
type SleepData struct {
	Summary  map[SleepPhase]SleepSummary `json:"summary"`
	Timeline []SleepEvent                `json:"data"`
}

// SleepLogEntry represents a single sleep
type SleepLogEntry struct {
	ID        int64               `json:"logId"`
	Duration  MillisecondDuration `json:"duration"`
	StartDate LocalDate           `json:"startTime"`
	EndDate   LocalDate           `json:"endTime"`
	Data      SleepData           `json:"levels"`
}

// GetSleepLogs returns all sleep logs for a specific date
func (client Client) GetSleepLogs(date time.Time) ([]SleepLogEntry, error) {
	var result struct {
		Logs []SleepLogEntry `json:"sleep"`
	}

	path := fmt.Sprintf("/1.2/user/-/sleep/date/%s", date.Format(dateLayout))
	err := client.getResource(path, nil, &result)
	if err != nil {
		return nil, err
	}

	return result.Logs, nil
}
