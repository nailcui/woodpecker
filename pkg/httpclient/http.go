package httpclient

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"woodpecker/logger"
)

func PostJson(url string, body string) error {
	client := &http.Client{
		// 超时时间
		Timeout: 5000 * time.Millisecond,
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	resp, err := client.Post(url, "application/json; charset=utf-8", strings.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	logger.Debug("http post", url, string(response))
	if err != nil {
		return err
	}
	return err
}
