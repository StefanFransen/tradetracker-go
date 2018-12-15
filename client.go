package tradetracker

import (
  "io/ioutil"
  "net/http"
  "bytes"
  "log"
  "io"
)

type Client struct {
  CustomerID      int64
  Passphrase      string
  AffiliateSiteID int64
  authenticated   bool

  Logger          *log.Logger
  HTTPClient      *http.Client
}

func (c *Client) execute(action string, payload io.Reader) (*Response, error) {
  var err error

  if !c.authenticated {
    c.authenticated = true
    authEnv := getAuthEnvelope(c.CustomerID, c.Passphrase)
    res, err := c.execute("https://ws.tradetracker.com/soap/affiliate/authenticate", authEnv)

    if err != nil {
      return nil, err
    }

    if res.StatusCode != http.StatusOK {
      return nil, res.error()
    }
  }

  req, err := http.NewRequest("POST", "https://ws.tradetracker.com/soap/affiliate", payload)
  if err != nil {
    return nil, err
  }

  // Setting the SOAP header manually because TradeTracker/SOAP only accepts SOAP in uppercase (Header.Set will lowercase the header)
  req.Header["SOAPAction"] = []string{ action }
  req.Header.Set("Content-Type", "text/xml; charset=UTF-8")

  if c.Logger != nil {
		c.Logger.Printf("tradetracker: %s %s %s\n", req.Method, req.URL, req.Body)
	}
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if c.Logger != nil {
		body, _ := ioutil.ReadAll(res.Body)
		c.Logger.Printf("tradetracker: %s\n", res.Status)
		c.Logger.Printf("tradetracker: %s\n", body)
		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	return &Response{res}, nil
}
