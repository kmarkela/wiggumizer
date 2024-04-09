package historyparser

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// type PostParams interface {
// 	map[string]string | map[string]interface
// }

// parse key=value
func parseKeqV(s string) (map[string]string, error) {
	params := make(map[string]string)

	kEqV, err := url.Parse(s)
	if err != nil {
		return params, err
	}

	queryParams := kEqV.Query()
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

	var result map[string]string
	var err error

	// empty Body (t=1)
	if len(r.Body) < 3 {
		return nil
	}

	// parse form data
	if strings.HasPrefix(r.ContentType, "application/x-www-form-urlencoded") {
		result, err = parseKeqV(fmt.Sprintf("?%s", r.Body))
		if err != nil {
			return err
		}
	}

	// parse json
	if r.ContentType == "application/json" {
		var data map[string]interface{} = make(map[string]interface{})
		var dataList []map[string]interface{}

		switch {
		case strings.HasPrefix(r.Body, "["):
			if err := json.Unmarshal([]byte(r.Body), &dataList); err != nil {
				return fmt.Errorf("error parsing JSON: %s", err.Error())
			}
			// TODO: Fuzzing only firs element in the list. update to fuzz all
			data = dataList[0]
			data["WG-data-in-slice"] = struct{}{}
		default:
			if err := json.Unmarshal([]byte(r.Body), &data); err != nil {
				return fmt.Errorf("error parsing JSON: %s", err.Error())
			}
		}

		result = parseJSON(data, "")
	}

	r.Parameters.Post = result

	return nil
}

func parseJSON(data map[string]interface{}, prefix string) map[string]string {
	result := make(map[string]string)

	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			// Nested object, recurse
			nestedResult := parseJSON(v, prefix+key+".")
			for nestedKey, nestedValue := range nestedResult {
				result[nestedKey] = nestedValue
			}
		default:
			// Leaf node, add to the result map
			result[prefix+key] = fmt.Sprintf("%v", value)
		}
	}

	return result
}
