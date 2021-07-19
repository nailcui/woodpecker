package notifier

import (
	"fmt"
	"woodpecker/core"
	"woodpecker/notifier/dingtalk"
)

var builders = map[string]func(template *NotifierTemplate) (Notifier, error){}

func init() {
	builders[dingtalk.Kind] = newDingtalkNotifier
}

type Notifier interface {
	GetKind() string
	GetName() string
	Send(string) error
	ParamsCheck() error
}

type NotifierSpec struct {
	Kind string      `json:"kind" yaml:"kind"`
	Spec interface{} `json:"spec" yaml:"spec"`
}

type NotifierTemplate struct {
	ApiVersion string        `json:"apiVersion" yaml:"apiVersion"`
	Kind       string        `json:"kind" yaml:"kind"`
	Metadata   core.Metadata `json:"metadata" yaml:"metadata"`
	Spec       NotifierSpec  `json:"spec" yaml:"spec"`
}

func NewNotifier(template *NotifierTemplate) (Notifier, error) {
	f := builders[template.Spec.Kind]
	if f == nil {
		return nil, fmt.Errorf("notifier kind: %s not found", template.Spec.Kind)
	}
	return f(template)
}

func newDingtalkNotifier(template *NotifierTemplate) (Notifier, error) {
	nf := dingtalk.Notifier{}
	err := core.Interface2Interface(template.Spec.Spec, &nf)
	if err != nil {
		return nil, err
	}
	nf.Name = template.Metadata.Name
	return &nf, nil
}
