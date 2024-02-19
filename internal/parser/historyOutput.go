package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kmarkela/Wiggumizeng/pkg/collections"
)

type BrowseHistory struct {
	HistoryItems []HistoryItem
	ListOfHosts  collections.Set
}

type HistoryItem struct {
	Time   string
	Host   string
	Path   string
	Method string
	Status int
	Req    HistoryReqRes
	Res    HistoryReqRes
}

type HistoryReqRes struct {
	Headers     string
	Body        string
	ContentType string
}

func parseReqRes(irr *InReqRes, hrr *HistoryReqRes) error {

	if err := irr.decodeBase64(); err != nil {
		return err
	}

	rs := strings.Split(irr.Value, "\r\n\r\n")

	hrr.Headers = rs[0]

	if len(rs) > 1 {
		hrr.Body = rs[1]
	}

	hrr.ContentType = getContentType(hrr.Headers)
	return nil

}

func getContentType(headerString string) string {
	lines := strings.Split(headerString, "\n")
	contentTypeRegex := regexp.MustCompile(`Content-Type:\s*(.*)`)

	for _, line := range lines {
		match := contentTypeRegex.FindStringSubmatch(line)
		if len(match) > 1 {
			return strings.TrimSpace(match[1])
		}
	}

	return ""
}

func (b *BrowseHistory) FilterByHost(hosts collections.Set) {
	filteredItems := []HistoryItem{}

	for _, item := range b.HistoryItems {

		if hosts.Contains(item.Host) {
			filteredItems = append(filteredItems, item)
		}
	}
	b.HistoryItems = filteredItems
	b.ListOfHosts = hosts
}

func (bh *BrowseHistory) parse(hx *historyXML) error {
	bh.ListOfHosts = collections.Set{}
	for _, item := range hx.ItemElements {

		i, err := prepareHItem(&item)
		if err != nil {
			fmt.Printf("can't parse request. from: %s. Err: %s\n", item.Time, err.Error())
			continue
		}
		bh.HistoryItems = append(bh.HistoryItems, i)
		bh.ListOfHosts.Add(i.Host)

	}

	if len(bh.HistoryItems) < 1 {
		return fmt.Errorf("empty history")
	}

	return nil
}

func prepareHItem(ii *InItem) (HistoryItem, error) {
	var hi = HistoryItem{
		Req: HistoryReqRes{},
		Res: HistoryReqRes{},
	}

	status, err := strconv.Atoi(ii.Status)
	if err != nil {
		return hi, fmt.Errorf("wrong stauts, status: %s", ii.Status)
	}

	// Asign values
	hi.Time = ii.Time
	hi.Host = ii.Protocol + "://" + ii.Host.Value + ":" + ii.Port
	hi.Path = ii.Path
	hi.Method = ii.Method
	hi.Status = status

	if err := parseReqRes(&ii.Request, &hi.Req); err != nil {
		return hi, err
	}

	if err := parseReqRes(&ii.Response, &hi.Res); err != nil {
		return hi, err
	}

	return hi, nil

}
