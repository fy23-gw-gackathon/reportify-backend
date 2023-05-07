package driver

import (
	"bytes"
	"io"
	"net/http"
)

type Http struct{}

func NewHttp() *Http {
	return &Http{}
}

func (h Http) Put(path string, jsonData []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
