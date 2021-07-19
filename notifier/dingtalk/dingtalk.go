package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
	"woodpecker/pkg/httpclient"
)

const Kind = "dingtalk"

type Notifier struct {
	Name        string `json:"name" yaml:"name"`
	AccessToken string `json:"accessToken" yaml:"accessToken"`
	Secret      string `json:"secret" yaml:"secret"`
}

func (n *Notifier) GetKind() string {
	return Kind
}

func (n *Notifier) GetName() string {
	return n.Name
}

func (n *Notifier) Send(message string) error {
	now := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%s&sign=%s",
		n.AccessToken, now, sign(now, n.Secret))
	postBody, err := json.Marshal(body{
		Msgtype: "markdown",
		Markdown: markdown{
			Title: "title",
			Text:  message,
		},
	})
	if err != nil {
		return err
	}
	return httpclient.PostJson(url, string(postBody))
}

func (n *Notifier) ParamsCheck() error {
	if n.AccessToken == "" {
		return fmt.Errorf("dingtald notifier accessToken not be null")
	}
	if n.Secret == "" {
		return fmt.Errorf("dingtald notifier secret not be null")
	}
	return nil
}

type markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type body struct {
	Msgtype  string   `json:"msgtype"`
	Markdown markdown `json:"markdown"`
}

func sign(now, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(now + "\n" + secret))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
