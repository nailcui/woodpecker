package core

import (
	"encoding/json"
	yj "github.com/ghodss/yaml"
	"gopkg.in/yaml.v2"
)

func Interface2Interface(source interface{}, target interface{}) error {
	out, err := yaml.Marshal(source)
	if err != nil {
		return err
	}
	jsonByte, err := yj.YAMLToJSON(out)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonByte, target)
}
