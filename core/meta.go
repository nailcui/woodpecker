package core

type TypeMeta struct {
	Kind string `json:"kind" yaml:"kind"`
}

type ObjectMeta struct {
	Name   string            `json:"name" yaml:"name"`
	Labels map[string]string `json:"labels" yaml:"labels"`
}
