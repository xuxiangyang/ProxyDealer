package proxy

import (
	"logx"
	"net/http"
	"net/url"
)

func TestGet(proxyAddress, testSite string, validFunc func(http.Response) bool) bool {
	proxyUrl, err := url.Parse("http://" + proxyAddress)
	if err != nil {
		logx.Info(proxy.Address, "not valid with error")
		logx.Error(err)
		return false
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	resp, err := client.Get(testSite)

	if err != nil || resp.StatusCode != 200 {
		return false
	}

	return validFunc(resp)
}
