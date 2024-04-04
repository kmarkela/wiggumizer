package historyparser

import (
	"fmt"
	"net/url"
	"strings"
)

// parse key=value
func parseKeqV(s string) (map[string]string, error) {
	params := make(map[string]string)

	// parse Get parameters
	parsedURL, err := url.Parse(s)
	if err != nil {
		return params, err
	}

	queryParams := parsedURL.Query()
	for k := range queryParams {
		params[k] = queryParams.Get(k)
	}

	return params, nil
}

func parseGetParam(r *HistoryReqRes, path string) error {

	params, err := parseKeqV(path)
	if err != nil {
		return err
	}
	r.Parameters.Get = params

	return nil
}

func parsePostParams(r *HistoryReqRes) error {

	// empty Body (t=1)
	if len(r.Body) < 3 {
		return nil
	}

	// starting with form. others content type will be added in new versions
	if !strings.HasPrefix(r.ContentType, "application/x-www-form-urlencoded") {
		return nil
	}

	params, err := parseKeqV(fmt.Sprintf("?%s", r.Body))
	if err != nil {
		return err
	}

	r.Parameters.Post = params

	return nil
}
