package proxy

import (
	"bytes"
	"io"
	"logx"
	"net/http"
	"net/url"
)

const (
	TIME_OUT = 30
)

func Get(proxyAddress, testSite string, validFunc func(http.Response) bool) bool {
	client, err := createProxyedClient(proxyAddress)
	if err != nil {
		return false
	}
	resp, err := client.Get(testSite)

	if err != nil || resp.StatusCode != 200 {
		return false
	}

	return validFunc(resp)
}

func PostWithBlankJson(proxyAddress, testSite string, validFunc func(http.Response) bool) bool {
	return TestPost(proxyAddress, testSite, "application/json", bytes.NewBuffer([]byte("{}")), validFunc)
}

func Post(proxyAddress, testSite, bodyType string, body io.Reader, validFunc func(http.Response) bool) bool {
	client, err := createProxyedClient(proxyAddress)
	if err != nil {
		return false
	}

	resp, err := client.Post(testSite, bodyType, body)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return validFunc
}

func createProxyedClient(proxyAddress string) (*http.Client, error) {
	proxyUrl, err := url.Parse("http://" + proxyAddress)
	if err != nil {
		logx.Info(proxy.Address, "not valid with error")
		logx.Error(err)
		return nil, err
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}, Timeout: TIME_OUT}
	return client, nil
}
