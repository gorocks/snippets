package yaml_test

import (
	"testing"

	yaml "gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

type TT struct {
	F int `yaml:"a,omitempty"`
	B int
}

func TestYaml(t *testing.T) {
	t1 := T{}
	if err := yaml.Unmarshal([]byte(data), &t1); err != nil {
		t.Error(err)
	}
	if d, err := yaml.Marshal(t1); err != nil {
		t.Error(err, d)
	}
	if d, _ := yaml.Marshal(&TT{B: 2}); string(d) != "b: 2\n" {
		t.Error(d)
	}
	if d, _ := yaml.Marshal(&TT{F: 1}); string(d) != "a: 1\nb: 0\n" {
		t.Error(d)
	}
}
