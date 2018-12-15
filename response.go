package tradetracker

import (
  "encoding/xml"
	"net/http"
)

// Response wraps a TradeTracker API response
type Response struct {
	*http.Response
}

// APIError holds data for a TradeTracker API error
type APIError struct {
	Response *Response
	Envelope FaultEnvelope
}

type FaultEnvelope struct {
  Body	FaultBody	`xml:"Body"`
}

type FaultBody struct {
  Fault	Fault	`xml:"Fault"`
}

type Fault struct {
  Code    string	`xml:"faultcode"`
  String	string	`xml:"faultstring"`
}

func (e *APIError) Error() string {
	// If there is a error inside the response we'll use that.
  if e.Envelope.Body.Fault.Code != "" || e.Envelope.Body.Fault.String != "" {
    return "tradetracker: " + e.Envelope.Body.Fault.Code + " " + e.Envelope.Body.Fault.String
  }

	return "tradetracker: " + e.Response.Status
}

func (res *Response) error() error {
	apiErr := &APIError{
		Response: res,
	}

  // try to decode into APIError struct
  d := xml.NewDecoder(res.Body)
  d.Entity = map[string]string{
    "copy": "©",
    "reg": "®",
  }
  err := d.Decode(&apiErr.Envelope)
	if err != nil {
		return err
	}

	return apiErr
}
