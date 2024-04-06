package fuzz

import (
	"bytes"
	"io"
	"net/url"
	"strings"
)

func (w *worker) encodeBody(ctype, word, param string, params map[string]string) io.Reader {

	params[param] = word

	var body io.Reader
	switch {
	case strings.HasPrefix(ctype, "application/x-www-form-urlencoded"):
		body = prepareForm(params)
	}

	return body

}

func prepareForm(params map[string]string) io.Reader {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	data := values.Encode()

	return bytes.NewBuffer([]byte(data))
}
