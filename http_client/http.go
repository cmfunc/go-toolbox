package httpclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ReqHttp struct{}

// GetWithURL http get method request
func (s *ReqHttp) GetWithURL(requestUrl string, paramsMap map[string]string) (res []byte, err error) {

	Url, err := url.Parse(requestUrl)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	for k, v := range paramsMap {
		params.Set(k, v)
	}

	Url.RawQuery = params.Encode()

	requestUrl = Url.String()

	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, errors.WithMessagef(err, "[MarketService] GetWithURL url:[%s] paramsMap:[%v]", requestUrl, paramsMap)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// http post json
func (s *ReqHttp) PostWithJSON(apiUrl string, in []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(in)

	request, err := http.NewRequest(http.MethodPost, apiUrl, buffer)
	if err != nil {
		return nil, errors.WithMessage(err, "[MarketService] PostWithJSON http post request generate failed")
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	cli := http.Client{}

	resp, err := cli.Do(request.WithContext(context.TODO()))
	if err != nil {
		return nil, errors.WithMessage(err, "[MarketService] PostWithJSON http post request send failed")
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "[MarketService] PostWithJSON...ioutil.ReadAll")
	}

	return respBodyBytes, nil

}

// HTTP POST 方法统一使用contentType 为“application/x-www-form-urlencoded”
func (s *ReqHttp) PostWithUrlencoded(apiUrl, resource string, keyValue map[string]string) ([]byte, error) {
	data := url.Values{}

	for k, v := range keyValue {
		data.Set(k, v)
	}

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		return nil, errors.WithMessagef(err, "[ReqHttp] PostWithUrlencoded apiUrl=[%s] keyValue=[%+v] ", apiUrl, keyValue)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("[ReqHttp] PostWithUrlencoded apiUrl=[%s] keyValue=[%+v] StatusCode=[%d]", apiUrl, keyValue, resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
