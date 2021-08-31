package checker

import (
	"fmt"
	"woodpecker/checker/http"
	"woodpecker/checker/tcp"
	"woodpecker/core"
)

var builders = map[string]func(template *CheckerTemplate) (Checker, error){}

func init() {
	builders[tcp.Kind] = newTcpChecker
	builders[http.Kind] = newHttpChecker
}

type Checker interface {
	Init()
	GetKind() string
	GetName() string
	GetCron() string
	Enabled() bool
	Check() error
	GetNotifier() string
}

type CheckerSpec struct {
	Kind string      `json:"kind" yaml:"kind"`
	Spec interface{} `json:"spec" yaml:"spec"`
}

type CheckerTemplate struct {
	ApiVersion string        `json:"apiVersion" yaml:"apiVersion"`
	Kind       string        `json:"kind" yaml:"kind"`
	Metadata   core.Metadata `json:"metadata" yaml:"metadata"`
	Spec       CheckerSpec   `json:"spec" yaml:"spec"`
}

func NewChecker(template *CheckerTemplate) (Checker, error) {
	f := builders[template.Spec.Kind]
	if f == nil {
		return nil, fmt.Errorf("checker kind: %s not found", template.Spec.Kind)
	}
	return f(template)
}

func newHttpChecker(template *CheckerTemplate) (Checker, error) {
	c := http.Checker{}
	err := core.Interface2Interface(template.Spec.Spec, &c)
	if err != nil {
		return nil, err
	}
	c.Name = template.Metadata.Name
	return &c, nil
}

func newTcpChecker(template *CheckerTemplate) (Checker, error) {
	c := tcp.Checker{}
	err := core.Interface2Interface(template.Spec.Spec, &c)
	if err != nil {
		return nil, err
	}
	c.Name = template.Metadata.Name
	return &c, nil
}
