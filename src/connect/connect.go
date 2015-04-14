package connect

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"logx"
	"net/http"
	"net/url"
	"time"
)

const (
	TestUrl = `http://211.155.88.207:9191/`
)

type TestResult struct {
	Proxy *url.URL
	Time  int
	Ok    bool
}

type responseJson struct {
	Success bool `json:success`
}

func Test(proxy *url.URL) *TestResult {
	getTime, getOk := testGet(proxy)
	postTime, postOk := testPost(proxy)
	return &TestResult{Proxy: proxy, Time: max(getTime, postTime), Ok: getOk && postOk}
}

func testGet(proxy *url.URL) (time int, ok bool) {
	getFunc := func(client *http.Client) (*http.Response, error) {
		return client.Get(TestUrl)
	}
	return proxyedResponseTest(proxy, getFunc, respValidFunc)
}

func testPost(proxy *url.URL) (time int, ok bool) {
	postFunc := func(client *http.Client) (*http.Response, error) {
		jsonStr := []byte(`{"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}`)
		return client.Post(TestUrl, "json", bytes.NewBuffer(jsonStr))
	}
	return proxyedResponseTest(proxy, postFunc, respValidFunc)
}

func respValidFunc(resp *http.Response) bool {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logx.Warn(resp.StatusCode)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Warn(err)
		return false
	}

	var rj responseJson
	err = json.Unmarshal(body, &rj)
	if err != nil {
		logx.Warn(err)
		return false
	}
	return rj.Success
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func proxyedResponseTest(proxy *url.URL, fetchFunc func(*http.Client) (*http.Response, error), validFunc func(*http.Response) bool) (int, bool) {
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	startTime := time.Now()
	resp, err := fetchFunc(client)
	endTime := time.Now()
	if err != nil || !validFunc(resp) {
		logx.Warn(err)
		return -1, false
	}
	return int(endTime.Sub(startTime).Seconds()), true
}
