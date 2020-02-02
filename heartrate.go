package fitbit

import (
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)


func (client Client) GetHeartrate(date time.Time) Heartrate {

	path := fmt.Sprintf("/1/user/-/activities/heart/date/%s/1d/1min", date.Format(dateLayout))
	json := client.getJson(path)

	activeDate := json.Get("activities-heart.0.dateTime").String()

	var Heartrate Heartrate

	json.Get("activities-heart-intraday.dataset").ForEach(func(key, value gjson.Result) bool {

		var entry HeartrateEntry

		parsedTime, _ := time.Parse("2006-01-0215:04:05", activeDate + value.Get("time").String())

		entry.Timestamp = parsedTime.UnixNano()
		entry.BPM = value.Get("value").Float()

		Heartrate.HeartrateEntry = append(Heartrate.HeartrateEntry, entry)

		return true // keep iterating
	})

	return Heartrate
}


type Heartrate struct {
	HeartrateEntry []HeartrateEntry
}

type HeartrateEntry struct {
	Timestamp int64
	BPM     float64
}


