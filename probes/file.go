package probes

import (
	"errors"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func init() {
	if ProbesPlugin == nil {
		ProbesPlugin = make(map[string]ProbesFactory)
	}
	ProbesPlugin["file"] = NewFile
}

// File storing probes in YAML
type File struct {
	MapProbes
}

type fileProbes []fileProbe
type fileProbe string

func NewFile(cfg map[string]interface{}) (Probes, error) {
	l := log.WithField("cfg", cfg)
	const KEY = "path"
	raw, ok := cfg[KEY]
	if !ok {
		err := fmt.Errorf(`"%s" key is mandatory`, KEY)
		l.WithError(err).Error()
		return nil, err
	}
	path, ok := raw.(string)
	if !ok {
		err := errors.New(`"file" key must be a string`)
		l.WithError(err).Error()
		return nil, err
	}
	return NewFileFromPath(path)
}

// NewFileFromYAML parse YAML
func NewFileFromYAML(raw []byte) (*File, error) {
	var probes fileProbes
	err := yaml.Unmarshal(raw, &probes)
	if err != nil {
		return nil, err
	}
	pp := make([]string, len(probes))
	for i, p := range probes {
		pp[i] = string(p)
	}
	f := &File{*NewMapProbes(pp)}
	return f, nil
}

// NewFileFromPath from a path
func NewFileFromPath(path string) (*File, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewFileFromYAML(raw)
}
