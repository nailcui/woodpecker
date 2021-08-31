package tcp

import (
	"fmt"
	"net"
	"time"
	"woodpecker/logger"
)

const (
	Kind string = "tcp"
)

type Checker struct {
	Name           string   `json:"name" yaml:"name"`
	Disable        bool     `json:"disable" yaml:"disable"`
	Notifier       string   `json:"notifier" yaml:"notifier"`
	ReportTimes    uint     `json:"reportTimes" yaml:"reportTimes"`
	Timeout        int64    `json:"timeout" yaml:"timeout"`
	Address        string   `json:"address" yaml:"address"`
	Cron           string   `json:"cron" yaml:"cron"`
	status         *CheckerStatus
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
	err := h.conclude()
	if err != nil {
		if h.status.Report(false) {
			return fmt.Errorf("%s address check 异常【%s】 %s", h.Name, h.Address, err.Error())
		} else {
			logger.Error(fmt.Sprintf("%s address check 异常屏蔽【%s】 %s", h.Name, h.Address, err.Error()))
		}
	} else {
		if h.status.Report(true) {
			return fmt.Errorf("%s address check 异常恢复【%s】", h.Name, h.Address)
		}
		logger.Info(fmt.Sprintf("%s address check success", h.Name))
	}
	return nil
}

func (h *Checker) GetNotifier() string {
	return h.Notifier
}

func (h *Checker) conclude() error {
	conn, err := net.DialTimeout("tcp", h.Address, time.Duration(h.Timeout) * time.Millisecond)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
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
