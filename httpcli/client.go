package httpcli

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// http cli封装成点式修改对象方法

type Client struct {
	header  map[string]string
	client  *http.Client
	request *http.Request
}

func newClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c *Client) SetContentTypeURLEncoded() *Client {
	c.header["Content-Type"] = "application/x-www-form-urlencoded"
	return c
}

func (c *Client) SetContentTypejson() *Client {
	c.header["Content-Type"] = "application/json;charset=UTF-8"
	return c
}

func (c *Client) Rest() {
	c.header = map[string]string{}
	c.request = nil
}

func (c *Client) Get(path string, param map[string]string) (respBody []byte, err error) {
	data := url.Values{}
	for k, v := range param {
		data.Set(k, v)
	}
	u, err := url.ParseRequestURI(path)
	if err != nil {
		return nil, err
	}
	u.Path = path
	c.request, err = http.NewRequest(http.MethodGet, u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	for key, value := range c.header {
		c.request.Header.Add(key, value)
	}
	c.request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := c.client.Do(c.request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http status code is not OK")
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) Post(path string, query map[string]string, reqbody []byte) (respBody []byte, err error) {
	data := url.Values{}
	for k, v := range query {
		data.Set(k, v)
	}
	u, err := url.ParseRequestURI(path)
	if err != nil {
		return nil, err
	}
	u.Path = path
	u.RawQuery = data.Encode()
	c.request, err = http.NewRequest(http.MethodGet, u.String(), bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}
	for key, value := range c.header {
		c.request.Header.Add(key, value)
	}
	c.request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := c.client.Do(c.request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http status code is not OK")
	}
	return ioutil.ReadAll(resp.Body)
}
