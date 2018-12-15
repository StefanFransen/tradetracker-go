package tradetracker

import (
  "encoding/xml"
  "net/http"
  "strings"
  //"time"

  "fmt"
)

type FeedFilter struct {
    AffiliateSiteID     int64
    ID                  int64
    CampaignID          int64
    CampaignCategoryID  int64
    AssignmentStatus    CampaignAssignmentStatus
}

type CampaignAssignmentStatus string

const (
  NotSignedUp CampaignAssignmentStatus = "notsignedup"
  Pending     CampaignAssignmentStatus = "pending"
  Accepted    CampaignAssignmentStatus = "accepted"
  Rejected    CampaignAssignmentStatus = "rejected"
  OnHold      CampaignAssignmentStatus = "onhold"
  SignedOut   CampaignAssignmentStatus = "signedout"
)

type FeedEnvelope struct {
  Feeds	[]Feed	`xml:"Body>getFeedsResponse>feeds>item"`
}

type Feed struct {
  ID            int64   `xml:"ID"`
  Name          string  `xml:"name"`
  Campaign      string  `xml:"campaign>name"`
  URL           string  `xml:"URL"`
  UpdatedAt     string  `xml:"updateDate"`
  ProductCount  int64   `xml:"productCount"`

  //UpdatedAt     *time.Time   `xml:"updateDate"` - Temperary disabled
}

type FeedGateway struct {
  Client      *Client
}

func (c *Client) Feed() *FeedGateway {
  gateway := FeedGateway{Client: c}

  return &gateway
}

func (c *FeedGateway) List(filters ...FeedFilter) (*[]Feed, error) {
  var options FeedFilter
  if len(filters) > 0 {
    options = filters[0]
  }
  options.AffiliateSiteID = c.Client.AffiliateSiteID

  fmt.Println(options)

  feedsEnv, err := getFeedsEnvelope(options)
  if err != nil {
    return nil, err
  }

  res, err := c.Client.execute("https://ws.tradetracker.com/soap/affiliate/getFeeds", strings.NewReader(feedsEnv))
  if err != nil {
    return nil, err
  }

  if res.StatusCode == http.StatusOK {
    var parseEnv *FeedEnvelope

    d := xml.NewDecoder(res.Body)
    d.Entity = map[string]string{
      "copy": "©",
      "reg": "®",
    }
    err = d.Decode(&parseEnv)
    if err != nil {
      return nil, err
    }

    return &parseEnv.Feeds, nil
  }

  return nil, res.error()
}
