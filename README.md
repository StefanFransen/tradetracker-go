TradeTracker Go
================

An unofficial Go client for [TradeTracker](https://tradetracker.com/). This package is still WIP, documentation, unit tests and comments will be added later on.

## Example
```
package main

import (
  "net/http/cookiejar"
  "net/http"
  "log"

  "github.com/StefanFransen/tradetracker-go"
)

func main() {
  cookieJar, _ := cookiejar.New(nil)

  c := &tradetracker.Client{
    CustomerID: 123456,
    Passphrase: "hello",
    AffiliateSiteID: 12345,

    HTTPClient: &http.Client{
      Jar: cookieJar,
    },
  }

  feeds, err := c.Feed().List()
  if err != nil {
    log.Fatalln(err)
  }

  for _, feed := range *feeds {
    log.Println(feed)
  }
}

```

## Notes
Inspired on [moneybird-go](https://github.com/dannyvankooten/moneybird-go)


## License

MIT Licensed. See the [LICENSE](LICENSE) file for details.
