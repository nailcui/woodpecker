package http

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"woodpecker/core"
	"woodpecker/logger"
)

const (
	Kind string = "http"
)

type Checker struct {
	Name           string   `json:"name" yaml:"name"`
	Disable        bool     `json:"disable" yaml:"disable"`
	Notifier       string   `json:"notifier" yaml:"notifier"`
	ReportTimes    uint     `json:"reportTimes" yaml:"reportTimes"`
	Timeout        int64    `json:"timeout" yaml:"timeout"`
	Url            string   `json:"url" yaml:"url"`
	Cron           string   `json:"cron" yaml:"cron"`
	MustContain    []string `json:"mustContain" yaml:"mustContain"`
	MustNotContain []string `json:"mustNotContain" yaml:"mustNotContain"`
	SuccessCode    []string `json:"successCode" yaml:"successCode"`
	status         *CheckerStatus
	client         *http.Client
}

func (h *Checker) Init() {
	// 默认连续告警次数2，超过后不再告警，恢复后有提醒
	if h.ReportTimes <= 0 {
		h.ReportTimes = 2
	}
	// 默认超时时间 5000ms
	if h.Timeout <= 0 {
		h.Timeout = 5000
	}
	h.status = NewCheckerStatus(h.ReportTimes)

	h.client = &http.Client{
		// 超时时间
		Timeout: time.Duration(h.Timeout) * time.Millisecond,
		Transport: &http.Transport{
			DisableKeepAlives:  true,
			DisableCompression: true,
			TLSClientConfig: &tls.Config{
				// 跳过证书校验
				InsecureSkipVerify: true,
			},
		},
	}
}

func (h *Checker) GetKind() string {
	return Kind
}

func (h *Checker) GetName() string {
	return h.Name
}

func (h *Checker) GetCron() string {
	return h.Cron
}

func (h *Checker) Enabled() bool {
	return !h.Disable
}

func (h *Checker) Check() error {
	err := h.conclude(h.get(h.Url))
	if err != nil {
		if h.status.Report(false) {
			return fmt.Errorf("%s url check 异常【%s】 %s", h.Name, h.Url, err.Error())
		} else {
			logger.Error(fmt.Sprintf("%s url check 异常屏蔽【%s】 %s", h.Name, h.Url, err.Error()))
		}
	} else {
		if h.status.Report(true) {
			return fmt.Errorf("%s url check 异常恢复【%s】", h.Name, h.Url)
		}
		logger.Info(fmt.Sprintf("%s url check success", h.Name))
	}
	return nil
}

func (h *Checker) GetNotifier() string {
	return h.Notifier
}

func (h *Checker) conclude(statusCode int, response string, err error) error {
	if err != nil {
		return err
	}
	if len(h.MustContain) > 0 {
		for i := range h.MustContain {
			if !strings.Contains(response, h.MustContain[i]) {
				return fmt.Errorf("must contain: %s, response: %s", h.MustContain[i], response)
			}
		}
	}
	if len(h.MustNotContain) > 0 {
		for i := range h.MustNotContain {
			if strings.Contains(response, h.MustNotContain[i]) {
				return fmt.Errorf("must not contain: %s, response: %s", h.MustNotContain[i], response)
			}
		}
	}
	if len(h.SuccessCode) > 0 {
		code := strconv.Itoa(statusCode)
		if !core.ContainsStr(h.SuccessCode, code) {
			return fmt.Errorf("stauts code: %s, response: %s", code, response)
		}
	}
	return nil
}

func (h *Checker) get(url string) (int, string, error) {
	resp, err := h.client.Get(url)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	return resp.StatusCode, string(response), nil
}

type CheckerStatus struct {
	// 错误的次数
	errorTimes uint
	// 每次失败允许的告警次数
	ReportTimes uint
}

func NewCheckerStatus(reportTimes uint) *CheckerStatus {
	return &CheckerStatus{ReportTimes: reportTimes}
}

func (cs *CheckerStatus) Report(status bool) bool {
	if status {
		if cs.errorTimes == 0 {
			return false
		}
		cs.errorTimes = 0
		return true
	} else {
		if cs.errorTimes < cs.ReportTimes {
			cs.errorTimes += 1
			return true
		} else {
			return false
		}
	}
}
