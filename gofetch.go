package gofetch

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

type Fetch struct{
	UserAgent string
	Proxy string
}

type Header struct {
	Name string
	Value string
}

func (s Fetch) Go (method string, url string, headers []Header, body string) (string string, err error) {
	request := gorequest.New()

	request.CustomMethod(method, url)

	for _, v := range headers {
		request.Set(v.Name, v.Value)
	}

	if s.UserAgent != "" {
		request.Set("User-Agent", s.UserAgent)
	}

	if s.Proxy != "" {
		request.Proxy(s.Proxy)
	}

	if body != "" {
		request.Send(body)
	}

	response, content, errs := request.End()

	if errs != nil {
		log.Fatal(errs)
		err = errs[0]
	}

	if response.Status != "200 OK" {
		// log.Fatal("Unexpected response status")
	}

	return content, err
}