package probes

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// File storing probes in YAML
type File struct {
	MapProbes
}

type fileProbes []fileProbe
type fileProbe string

// NewFileFromYAML parse YAML
func NewFileFromYAML(raw []byte) (*File, error) {
	var probes fileProbes
	err := yaml.Unmarshal(raw, &probes)
	if err != nil {
		return nil, err
	}
	f := &File{*NewMapProbes()}
	for _, p := range probes {
		f.probes[string(p)] = new(interface{})
	}
	return f, nil
}

// NewFile from a path
func NewFile(path string) (*File, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewFileFromYAML(raw)
}
