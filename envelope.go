package tradetracker

import (
  "text/template"
  "bytes"
  "strings"
  "fmt"
  "io"
)

const getFeeds = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="https://ws.tradetracker.com/soap/affiliate" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <SOAP-ENV:Body>
    <ns1:getFeeds>
      <affiliateSiteID xsi:type="xsd:nonNegativeInteger">{{.AffiliateSiteID}}</affiliateSiteID>
      <options xsi:type="ns1:FeedFilter">
        {{ if (eq .ID 0)}}<ID xsi:nil="true"/>{{ else }}<ID xsi:type="xsd:nonNegativeInteger">{{.ID}}</ID>{{ end }}
        {{ if (eq .CampaignID 0)}}<campaignID xsi:nil="true"/>{{ else }}<campaignID xsi:type="xsd:nonNegativeInteger">{{.CampaignID}}</campaignID>{{ end }}
        {{ if (eq .CampaignCategoryID 0)}}<campaignCategoryID xsi:nil="true"/>{{ else }}<campaignCategoryID xsi:type="xsd:nonNegativeInteger">{{.CampaignCategoryID}}</campaignCategoryID>{{ end }}
        {{ if (eq .AssignmentStatus "")}}<assignmentStatus xsi:nil="true"/>{{ else }}<assignmentStatus xsi:type="ns1:CampaignAssignmentStatus">{{.AssignmentStatus}}</assignmentStatus>{{ end }}
      </options>
    </ns1:getFeeds>
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

func getAuthEnvelope(customerID int64, passphrase string) (io.Reader) {
	return strings.NewReader(fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<SOAP-ENV:Envelope xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:ns1=\"https://ws.tradetracker.com/soap/affiliate\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n  <SOAP-ENV:Body>\n    <ns1:authenticate>\n      <customerID xsi:type=\"xsd:nonNegativeInteger\">%d</customerID>\n      <passphrase xsi:type=\"xsd:normalizedString\">%s</passphrase>\n      <sandbox xsi:nil=\"true\"/>\n      <locale xsi:nil=\"true\" xsi:type=\"ns1:Locale\"/>\n      <demo xsi:nil=\"true\"/>\n    </ns1:authenticate>\n  </SOAP-ENV:Body>\n</SOAP-ENV:Envelope>", customerID, passphrase))
}

func getCampaignsEnvelope(affiliateSiteID int64) (io.Reader) {
	return strings.NewReader(fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<SOAP-ENV:Envelope xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:ns1=\"https://ws.tradetracker.com/soap/affiliate\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n  <SOAP-ENV:Body>\n    <ns1:getCampaigns>\n      <affiliateSiteID xsi:type=\"xsd:nonNegativeInteger\">%d</affiliateSiteID>\n      <options xsi:type=\"ns1:CampaignFilter\">\n        <ID xsi:nil=\"true\"/>\n        <query xsi:nil=\"true\"/>\n        <campaignCategoryID xsi:nil=\"true\"/>\n        <assignmentStatus xsi:type=\"ns1:CampaignAssignmentStatus\">accepted</assignmentStatus>\n        <policySearchEngineMarketingStatus xsi:nil=\"true\"/>\n        <policyEmailMarketingStatus xsi:nil=\"true\"/>\n        <policyCashbackStatus xsi:nil=\"true\"/>\n        <policyDiscountCodeStatus xsi:nil=\"true\"/>\n        <limit xsi:nil=\"true\"/>\n        <offset xsi:nil=\"true\"/>\n        <sort xsi:nil=\"true\"/>\n        <sortDirection xsi:nil=\"true\"/>\n        <excludeInfo xsi:nil=\"true\"/>\n      </options>\n    </ns1:getCampaigns>\n  </SOAP-ENV:Body>\n</SOAP-ENV:Envelope>", affiliateSiteID))
}

func getFeedsEnvelope(options FeedFilter) (string, error) {
  t, err := template.New("getFeeds").Parse(getFeeds)
  if err != nil {
    return "", err
  }

  var result bytes.Buffer
  err = t.Execute(&result, options)
  if err != nil {
    return "", err
  }

  return result.String(), nil
}

/*func GetFeedsEnvelopeTwo(options FeedFilter) (io.Reader) {
  return strings.NewReader("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<SOAP-ENV:Envelope xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:ns1=\"https://ws.tradetracker.com/soap/affiliate\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n  <SOAP-ENV:Body>\n    <ns1:getFeeds>\n      <affiliateSiteID xsi:type=\"xsd:nonNegativeInteger\">92543</affiliateSiteID>\n      <options xsi:type=\"ns1:FeedFilter\">\n        <ID xsi:nil=\"true\"/>\n        <campaignID xsi:nil=\"true\"/>\n        <campaignCategoryID xsi:nil=\"true\"/>\n        <assignmentStatus xsi:type=\"ns1:CampaignAssignmentStatus\">accepted</assignmentStatus>\n      </options>\n    </ns1:getFeeds>\n  </SOAP-ENV:Body>\n</SOAP-ENV:Envelope>")
}*/

  //getFeeds payload
	//payload := strings.NewReader("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<SOAP-ENV:Envelope xmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:ns1=\"https://ws.tradetracker.com/soap/affiliate\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:SOAP-ENC=\"http://schemas.xmlsoap.org/soap/encoding/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n  <SOAP-ENV:Body>\n    <ns1:getFeeds>\n      <affiliateSiteID xsi:type=\"xsd:nonNegativeInteger\">92543</affiliateSiteID>\n      <options xsi:type=\"ns1:FeedFilter\">\n        <ID xsi:nil=\"true\"/>\n        <campaignID xsi:nil=\"true\"/>\n        <campaignCategoryID xsi:nil=\"true\"/>\n        <assignmentStatus xsi:type=\"ns1:CampaignAssignmentStatus\">accepted</assignmentStatus>\n      </options>\n    </ns1:getFeeds>\n  </SOAP-ENV:Body>\n</SOAP-ENV:Envelope>")
