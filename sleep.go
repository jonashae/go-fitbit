package fitbit

import (
	"fmt"
	"github.com/tidwall/gjson"
	"sort"
	"time"
)


func (client Client) GetSleep(date time.Time) Sleep {

	path := fmt.Sprintf("/1.2/user/-/sleep/date/%s", date.Format(dateLayout))
	json := client.getJson(path)

	var Sleep Sleep

	Sleep.GetSleepEvents(json,"sleep.0.levels.data", "data")
	Sleep.GetSleepEvents(json,"sleep.0.levels.shortData", "short")

	sort.Slice(Sleep.SleepEntry, func(i, j int) bool {
		return Sleep.SleepEntry[i].Timestamp < Sleep.SleepEntry[j].Timestamp
	})

	return Sleep
}

func (Sleep *Sleep) GetSleepEvents(json gjson.Result, path string, tag string){
	json.Get(path).ForEach(func(key, value gjson.Result) bool {
		var entry SleepEntry


		parsedTime, err := time.Parse("2006-01-02T15:04:05.000", value.Get("dateTime").String())
		if err != nil {
			panic(err)
		}

		entry.Duration = int(value.Get("seconds").Int())
		entry.Timestamp = parsedTime.UnixNano()
		entry.Level = value.Get("level").String()
		entry.Tag = tag

		Sleep.SleepEntry = append(Sleep.SleepEntry, entry)

		return true // keep iterating
	})

}

type Sleep struct {
	SleepEntry []SleepEntry
}

type SleepEntry struct {
	Timestamp int64
	Level     string
	Duration  int
	Tag string
}






