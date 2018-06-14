package got

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

var (
	HttpClient *http.Client
)

func init() {
	HttpClient = &http.Client{
		Transport: &http.Transport{
			Dial: func(netr, addr string) (net.Conn, error) {
				conn, e := net.DialTimeout(netr, addr, time.Second*6)
				if e != nil {
					return nil, e
				}
				conn.SetDeadline(time.Now().Add(time.Second * 6))
				return conn, nil
			},
			MaxIdleConns:          200,
			ResponseHeaderTimeout: time.Second * 6,
		},
		Timeout: 6 * time.Second,
	}
}

func HttpGetJson(url string) (*simplejson.Json, error) {
	res, e := HttpClient.Get(url)
	if e != nil {
		fmt.Println("Request %s error %s", url, e.Error())
		return simplejson.New(), e
	}

	if res.StatusCode != 200 {

		fmt.Println("Request %d %s error ", res.StatusCode, url)
		return simplejson.New(), e
	}

	b, e := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if e != nil {
		fmt.Println("Request %s read body error %s", url, e.Error())
		return simplejson.New(), e
	}

	json, e := simplejson.NewJson(b)
	if e != nil {
		fmt.Println("Json %s ", e.Error())
		return simplejson.New(), e
	}
	return json, nil
}

func HttpPostGetJson(url string, form url.Values) (*simplejson.Json, error) {

	res, e := HttpClient.PostForm(url, form)
	if e != nil {
		fmt.Println("Request %s error %s", url, e.Error())
		return simplejson.New(), e
	}

	if res.StatusCode != 200 {
		fmt.Println("Request %d %s error ", res.StatusCode, url)
		return simplejson.New(), e
	}

	b, e := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if e != nil {
		fmt.Println("Request %s read body error %s", url, e.Error())
		return simplejson.New(), e
	}

	json, e := simplejson.NewJson(b)
	if e != nil {
		fmt.Println("Json %s ", e.Error())
		return simplejson.New(), e
	}
	return json, nil
}
