package fuzz

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"strings"
)

func (w *worker) encodeBody(ctype, word, param string, params map[string]string) (io.Reader, error) {

	params[param] = word

	var body io.Reader
	switch {
	case strings.HasPrefix(ctype, "application/x-www-form-urlencoded"):
		body = prepareForm(params)
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
