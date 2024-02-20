package search

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Method         = "Method"
	ReqHeader      = "ReqHeader"
	ReqContentType = "ReqContentType"
	ReqBody        = "ReqBody"
	ResHeader      = "ResHeader"
	ResContentType = "ResContentType"
	ResBody        = "ResBody"
)

// cehck if search is negatiove (`!`)
func checkNegative(input string) (string, sMatch, error) {

	// check if it is negative (starts with `!`)
	input = strings.TrimSpace(input)
	neg := strings.HasPrefix(input, "!")

	if neg {
		input = strings.TrimSpace(strings.TrimPrefix(input, "!"))
	}

	//
	sParts := strings.SplitN(input, " ", 2)
	if len(sParts) < 2 {
		return "", sMatch{}, errors.New("invalid input")
	}

	sm := sMatch{
		value:    sParts[1],
		negative: neg,
	}
	return sParts[0], sm, nil
}

func parseInput(input string) sParam {
	var sp = sParam{
		req: sReg{},
		res: sReg{},
	}
	parseAnd := strings.Split(input, "&")
	for _, v := range parseAnd {

		op, match, err := checkNegative(v)
		if err != nil {
			fmt.Printf("Error %s in %s\n", err.Error(), v)
			continue
		}

		switch op {
		case Method:
			sp.method = match
		case ReqHeader:
			sp.req.header = append(sp.req.header, match)
		case ReqContentType:
			sp.req.contentType = append(sp.req.contentType, match)
		case ReqBody:
			sp.req.body = append(sp.req.body, match)
		case ResHeader:
			sp.res.header = append(sp.res.header, match)
		case ResContentType:
			sp.res.contentType = append(sp.res.contentType, match)
		case ResBody:
			sp.res.body = append(sp.res.body, match)
		default:
			fmt.Printf("unsupported search field: %s. Use help (in menu) for list of supported field.", op)
			continue
		}
	}
	return sp
}
