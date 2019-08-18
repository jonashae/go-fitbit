# go-fitbit
A library for exporting your tracker data from the Fitbit. 
It's meant to be used for extracting raw tracker data and **not** as a backup tool. It does not cover all available fields, but rather aims to capture the most granular data available.

## Status
- [ ] Activities
- [x] Heart rate
- [ ] Location
- [ ] Nutrition
- [x] Sleep
- [ ] Weight

## Prerequisites
To use this library, you need to create an OAuth2 application on [dev.fitbit.com](https://dev.fitbit.com/apps/new). To get the full data, you should mark the application as **Personal** type.

The Fitbit authorization process requires that the callback URL used in the OAuth2 application to be available for the initial token request (you can put for example *localhost:8080* there).

## Installation
```bash
go get github.com/byonchev/go-fitbit
```

## Usage
```golang
package main

import (
	"fmt"
	"time"

	"github.com/byonchev/go-fitbit"
)

func main() {
	config := fitbit.Config{
		// Port for the initial callback URL server
		CallbackServerPort: 8080,
		// OAuth2 credentials
		ClientID:     "XXXXXXXX",
		ClientSecret: "XXXXXXXXXXXXXXXXXXXXXXXX",
		// Scopes that you want to access
		Scopes: []fitbit.Scope{fitbit.Sleep, fitbit.HeartRate},
		// Path to the token that will be stored on your machine
		TokenPath: "token.json",
	}

	client := fitbit.NewClient(config)

	sleepLogs, err := client.GetSleepLogs(time.Now().Add(-24 * time.Hour))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", sleepLogs)
}
```

**Note**: Keep in mind that the API rate limit is 150 requests per hour.

## Initial run
When running for the first time (or adding more scopes), you'll have to accept the scopes manually in a browser.
You'll get a message like this in the stdout:
```bash
2019/08/17 23:32:57 Visit URL: https://www.fitbit.com/oauth2/authorize?client_id=XXXXXX&response_type=code&scope=sleep+heartrate&state=XXXX
```
The library creates a temporary HTTP server on the **CallbackServerPort**. Once you've accepted the requested scopes, the browser will redirect to the callback URL, which will be handled by the temporary HTTP server. This captures the initial token and stores it on the **TokenPath**.

Once the initial setup is done, the library takes care of the token refreshing and *should* not require manual intervention.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)