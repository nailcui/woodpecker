package core

type ResourceFile struct {
	FilePath    string
	Fingerprint string
	Enable      bool
}

type Resource struct {
	ApiVersion string      `json:"apiVersion" yaml:"apiVersion"`
	Kind       string      `json:"kind" yaml:"kind"`
	Metadata   Metadata    `json:"metadata" yaml:"metadata"`
	Spec       interface{} `json:"spec" yaml:"spec"`
}
