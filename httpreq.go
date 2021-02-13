package netflixcheck

import (
	"bytes"
	"compress/gzip"
	"net/http"
)

var baseHeader = http.Header{
	"User-Agent":                []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:85.0) Gecko/20100101 Firefox/85.0"},
	"Accept":                    []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
	"Accept-Encoding":           []string{"gzip"},
	"Connection":                []string{"keep-alive"},
	"Upgrade-Insecure-Requests": []string{"1"},
}

const host = "https://www.netflix.com"

func get(cl *http.Client, cookies []*http.Cookie, path string) ([]byte, error) {
	if cl == nil {
		cl = http.DefaultClient
	}

	req, err := http.NewRequest("GET", host+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header = baseHeader

	if cookies != nil {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buff, err := decompressResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func decompressResponseBody(resp *http.Response) (*bytes.Buffer, error) {
	buff := new(bytes.Buffer)
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}

		buff.ReadFrom(reader)
	default:
		buff.ReadFrom(resp.Body)
	}

	return buff, nil
}
