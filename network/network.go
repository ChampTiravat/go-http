package network

import (
	"net/http"
	"time"
)

const (
	StatusOK                  int = http.StatusOK
	StatusCreated             int = http.StatusCreated
	StatusBadRequest          int = http.StatusBadRequest
	StatusUnauthorized        int = http.StatusUnauthorized
	StatusInternalServerError int = http.StatusInternalServerError
)

type (
	Body map[string]string

	Headers map[string]string

	Response struct {
		StatusCode  int
		RawResponse *[]byte
	}

	Request struct {
		HTTPClient      *http.Client
		HTTPMethod      string
		Endpoint        string
		ReqHeaders      *Headers
		ReqBody         *Body
		RetryDuration   *time.Duration
		TimeoutDuration *time.Duration
		Response        *Response
	}
)

func NewRequest() *Request {
	r := Request{}

	r.HTTPClient = &http.Client{}

	return &r
}

func (rPtr *Request) Method(httpMethod string) *Request {
	(*rPtr).HTTPMethod = httpMethod
	return rPtr
}

func (rPtr *Request) URL(url string) *Request {
	(*rPtr).Endpoint = url
	return rPtr
}

func (rPtr *Request) Headers(headers *Headers) *Request {
	(*rPtr).ReqHeaders = headers
	return rPtr
}

func (rPtr *Request) Body(body *Body) *Request {
	(*rPtr).ReqBody = body
	return rPtr
}

func (rPtr *Request) Submit() *Request {
	return rPtr
}

func (rPtr *Request) Timeout(waitTime *time.Duration) *Request {
	(*rPtr).TimeoutDuration = waitTime
	return rPtr
}

func (rPtr *Request) Retry(waitTime *time.Duration) *Request {
	(*rPtr).RetryDuration = waitTime
	return rPtr
}

func (rPtr *Request) BindJSON(destination *any) (int, error) {
	return (*rPtr).Response.StatusCode, nil
}

func (rPtr *Request) BindText(destination *string) (int, error) {
	return (*rPtr).Response.StatusCode, nil
}
