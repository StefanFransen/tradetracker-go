package tradetracker

import (
  "encoding/xml"
  "net/http"
)

type CampaignEnvelope struct {
  Campaigns	[]Campaign	`xml:"Body>getCampaignsResponse>campaigns>item"`
}

type Campaign struct {
  ID          int64
  Name        string  `xml:"name"`
  Description string  `xml:"info>shopDescription"`
  DisplayURL  string  `xml:"URL"`
  TrackingURL string  `xml:"info>trackingURL"`
  Logo        string  `xml:"info>imageURL"`
}

type CampaignGateway struct {
  CampaignID  string
  Client      *Client
}

func (c *Client) Campaign(campaignID ...string) *CampaignGateway {
  gateway := CampaignGateway{Client: c}

  if len(campaignID) > 0 {
    gateway.CampaignID = campaignID[0]
  }

  return &gateway
}

func (c *CampaignGateway) List() (*[]Campaign, error) {
  authEnv := getCampaignsEnvelope(c.Client.AffiliateSiteID)
  res, err := c.Client.execute("https://ws.tradetracker.com/soap/affiliate/getCampaigns", authEnv)
  if err != nil {
    return nil, err
  }

  if res.StatusCode == http.StatusOK {
    var campaignEnv *CampaignEnvelope

    err := xml.NewDecoder(res.Body).Decode(&campaignEnv)
    if err != nil {
      return nil, err
    }

    return &campaignEnv.Campaigns, nil
  }

  return nil, res.error()
}
