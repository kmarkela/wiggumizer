package parser

import (
	"encoding/base64"
	"encoding/xml"
)

type historyXML struct {
	XMLName      xml.Name `xml:"items"`
	ItemElements []InItem `xml:"item"`
}

type InItem struct {
	Time     string   `xml:"time"`
	URL      string   `xml:"url"`
	Host     InHost   `xml:"host"`
	Port     string   `xml:"port"`
	Protocol string   `xml:"protocol"`
	Method   string   `xml:"method"`
	Path     string   `xml:"path"`
	Request  InReqRes `xml:"request"`
	Status   string   `xml:"status"`
	MimeType string   `xml:"mimetype"`
	Response InReqRes `xml:"response"`
}

type InHost struct {
	Value string `xml:",chardata"`
	IP    string `xml:"ip,attr"`
}

type InReqRes struct {
	Base64 bool   `xml:"base64,attr"`
	Value  string `xml:",chardata"`
}

func (r *InReqRes) decodeBase64() (err error) {

	// Return if not encoded
	if !r.Base64 {
		return nil
	}

	stringBytes, err := base64.StdEncoding.DecodeString(r.Value)
	if err != nil {
		return err
	}

	r.Value = string(stringBytes)

	return nil
}

// unmarshal XML
func (hx *historyXML) parseHistory(data []byte) error {

	if err := xml.Unmarshal(data, hx); err != nil {
		return err
	}
	return nil
}
