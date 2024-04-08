package fuzz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func (w *worker) encodeBody(ctype, word, param string, params map[string]string) (io.Reader, error) {

	nParams := make(map[string]string, len(params))
	for key, value := range params {
		nParams[key] = value
	}
	nParams[param] = word

	var body io.Reader
	switch {
	case strings.HasPrefix(ctype, "application/x-www-form-urlencoded"):
		body = prepareForm(nParams)
	case ctype == "application/json":
		// fmt.Printf("PARAMS: %s\nword: %s\nparam: %s\n========", params, word, param)
		body = prepareJSON(nParams)
	}

	if body == nil {
		return nil, fmt.Errorf("encoder is not implemented")
	}

	return body, nil

}

func prepareForm(params map[string]string) io.Reader {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	data := values.Encode()

	return bytes.NewBuffer([]byte(data))
}

func isSlice(str string) bool {
	pattern := `^\[.*\]$`
	match, _ := regexp.MatchString(pattern, str)
	return match
}

func isObj(str string) bool {
	pattern := `^\{.*\}$`
	match, _ := regexp.MatchString(pattern, str)
	return match
}

func prepareJSON(data map[string]string) io.Reader {
	var sliceValue []interface{}
	var objValue map[string]interface{}
	jsonData := make(map[string]interface{})

	for key, value := range data {
		// Split the key into parts
		keys := strings.Split(key, ".")

		// Traverse the keys to set the value in jsonData
		temp := jsonData
		for i := 0; i < len(keys)-1; i++ {
			if _, ok := temp[keys[i]]; !ok {
				temp[keys[i]] = make(map[string]interface{})
			}
			temp = temp[keys[i]].(map[string]interface{})
		}

		// check if int
		if v, err := strconv.Atoi(value); err == nil {
			temp[keys[len(keys)-1]] = v
			continue
		}

		// check if slice
		if isSlice(value) {
			json.Unmarshal([]byte(value), &sliceValue)
			temp[keys[len(keys)-1]] = sliceValue
			continue
		}

		// check if obj
		if isObj(value) {
			json.Unmarshal([]byte(value), &objValue)
			temp[keys[len(keys)-1]] = objValue
			continue
		}

		temp[keys[len(keys)-1]] = value
	}

	d, err := json.Marshal(jsonData)
	if err != nil {
		// TODO: log it in verbose
		return nil
	}

	return bytes.NewBuffer(d)
}
