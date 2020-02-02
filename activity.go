package fitbit

import (
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)


func (client Client) GetSteps(date time.Time) Activity {
	var Activity Activity
	Activity.getActivity(client,"steps",date)
	return Activity
}

func (client Client) GetDistance(date time.Time) Activity {
	var Activity Activity
	Activity.getActivity(client,"distance",date)
	return Activity
}

func (client Client) GetFloors(date time.Time) Activity {
	var Activity Activity
	Activity.getActivity(client,"floors",date)
	return Activity
}

func (client Client) GetElevation(date time.Time) Activity {
	var Activity Activity
	Activity.getActivity(client,"elevation",date)
	return Activity
}

func (client Client) GetCalories(date time.Time) Activity {
	var Activity Activity
	Activity.getActivity(client,"calories",date)
	return Activity
}


func (Activity *Activity) getActivity(client Client, kind string, date time.Time,) {
	Activity.Tag = kind
	path := fmt.Sprintf("/1/user/-/activities/" + kind + "/date/"+date.Format(dateLayout)+"/1d")
	json := client.getJson(path)

	activeDate := json.Get("activities-"+kind+".0.dateTime").String()

	json.Get("activities-"+kind+"-intraday.dataset").ForEach(func(key, value gjson.Result) bool {

		var entry ActiveEntry

		parsedTime, _ := time.Parse("2006-01-0215:04:05", activeDate + value.Get("time").String())

		entry.Timestamp = parsedTime.UnixNano()
		entry.Value = value.Get("value").Float()

		Activity.ActiveEntry = append(Activity.ActiveEntry, entry)

		return true // keep iterating
	})

}

type Activity struct {
	Tag string
	ActiveEntry []ActiveEntry
}

type ActiveEntry struct {
	Timestamp int64
	Value     float64
}


